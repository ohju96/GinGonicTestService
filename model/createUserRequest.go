package model

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Code int    `json:"code"`
}
