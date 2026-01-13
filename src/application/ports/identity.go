package ports

import "github.com/lautarok/hexa/src/application/domain"

type IdentityPort interface {
	NewToken(identity *domain.Identity) (string, int64, error)
	ParseToken(token string) (*domain.Identity, error)
}
