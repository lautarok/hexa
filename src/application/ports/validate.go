package ports

type ValidationPort interface {
	Struct(structure any) error
}
