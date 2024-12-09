/*
@Time : 2024/12/9 15:23
@Author : linx
@File : authInterface.go
@dsc: 验证功能的接口
*/

package auth

type Auth interface {
	UserRegister() (string, error)
	UserLogin() (string, error)
}
