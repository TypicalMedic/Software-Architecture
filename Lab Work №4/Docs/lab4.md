 ## REST API для сервиса работы со встречами
 Были определены следующие методы:
 - *Посмотреть расписание преподавателя (GET)*
 - *Посмотреть встречи с определенным студентом (GET)*
 - *Посмотреть встречу (GET)*
 - *Назначить встречу (POST)*
 - *Отменить встречу (DELETE)*
 - *Перенести встречу (PUT)*
 - *Обновить информацию о встрече (PUT)*
 - *Подключить облачное хранилище (POST)*
 - *Отключить облачное хранилище (DELETE)*
 
 ## Посмотреть расписание преподавателя (GET)
 
### URL: `/meetings`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, для которого получается
   расписание
 
		 {
			 "professor_id": 12345678
		 }

###  **Ответы:**

#### HTTP: 200 - операция успешна.
Заголовки ответа:

 - Content-Type: application/json
 
<details>
  <summary> Тело ответа (список всех предстоящих встреч с информацией о них)</summary>
    

    {
	    "meetings":[
		    {
			    "id": 1,
			    "name": "this is a very important meeting",
			    "description": "there are provided some details",
			    "meeting_time": "2024-01-30T12:35:24.000Z", // время в формате RFC3339
			    "student_participant": {
				    "id": 1,
				    "name": "Ivan",
				    "surname": "Ivanov",
				    "middlename": "Ivanovich",
				    "cource": 3,
				    "project_theme": "Theme name"
			    },
			    "is_online": true
		    },
		    {
			    "id": 2,
			    "name": "this is a not very important meeting",
			    "description": "there are provided some details",
			    "meeting_time": "2024-01-30T16:35:24.000Z", // время в формате RFC3339
			    "student_participant": {
				    id: 2,
				    "name": "Pavel",
				    "surname": "Pavlov",
				    "middlename": "Pavlovich",
				    "cource": 2,
				    "project_theme": "Theme name"
			    },
			    "is_online": false
		    }
	    ]
    }
 
</details> 

#### HTTP: 401 - не авторизован. Для получения запрашиваемого ответа нужна аутентификация.

#### HTTP: 400 - неправильный запрос. ID пользователя неправильного формата. 

#### HTTP: 404 - страница не найдена. ID пользователя  не найден.


##  Посмотреть встречи с определенным студентом (GET)

### URL: `/meetings/filter?student={student_id}`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя
### **Параметры запроса:**

 - `student_id` (int) - ID студента, с которым назначены встречи

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, для которого получается
   расписание
 
		 {
			 "professor_id": 12345678
		 }

###  **Ответы:**

#### HTTP: 200 - операция успешна.
Заголовки ответа:

 - Content-Type: application/json
 
<details>
  <summary> Тело ответа (список всех предстоящих встреч с информацией о них)</summary>
    

    {
	    "meetings":[
		    {
			    "id": 1,
			    "name": "this is a very important meeting",
			    "description": "there are provided some details",
			    "meeting_time": "2024-01-30T12:35:24.000Z", // время в формате RFC3339
			    "student_participant": {
				    "id": 1,
				    "name": "Ivan",
				    "surname": "Ivanov",
				    "middlename": "Ivanovich",
				    "cource": 3,
				    "project_theme": "Theme name"
			    },
			    "is_online": true
		    },
		    {
			    "id": 3,
			    "name": "this is a not very important meeting",
			    "description": "there are provided some details",
			    "meeting_time": "2024-02-20T16:35:24.000Z", // время в формате RFC3339
			    "student_participant": {
				    "id": 1,
				    "name": "Ivan",
				    "surname": "Ivanov",
				    "middlename": "Ivanovich",
				    "cource": 3,
				    "project_theme": "Theme name"
			    },
			    "is_online": false
		    }
	    ]
    }
 
</details> 

#### HTTP: 401 - не авторизован. Для получения запрашиваемого ответа нужна аутентификация.

