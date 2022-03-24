package model

import "time"

type Administrator struct {
	EmployeeID   string
	FirstName    string
	MiddleName   string
	Lastname     string
	PhoneNumber  int64
	Address      ResidentialInfo
	Role         string
	DOJ          time.Time
	DOL          time.Time
	IsDelete     bool
	IsMarried    MaritialStatus
	PayStrucrure AdministratorPayStructure
}

type AdministratorPayStructure struct {
	EmployeeID   string
	FirstName    string
	MiddleName   string
	Lastname     string
	PayStrucrure PayStrucrure
}

type ResidentialInfo struct {
	HouseNumber int64
	Street      string
	Lane        string
	LandMark    string
	Pincode     int64
}

type PayStrucrure struct {
	Salary             Salary
	YearsOfService     float64
	LastHikePercentage float64
}

type Salary struct {
	TotalPay       float64
	GrossPay       float64
	PFContribution float64
	Tax            Tax
}

type Tax struct {
	TotalTaxableIncome float64
	TaxBracket         float64
	TotalTax           float64
	TaxPaid            float64
	TaxRemaining       float64
}

type MaritialStatus struct {
	IsMarried        bool
	SpouseFirstName  string
	SpouseMiddleName string
	SpouseLastName   string
	Address          ResidentialInfo
}
