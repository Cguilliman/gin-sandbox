package chat

import (
	// "gopkg.in/gin-gonic/gin.v1"
	"github.com/Cguilliman/gin-sandbox/common"
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
	User   users.UserModel
	UserID uint
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

func AllRooms() ([]RoomModel, error) {
	db := common.GetDB()
	var models []RoomModel
	err := db.Find(&models).Error
	return models, err
}

func GetRoom(id string) (RoomModel, error) {
	db := common.GetDB()
	var model RoomModel
	err := db.Where("id = ?", id).First(&model).Error
	return model, err
}
