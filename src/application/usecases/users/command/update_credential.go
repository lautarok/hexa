package command

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
)

type UpdateCredentialUsecase struct {
	credentialsRepository models.ICredentialsRepository
}

type UpdateCredentialUsecaseDeps struct {
	CredentialsRepository models.ICredentialsRepository
}

func NewUpdateCredentialsUsecase(deps *UpdateCredentialUsecaseDeps) *UpdateCredentialUsecase {
	return &UpdateCredentialUsecase{
		credentialsRepository: deps.CredentialsRepository,
	}
}

type UpdateCredentialUsecaseInput struct {
	UserID   uuid.UUID
	Username string
}

func (usecase *UpdateCredentialUsecase) UpdateCredential(ctx context.Context, input *UpdateCredentialUsecaseInput) (*models.Credential, *models.AppError) {
	credential, err := usecase.credentialsRepository.UpdateOne(ctx, &models.Credential{
		UserID:   input.UserID,
		Username: input.Username,
	})
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return credential, nil
}
