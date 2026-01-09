package usecases

import (
	"context"
	"log"

	"github.com/lautarok/hexa/src/app/domain"
	dto "github.com/lautarok/hexa/src/app/dtos"
	"github.com/lautarok/hexa/src/app/ports"
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
	Username string
	Password string
}

func (usecase *SignupUsecase) Signup(ctx context.Context, input *SignupUsecaseInput) (*dto.TokenOutputDto, *domain.AppError) {
	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		insertedUser, repoErr := usecase.usersRepository.CreateOne(ctx, &domain.User{
			Name:    input.Name,
			Surname: input.Surname,
		})

		if repoErr != nil {
			log.Println("!!!!!!!!!!!!")
			return repoErr
		}

		_, repoErr = usecase.credentialsRepository.CreateOne(ctx, &domain.Credential{
			Username:     input.Username,
			PasswordHash: input.Password,
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
		return nil, &domain.AppError{
			Code:    "InternalError",
			Message: "Internal error: " + err.Error(),
		}
	}

	return &dto.TokenOutputDto{
		Token: "Hola mundo",
	}, nil
}
