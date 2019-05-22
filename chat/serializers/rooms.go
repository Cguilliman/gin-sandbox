package serializers

import (
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/Cguilliman/gin-sandbox/chat"
)

type RoomSerializer struct {
    C *gin.Context
    chat.RoomModel
}

type RoomsSerializer struct {
    C     *gin.Context
    Rooms []chat.RoomModel
}

type RoomResponse struct {
    ID        uint   `json:"-"`
    // Slug      string          `json:"slug"`
    Title     string `json:"title"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
    // TODO: mb add more fields
}

func (s *RoomSerializer) Response() RoomResponse {
    response := RoomResponse{
        ID: s.ID,
        // Slug: s.Slug,
        Title:     s.Title,
        CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
        UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
    }
    return response
}

func (s *RoomsSerializer) Response() []RoomResponse {
    response := []RoomResponse{}

    for _, room := range s.Rooms {
        serializer := RoomSerializer{s.C, room}
        response = append(response, serializer.Response())
    }

    return response
}
