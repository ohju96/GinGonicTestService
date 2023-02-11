package dto

type CreateUserReq struct {
	Uid      string `json:"uid"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRes struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}
