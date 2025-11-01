package mapper

import (
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func ToExamTypeViewModel(et models.ExamType) viewmodels.ExamTypeViewModel {
	return viewmodels.ExamTypeViewModel{
		ID:          et.ID,
		Name:        et.Name,
		Description: et.Description,
	}
}

func ToExamTypeViewModelList(examTypes []models.ExamType) []viewmodels.ExamTypeViewModel {
	vms := make([]viewmodels.ExamTypeViewModel, len(examTypes))
	for i, et := range examTypes {
		vms[i] = ToExamTypeViewModel(et)
	}
	return vms
}
