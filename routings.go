package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    
    "github.com/Cguilliman/gin-sandbox/users"
    "github.com/Cguilliman/gin-sandbox/articles"
    // "github.com/Cguilliman/gin-sandbox/common"
    // "github.com/Cguilliman/gin-sandbox/chat"
)

func initRoutings(engine *gin.Engine) {
    v1 := engine.Group("/api")
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
}
