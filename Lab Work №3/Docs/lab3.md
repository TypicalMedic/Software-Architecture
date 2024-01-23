Для работы был выбран прецедент "Дать задание"
## Диаграммы компонентов
Были развернуты контейнеры:
#### Сервис управления проектной деятельностью студентов
![Сервис управления проектной деятельностью студентов](https://github.com/TypicalMedic/Software-Architecture/blob/main/Lab%20Work%20%E2%84%963/Docs/c4%20is.png?raw=true)
#### Сервис чат-бота Telegram
![Сервис чат-бота Telegram](https://github.com/TypicalMedic/Software-Architecture/blob/main/Lab%20Work%20%E2%84%963/Docs/c4%20tg.png?raw=true)
#### Сервис интеграций
![Сервис интеграций](https://github.com/TypicalMedic/Software-Architecture/blob/main/Lab%20Work%20%E2%84%963/Docs/c4%20intgr.png?raw=true)

## Диаграмма последовательностей
![Диаграмма последовательностей](https://github.com/TypicalMedic/Software-Architecture/blob/main/Lab%20Work%20%E2%84%963/Docs/seq.png?raw=true)

## Модель БД в виде диаграммы классов
![БД](https://github.com/TypicalMedic/Software-Architecture/blob/main/Lab%20Work%20%E2%84%963/Docs/db.png?raw=true)
## Использование принципов KISS, YAGNI, DRY и SOLID в коде
Для реализации системы используется в основном язык Go.
### KISS
Код легко читается и не использует какие-то незаурядные методы


    type  TaskManager  struct {}
        
    func (tm *TaskManager) AddTaskToCloud(task Task) Task {    
	    projectFolderId  := projectManager.GetProjectFolderId(task.ProjectId)    
	    task.TaskFolderId  = integration.api.cloud.AddTaskFolder(projectFolderId)    
	    task.TaskFileId  = integration.api.cloud.AddTaskFile(task.TaskFolderId)    
	    return task   
    }
    
    func (tm *TaskManager) CreateTask(task Task) {    
	    database.InsertTask(task)    
    }
    
    func (tm *TaskManager) EditTask(task Task, editedTask Task) {    
	    database.UpdateTask(task.Id, editedTask)    
    }
    
    func (tm *TaskManager) DeleteTask(task Task) {    
	    database.DeleteTask(task.Id)    
    }
    
    func (tm *TaskManager) GetTask(id int) Task {    
	    return database.GetTask(id)    
    }  
      
    //...

### YAGNI
Для выбранного прецедента нам требуются только функции создания задания и папки к нему, поэтому другие операции нам не нужны.
	     
    type  ITask  interface {    
	    CreateTask(task Task)    
	    AddTaskToCloud(task Task) Task    
		    //EditTask(task Task, editedTask Task)    
		    //DeleteTask(task Task)    
		    //GetTask(id int) Task
    }

### DRY 
Чтобы не повторять один и тот же код в разных местах он выносится в функции.
  

    type  TaskManager  struct {} 
         
    func (tm *TaskManager) AddTaskToCloud(task Task) Task {    
	    projectFolderId  := projectManager.GetProjectFolderId(task.ProjectId)    
	    task.TaskFolderId  = integration.api.cloud.AddTaskFolder(projectFolderId)    
	    task.TaskFileId  = integration.api.cloud.AddTaskFile(task.TaskFolderId)    
	    return task    
    }
    
    func (tm *TaskManager) CreateTask(task Task) {    
	    database.InsertTask(task)    
    }

### SRP
Структура (класс) TaskManager имеет единственную ответственность по работе с заданиями для студентов, источником изменений является только преподаватель, а от интеграции класс не зависит, т.к. её реализация скрыта за интерфейсом.
  

    type  TaskManager  struct {}
        
    func (tm *TaskManager) AddTaskToCloud(task Task) Task {    
	    projectFolderId  := projectManager.GetProjectFolderId(task.ProjectId)    
	    task.TaskFolderId  = integration.api.cloud.AddTaskFolder(projectFolderId)    
	    task.TaskFileId  = integration.api.cloud.AddTaskFile(task.TaskFolderId)    
	    return task   
    }
    
    func (tm *TaskManager) CreateTask(task Task) {    
	    database.InsertTask(task)    
    }
    
    func (tm *TaskManager) EditTask(task Task, editedTask Task) {    
	    database.UpdateTask(task.Id, editedTask)    
    }
    
    func (tm *TaskManager) DeleteTask(task Task) {    
	    database.DeleteTask(task.Id)    
    }
    
    func (tm *TaskManager) GetTask(id int) Task {    
	    return database.GetTask(id)    
    }  
      
    //...

### OCP
В языке Go не предусмотрено наследование, но даже если и было, в проекте нет классов, которые могут расширяться другими классами, поэтому показать данный метод не предоставляется возможным.

### Принцип подстановки Барбары Лисков
В языке Go не предусмотрено наследование, но даже если и было, в проекте нет классов, которые наследуются другими классами, поэтому показать данный метод не предоставляется возможным.

### ISP
Вместо одного большого интерфейса для сервиса управления проектной деятельностью студентов используются множество специализированных интерфейсов (в рамках прецедента интерфейсы работы с заданиями и проектами)
  

    type  ITask  interface {    
	    CreateTask(task Task)    
	    AddTaskToCloud(task Task) Task    
	    EditTask(task Task, editedTask Task)    
	    DeleteTask(task Task)    
	    GetTask(id int) Task    
	    //...
    }  

    type  IProject  interface {    
	    GetProjectFolderId(projectId int)    
	    AddProject(project Project)    
	    EditProject(project Project, editedProject Project)    
	    //...
    }

### DIP
Вся реализация скрыта за абстракциями, а именно интерфейсами, и мы в любой момент можем изменить реализацию без изменения кода более высокого уровня. Приложение зависит от методов, предоставленных интерфейсами, а не конкретной реализации.
  

    func  main() {    
	    // интерфейс    
	    var  taskMng ITask    
	    // реализация    
	    taskMng  = taskmng.InitTaskManager()    
	    testTask  := Task{}    
	    testTask  = taskMng.AddTaskToCloud(testTask)    
	    taskMng.CreateTask(testTask)    
    }      
	//...
    type  ITask  interface {    
	    CreateTask(task Task)    
	    AddTaskToCloud(task Task) Task    
	    EditTask(task Task, editedTask Task)    
	    DeleteTask(task Task)    
	    GetTask(id int) Task    
    }


