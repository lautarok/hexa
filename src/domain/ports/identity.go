package ports

import "github.com/lautarok/hexa/src/domain/models"

type IdentityPort interface {
	NewToken(identity *models.Identity) (string, int64, error)
	ParseToken(token string) (*models.Identity, error)
}
