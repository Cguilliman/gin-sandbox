package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Cguilliman/gin-sandbox/articles"
	"github.com/Cguilliman/gin-sandbox/chat"
	"github.com/Cguilliman/gin-sandbox/common"
	"github.com/Cguilliman/gin-sandbox/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()

	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})

	db.AutoMigrate(&chat.RoomModel{})
	db.AutoMigrate(&chat.RoomUserModel{})
	db.AutoMigrate(&chat.MessageModel{})
}

func initDatabase() *gorm.DB {
	db := common.Init() // initialize database
	Migrate(db)         // make migrations
	return db
}

func testDbWorking(db *gorm.DB) {
	// test 1 to 1
	tx1 := db.Begin()

	user1 := users.UserModel{
		Username: "AAAAAAAAAAAAAAAA",
		Email:    "aaaa@g.cn",
		Bio:      "hehddeda",
		Image:    nil,
	}
	tx1.Save(&user1)
	tx1.Commit()
	// fmt.Println(user1)

	db.Save(&chat.RoomUserModel{
		UserID: user1.ID,
	})
	var roomUser1 chat.RoomUserModel
	db.Where(&chat.RoomUserModel{
		UserID: user1.ID,
	}).First(&roomUser1)

	chat1 := chat.RoomModel{
		Title: "testChat",
		Slug: "testChat",
		Users: []chat.RoomUserModel{roomUser1},
	}
	tx1.Save(&chat1)
	tx1.Commit()
	fmt.Println("++++++++++++++++", chat1.Users)
}
