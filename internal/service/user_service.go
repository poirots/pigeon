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

func GetUserList(uuid string) []model.User {
	db := dao.GetDB()
	var queryUser *model.User
	db.First(&queryUser, "uuid = ?", uuid)
	var nullId int32 = 0
	if nullId == queryUser.Id {
		return nil
	}

	var queryUsers []model.User
	db.Raw("SELECT u.username, u.uuid, u.avatar FROM user_friends AS uf JOIN users AS u ON uf.friend_id = u.id WHERE uf.user_id = ?", queryUser.Id).Scan(&queryUsers)

	return queryUsers
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
