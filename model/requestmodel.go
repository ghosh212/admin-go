package model

import "time"

type AdminRequest struct {
	FirstName   string
	MiddleName  string
	Lastname    string
	PhoneNumber int64
	Address     ResidentialInfo
	Role        string
	DOJ         time.Time
	DOL         time.Time
}

type SearchAdmin struct {
	FirstName  string
	EmployeeID string
}

type UpdateAdmin struct {
	AdminRequest AdminRequest
	EmployeeId   string
}

type UpdateAdminRequest struct {
	EmployeeId   string
	PayStructure AdministratorPayStructure
}
