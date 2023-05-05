package service

import (
	"errors"
	"github.com/google/uuid"
	"pigeon/internal/dao"
	"pigeon/internal/model"
	"time"
)

type userService struct {
}

var UserService = new(userService)

func (u *userService) Register(user *model.User) error {
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

func (u *userService) GetUserList(uuid string) []model.User {
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

func (u *userService) ListAllUser() []model.User {
	db := dao.GetDB()
	var queryUsers []model.User
	db.Raw("select id,uuid,username,nickname,avatar,email,create_at,update_at from users").Scan(&queryUsers)
	return queryUsers
}

func (u *userService) GetUserByUserName(queryUser *model.User) *model.User {
	db := dao.GetDB()
	var resultUser *model.User
	db.First(&resultUser, "username = ?", queryUser.Username)
	return resultUser
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
