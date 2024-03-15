package user

import "time"

// 创建用户请求结构体
type CreateUserRequest struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Password2 string    `json:"password2"`
	Gender    string    `json:"gender"`
	Birth     time.Time `json:"birth"`
	Url       string    `json:"url"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
}

// 创建用户响应
type CreateUserResponse struct {
	Username string `json:"username"`
}

// cookie 登录请求
type CookieLoginRequest struct {
	Token string `json:"Token"`
}

// 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	UserId   uint   `json:"userId"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"isadmin"`
}

type GetInfoResponse struct {
	Username string    `josn:"username"`
	Gender   string    `josn:"gender"`
	Birth    time.Time `josn:"birth"`
	Email    string    `josn:"email"`
	Address  string    `josn:"address"`
}

type UpdateUserRequest struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Password2 string    `json:"password2"`
	Gender    string    `json:"gender"`
	Birth     time.Time `json:"birth"`
	Url       string    `json:"url"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
}
