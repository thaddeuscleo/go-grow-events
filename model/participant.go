package model

import (
	"database/sql"
	"time"
)

type Participant struct {
	ID					int				`json:"id"`
	Name				string			`json:"userName"`
	Email				string			`json:"userEmail"`
	PhoneNo				string			`json:"phoneNo"`
	RequestedSeat		int				`json:"requestedSeat"`
	SessionID			int				`json:"sessionID"`
	IsScanned			int				`json:"isScanned"`
	Reasons				string			`json:"reasons"`
	QRCode				string			`json:"qrCode"`
	RegistrationCode	string			`json:"registrationCode"`
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
	DeletedAt			sql.NullTime	`json:"deletedAt"`
}

type RegisterParticipantRequest struct {
	Name				string			`json:"name" binding:"required"`
	Email				string			`json:"email" binding:"required"`
	PhoneNo				string			`json:"phoneNo" binding:"required"`
	RequestedSeat		int				`json:"requestedSeat" binding:"required"`
	Reasons				string			`json:"reasons"`
	SessionID			int				`json:"sessionID" binding:"required"`
}

