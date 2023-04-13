package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pigeon/internal/dao"
	"github.com/pigeon/internal/model"
	"time"
)

func Register(user *model.User) error {
	db := dao.GetDB()
	var count int64
	db.Model(&user).Where("username", user.Username).Count(&count)
	if count > 0 {
		return errors.New("username in used")
	}
	user.Uuid = uuid.New().String()
	user.CreateAt = time.Now()
	db.Create(&user)
	return nil
}

func GetUserDetail(uid string) model.User {
	var u *model.User
	db := dao.GetDB()
	db.Select("uuid", "username", "nickname", "avatar", "create_at").First(&u, "uuid = ?", uid)
	return *u
}

func FindUserByUserName(name string) model.User {
	var u *model.User
	db := dao.GetDB()
	db.Select("uuid", "username", "nickname", "avatar", "create_at").First(&u, "username = ?", name)
	return *u
}
