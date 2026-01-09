package usecases

import (
	"context"

	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
)

type SignupUsecase struct {
	credentialsRepository domain.ICredentialsRepository
	usersRepository       domain.IUsersRepository
	persistenceAdapter    ports.PersistencePort
}

type SignupUsecaseDeps struct {
	CredentialsRepository domain.ICredentialsRepository
	UsersRepository       domain.IUsersRepository
	PersistenceAdapter    ports.PersistencePort
}

func NewSignupUsecase(deps *SignupUsecaseDeps) *SignupUsecase {
	return &SignupUsecase{
		credentialsRepository: deps.CredentialsRepository,
		usersRepository:       deps.UsersRepository,
		persistenceAdapter:    deps.PersistenceAdapter,
	}
}

type SignupUsecaseInput struct {
	Name     string
	Surname  string
	Email    string
	Username string
	Password string
}

func (usecase *SignupUsecase) Signup(ctx context.Context, input *SignupUsecaseInput) (string, *domain.AppError) {
	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		insertedUser, repoErr := usecase.usersRepository.CreateOne(ctx, &domain.User{
			Name:    input.Name,
			Surname: input.Surname,
		})

		if repoErr != nil {
			return repoErr
		}

		_, repoErr = usecase.credentialsRepository.CreateOne(ctx, &domain.Credential{
			Username: input.Username,
			Password: input.Password,
			Email:    input.Email,
			User: &domain.User{
				ID: insertedUser.ID,
			},
		})

		if repoErr != nil {
			return repoErr
		}

		return nil
	})

	if err != nil {
		if usecase.persistenceAdapter.IsUniqueViolation(err) {
			return "", errors.NewAlreadyExistsError("Email or username already exists")
		}

		return "", errors.NewInternalError(err)
	}

	return "hardcoded_token_xd", nil
}