#### HTTP: 400 - неправильный запрос. ID пользователя или студента неправильного формата. 

#### HTTP: 404 - страница не найдена. ID пользователя или студента не найдены.

 
 ##   Посмотреть встречу преподавателя (GET)
 
 
### URL: `/meeting/{meeting_id}`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя
 
### **Параметры запроса:**

 - `meeting_id` (int) - ID  встречи

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, для которого получается
   встреча
 
		 {
			 "professor_id": 12345678
		 }

###  **Ответы:**

#### HTTP: 200 - операция успешна.
Заголовки ответа:

 - Content-Type: application/json
 
<details>
  <summary> Тело ответа (список всех предстоящих встреч с информацией о них)</summary>
    

    {
	    "id": 1,
	    "name": "this is a very important meeting",
	    "description": "there are provided some details",
	    "meeting_time": "2024-01-30T12:35:24.000Z", // время в формате RFC3339
	    "student_participant": {
		    "id": 1,
		    "name": "Ivan",
		    "surname": "Ivanov",
		    "middlename": "Ivanovich",
		    "cource": 3,
		    "project_theme": "Theme name"
	    },
	    "is_online": true
    }
 
</details> 

#### HTTP: 401 - не авторизован. Для получения запрашиваемого ответа нужна аутентификация.

#### HTTP: 400 - неправильный запрос. ID встречи неправильного формата. 

#### HTTP: 404 - страница не найдена. ID встречи или преподавателя не найдены.

 
 ##    Назначить встречу (POST)
 
