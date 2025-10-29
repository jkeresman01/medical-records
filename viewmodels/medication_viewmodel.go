package viewmodels

type MedicationViewModel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
}
