package main

import (
    "fmt"

    "github.com/jinzhu/gorm"

    "github.com/Cguilliman/gin-sandbox/users"
    "github.com/Cguilliman/gin-sandbox/common"
    "github.com/Cguilliman/gin-sandbox/articles"
    "github.com/Cguilliman/gin-sandbox/chat"
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
    Migrate(db) // make migrations
    return db
}

func testDbWorking(db *gorm.DB) {
    // test 1 to 1
    tx1 := db.Begin()
    userA := users.UserModel{
        Username: "AAAAAAAAAAAAAAAA",
        Email:    "aaaa@g.cn",
        Bio:      "hehddeda",
        Image:    nil,
    }
    tx1.Save(&userA)
    tx1.Commit()
    fmt.Println(userA)

    db.Save(&articles.ArticleUserModel{
        UserModelID: userA.ID,
    })
    var userAA articles.ArticleUserModel
    db.Where(&articles.ArticleUserModel{
        UserModelID: userA.ID,
    }).First(&userAA)
    fmt.Println(userAA)
}
