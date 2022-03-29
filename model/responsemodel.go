package model

type AdminCreateUpdateResponse struct {
	EmployeeID string `json:"employeeID"`
	FirstName  string `json:"firstName"`
	Lastname   string `json:"lastName"`
	Message    string `json:"message"`
}

type AdministratorSearchResponse struct {
	EmployeeID  string          `json:"employeeID"`
	FirstName   string          `json:"firstName"`
	MiddleName  string          `json:"middleName"`
	Lastname    string          `json:"lastName"`
	PhoneNumber int64           `json:"phoneNumber"`
	Address     ResidentialInfo `json:"address"`
	Role        string          `json:"role"`
	DOJ         int64           `json:"doj"`
	DOL         int64           `json:"dol,omitempty"`
	Status      string          `json:"status"`
}
