package model

type AdminCreation struct {
	EmployeeID  string
	FirstName   string
	Lastname    string
	PhoneNumber int64
	Message     string
	Error       error
}

type AdminPayUpdate struct {
	EmployeeID string
	FirstName  string
	Lastname   string
	Message    string
	Error      error
}
