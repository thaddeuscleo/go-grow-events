package http

import (
	"go-grow-events/model"
	"go-grow-events/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventUsecase usecase.EventUsecase
}

func NewEventHandler(eventUsecase usecase.EventUsecase) *EventHandler {
	return &EventHandler{eventUsecase}
}

func (h *EventHandler) RegisterParticipant(ctx *gin.Context) {
	var participantRequest model.RegisterParticipantRequest

	err := ctx.ShouldBindJSON(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode":    "42201",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	participant, err := h.eventUsecase.PostRegisterSession(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "40001",
			"responseMessage": "Usecase PostRegisterSession is not working properly: " + err.Error(),
		})
		return
	}

	if participant.SessionID == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"responseCode":     "20000",
			"responseMessage":  "Participant has been registered successfully",
			"name":             participant.Name,
			"email":            participant.Email,
			"phoneNo":          participant.PhoneNo,
			"sessionID":        participant.SessionID,
			"sessionName":      "GROW Center Anniversary 1st Service",
			"scanStatus":       "Not Scanned",
			"requestedSeat":    participant.RequestedSeat,
			"registrationCode": participant.RegistrationCode,
			"qrCode":           participant.QRCode,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode":     "20000",
		"responseMessage":  "Participant has been registered successfully",
		"name":             participant.Name,
		"email":            participant.Email,
		"phoneNo":          participant.PhoneNo,
		"sessionID":        participant.SessionID,
		"sessionName":      "GROW Center Anniversary 2nd Service",
		"scanStatus":       "Not Scanned",
		"requestedSeat":    participant.RequestedSeat,
		"reasons":          participant.Reasons,
		"registrationCode": participant.RegistrationCode,
		"qrCode":           participant.QRCode,
	})
}

func (h *EventHandler) VerifyParticipant(ctx *gin.Context) {
	var participantRequest model.VerifyParticipantRequest

	err := ctx.ShouldBindJSON(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode":    "42202",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	participant, err := h.eventUsecase.PostVerifySession(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "40002",
			"responseMessage": "Usecase PostRegisterUser is not working properly: " + err.Error(),
		})
		return
	}

	if participant.SessionID == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"responseCode":     "20000",
			"responseMessage":  "Participant has been verified to enter the service",
			"name":             participant.Name,
			"email":            participant.Email,
			"phoneNo":          participant.PhoneNo,
			"sessionID":        participant.SessionID,
			"sessionName":      "GROW Center Anniversary 1st Service",
			"totalScanned":     participant.IsScanned,
			"requestedSeat":    participant.RequestedSeat,
			"registrationCode": participant.RegistrationCode,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode":     "20000",
		"responseMessage":  "Participant has been registered successfully",
		"name":             participant.Name,
		"email":            participant.Email,
		"phoneNo":          participant.PhoneNo,
		"sessionID":        participant.SessionID,
		"sessionName":      "GROW Center Anniversary 2nd Service",
		"totalScanned":     participant.IsScanned,
		"requestedSeat":    participant.RequestedSeat,
		"registrationCode": participant.RegistrationCode,
	})
}

func (h *EventHandler) ViewBooking(ctx *gin.Context) {
	var participantRequest *model.ViewBookingRequest

	err := ctx.ShouldBindJSON(&participantRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode":    "42203",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	participant, err := h.eventUsecase.PostViewBooking(participantRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "40003",
			"responseMessage": "Usecase PostViewBooking is not working properly: " + err.Error(),
		})
		return
	}

	if participant.SessionID == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"responseCode":     "20000",
			"responseMessage":  "Request has been proceeded successfully",
			"name":             participant.Name,
			"email":            participant.Email,
			"phoneNo":          participant.PhoneNo,
			"sessionID":        participant.SessionID,
			"sessionName":      "GROW Center Anniversary 1st Service",
			"scanStatus":       "Not Scanned",
			"requestedSeat":    participant.RequestedSeat,
			"registrationCode": participant.RegistrationCode,
			"qrCode":           participant.QRCode,
			"additionalInfo":   "Session 1: 10.30 Open Gate 09.45 Please come before 10.25",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode":     "20000",
		"responseMessage":  "Request has been proceeded successfully",
		"name":             participant.Name,
		"email":            participant.Email,
		"phoneNo":          participant.PhoneNo,
		"sessionID":        participant.SessionID,
		"sessionName":      "GROW Center Anniversary 2nd Service",
		"scanStatus":       "Not Scanned",
		"requestedSeat":    participant.RequestedSeat,
		"reasons":          participant.Reasons,
		"registrationCode": participant.RegistrationCode,
		"qrCode":           participant.QRCode,
		"additionalInfo":   "Session 2: 17.00 Open Gate 16.30 Please come before 16.55",
	})
}

func (h *EventHandler) SessionInfo(ctx *gin.Context) {
	var adminRequest *model.SessionInfoRequest

	err := ctx.ShouldBindUri(&adminRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode":    "42204",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	session, err := h.eventUsecase.GetSessionInfo(*adminRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode":    "40004",
			"responseMessage": "Usecase GetSessionInfo is not working properly: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode":         "20000",
		"responseMessage":      "Request has been proceeded successfully",
		"sessionID":            session.ID,
		"sessionName":          session.Name,
		"sessionTime":          session.Time,
		"sessionCapacity":      1750,
		"sessionBookedCount":   session.FilledCapacity,
		"sessionAttendedCount": session.ScannedSeat,
		"sessionAvailableSeat": session.UnscannedSeat,
	})
}

func (h *EventHandler) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": "200",
	})
}
