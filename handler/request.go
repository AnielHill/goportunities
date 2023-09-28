package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Create Opening

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	// <= (Igual ou menor que:)
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("Company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("Location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("Link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}
