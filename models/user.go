package models



type UserReq struct {
	Name string `json:"name"`
	Age int32 	`json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string	`json:"gender"`
}

type UserRes struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int32 	`json:"age"`
	Email string `json:"email"`
	Token string `json:"token"`
	Gender string	`json:"gender"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type UserLogin struct{
	EmailOrUsername string `json:"email_username"`
	Password string `json:"password"`
}

type LoginRes struct{
	ID string `json:"id"`
	Description string `json:"username"`
}
