package command

import (
	"context"
	"strings"

	"github.com/lautarok/hexa/src/domain/errors"
	"github.com/lautarok/hexa/src/domain/models"
	"github.com/lautarok/hexa/src/domain/ports"
	"github.com/lautarok/hexa/src/pkg"
)

type SignupUsecase struct {
	credentialsRepository models.ICredentialsRepository
	usersRepository       models.IUsersRepository
	rolesRepository       models.IRolesRepository
	persistenceAdapter    ports.PersistencePort
	identityAdapter       ports.IdentityPort
	passwordAdapter       ports.PasswordPort
}

type SignupUsecaseDeps struct {
	CredentialsRepository models.ICredentialsRepository
	UsersRepository       models.IUsersRepository
	RolesRepository       models.IRolesRepository
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
	User  *models.User
}

func (usecase *SignupUsecase) Signup(ctx context.Context, input *SignupUsecaseInput) (*SignupUsecaseOutput, *models.AppError) {
	var insertedUser *models.User
	var repoErr error

	err := usecase.persistenceAdapter.Transaction(ctx, func(ctx context.Context) error {
		role, err := usecase.rolesRepository.GetOneBySlug(ctx, "common:user")
		if err != nil {
			return err
		}

		insertedUser, repoErr = usecase.usersRepository.CreateOne(ctx, &models.User{
			Name:    pkg.NormalizeName(input.Name),
			Surname: pkg.NormalizeName(input.Surname),
			Role: models.Role{
				ID: role.ID,
			},
		})

		if repoErr != nil {
			return repoErr
		}

		password, err := usecase.passwordAdapter.Hash(input.Password)
		if err != nil {
			return err
		}

		_, repoErr = usecase.credentialsRepository.CreateOne(ctx, &models.Credential{
			Username: input.Username,
			Password: string(password),
			Email:    input.Email,
			UserID:   insertedUser.ID,
		})

		if repoErr != nil {
			return repoErr
		}

		return nil
	})

	if err != nil {
		if usecase.persistenceAdapter.IsUniqueViolation(err) {
			errStr := err.Error()
			if strings.Contains(errStr, "username") {
				return nil, errors.NewAlreadyExistsError("Username already exists")
			}
			return nil, errors.NewAlreadyExistsError("Email already exists")
		}

		return nil, errors.NewInternalError(err)
	}

	user, err := usecase.usersRepository.FindOneByID(ctx, insertedUser.ID)
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	token, exp, err := usecase.identityAdapter.NewToken(&models.Identity{
		SubUserID: insertedUser.ID,
		UserID:    insertedUser.ID,
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
