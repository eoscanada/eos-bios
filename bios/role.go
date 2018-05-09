package bios

type Role int

const (
	RoleBootNode = Role(iota)
	RoleABP
	RoleParticipant
)
