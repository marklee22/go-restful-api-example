package model

// swagger:model Company
type Company struct {
	// The name for this company
	// required: true
	Name string `json:"name"`

	// The telephone number of the company
	// required: true
	Tel string `json:"tel"`

	// The email of the company
	// required: true
	Email string `json:"email"`
}
