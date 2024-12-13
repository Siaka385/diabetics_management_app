package api

type MedicationHandler struct {
	service *services.MedicalService
}

// handler for managing new medications
func NewMedicalHandler(services *services.MedicalService) *MedicalHandler {
	return &MedicationHandler{service: services}
}
