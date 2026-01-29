package ports

type PasswordPort interface {
	Hash(password string) (string, error)
	Compare(hash string, password string) bool
}
