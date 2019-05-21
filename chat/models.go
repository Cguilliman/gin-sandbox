package chat

import (
	// "github.com/Cguilliman/gin-sandbox/common"
	"github.com/Cguilliman/gin-sandbox/users"
	"github.com/jinzhu/gorm"
)

type RoomModel struct {
	gorm.Model
	Slug     string `gorm:"unique_index"`
	Title    string
	Messages []MessageModel  `gorm:"ForeignKey:RoomID"`
	Users    []RoomUserModel `gorm:"many2many:room_users;"`
}

type RoomUserModel struct {
	gorm.Model
	UserID uint
	User   users.UserModel
	Rooms  []RoomModel `gorm:"many2many:room_users;"`
}

type MessageModel struct {
	gorm.Model
	Message    string `gorm:"size:2048"`
	FromUser   users.UserModel
	FromUserID uint
	ToUser     users.UserModel
	ToUserID   uint
	Room       RoomModel
	RoomID     uint
}
