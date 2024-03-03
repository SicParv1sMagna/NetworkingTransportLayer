package delivery

import (
	"errors"
	"fmt"
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

	var segments []*model.Segment

	for idx, segment := range segmentedMessage {
		segments = append(segments, &model.Segment{
			ID:            time.Now(),
			TotalSegments: uint(len(segmentedMessage)),
			SenderName:    message.SenderName,
			SegmentNumber: uint(idx + 1),
			Payload:       segment,
		})
	}
}

func (h *Handler) TransferSegments(c *gin.Context) {
	var segment model.Segment

	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, errors.New("ошибка при трансфере сегментов").Error())
		log.Println(err)
		return
	}

	fmt.Println(segment)
}
