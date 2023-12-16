package usecase

import (
	"errors"
	"fmt"
	"go-grow-events/model"
	"go-grow-events/repository"
	"go-grow-events/util"
	"strings"
)

type EventUsecase interface {
	PostRegisterSession(request *model.RegisterParticipantRequest) (*model.Participant, error)
	PostVerifySession(request *model.VerifyParticipantRequest) (*model.Participant, error)
	PostViewBooking(request *model.ViewBookingRequest) (*model.Participant, error)

	GetSessionInfo(request model.SessionInfoRequest) (*model.Session, error)
}

type eventUsecase struct {
	repo repository.BaseRepository
}

func NewEventUsecase(repo repository.BaseRepository) *eventUsecase {
	return &eventUsecase{repo}
}

func (e *eventUsecase) PostRegisterSession(request *model.RegisterParticipantRequest) (*model.Participant, error) {
	participant := model.Participant{}
	participant.Name = request.Name
	
	participantInputtedEmail := strings.ToLower(request.Email)
	checkEmailString, err := util.EmailStringRegex(participantInputtedEmail)
	if !checkEmailString {
		return nil, err
	}

	checkEmailAvailable, err := e.repo.FindParticipantByEmail(participantInputtedEmail)
	if err != nil {
		return &participant, err
	}

	if checkEmailAvailable.ID != 0 {
		return &participant, errors.New("user with this email is already exist")
	}

	participant.Email = participantInputtedEmail

	participantInputtedPhoneNo := request.PhoneNo
	checkPhoneNoAvailable, err := e.repo.FindParticipantByPhoneNo(participantInputtedPhoneNo) // new changes to check phone number
	if err != nil {
		return &participant, err
	}

	if checkPhoneNoAvailable.ID != 0 {
		return &participant, errors.New("user with this phone number is already exist")
	}

	participant.PhoneNo = participantInputtedPhoneNo // new changes to check phone number
	participant.IsScanned = 0

	if request.SessionID > 2 {
		return &participant, errors.New("no sessions existed on that ID")
	}

	if request.SessionID == 0 {
		return &participant, errors.New("no sessions existed on that ID")
	}

	/*if request.SessionID == 1 {
		session, err := e.repo.FindSessionBySessionID(request.SessionID)
		if err != nil {
			return &participant, err
		}

		if request.RequestedSeat > 4 {
			return &participant, errors.New("maximum number of seat is 4")
		}
		
		participant.RequestedSeat = request.RequestedSeat
		
		if session.EmptyCapacity == 0 {
			return &participant, errors.New("there is no empty seat available anymore")
		}

		session.FilledCapacity += participant.RequestedSeat
		session.EmptyCapacity -= participant.RequestedSeat
		session.UnscannedSeat += participant.RequestedSeat
		participant.Reasons = ""
		participant.SessionID = 1

		newParticipant, err := e.repo.CreateParticipantToDB(&participant)
		if err != nil {
			return newParticipant, err
		}
		
		participant.RegistrationCode = fmt.Sprintf("GC%04d", participant.ID)
		participant.QRCode, err = util.GenerateQRCode(participant.RegistrationCode)
		if err != nil {
			return newParticipant, err
		}

		updatedParticipant, err := e.repo.UpdateParticipantToDB(&participant)
		if err != nil {
			return updatedParticipant, err
		}

		_, err = e.repo.UpdateSessionToDB(session)
		if err != nil {
			return &participant, err
		}

		return newParticipant, nil
	}

	if request.RequestedSeat > 2 {
		return &participant, errors.New("maximum number of seat is 2")
	}

	participant.RequestedSeat = request.RequestedSeat

	session, err := e.repo.FindSessionBySessionID(request.SessionID)
	if err != nil {
		return &participant, err
	}
	
	session.FilledCapacity += participant.RequestedSeat
	session.EmptyCapacity -= participant.RequestedSeat
	session.UnscannedSeat += participant.RequestedSeat

	participantInputtedReasons := request.Reasons
	if participantInputtedReasons == "" {
		return &participant, errors.New("reasons cannot be empty")
	}

	participant.Reasons = participantInputtedReasons
	participant.SessionID = 2

	newParticipant, err := e.repo.CreateParticipantToDB(&participant)
	if err != nil {
		return newParticipant, err
	}

	participant.RegistrationCode = fmt.Sprintf("GC%04d", participant.ID)
	participant.QRCode, err = util.GenerateQRCode(participant.RegistrationCode)
	if err != nil {
		return newParticipant, err
	}

	updatedParticipant, err := e.repo.UpdateParticipantToDB(&participant)
	if err != nil {
		return updatedParticipant, err
	}

	_, err = e.repo.UpdateSessionToDB(session)
	if err != nil {
		return &participant, err
	}

	return newParticipant, nil*/
	session, err := e.repo.FindSessionBySessionID(request.SessionID)
	if err != nil {
		return &participant, err
	}

	if request.RequestedSeat > 4 {
		return &participant, errors.New("maximum number of seat is 4")
	}
		
	participant.RequestedSeat = request.RequestedSeat
		
	if session.EmptyCapacity == 0 {
		return &participant, errors.New("there is no empty seat available anymore")
	}

	session.FilledCapacity += participant.RequestedSeat
	session.EmptyCapacity -= participant.RequestedSeat
	session.UnscannedSeat += participant.RequestedSeat
	participant.Reasons = ""
	participant.SessionID = request.SessionID

	newParticipant, err := e.repo.CreateParticipantToDB(&participant)
	if err != nil {
		return newParticipant, err
	}
		
	participant.RegistrationCode = fmt.Sprintf("GC%04d", participant.ID)
	participant.QRCode, err = util.GenerateQRCode(participant.RegistrationCode)
	if err != nil {
		return newParticipant, err
	}

	updatedParticipant, err := e.repo.UpdateParticipantToDB(&participant)
	if err != nil {
		return updatedParticipant, err
	}

	_, err = e.repo.UpdateSessionToDB(session)
	if err != nil {
		return &participant, err
	}

	return updatedParticipant, nil
}

