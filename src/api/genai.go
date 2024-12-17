package api

import (
	"context"
	"diawise/src/services"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type AIHealthAnalyser struct {
	client *genai.Client
}

func NewAIHealthAnalyser() (*AIHealthAnalyser, error) {
	// Get API key from environment variable
	apiKey := os.Getenv("GOOGLE_AI_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set GOOGLE_AI_API_KEY environment variable")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create ai client: %v", err)
	}
	return &AIHealthAnalyser{client: client}, nil
}

func (a *AIHealthAnalyser) Close() {
	a.client.Close()
}

func (a *AIHealthAnalyser) DietProfile(mealEntry *services.MealLogEntry) (*services.DietProfile, error) {
	fmt.Println("Calling gen ai client...")
	ctx := context.Background()
	model := a.client.GenerativeModel("gemini-2.0-flash-exp")

	prompt := fmt.Sprintf(`Provide a diet profile analysis for:
	- Food name: %s
	- Weight: %.2f g
	- Proportion of plate: %.2f

	Provide the response in the format:
	CaloriesIntake     float64
	CarbIntake         float64
	ProteinIntake      float64
	FatIntake          float64
	SugarConsumption   float64
	ProcessedFoodRatio float64`, mealEntry.FoodItem, mealEntry.Weight, mealEntry.Proportion)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate diet profile analysis: %v", err)
	}
	var respStr string
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		respStr = fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	}
	var dietProfile services.DietProfile
	dietProfile.UserID = mealEntry.UserID
	dietProfile.FoodName = mealEntry.FoodItem
	dietProfile.MealType = mealEntry.MealType
	err = dietProfile.ParseDietProfileString(respStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse diet profile analysis: %v", err)
	}
	return &dietProfile, nil
}
