/*
@Time : 2024/12/9 15:26
@Author : linx
@File : modelInterface.go
@dsc:
*/

package auth

type User interface {
	EncryptPassword() error
	CreateUser() (any, error)
	GetUser() (any, error)
	GetPassword() string
}
