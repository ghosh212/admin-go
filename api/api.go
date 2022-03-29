package api

import (
	"admin-go/model"
	"math/rand"
	"strconv"
	"time"
)

//implementation of Enum kindy
const (
	CREATED_SUCCESSFULLY     string = "CREATED SUCCESSfULLY"
	NOT_CREATED_SUCCESSFULLY        = "ADMIN ACCOUNT NOT CREATED"
	ACTIVE_STATUS                   = "ADMIN ACCOUNT IS ACTIVE"
	UNACTIVE_STATUS                 = "ADMIN STATUS NOT ACTIVE"
)

func CreateSearchAdminResponse(resp model.Administrator) model.AdministratorSearchResponse {
	retval := model.AdministratorSearchResponse{}
	retval.EmployeeID = resp.EmployeeID
	retval.FirstName = resp.FirstName
	retval.MiddleName = resp.MiddleName
	retval.Lastname = resp.Lastname
	retval.PhoneNumber = resp.PhoneNumber
	retval.Address = resp.Address
	retval.Role = resp.Role
	retval.DOJ = resp.DOJ
	if resp.DOL != 0 {
		retval.DOJ = resp.DOL
		retval.Status = ACTIVE_STATUS
	} else {
		retval.Status = UNACTIVE_STATUS
	}
	return retval
}

func CreateUpdateAdminResponse(resp model.Administrator) model.AdminCreateUpdateResponse {
	retVal := model.AdminCreateUpdateResponse{}
	if (model.Administrator{} != resp) {
		retVal.EmployeeID = resp.EmployeeID
		retVal.FirstName = resp.FirstName
		retVal.Lastname = resp.Lastname
		retVal.Message = CREATED_SUCCESSFULLY
	} else {
		retVal.EmployeeID = resp.EmployeeID
		retVal.FirstName = resp.FirstName
		retVal.Lastname = resp.Lastname
		retVal.Message = NOT_CREATED_SUCCESSFULLY
	}
	return retVal
}

func CreateDataForAdminUpdate(request model.AdminUpdateRequest) model.Administrator {
	retval := model.Administrator{}
	//to-do
	return retval
}

func CreateNewPayData(request model.UpdateAdminPayRequest) model.Administrator {
	retval := model.Administrator{
		PayStrucrure: model.PayStruture{
			Salary: model.Salary{
				TotalPay:       request.PayStructure.Salary.TotalPay,
				GrossPay:       request.PayStructure.Salary.GrossPay,
				PFContribution: request.PayStructure.Salary.PFContribution,
				Tax: model.Tax{
					TotalTaxableIncome: request.PayStructure.Salary.Tax.TotalTaxableIncome,
					TaxBracket:         request.PayStructure.Salary.Tax.TaxBracket,
				},
			},
		},
	}
	return retval
}

func GenerateEmployeeID() string {
	rand.Seed(time.Now().Unix())
	min := 0
	max := 100
	employeeID := rand.Intn(max-min+1) + min
	return strconv.Itoa(employeeID) //convert employeeID from int to string
}
