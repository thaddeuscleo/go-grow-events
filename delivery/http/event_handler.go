package http

import (
	"bytes"
	"encoding/base64"
	"go-grow-events/model"
	"go-grow-events/usecase"
	"image/png"
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventUsecase usecase.EventUsecase
}

func NewEventHandler (eventUsecase usecase.EventUsecase) *EventHandler {
	return &EventHandler{eventUsecase}
}

func (h *EventHandler) RegisterParticipant(ctx *gin.Context) {
	var participantRequest model.RegisterParticipantRequest

	err := ctx.ShouldBindJSON(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42201",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	participant, err := h.eventUsecase.PostRegisterSession(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40001",
			"responseMessage": "Usecase PostRegisterUser is not working properly: " + err.Error(),
		})
		return
	}

	encodedToQR := "user id: " + strconv.Itoa(participant.ID)
	imagedQrCode, err := qrcode.New(encodedToQR, qrcode.Medium)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"responseCode": "50001",
			"responseMessage": "QR Code does not work properly: " + err.Error(),
		})
		return
	}

	buffer := new(bytes.Buffer)
	png.Encode(buffer, imagedQrCode.Image(256))
	base64QRCode := base64.StdEncoding.EncodeToString(buffer.Bytes())

	if participant.SessionID == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"responseCode": "20000",
			"responseMessage": "Participant has been registered successfully",
			"name": participant.Name,
			"email": participant.Email,
			"phoneNo": participant.PhoneNo,
			"sessionID": participant.SessionID,
			"sessionName": "GROW Center Anniversary 1st Service",
			"scanStatusID": participant.IsScanned,
			"scanStatus": "Not Scanned",
			"requestedSeat": participant.RequestedSeat,
			"qrCode": base64QRCode,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Participant has been registered successfully",
		"name": participant.Name,
		"email": participant.Email,
		"phoneNo": participant.PhoneNo,
		"sessionID": participant.SessionID,
		"sessionName": "GROW Center Anniversary 2nd Service",
		"scanStatusID": participant.IsScanned,
		"scanStatus": "Not Scanned",
		"requestedSeat": participant.RequestedSeat,
		"reasons": participant.Reasons,
		"qrCode": base64QRCode,
	})
}