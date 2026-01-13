package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/lautarok/hexa/src/application/domain"
	"github.com/lautarok/hexa/src/application/errors"
	"github.com/lautarok/hexa/src/application/ports"
)

type SignupUsecase struct {
	credentialsRepository domain.ICredentialsRepository
	usersRepository       domain.IUsersRepository
	rolesRepository       domain.IRolesRepository
	persistenceAdapter    ports.PersistencePort
	identityAdapter       ports.IdentityPort
}

type SignupUsecaseDeps struct {
	CredentialsRepository domain.ICredentialsRepository
	UsersRepository       domain.IUsersRepository
	RolesRepository       domain.IRolesRepository
	PersistenceAdapter    ports.PersistencePort
	IdentityAdapter       ports.IdentityPort
}

func NewSignupUsecase(deps *SignupUsecaseDeps) *SignupUsecase {
	return &SignupUsecase{
		credentialsRepository: deps.CredentialsRepository,
		usersRepository:       deps.UsersRepository,
		rolesRepository:       deps.RolesRepository,
		persistenceAdapter:    deps.PersistenceAdapter,
		identityAdapter:       deps.IdentityAdapter,
	}
}

type SignupUsecaseInput struct {
	Name     string
	Surname  string
	Email    string
	Username string
	Password string
	RoleID   uuid.UUID
}

type SignupUsecaseOutput struct {
	Token string
	Exp   int64
}

func (usecase *SignupUsecase) Signup(ctx context.Context, input *SignupUsecaseInput) (*SignupUsecaseOutput, *domain.AppError) {
	var insertedUser *domain.User
	var insertedCredential *domain.Credential
	var repoErr error

	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		insertedUser, repoErr = usecase.usersRepository.CreateOne(ctx, &domain.User{
			Name:    input.Name,
			Surname: input.Surname,
			Role: &domain.Role{
				ID: input.RoleID,
			},
		})

		if repoErr != nil {
			return repoErr
		}

		insertedCredential, repoErr = usecase.credentialsRepository.CreateOne(ctx, &domain.Credential{
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
			return nil, errors.NewAlreadyExistsError("Email or username already exists")
		}

		return nil, errors.NewInternalError(err)
	}

	token, exp, err := usecase.identityAdapter.NewToken(&domain.Identity{
		SubUserID: insertedUser.ID,
		UserID:    insertedUser.ID,
		Email:     insertedCredential.Email,
		Username:  insertedCredential.Username,
	})

	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return &SignupUsecaseOutput{
		Token: token,
		Exp:   exp,
	}, nil
}
