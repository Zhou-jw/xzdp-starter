package db

import (
	"time"

	"github.com/Zhou-jw/xzdp-starter/src/pkg/constants"
)

type User struct {
	ID           int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"` // 对应表的id（主键自增）
	Phone        string    `json:"phone" gorm:"column:phone;unique"`             // 对应表的phone（唯一索引）
	Password     string    `json:"password" gorm:"column:password"`              // 对应表的password
	NickName     string    `json:"nick_name" gorm:"column:nick_name"`            // 对应表的nick_name
	Avatar       string    `json:"avatar" gorm:"column:icon"`                    // 结构体avatar ↔ 表icon（字段名不同用tag映射）
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"` // 自动填充创建时间
	UpdateTime   time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"` // 自动填充更新时间
}

func (User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user info
func CreateUser(user *User) (int64, error) {
	err := DB.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

// return user id, whether created new user, error
func CreateUserIfNotExist(phone string, hashedPwd string) (int64, bool, error) {
	var user User

	err := DB.Where("phone = ?", phone).First(&user).Error
	if err == nil {
		// 用户已存在，返回现有ID
		return user.ID, false, nil
	}

	newUser := &User{
		Phone:    phone,
		Password: hashedPwd, // 注意：必须传入加密后的密码，不要传明文
	}

	if err := DB.Create(newUser).Error; err != nil {
		return 0, false, err
	}

	return newUser.ID, true, nil
}