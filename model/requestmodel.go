package model

type AdminRequest struct {
	FirstName   string
	MiddleName  string
	Lastname    string
	PhoneNumber int64
	Address     ResidentialInfo
	Role        string
	DOJ         Time
	DOL         Time
	IsDelete    bool
}

type SearchAdminRequest struct {
	FirstName      string `json:"firstname"`
	EmployeeNumber string `json:"employeeNumber"`
}
