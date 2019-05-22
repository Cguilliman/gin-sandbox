package routers

import "gopkg.in/gin-gonic/gin.v1"

func ChatRenderRegister(router *gin.RouterGroup) {
    router.GET("rooms", RoomRender)
    router.GET("room/:id", ChatRender)
}
