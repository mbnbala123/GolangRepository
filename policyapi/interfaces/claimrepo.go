package interfaces

import "policymanagement/policyapi/models"

type IClaimRepo interface {
	Save() (int64, error)
	GetByClaimID(id uint) (*models.Claim, error)
	GetAllClaim() ([]*models.Claim, error)
	UpdateClaim(claimID uint, claimAmount int) (*models.Claim, error)
}
