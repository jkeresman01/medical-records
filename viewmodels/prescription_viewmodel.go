package viewmodels

type PrescriptionViewModel struct {
	ID             uint   `json:"id"`
	PatientID      uint   `json:"patient_id"`
	MedicationID   uint   `json:"medication_id"`
	PatientName    string `json:"patient_name"`
	MedicationName string `json:"medication_name"`
	Dosage         string `json:"dosage"`
	Frequency      string `json:"frequency"`
}
