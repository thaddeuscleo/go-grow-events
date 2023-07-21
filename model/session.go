package model

import (
	"database/sql"
	"time"
)

type Session struct {
	ID					int
	Name				string			`json:"sessionName"`
	Time				string			`json:"sessionTime"`
	EmptyCapacity		int				`json:"sessionCapacity"`
	FilledCapacity		int				`json:"filledCapacity"`
	ScannedSeat			int				`json:"scannedSeat"`
	UnscannedSeat		int				`json:"unscannedSeat"`
	CreatedAt 			time.Time		`json:"createdAt"`
	UpdatedAt 			time.Time		`json:"updatedAt"`
	DeletedAt			sql.NullTime	`json:"deletedAt"`
}

type CreateSessionRequest struct {
	Name				string			`json:"sessionName"`
	Time				string			`json:"sessionTime"`
	EmptyCapacity		int				`json:"sessionCapacity"`
}