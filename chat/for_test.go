package chat

import (
    "github.com/Cguilliman/gin-sandbox/common"
    "github.com/Cguilliman/gin-sandbox/users"
)

func createRoom() {
    db := common.GetDB()
    // tx1 := db.Begin()

    user1 := users.UserModel{
        Username: "AAAAAAAAAAAAAAAA",
        Email:    "aaaa@g.cn",
        Bio:      "hehddeda",
        Image:    nil,
    }
    db.Create(&user1)
    // tx1.Commit()
    // fmt.Println(user1)

    db.Create(&RoomUserModel{
        UserID: user1.ID,
    })
    var roomUser1 RoomUserModel
    db.Where(&RoomUserModel{
        UserID: user1.ID,
    }).First(&roomUser1)

    chat1 := RoomModel{
        Title: "testChat",
        Slug: "testChat",
        Users: []RoomUserModel{roomUser1},
    }
    db.Create(&chat1)
    // tx1.Commit()
    fmt.Println("++++++++++++++++", chat1.Users)
}
