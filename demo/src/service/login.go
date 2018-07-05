package service

import (
	"context"
	"model"
	"repository"
)

type LoginService interface {
	Login(ctx context.Context, name string) (*model.Account, error)
}

type loginFlow struct {
	accountRepo repository.AccountRepository
}

func NewLoginService(ar repository.AccountRepository) LoginService {
	return &loginFlow{ar}
}

func (l *loginFlow) Login(ctx context.Context, name string) (*model.Account, error) {
	res, err := l.accountRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}