### URL: `/meeting/add`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, который организует встречу
  - `name (string)` - название встречи
  - `description (string)` - описание встречи
  - `meeting_time (string)` -  время встречи в формате RFC3339
  - `student_participant_id (int)` - ID студента, с кем будет встреча
  - `is_online (bool)` - встреча будет онлайн или нет
 
		{
		    "professor_id": 1,
		    "name": "this is a very important meeting",
		    "description": "there are provided some details",
		    "meeting_time": "2024-01-30T12:35:24.000Z", 
		    "student_participant_id":  1,
		    "is_online": true
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело запроса неправильного формата. 

#### HTTP: 404 - ID пользователя  не найден.

 ##     Отменить встречу (DELETE)
 
### URL: `/meeting/{meeting_id}/delete`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Параметры запроса:**

 - `meeting_id` (int) - ID  встречи

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, встреча которого отменяется
 
		{
		    "professor_id": 1
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело или параметры запроса неправильного формата. 

#### HTTP: 404 - ID пользователя  не найден.

 ##      Перенести встречу (PUT)
 
### URL: `/meeting/{meeting_id}/reschedule`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Параметры запроса:**

 - `meeting_id` (int) - ID  встречи

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, встреча которого переносится
 - `new_meeting_time` (string) -  новое время встречи в формате RFC3339
 
		{
		    "professor_id": 1,
		    "new_meeting_time": "2024-01-30T12:35:24.000Z"
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело или параметры запроса неправильного формата. 

#### HTTP: 404 -  ID пользователя  не найден.
 
 ## Обновить информацию о встрече (PUT)
 
### URL: `/meeting/{meeting_id}/update`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Параметры запроса:**

 - `meeting_id` (int) - ID  встречи

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, который организует встречу
 
 Измененная информация о встрече:
  - `name (string)` - название встречи
  - `description (string)` - описание встречи
  - `student_participant_id (int)` - ID студента, с кем будет встреча
  - `is_online (bool)` - встреча будет онлайн или нет
 
		{
		    "professor_id": 1,
		    "name": "this is a very important meeting",
		    "description": "there are provided some details",
		    "student_participant_id":  1,
		    "is_online": true
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело или параметры запроса неправильного формата. 

#### HTTP: 404 - ID пользователя  не найден.

 ##        Подключить календарь (POST)
 
### URL: `/integrations/calendar/connect`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, который подключает календарь 
 -  `email` (string) - почта преподавателя, который подключает календарь 
 
		{
		    "professor_id": 1,
		    "email": "email@domain.com"
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело или параметры запроса неправильного формата. 


#### HTTP: 404 - ID пользователя  не найден.
 
 ##         Отключить календарь  (DELETE)
 
### URL: `/integrations/calendar/disconnect`

### **Заголовки запроса:**

 - authorization_token - токен авторизации пользователя

### **Тело запроса:**

 - `professor_id` (int) - ID преподавателя, который отключает календарь 
 
		{
		    "professor_id": 1
	    }

###  **Ответы:**

#### HTTP: 200 - операция успешна.

#### HTTP: 401 - не авторизован. Для выполнения запроса нужна аутентификация.

#### HTTP: 400 - неправильный запрос. Тело или параметры запроса неправильного формата. 


#### HTTP: 404 - ID пользователя  не найден.

# Реализация API

Реализация проводилась на языке Go, основной код обработки представлен ниже:
 
<details>
  <summary> Код</summary>
    

    func (h *Handler) Meetings(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization_token")
	err := h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result := struct {
		Meetings []Meeting `json:"meetings"`
	}{
		h.App.Data.GetProfessorMeetings(rb.ProfessorId),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	}
	
	func (h *Handler) MeetingsStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stud, err := strconv.Atoi(vars["student_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("authorization_token")
	err = h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	result := struct {
		Meetings []Meeting `json:"meetings"`
	}{
		h.App.Data.GetProfessorMeetingsStudent(rb.ProfessorId, stud),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	}

	func (h *Handler) MeetingById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	meeting, err := strconv.Atoi(vars["meeting_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("authorization_token")
	err = h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := h.App.Data.MeetingById(meeting)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	}

	func (h *Handler) MeetingAdd(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization_token")
	err := h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	m := Meeting{
		Name:        rb.Name,
		Description: rb.Description,
		Time:        rb.Time,
		IsOnline:    rb.IsOnline,
	}
	h.App.Data.MeetingAdd(m, rb.ProfessorId, rb.StudentId)

	h.App.Calendar.AddMeeting(m)
	w.WriteHeader(http.StatusOK)
	}

	func (h *Handler) MeetingDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meeting, err := strconv.Atoi(vars["meeting_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("authorization_token")
	err = h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = h.App.Data.MeetingDelete(meeting)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.App.Calendar.DeleteMeeting(meeting)
	w.WriteHeader(http.StatusOK)
	}

	func (h *Handler) MeetingReschedule(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	meeting, err := strconv.Atoi(vars["meeting_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("authorization_token")
	err = h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = h.App.Data.MeetingReschedule(meeting, rb.NewTime)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.App.Calendar.RescheduleMeeting(meeting, rb.NewTime)
	w.WriteHeader(http.StatusOK)
	}

	func (h *Handler) MeetingUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meeting, err := strconv.Atoi(vars["meeting_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authHeader := r.Header.Get("authorization_token")
	err = h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	m := Meeting{
		Id:          meeting,
		Name:        rb.Name,
		Description: rb.Description,
		Time:        rb.Time,
		IsOnline:    rb.IsOnline,
	}
	err = h.App.Data.MeetingUpdate(m, rb.StudentId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	h.App.Calendar.UpdateMeeting(m, rb.StudentId)
	w.WriteHeader(http.StatusOK)
	}

	func (h *Handler) CalendarDisconnect(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization_token")
	err := h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	h.App.Calendar.DisonnectService(rb.ProfessorId)
	w.WriteHeader(http.StatusOK)
	}

	func (h *Handler) CalendarConnect(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization_token")
	err := h.App.Auth.CheckAuth(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var rb RequestBody
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&rb)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.Data.FindUser(rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	err = h.App.Auth.CheckUserIdentity(authHeader, rb.ProfessorId)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	h.App.Calendar.ConnectToService(rb.Email, rb.ProfessorId)
	w.WriteHeader(http.StatusOK)
	}
	
	func main() {
	r := mux.NewRouter()
	handler := Handler{
		App: &App{
			Data:     &Data{},
			Calendar: &Calendar{},
			Auth:     &Authentification{},
		}}
	r.HandleFunc("/meetings", handler.Meetings).Methods("GET")
	r.HandleFunc("/meetings/filter", handler.MeetingsStudent).Methods("GET").Queries("student", "{student_id}")
	r.HandleFunc("/meeting/{meeting_id}", handler.MeetingById).Methods("GET")
	r.HandleFunc("/meeting/add", handler.MeetingAdd).Methods("POST")
	r.HandleFunc("/meeting/{meeting_id}/delete", handler.MeetingDelete).Methods("DELETE")
	r.HandleFunc("/meeting/{meeting_id}/reschedule", handler.MeetingReschedule).Methods("PUT")
	r.HandleFunc("/meeting/{meeting_id}/update", handler.MeetingUpdate).Methods("PUT")
	r.HandleFunc("/integrations/calendar/disconnect", handler.CalendarDisconnect).Methods("DELETE")
	r.HandleFunc("/integrations/calendar/connect", handler.CalendarConnect).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowCredentials:    true,
		AllowedMethods:      []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:      []string{"*"},
		AllowPrivateNetwork: true,
	})

	handler1 := c.Handler(r)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", handler1))
	}
 
</details> 

  # Тестирование API с Postman
 
  Для каждого метода были созданы запросы, проверяющие корректность возвращаемых кодов. Для проверки корректности кодов были написаны тесты. Большинство запросов проверяют одни и те же ошибки (такие как остуствие авторизации, id пользователя), поэтому они будут опущены после первого раза.
  
## Посмотреть расписание преподавателя (GET)

### Корректный запрос

 1. Строка запроса
    
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/111.png)
 
 2. Заголовки и параметры запроса
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png)
 
 3. Ответ
    
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/114.png)
 
 4. Код автотестов
    
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/115.png)
 
 5. Результат тестов
     
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/116.png)

### Нет авторизации
 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/111.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/122.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png)

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/124.png)

 4. Код автотестов 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/125.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/126.png)

### Пользователя не существует

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/111.png)
 
 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/133.png)

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/134.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/135.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/136.png)

### Некорректное тело запроса

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/111.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/143.png)

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/144.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/145.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/146.png)

