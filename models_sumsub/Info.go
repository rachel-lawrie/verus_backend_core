package models_sumsub

type FixedInfo struct {
	FirstName  string       `json:"firstName,omitempty"`
	MiddleName string       `json:"middleName,omitempty"`
	LastName   string       `json:"lastName,omitempty"`
	Dob        string       `json:"dob,omitempty"` //yyyy-mm-dd format
	Gender     string       `json:"gender,omitempty"`
	Country    string       `json:"country,omitempty"`
	Addresses  []RawAddress `json:"addresses,omitempty"`
}

// Address represents the address fields
type RawAddress struct {
	Line1      string `json:"street,omitempty"`
	Line2      string `json:"flatNumber,omitempty"`
	City       string `json:"town"`
	Region     string `json:"state"`
	PostalCode string `json:"postCode"`
	Country    string `json:"country"`
}
