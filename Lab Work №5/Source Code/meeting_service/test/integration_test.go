package test

import (
	models "PAPS-lab5/meeting_service/api/handlers"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const host = "http://0.0.0.0:8081"
const tm = "01.05.2024 14:00"

func TestAddMeeting(t *testing.T) {
	timeM, _ := time.Parse("02.01.2006 15:04", tm)
	reqb := models.AddMeetingRequestBody{
		Name:        "Test",
		Description: "test 1",
		MeetingTime: timeM,
		StudentId:   1,
		IsOnline:    false,
		BaseRequestBody: models.BaseRequestBody{
			ProfessorId: 1,
		},
	}
	requestBody, err := json.Marshal(reqb) //тело запроса
	if err != nil {
		log.Print(err)
		requestBody = nil
	}
	request, err := http.NewRequest("POST", host+"/meeting/add", bytes.NewBuffer(requestBody)) //создаем запрос
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	defer res.Body.Close()
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestGetMeeting(t *testing.T) {
	id := 1
	request, err := http.NewRequest("GET", fmt.Sprint(host, "/meeting/", id), nil) //создаем запрос
	if err != nil {
		log.Print(err)
		t.Fail()
	}

	q := request.URL.Query()
	q.Add("prof", "1")
	request.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print(err)
		t.Fail()
	}
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var meetResp struct {
		Meeting models.Meeting `json:"meeting"`
	}
	decoder := json.NewDecoder(res.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&meetResp)
	if err != nil {
		log.Print(err)
		t.Fail()
	}

	timeM, _ := time.Parse("02.01.2006 15:04", tm)
	expected := models.AddMeetingRequestBody{
		Name:        "Test",
		Description: "test 1",
		MeetingTime: timeM,
		StudentId:   1,
		IsOnline:    false,
		BaseRequestBody: models.BaseRequestBody{
			ProfessorId: 1,
		},
	}
	actual := meetResp.Meeting
	assert.Equal(t, id, meetResp.Meeting.Id)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.IsOnline, actual.IsOnline)
	assert.Equal(t, expected.MeetingTime, actual.MeetingTime)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProfessorId, 1)
	assert.Equal(t, expected.StudentId, actual.StudentParticipantId)
}
