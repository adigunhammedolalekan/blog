package service

import (
	"context"
	account "github.com/adigunhammedolalekan/blog/account-service/proto/account"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"net/http"
)

type ApiService struct {

	accountClient account.AccountServiceClient
}

func NewApiService(client client.Client) *ApiService {

	return &ApiService{
		accountClient: account.NewAccountServiceClient("service.account", client),
	}
}

func (api *ApiService) NewAccount(req *restful.Request, res *restful.Response) {

	accountPayload := &account.Account{}
	err := req.ReadEntity(accountPayload)
	if err != nil {
		respondWithError(res, err)
		return
	}


	response, err := api.accountClient.CreateAccount(context.Background(), accountPayload)
	if err != nil {
		respondWithError(res, err)
		return
	}

	res.WriteEntity(response)
}

func (api *ApiService) GetAccount(req *restful.Request, res *restful.Response) {

	response, err := api.accountClient.GetAccount(context.Background(), &account.GetAccountRequest{
		UserId: req.PathParameter("id"),
	})
	if err != nil {
		respondWithError(res, err)
		return
	}

	res.WriteEntity(response)
}

func (api *ApiService) Hello(req *restful.Request, res *restful.Response)  {
	res.WriteEntity(map[string]interface{} {"hello" : "World"})
}

func respondWithError(res *restful.Response, err error) {

	e := errors.Parse(err.Error())
	if err != nil {
		res.WriteError(int(e.Code), err)
		return
	}

	res.WriteError(http.StatusInternalServerError, err)
}