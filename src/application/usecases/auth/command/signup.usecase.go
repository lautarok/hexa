package command

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
	passwordAdapter       ports.PasswordPort
}

type SignupUsecaseDeps struct {
	CredentialsRepository domain.ICredentialsRepository
	UsersRepository       domain.IUsersRepository
	RolesRepository       domain.IRolesRepository
	PersistenceAdapter    ports.PersistencePort
	IdentityAdapter       ports.IdentityPort
	PasswordAdapter       ports.PasswordPort
}

func NewSignupUsecase(deps *SignupUsecaseDeps) *SignupUsecase {
	return &SignupUsecase{
		credentialsRepository: deps.CredentialsRepository,
		usersRepository:       deps.UsersRepository,
		rolesRepository:       deps.RolesRepository,
		persistenceAdapter:    deps.PersistenceAdapter,
		identityAdapter:       deps.IdentityAdapter,
		passwordAdapter:       deps.PasswordAdapter,
	}
}

type SignupUsecaseInput struct {
	Name     string
	Surname  string
	Email    string
	Username string
	Password string
}

type SignupUsecaseOutput struct {
	Token string
	Exp   int64
	User  *domain.User
}

func (usecase *SignupUsecase) Signup(ctx context.Context, input *SignupUsecaseInput) (*SignupUsecaseOutput, *domain.AppError) {
	var insertedUser *domain.User
	var insertedCredential *domain.Credential
	var repoErr error

	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		roleId, err := uuid.Parse("cf49f2ea-8605-41a1-aa6b-3bc35129031e")
		if err != nil {
			return err
		}

		insertedUser, repoErr = usecase.usersRepository.CreateOne(ctx, &domain.User{
			Name:    input.Name,
			Surname: input.Surname,
			Role: &domain.Role{
				ID: roleId,
			},
		})

		if repoErr != nil {
			return repoErr
		}

		password, err := usecase.passwordAdapter.Hash(input.Password)
		if err != nil {
			return err
		}

		insertedCredential, repoErr = usecase.credentialsRepository.CreateOne(ctx, &domain.Credential{
			Username: input.Username,
			Password: string(password),
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

	user, err := usecase.usersRepository.FindOneByID(ctx, insertedUser.ID)
	if err != nil {
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
		User:  user,
	}, nil
}
