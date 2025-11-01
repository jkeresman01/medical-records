package mapper

import (
	"github.com/jkeresman01/medical-records/models"
	"github.com/jkeresman01/medical-records/viewmodels"
)

func ToMedicationViewModel(m models.Medication) viewmodels.MedicationViewModel {
	return viewmodels.MedicationViewModel{
		ID:           m.ID,
		Name:         m.Name,
		Manufacturer: m.Manufacturer,
	}
}

func ToMedicationViewModelList(medications []models.Medication) []viewmodels.MedicationViewModel {
	vms := make([]viewmodels.MedicationViewModel, len(medications))
	for i, m := range medications {
		vms[i] = ToMedicationViewModel(m)
	}
	return vms
}
