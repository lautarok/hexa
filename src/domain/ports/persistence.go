package ports

import "context"

type PersistencePort interface {
	Transaction(
		ctx context.Context,
		function func(ctx context.Context) error,
	) error
	IsUniqueViolation(err error) bool
	IsErrNotFound(err error) bool
}
