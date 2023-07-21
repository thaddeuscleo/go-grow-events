package usecase

import (
	"errors"
	"fmt"
	"go-grow-events/model"
	"go-grow-events/repository"
	"go-grow-events/util"
)

type EventUsecase interface {
	PostRegisterSession(request *model.RegisterParticipantRequest) (*model.Participant, error)
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
	
	participantInputtedEmail := request.Email
	checkEmailString, err := util.EmailStringRegex(participantInputtedEmail)
	if !checkEmailString {
		return &participant, err
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
	checkPhoneNoAvailable, err := e.repo.FindParticipantByEmail(participantInputtedPhoneNo)
	if err != nil {
		return &participant, err
	}

	if checkPhoneNoAvailable.ID != 0 {
		return &participant, errors.New("user with this phone number is already exist")
	}

	participant.PhoneNo = request.PhoneNo
	// Bikin is scanned jadi 0 v
	participant.IsScanned = 0


	/*
	Bikin kode qr codenya, isi qr codenya itu kode RI0930, simpen aja kode bookingnya di db
	simpen kode qr code di db
	cek booking by PHONE NUMBER - isi qr, ibadahnya jam brp, kode booking

	wktu scan qr, is scanned bakal nambah 1 kali
	lalu cek klo is scannednya = requested seat, klo udah gabisa scan lagi
	
	bikin admin
	terus tambah di db yg udah ngescan brapa sama yg belum scan brapa
	*/


	if request.SessionID > 2 {
		return &participant, errors.New("no sessions existed on that ID")
	}

	if request.SessionID == 0 {
		return &participant, errors.New("no sessions existed on that ID")
	}

	if request.SessionID == 1 {
		session, err := e.repo.FindSessionBySessionID(request.SessionID)
		if err != nil {
			return &participant, err
		}

		if request.RequestedSeat > 6 {
			return &participant, errors.New("maximum number of seat is 6")
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

		participant.RegistrationCode = fmt.Sprintf("GCANNIV202301%06d", participant.ID)
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

	participant.RegistrationCode = fmt.Sprintf("GCANNIV202302%06d", participant.ID)
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