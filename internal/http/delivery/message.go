package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/model"
	"github.com/gin-gonic/gin"
)

// SendMessage отправляет сообщение на сервер
// @Summary отправка сообщения
// @Description Этот эндпоит отправляет сообщение на сервер
// @Accept json
// @Produce json
// @Param message body model.Message true "Сообщение для отправки"
// @Success 200 {object} model.MessageResponse "Успешная отправка"
// @Failure 400 {object} model.ErrorResponse "Ошибка запроса"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /http/send [post]
func (h *Handler) SendMessage(c *gin.Context) {
	var message model.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("ошибка парсинга JSON").Error()})
		log.Println(err)
		return
	}

	segmentedMessage := h.UseCase.MessageSegmentation(message.StringMessage)

	fmt.Println("message", segmentedMessage)
	for idx, segment := range segmentedMessage {
		marshalledSegment, err := json.Marshal(model.Segment{
			ID:            message.Time,
			TotalSegments: uint(len(segmentedMessage)),
			SenderName:    message.SenderName,
			SegmentNumber: uint(idx + 1),
			Payload:       segment,
			HadError:      false,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		log.Println("Number: ", idx+1, "Payload: ", segment, "Message: ", message.StringMessage)

		resp, err := http.Post("http://localhost:8081/api/channel/code", "application/json", bytes.NewBuffer(marshalledSegment))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("ошибка при кодировании сообщения").Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Сообщение успешно отправлено"})
}

// TransferMessage передает полученные с канального уровня сегменты в кафку
// @Summary передача сегментов в кафку
// @Accept json
// @Produce json
// @param message body model.Segment true "Сегмент"
// @Success 200 {object} model.MessageResponse "Успешно получено и отправлено в кафку"
// @Failure 400 {object} model.ErrorResponse "Ошибка запроса"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /http/transfer [post]
func (h *Handler) TransferSegments(c *gin.Context) {
	var segment model.Segment

	if err := c.ShouldBindJSON(&segment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("ошибка при трансфере сегментов").Error()})
		log.Println(err)
		return
	}

	jsonSegment, err := json.Marshal(segment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("ошибка при парсинге в JSON").Error()})
		log.Println(err)
		return
	}

	if err = h.Producer.SendReport("forum-topic", string(jsonSegment)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("ошибка при отправке в kafka").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "сегмент успешно получен"})
}
