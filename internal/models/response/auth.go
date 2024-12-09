/*
@Time : 2024/12/9 16:09
@Author : linx
@File : auth.go
@dsc:
*/

package response

type RegisterResp struct {
	Token string `json:"token"`
}

type LoginResp struct {
	Token string `json:"token"`
}
