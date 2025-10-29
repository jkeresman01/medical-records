package viewmodels

type PatientViewModel struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DOB       string `json:"dob"`
}

type PatientDetailViewModel struct {
	ID            uint                    `json:"id"`
	FirstName     string                  `json:"first_name"`
	LastName      string                  `json:"last_name"`
	DOB           string                  `json:"dob"`
	Prescriptions []PrescriptionViewModel `json:"prescriptions"`
	Exams         []ExamViewModel         `json:"exams"`
}
