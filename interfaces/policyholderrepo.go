package interfaces

type IPolicyHolderRepo interface {
	AddPolicyDetails() (bool, error)
}
