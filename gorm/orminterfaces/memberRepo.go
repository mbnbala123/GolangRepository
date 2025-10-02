package orminterfaces

import "policymanagement/gorm/DBStore"

type MemberRepo interface {
	SaveMember() (bool, error)
	GetAllMembers() ([]*DBStore.Member, error)
}
