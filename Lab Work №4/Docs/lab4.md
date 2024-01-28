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
    

    {
	    ------
	}
 
</details> 

  # Тестирование API с Postman
 
  Для каждого метода были созданы запросы, проверяющие корректность возвращаемых кодов. Для проверки корректности кодов были написаны тесты. Большинство запросов проверяют одни и те же ошибки (такие как остуствие авторизации, id пользователя), поэтому они будут опущены после первого раза.
  

## Посмотреть расписание преподавателя (GET)

### Корректный запрос

 1. Строка запроса
 ![enter image description here](111)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](113)
 4. Ответ 
 ![enter image description here](114)
 5. Код автотестов 
 ![enter image description here](115)
 6. Результат тестов 
 ![enter image description here](116)

### Нет авторизации
 1. Строка запроса
 ![enter image description here](111)
 3. Заголовки и параметры запроса 
 ![enter image description here](122) 
 ![enter image description here](113)
 4. Ответ 
 ![enter image description here](124)
 5. Код автотестов 
 ![enter image description here](125)
 6. Результат тестов 
 ![enter image description here](126)
### Пользователя не существует
 1. Строка запроса
 ![enter image description here](111)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](133)
 4. Ответ 
 ![enter image description here](134)
 5. Код автотестов 
 ![enter image description here](135)
 6. Результат тестов 
 ![enter image description here](136)
### Некорректное тело запроса
 1. Строка запроса
 ![enter image description here](111)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](143)
 4. Ответ 
 ![enter image description here](144)
 5. Код автотестов 
 ![enter image description here](145)
 6. Результат тестов 
 ![enter image description here](146)
### Получаются данные не авторизированного пользователя
 1. Строка запроса
 ![enter image description here](111)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](153)
 4. Ответ 
 ![enter image description here](154)
 5. Код автотестов 
 ![enter image description here](155)
 6. Результат тестов 
 ![enter image description here](156)

## Посмотреть встречи с определенным студентом (GET)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](211)
 3. Заголовки и параметры запроса 
 ![enter image description here](212) 
 ![enter image description here](112) 
 ![enter image description here](113)
 4. Ответ 
 ![enter image description here](215)
 5. Код автотестов 
 ![enter image description here](216)
 6. Результат тестов 
 ![enter image description here](217)
 
### Некорректные параметры запроса

 1. Строка запроса
 ![enter image description here](221)
 3. Заголовки и параметры запроса 
 ![enter image description here](222) 
 ![enter image description here](112) 
 ![enter image description here](113)
 4. Ответ 
 ![enter image description here](225)
 5. Код автотестов 
 ![enter image description here](226)
 6. Результат тестов 
 ![enter image description here](227)
 
## Посмотреть встречу (GET)

### Корректный запрос

 1. Строка запроса
 ![enter image description here](311)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](113) 
 4. Ответ 
 ![enter image description here](314)
 5. Код автотестов 
 ![enter image description here](315)
 6. Результат тестов 
 ![enter image description here](316)
 
### Встречи не существует

 1. Строка запроса
 ![enter image description here](321)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](113) 
 4. Ответ 
 ![enter image description here](324)
 5. Код автотестов 
 ![enter image description here](325)
 6. Результат тестов 
 ![enter image description here](326)

## Назначить встречу (POST)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](411)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](413) 
 4. Ответ 
 ![enter image description here](414)
 5. Код автотестов 
 ![enter image description here](415)
 6. Результат тестов 
 ![enter image description here](416)

## Отменить встречу (DELETE)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](511)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](113) 
 4. Ответ 
 ![enter image description here](514)
 5. Код автотестов 
 ![enter image description here](515)
 6. Результат тестов 
 ![enter image description here](516)

## Перенести встречу (PUT)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](611)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](613) 
 4. Ответ 
 ![enter image description here](614)
 5. Код автотестов 
 ![enter image description here](615)
 6. Результат тестов 
 ![enter image description here](616)

## Обновить информацию о встрече (PUT)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](711)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](713) 
 4. Ответ 
 ![enter image description here](714)
 5. Код автотестов 
 ![enter image description here](715)
 6. Результат тестов 
 ![enter image description here](716)

## Подключить облачное хранилище (POST)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](811)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](813) 
 4. Ответ 
 ![enter image description here](814)
 5. Код автотестов 
 ![enter image description here](815)
 6. Результат тестов 
 ![enter image description here](816)

## Отключить облачное хранилище (DELETE)

### Корректный запрос
 1. Строка запроса
 ![enter image description here](911)
 3. Заголовки и параметры запроса 
 ![enter image description here](112) 
 ![enter image description here](113) 
 4. Ответ 
 ![enter image description here](914)
 5. Код автотестов 
 ![enter image description here](915)
 6. Результат тестов 
 ![enter image description here](916)