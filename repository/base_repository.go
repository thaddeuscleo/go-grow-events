package repository

import (
	"go-grow-events/model"

	"gorm.io/gorm"
)

type BaseRepository interface {
    // User Account Management
    CreateParticipantToDB(participant *model.Participant) (*model.Participant, error)
    FindParticipantByEmail(email string) (*model.Participant, error)
    FindParticipantByPhoneNo(phoneNo string) (*model.Participant, error)
    FindParticipantByCode(code string) (*model.Participant, error)
    UpdateParticipantToDB(participant *model.Participant) (*model.Participant, error)
	
	UpdateSessionToDB(session *model.Session) (*model.Session, error)
	FindSessionBySessionID(sessionID int) (*model.Session, error)
}

type baseRepository struct {
    db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *baseRepository {
    return &baseRepository{db}
}

func (r *baseRepository) CreateParticipantToDB(participant *model.Participant) (*model.Participant, error) {
    err := r.db.Create(&participant).Error
    if err != nil {
        return participant, err
    }

    return participant, nil
}

func (r *baseRepository) FindParticipantByEmail(email string) (*model.Participant, error) {
    var participant *model.Participant
    err := r.db.Where("email = ?", email).Find(&participant).Error
    if err != nil {
        return participant, err
    }

    return participant, nil
}

func (r *baseRepository) FindParticipantByPhoneNo(phoneNo string) (*model.Participant, error) {
    var participant *model.Participant
    err := r.db.Where("email = ?", phoneNo).Find(&participant).Error
    if err != nil {
        return participant, err
    }

    return participant, nil
}

func (r *baseRepository) FindParticipantByCode(code string) (*model.Participant, error) {
    var participant * model.Participant
    err := r.db.Where("registration_code = ?", code).Find(&participant).Error
    if err != nil {
        return participant, err
    }

    return participant, nil 
}

func (r *baseRepository) UpdateParticipantToDB(participant *model.Participant) (*model.Participant, error) {
    err := r.db.Save(&participant).Error
    if err != nil {
        return participant, err
    }

    return participant, nil
}

func (r *baseRepository) UpdateSessionToDB(session *model.Session) (*model.Session, error) {
    err := r.db.Save(&session).Error
    if err != nil {
        return session, err
    }

    return session, nil
}

func (r *baseRepository) FindSessionBySessionID(sessionID int) (*model.Session, error) {
    var session *model.Session
    err := r.db.Where("id = ?", sessionID).Find(&session).Error
    if err != nil {
        return session, err
    }

    return session, nil
}