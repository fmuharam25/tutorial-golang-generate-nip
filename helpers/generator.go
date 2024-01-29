package helpers

import (
	"errors"
	"fmt"
)

// Create type participant
type Participant struct {
	Gender  string
	Date    string
	StartId int
}

func CreateNIP(participant Participant) (string, error) {

	var nip string
	var semester string = checkSemester(participant.Date)
	var gender string = checkGender(participant.Gender)
	var id string = fmt.Sprintf("%05d", participant.StartId)

	if semester == " " {
		return "", errors.New("Date not valid, use yyy-mm-dd format")
	}

	if gender == " " {
		return "", errors.New("Gender not valid")
	}

	nip = "AR" + gender + semester + "-" + id
	return nip, nil

}

func GenerateNIP(participant Participant, count int) ([]string, error) {
	nips := []string{}

	for i := 0; i < count; i++ {
		nip, err := CreateNIP(Participant{participant.Gender, participant.Date, participant.StartId + i})

		if err != nil {
			return []string{}, errors.New("Error generate nip")
		}

		nips = append(nips, nip)
	}
	return nips, nil

}

func CreateNextNIP(nip string) (string, error) {
	id := checkId(nip)
	prefix := checkPrefix(nip)
	nextId := fmt.Sprintf("%05d", id+1)

	if id == -1 || prefix == " " {
		return " ", errors.New("NIP not valid")
	}

	return prefix + "-" + nextId, nil
}

func GenerateNextNIP(nip string, count int) ([]string, error) {
	var nextId, nextNip string
	nips := []string{}
	id := checkId(nip)
	prefix := checkPrefix(nip)

	if id == -1 || prefix == " " {
		return []string{}, errors.New("NIP not valid")
	}

	for i := 1; i < count; i++ {
		nextId = fmt.Sprintf("%05d", id+i)
		nextNip = prefix + "-" + nextId
		nips = append(nips, nextNip)
	}
	return nips, nil

}
