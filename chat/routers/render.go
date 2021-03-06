package routers

import (
    "errors"
    "net/http"

    "gopkg.in/gin-gonic/gin.v1"
    
    "github.com/Cguilliman/gin-sandbox/common"
    "github.com/Cguilliman/gin-sandbox/chat"
    "github.com/Cguilliman/gin-sandbox/chat/serializers"
)

func RoomRender(c *gin.Context) {
    // createRoom()
    rooms, err := chat.AllRooms()
    if err != nil {
        c.JSON(
            http.StatusNotFound,
            common.NewError("rooms", errors.New("Invalid")),
        )
        return 
    }
    serializer := serializers.RoomsSerializer{c, rooms}
    c.HTML(
        http.StatusOK, "rooms.tmpl",
        gin.H{"rooms": serializer.Response()},
    )
}

func ChatRender(c *gin.Context) {
    id := c.Param("id")
    room, err := chat.GetRoom(id)
    if err != nil {
        c.JSON(
            http.StatusNotFound,
            common.NewError("chat", errors.New("Invalid")),
        )
        return 
    }
    serializer := serializers.RoomSerializer{c, room}
    c.HTML(
        http.StatusOK, "chat.tmpl",
        gin.H{"chat": serializer.Response()},
    )
}
