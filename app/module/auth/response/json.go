package response

type LoginResponse struct {
	Token     *string `json:"token"`
	Type      string  `json:"type"`
	ExpiresAt string  `json:"expires_at"`
	TimeNow   string  `json:"time_now"`
}

type RegisterResponse struct {
	EmployeeID   uint64 `json:"employee_id"`
	EmployeeCode string `json:"employee_code"`
}
