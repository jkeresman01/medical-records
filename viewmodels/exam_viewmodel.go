package viewmodels

type ExamViewModel struct {
	ID           uint   `json:"id"`
	PatientName  string `json:"patient_name"`
	ExamTypeName string `json:"exam_type_name"`
	Result       string `json:"result"`
	CreatedAt    string `json:"created_at"`
}
