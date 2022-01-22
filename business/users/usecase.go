package users

import (
	"CalFit/app/middlewares"
	"CalFit/business/paginations"
	"CalFit/exceptions"
	"CalFit/helpers"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

type UsersUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JWTAuth        *middlewares.ConfigJWT
}

func NewUsersUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UsersUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		JWTAuth:        jwtAuth,
	}
}

func (uu *UsersUsecase) LoginOAuth(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := uu.Repo.LoginOAuth(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	res.Token, _ = uu.JWTAuth.GenerateToken(res.Id, res.Email, true, false, false)
	return res, nil
}

func (uu *UsersUsecase) Register(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	var err error
	users.Password, err = helpers.Hash(users.Password)
	if err != nil {
		return Domain{}, err
	}
	res, err := uu.Repo.Register(ctx, users)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (uu *UsersUsecase) Login(ctx context.Context, users Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()
	if users.Email == "" || users.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := uu.Repo.GetByUsername(ctx, users.Email)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(users.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	res.Token, _ = uu.JWTAuth.GenerateToken(res.Id, res.Email, true, false, false)
	return res, nil
}

func (uu *UsersUsecase) GetAll(ctx context.Context, pagination paginations.Domain) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	return uu.Repo.GetAll(ctx, pagination)
}

func (uu *UsersUsecase) CountAll(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	return uu.Repo.CountAll(ctx)
}

func (uu *UsersUsecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	return uu.Repo.GetById(ctx, id)
}

func (uu *UsersUsecase) Update(ctx context.Context, id string, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uu.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return uu.Repo.Update(ctx, id, domain)
}
