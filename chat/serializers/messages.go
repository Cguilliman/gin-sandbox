package serializers

import (
    "gipkg.in/gin-gonic/gin.v1"
    "github.com/Cguilliman/gin-sandbox/chat"
)


type MessageSerializer struct {
    C *gin.Context
    chat.MessageModel
}

type MessagesSerializer struct {
    C *gin.Context
    Messages []chat.MessagesModel
}

type MessageResponse struct {
    ID        uint   `json:"-"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
    Message   string `json:"message"`
    // TODO: add more fields
}

func (s *MessageSerializer) Response() MessageResponse {
    response := MessageResponse{
        ID:        s.ID,
        CreatedAt: s.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
        UpdatedAt: s.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
        Message:   s.Message
    }
    // TODO: add fields getters
    return response
}

func (s *MessagesSerializer) Response() []MessageResponse {
    response := []MessageResponse{}
    for _, message := range s.Messages {
        serializer := MessageSerializer{s.C, message}
        response = append(response, serializer.Response())
    }
    return response
}
