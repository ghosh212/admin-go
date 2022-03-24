package base

import (
	"admin-go/model"
	"context"
)

type Service interface {
	SearchAdmin(ctx context.Context, request model.SearchAdmin) (model.Administrator, error)
	CreateAdmin(ctx context.Context, request model.AdminRequest) (model.AdminCreation, error)
	UpdateAdmin(ctx context.Context, req model.UpdateAdmin) (model.AdminCreation, error)
	UpdateAdminSalary(ctx context.Context, req model.UpdateAdminRequest) (model.AdminPayUpdate, error)
}
