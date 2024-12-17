package api

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	auth "diawise/src/auth"

	"gorm.io/gorm"
)

// Room model for database persistence
type Room struct {
	gorm.Model
	RoomID         string `gorm:"unique;index"`
	Name           string
	ActiveMessages []string               `gorm:"-"`
	Members        map[string]chan string `gorm:"-"`
	mu             sync.RWMutex           `gorm:"-"`
}

// Message represents a chat message (not persisted)
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

var (
	rooms   = make(map[string]*Room)
	roomsMu sync.RWMutex
)

func Support(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// user, ok := auth.GetUserFromContext(r)
		_, ok := auth.GetUserFromContext(r)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Log user details (optional)
		// fmt.Printf("Authenticated user: %+v\n", user.Name)
		// fmt.Printf("Authenticated user ID: %+v\n", user.ID)


		if err := tmpl.ExecuteTemplate(w, "support.html", Data); err != nil {
			InternalServerErrorHandler(w)
			return
		}
	}
}

func CreateRoom(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var roomData struct {
			Name string `json:"name"`
		}

		err := json.NewDecoder(r.Body).Decode(&roomData)
		if err != nil {
			http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
			return
		}

		room := createRoom(db, roomData.Name)
		if room == nil {
			http.Error(w, `{"error": "Failed to create room"}`, http.StatusInternalServerError)
			return
		}

		// Return a successful response as JSON
		json.NewEncoder(w).Encode(map[string]string{
			"roomId": room.RoomID,
			"name":   room.Name,
		})
	}
}

// In the listRoomsHandler function, update the response to include a delete button:
func ListRooms(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := listRooms(db)
		if err != nil {
			http.Error(w, "Failed to list rooms", http.StatusInternalServerError)
			return
		}

		var roomList []map[string]string
		for _, rm := range rooms {
			roomList = append(roomList, map[string]string{
				"roomId": rm.RoomID,
				"name":   rm.Name,
			})
		}

		// Add delete buttons to each room
		for _, room := range roomList {
			room["deleteButton"] = fmt.Sprintf(`<button onclick="deleteRoom('%s')">Delete</button>`, room["roomId"])
		}

		// Return room list with delete buttons
		json.NewEncoder(w).Encode(roomList)
	}
}

func JoinRoom(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomID := r.URL.Query().Get("roomId")
		userID := r.URL.Query().Get("userId")

		roomsMu.RLock()
		room, exists := rooms[roomID]
		roomsMu.RUnlock()

		if !exists {
			// Try to fetch from database
			var dbRoom Room
			result := db.Where("room_id = ?", roomID).First(&dbRoom)
			if result.Error != nil {
				http.Error(w, "Room not found", http.StatusNotFound)
				return
			}

			// Recreate room in memory
			room = &Room{
				RoomID:         dbRoom.RoomID,
				Name:           dbRoom.Name,
				Members:        make(map[string]chan string),
				ActiveMessages: []string{},
			}

			roomsMu.Lock()
			rooms[roomID] = room
			roomsMu.Unlock()
		}

		channel := room.join(userID)

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-open")

		f, _ := w.(http.Flusher)

		// Send initial connection message
		fmt.Fprintf(w, "data: %s\n\n", `{"type":"system","content":"Joined room"}`)
		f.Flush()

		for msg := range channel {
			fmt.Fprintf(w, "data: %s\n\n", msg)
			f.Flush()
		}
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var messageData struct {
		RoomID  string `json:"roomId"`
		UserID  string `json:"userId"`
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	roomsMu.RLock()
	room, exists := rooms[messageData.RoomID]
	roomsMu.RUnlock()

	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	room.broadcast(messageData.UserID, messageData.Message)
	w.WriteHeader(http.StatusOK)
}

// Add a new handler to delete rooms:
func DeleteRoom(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomID := r.URL.Query().Get("roomId")
		fmt.Println("IRD: ", roomID)

		// Delete room from the database
		result := db.Where("room_id = ?", roomID).Delete(&Room{})
		if result.Error != nil {
			http.Error(w, "Failed to delete room", http.StatusInternalServerError)
			return
		}

		// Remove from in-memory rooms
		roomsMu.Lock()
		delete(rooms, roomID)
		roomsMu.Unlock()

		w.WriteHeader(http.StatusOK)
	}
}

// Generate a random pseudo name
func generatePseudoName() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return "Anonymous"
	}
	return "" + hex.EncodeToString(b)
}

// Create a new room
func createRoom(db *gorm.DB, name string) *Room {
	roomID := generatePseudoName()
	room := &Room{
		RoomID:         roomID,
		Name:           name,
		Members:        make(map[string]chan string),
		ActiveMessages: []string{},
	}

	// Save to database
	result := db.Create(room)
	if result.Error != nil {
		log.Printf("Error saving room: %v", result.Error)
		return nil
	}

	roomsMu.Lock()
	rooms[roomID] = room
	roomsMu.Unlock()

	return room
}

func (r *Room) join(userID string) chan string {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Create a channel for the user
	channel := make(chan string, 100)
	r.Members[userID] = channel

	// Send a system message for the user joining the room
	joinMessage := fmt.Sprintf("User %s joined the room", userID)
	systemMsg := Message{
		Sender:  "System",
		Content: joinMessage,
	}
	// Convert message to JSON and send to the new user
	jsonMsg, _ := json.Marshal(systemMsg)
	channel <- string(jsonMsg)

	// Send recent messages to the new member
	for _, msg := range r.ActiveMessages {
		channel <- msg
	}

	return channel
}

// broadcast a message to room members
func (r *Room) broadcast(senderID, message string) {
	r.mu.Lock()

	// Store recent message
	msg := Message{
		Sender:  senderID,
		Content: message,
	}

	jsonMsg, _ := json.Marshal(msg)
	msgStr := string(jsonMsg)

	// Keep only last 50 messages
	r.ActiveMessages = append(r.ActiveMessages, msgStr)
	if len(r.ActiveMessages) > 50 {
		r.ActiveMessages = r.ActiveMessages[1:]
	}

	r.mu.Unlock()

	r.mu.RLock()
	defer r.mu.RUnlock()

	for memberID, ch := range r.Members {
		if memberID != senderID {
			ch <- msgStr
		}
	}
}

// List all rooms
func listRooms(db *gorm.DB) ([]Room, error) {
	var rooms []Room
	result := db.Find(&rooms)
	return rooms, result.Error
}
