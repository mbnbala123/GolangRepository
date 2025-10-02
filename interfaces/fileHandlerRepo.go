package interfaces

import "policymanagement/models"

type IFileHandlerRepo interface {
	SaveToFile(fileName string, headers []string, vehicleModels []*models.Vehicle) (bool, error)
}
