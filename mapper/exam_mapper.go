package mapper

import (
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func ToExamViewModel(e models.Exam) viewmodels.ExamViewModel {
	return viewmodels.ExamViewModel{
		ID:           e.ID,
		PatientName:  e.Patient.FirstName + " " + e.Patient.LastName,
		ExamTypeName: e.ExamType.Description,
		Result:       e.Result,
		CreatedAt:    e.CreatedAt.Format("2006-01-02"),
	}
}

func ToExamViewModelList(exams []models.Exam) []viewmodels.ExamViewModel {
	vms := make([]viewmodels.ExamViewModel, len(exams))
	for i, e := range exams {
		vms[i] = ToExamViewModel(e)
	}
	return vms
}
