package test

import (
	models "PAPS-lab5/user_service/api/handlers"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const host = "http://0.0.0.0:8080"
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
