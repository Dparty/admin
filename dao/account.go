package dao

import (
	gDao "github.com/Dparty/dao"
	"gorm.io/gorm"
)

var adminAccountRespository *AdminAccountRespository

func GetAdminAccountRespository() *AdminAccountRespository {
	if adminAccountRespository == nil {
		adminAccountRespository = NewAccountAdminRespository()
	}
	return adminAccountRespository
}

type AdminAccount struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
	Salt     []byte `json:"salt"`
	Role     string `json:"role"`
}

type AdminAccountRespository struct {
	db *gorm.DB
}

func NewAccountAdminRespository() *AdminAccountRespository {
	return &AdminAccountRespository{db: gDao.GetDBInstance()}
}
