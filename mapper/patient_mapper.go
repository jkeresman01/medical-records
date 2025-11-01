package mapper

import (
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func ToPatientViewModel(p models.Patient) viewmodels.PatientViewModel {
	return viewmodels.PatientViewModel{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		DOB:       p.DOB,
	}
}

func ToPatientViewModelList(patients []models.Patient) []viewmodels.PatientViewModel {
	vms := make([]viewmodels.PatientViewModel, len(patients))
	for i, p := range patients {
		vms[i] = ToPatientViewModel(p)
	}
	return vms
}

func ToPatientDetailViewModel(p models.Patient) viewmodels.PatientDetailViewModel {
	return viewmodels.PatientDetailViewModel{
		ID:            p.ID,
		FirstName:     p.FirstName,
		LastName:      p.LastName,
		DOB:           p.DOB,
		Prescriptions: ToPrescriptionViewModelList(p.Prescriptions),
		Exams:         ToExamViewModelList(p.Exams),
	}
}
