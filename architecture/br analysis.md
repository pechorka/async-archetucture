# Task-tracker requirements

1) Requirement: Таск-трекер должен быть отдельным дашбордом и доступен всем сотрудникам компании UberPopug Inc.
Analysis:
Actor - Any user
Command - Login to task tracker dashboard
Data - User login and password
Event - Task.Logined

2) Requirement: Новые таски может создавать кто угодно (администратор, начальник, разработчик, менеджер и любая другая роль). У задачи должны быть описание, статус (выполнена или нет) и попуг, на которого заассайнена задача
Analysis:
Actor - Any user
Command - Create task
Data - Task
Event - Task.Created

3) Requirement: Менеджеры или администраторы должны иметь кнопку «заассайнить задачи», которая возьмёт все открытые задачи и рандомно заассайнит каждую на любого из сотрудников. Не успел закрыть задачу до реассайна — сорян, делай следующую.
Analysis:
Actor - User with role admin or manager
Command - Assign tasks
Data - Nothing
Event - Task.Assigned

4) Requirement: Каждый сотрудник должен иметь возможность видеть в отдельном месте список заассайненных на него задач
Analysis:
Actor - Any user
Command - List tasks
Data - User ID
Event - Task.Listed (?)

5) Requirement: цены на задачу определяется единоразово, в момент появления в системе (можно с минимальной задержкой)
Analysis:
Actor - "Task.Created" event
Command - Appraise task
Data - task ID
Event - Task.Appraised

6) Requirement: Каждый сотрудник должен иметь возможность отметить задачу выполненной.
Analysis:
Actor - Any user
Command - Complete task
Data - task ID
Event - Task.Completed

# Accounting requirements

1) Requirement: Аккаунтинг должен быть в отдельном дашборде и доступен кому угодно
Analysis:
Actor - Any user
Command - Login to accounting dashboard
Data - User login and password
Event - Accounting.Logined

2) Requirement: У каждого из сотрудников должен быть свой счёт, который показывает, сколько за сегодня он получил денег. У счёта должен быть аудитлог того, за что были списаны или начислены деньги, с подробным описанием каждой из задач.
Analysis:
Actor - "Task.Assigned" event
Command - Add audit log to user account
Data - User ID, task cost, task description
Event - Accounting.LogEntryAdded

3) Requirement: деньги списываются сразу после ассайна на сотрудника, а начисляются после выполнения задачи.
Analysis:
Actor - "Task.Completed" event
Command - Add audit log to user account
Data - User ID, task cost, task description
Event - Accounting.LogEntryAdded