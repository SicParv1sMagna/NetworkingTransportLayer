package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendMessage(c *gin.Context) {
	var message model.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("ошибка парсинга JSON").Error())
		log.Println(err)
		return
	}

	segmentedMessage := h.UseCase.MessageSegmentation(message.StringMessage)

	for idx, segment := range segmentedMessage {
		marshalledSegment, err := json.Marshal(model.Segment{
			ID:            time.Now(),
			TotalSegments: uint(len(segmentedMessage)),
			SenderName:    message.SenderName,
			SegmentNumber: uint(idx + 1),
			Payload:       segment,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.New("ошибка при кодировании сообщений").Error())
			return
		}

		log.Println("Number: ", idx+1, "Payload: ", segment, "Message: ", message.StringMessage)

		resp, err := http.Post("http://localhost:8081/channel/code", "application/json", bytes.NewBuffer(marshalledSegment))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, errors.New("ошибка при кодировании сообщения").Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Сообщение успешно отправлено"})
}

func (h *Handler) TransferSegments(c *gin.Context) {
	var segment model.Segment

	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("ошибка при трансфере сегментов").Error())
		log.Println(err)
		return
	}
}
