package orminterfaces

type MemberRepo interface {
	SaveMember() (bool, error)
}
