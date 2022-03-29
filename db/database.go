package db

import (
	"admin-go/model"
	"context"
)

type Repository interface {
	SearchAdmin(ctx context.Context, request model.SearchAdmin) (model.Administrator, error)
	CreateAdmin(ctx context.Context, request model.Administrator) (model.Administrator, error)
	UpdateAdmin(ctx context.Context, req model.Administrator) (model.Administrator, error)
	UpdateAdminSalary(ctx context.Context, req model.Administrator) (model.Administrator, error)
}
