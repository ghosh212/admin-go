package model

type AdminCreationRequest struct {
	FirstName   string          `json:"firstName"`
	MiddleName  string          `json:"middleName"`
	Lastname    string          `json:"lastName"`
	PhoneNumber int64           `json:"phoneNumber"`
	Address     ResidentialInfo `json:"address"`
	Role        string          `json:"role"`
	DOJ         int64           `json:"doj"`
}

type SearchAdmin struct {
	FirstName  string `json:"firstName"`
	EmployeeID string `json:"employeeID"`
}

type AdminUpdateFields struct {
	Lastname    string          `json:"lastName"`
	PhoneNumber int64           `json:"phoneNumber"`
	Address     ResidentialInfo `json:"address"`
	Role        string          `json:"role"`
	DOL         int64           `json:"dol"`
	IsDelete    bool            `json:"isDelete"`
	IsMarried   MaritialStatus  `json:"isMarried"`
}

type AdminUpdateRequest struct {
	AdminRequest AdminUpdateFields `json:"adminRequest"`
	EmployeeId   string            `json:"employeeId"`
}

type UpdateAdminPayRequest struct {
	FirstName    string      `json:"firstName"`
	EmployeeId   string      `json:"employeeID"`
	PayStructure PayStruture `json:"payStructure"`
}
