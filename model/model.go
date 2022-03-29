package model

type Administrator struct {
	EmployeeID   string
	FirstName    string
	MiddleName   string
	Lastname     string
	PhoneNumber  int64
	Address      ResidentialInfo
	Role         string
	DOJ          int64
	DOL          int64
	IsDelete     bool
	IsMarried    MaritialStatus
	PayStrucrure PayStruture
}

type ResidentialInfo struct {
	HouseNumber int64  `json:"houseNumber"`
	Street      string `json:"street"`
	Lane        string `json:"lane"`
	LandMark    string `json:"landmark"`
	Pincode     int64  `json:"pincode"`
}

type PayStruture struct {
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
	IsMarried        bool            `json:"isMarried"`
	SpouseFirstName  string          `json:"spouceFirstName"`
	SpouseMiddleName string          `json:"spouceMiddleName"`
	SpouseLastName   string          `json:"spouceLastsName"`
	Address          ResidentialInfo `json:"address"`
}
