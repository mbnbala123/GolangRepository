package interfaces

import "policymanagement/models"

type IVehicleRepo interface {
	Save() (bool, error)
	GetByID(id string) (*models.Vehicle, error)
	GetAll() ([]*models.Vehicle, error)
	Update(id string, color string) (*models.Vehicle, error)
	Delete(id string) (bool, error)
}
