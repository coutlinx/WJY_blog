/*
@Time : 2024/12/9 16:08
@Author : linx
@File : auth.go
@dsc:
*/

package request

type RegisterReq struct {
	Email    string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
