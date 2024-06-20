package request

type LoginRequest struct {
	EmployeeCode string `json:"employee_code" example:"2210001" validate:"required"`
	Password     string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
}

type RegisterRequest struct {
	EmployeeCode string `json:"employee_code" example:"22100001" validate:"required"`
	Password     string `json:"password" example:"12345678" validate:"required,min=8,max=255"`
}
