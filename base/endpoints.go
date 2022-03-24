package base

import (
	"admin-go/model"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

var (
	ErrBadRequest = errors.New("BAD REQUEST")
	ErrInsertFail = errors.New("INSERT INTO SB NOT SUCCESSFULL")
	ErrNotFound   = errors.New("NOT FOUND")
)

type Endpoint struct {
	SearchAdmin       endpoint.Endpoint
	CreateAdmin       endpoint.Endpoint
	UpdateAdmin       endpoint.Endpoint
	UpdateAdminSalary endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoint {
	return Endpoint{
		SearchAdmin:       makeSearchAdminEndpoint(s),
		CreateAdmin:       makeCreateAdminEndpoint(s),
		UpdateAdmin:       makeUpdateAdmin(s),
		UpdateAdminSalary: makeUpdateAdminSalaryEndpoint(s),
	}
}

func makeSearchAdminEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.SearchAdmin)
		if !ok {
			return err, ErrBadRequest
		}
		return s.SearchAdmin(ctx, req)
	}
}

func makeCreateAdminEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.AdminRequest)
		if !ok {
			return err, ErrBadRequest
		}
		return s.CreateAdmin(ctx, req)
	}
}

func makeUpdateAdmin(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.UpdateAdmin)
		if !ok {
			return err, ErrBadRequest
		}
		return s.UpdateAdmin(ctx, req)
	}
}

func makeUpdateAdminSalaryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(model.UpdateAdminRequest)
		if !ok {
			return err, ErrBadRequest
		}
		return s.UpdateAdminSalary(ctx, req)
	}
}
