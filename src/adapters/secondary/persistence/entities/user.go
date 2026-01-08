package entities

type User struct {
	ID      string `bun:"type:uuid,pk,default:generate_uuid_v4()"`
	Name    string
	Surname string
}
