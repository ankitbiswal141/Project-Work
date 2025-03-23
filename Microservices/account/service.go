package account

import (
	"context"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount (ctx context.Context, name string) (*Account,error)
	GetAccount  (ctx context.Context, id string) (*Account,error)
	GetAccounts (ctx context.Context, skip ,take uint64) ([]Account,error)
}

type Account struct{
	ID   	string `json:"id"`
	Name string `json:"name"`
}

type accountService struct{
	repository Repository
}

func newService(r Repository) Service{
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context,name string) (*Account,error){
	acc:= &Account{
		Name: name,
		ID: ksuid.New().String(),
	}
	if err:= s.repository.PutAccount(ctx,*acc); err != nil {
		return nil,err
	}

	return acc,nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account,error) {
	return s.repository.GetAccountByID(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip, take uint64) ([]Account,error){
	if take >100 || (skip==0 && take==0){
		take=100
	}

	s.repository.ListAccounts(ctx	,take,skip)
}