func (e *eventUsecase) PostVerifySession(request *model.VerifyParticipantRequest) (*model.Participant, error) {
	registrationCode := request.RegistrationCode
	
	existingParticipant, err := e.repo.FindParticipantByCode(registrationCode)
	if err != nil {
		return nil, errors.New("registration code is not found")
	}

	if existingParticipant.IsScanned == existingParticipant.RequestedSeat {
		return existingParticipant, errors.New("already scan all the requested seats")
	}
	existingParticipant.IsScanned += 1

	currentSession, err := e.repo.FindSessionBySessionID(existingParticipant.SessionID)
	if err != nil {
		return existingParticipant, errors.New("session is not found")
	}

	currentSession.ScannedSeat += 1
	currentSession.UnscannedSeat -= 1

	updatedParticipant, err := e.repo.UpdateParticipantToDB(existingParticipant)
	if err != nil {
		return existingParticipant, err
	}

	_, err = e.repo.UpdateSessionToDB(currentSession)
	if err != nil {
		return existingParticipant, err
	}

	return updatedParticipant, nil
}

func (e *eventUsecase) PostViewBooking(request *model.ViewBookingRequest) (*model.Participant, error) {
	booking := request.Booking

	participant, err := e.repo.FindParticipantByMultipleFilter(booking)
	if err != nil {
		return participant, err
	}

	return participant, nil
}

func (e *eventUsecase) GetSessionInfo(request model.SessionInfoRequest) (*model.Session, error) {
	session, err := e.repo.FindSessionBySessionID(request.ID)
	if err != nil {
		return session, err
	}

	if session.ID == 0 {
		return session, errors.New("session does not exist")
	}

	return session, nil
}
