package models

import "time"

type PolicyHolder struct {
	PolicyNumber   string
	FirsName       string
	LastName       string
	DOB            time.Time
	AddressDetails Address
	Gender         GenderType
	Phone          string
	Email          string
	FromDate       time.Time
	ToDate         time.Time
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Zipcode      string
}

type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

var policyHolderMap = make(map[string]*PolicyHolder)

func (p *PolicyHolder) AddPolicyDetails() (bool, error) {
	policyHolderMap[p.PolicyNumber] = p
	return true, nil
}
