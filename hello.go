package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/gin-gonic/gin.v1"

	"github.com/Cguilliman/gin-sandbox/articles"
	"github.com/Cguilliman/gin-sandbox/chat"
	"github.com/Cguilliman/gin-sandbox/common"
	"github.com/Cguilliman/gin-sandbox/users"
)

func Migrate(db *gorm.DB) { // TODO: remove migration to another place
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

func main() {
	db := common.Init() // initialize database
	Migrate(db)
	defer db.Close()

	engine := gin.Default()

	v1 := engine.Group("/api") // TODO: remove routing registration to another place
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))

	articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := engine.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// testDbWorking(db)

	engine.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
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
