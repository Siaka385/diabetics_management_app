package api

import "diawise/internal/services"

type MedicationHandler struct {
	service *services.MedicationService
}

// handler for managing new medications
func NewMedicalHandler(services *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: services}
}
