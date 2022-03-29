package base

import (
	"admin-go/api"
	db "admin-go/db"
	"admin-go/model"
	"context"
	"log"
	"time"
)

//rpc methods
//endpoint methods exposed to user
type Service interface {
	SearchAdmin(ctx context.Context, request model.SearchAdmin) (model.AdministratorSearchResponse, error)
	CreateAdmin(ctx context.Context, request model.AdminCreationRequest) (model.AdminCreateUpdateResponse, error)
	UpdateAdmin(ctx context.Context, req model.AdminUpdateRequest) (model.AdminCreateUpdateResponse, error)
	UpdateAdminSalary(ctx context.Context, req model.UpdateAdminPayRequest) (model.AdminCreateUpdateResponse, error)
}

type mongodb struct {
	repository db.Repository
	log        log.Logger
}

//function to create the service
func NewService(rep db.Repository, logger log.Logger) Service {
	return &mongodb{
		repository: rep,
		log:        logger,
	}
}

func (s mongodb) SearchAdmin(ctx context.Context, request model.SearchAdmin) (model.AdministratorSearchResponse, error) {
	//log := log.Print("")
	search := model.SearchAdmin{
		FirstName:  request.FirstName,
		EmployeeID: request.EmployeeID,
	}

	resp, err := s.repository.SearchAdmin(ctx, search)
	if err != nil {
		return model.AdministratorSearchResponse{}, ErrNotFound
	}
	return api.CreateSearchAdminResponse(resp), nil

}

func (s mongodb) CreateAdmin(ctx context.Context, request model.AdminCreationRequest) (model.AdminCreateUpdateResponse, error) {
	//log

	createAdmin := model.Administrator{
		EmployeeID:  api.GenerateEmployeeID(),
		FirstName:   request.FirstName,
		MiddleName:  request.MiddleName,
		Lastname:    request.Lastname,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		Role:        request.Role,
		DOJ:         time.Now().Unix(),
	}
	resp, err := s.repository.CreateAdmin(ctx, createAdmin)
	if err != nil {
		return model.AdminCreateUpdateResponse{}, ErrInsertFail
	}

	return api.CreateUpdateAdminResponse(resp), nil

}
func (s mongodb) UpdateAdmin(ctx context.Context, req model.AdminUpdateRequest) (model.AdminCreateUpdateResponse, error) {
	//log
	updateAdmin := api.CreateDataForAdminUpdate(req)
	resp, err := s.repository.CreateAdmin(ctx, updateAdmin)
	if err != nil {
		return model.AdminCreateUpdateResponse{}, ErrOccuredDuringUpdation
	}
	return api.CreateUpdateAdminResponse(resp), nil
}

func (s mongodb) UpdateAdminSalary(ctx context.Context, req model.UpdateAdminPayRequest) (model.AdminCreateUpdateResponse, error) {
	//log
	//search for a employee
	searchAdmin := model.SearchAdmin{
		EmployeeID: req.EmployeeId,
		FirstName:  req.FirstName,
	}
	var resp model.Administrator
	var er error
	_, err := s.SearchAdmin(ctx, searchAdmin)
	if err != nil {
		return model.AdminCreateUpdateResponse{}, ErrNotFound
	} else {
		createNewPayData := api.CreateNewPayData(req)
		resp, er = s.repository.UpdateAdminSalary(ctx, createNewPayData)
		if er != nil {
			return model.AdminCreateUpdateResponse{}, ErrOccuredDuringUpdation
		}
	}
	return api.CreateUpdateAdminResponse(resp), nil

}

//to-do
//calculateTax and insert into db func
//promotion opf admin
