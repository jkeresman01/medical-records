package mapper

import (
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func ToPrescriptionViewModel(p models.Prescription) viewmodels.PrescriptionViewModel {
	return viewmodels.PrescriptionViewModel{
		ID:             p.ID,
		PatientName:    p.Patient.FirstName + " " + p.Patient.LastName,
		MedicationName: p.Medication.Name,
		Dosage:         p.Dosage,
		Frequency:      p.Frequency,
	}
}

func ToPrescriptionViewModelList(prescriptions []models.Prescription) []viewmodels.PrescriptionViewModel {
	vms := make([]viewmodels.PrescriptionViewModel, len(prescriptions))
	for i, p := range prescriptions {
		vms[i] = ToPrescriptionViewModel(p)
	}
	return vms
}

