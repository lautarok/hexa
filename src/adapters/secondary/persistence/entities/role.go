package entities

type Role struct {
	ID string `bun:"ID,type:uuid,notnull,pk,default:uuid_generate_v4()"`
}