### Получаются данные не авторизированного пользователя
 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/111.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/153.png)

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/154.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/155.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/156.png)

## Посмотреть встречи с определенным студентом (GET)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/211.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/212.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png)

 3. Ответ 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/215.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/216.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/217.png)
 
### Некорректные параметры запроса

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/221.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/222.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png)

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/225.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/226.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/227.png)
 
## Посмотреть встречу (GET)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/311.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/314.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/315.png)

 5. Результат тестов

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/316.png)
 
### Встречи не существует

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/321.png)
 
 2. Заголовки и параметры запроса 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/324.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/325.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/326.png)

## Назначить встречу (POST)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/411.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/413.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/414.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/415.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/416.png)

## Отменить встречу (DELETE)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/511.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/514.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/515.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/516.png)

## Перенести встречу (PUT)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/611.png)

 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/613.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/614.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/615.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/616.png)

## Обновить информацию о встрече (PUT)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/711.png)

 2. Заголовки и параметры запроса 
 
 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/713.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/714.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/715.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/716.png)

## Подключить облачное хранилище (POST)

### Корректный запрос

 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/811.png)
 
 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/813.png) 

 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/814.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/815.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/816.png)

## Отключить облачное хранилище (DELETE)

### Корректный запрос
 
 1. Строка запроса

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/911.png)
 
 2. Заголовки и параметры запроса 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/112.png) 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/113.png) 
 
 3. Ответ 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/914.png)

 4. Код автотестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/915.png)

 5. Результат тестов 

 ![enter image description here](https://github.com/TypicalMedic/Software-Architecture/blob/LabWork4/Lab%20Work%20%E2%84%964/Docs/img/916.png)
