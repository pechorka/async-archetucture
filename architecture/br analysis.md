# Task-tracker requirements

1) Requirement: Таск-трекер должен быть отдельным дашбордом и доступен всем сотрудникам компании UberPopug Inc.
Analysis:
Actor - Any user
Command - Login to task tracker dashboard
Data - User auth info
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
Data - Tasks with status open
Event - Task.Assigned

4) Requirement: Каждый сотрудник должен иметь возможность видеть в отдельном месте список заассайненных на него задач
Analysis:
Actor - Any user
Command - List tasks
Data - User
Event - Task.Listed 

5) Requirement: цены на задачу определяется единоразово, в момент появления в системе (можно с минимальной задержкой)
Analysis:
Actor - "Task.Created" event
Command - Appraise task
Data - Task
Event - Task.Appraised

6) Requirement: Каждый сотрудник должен иметь возможность отметить задачу выполненной.
Analysis:
Actor - Any user
Command - Complete task
Data - Task
Event - Task.Completed

# Accounting requirements

1) Requirement: Аккаунтинг должен быть в отдельном дашборде и доступен кому угодно
Analysis:
Actor - Any user
Command - Login to accounting dashboard
Data - UserAuthInfo
Event - Accounting.Logined

2) Requirement: У каждого из сотрудников должен быть свой счёт, который показывает, сколько за сегодня он получил денег. У счёта должен быть аудитлог того, за что были списаны или начислены деньги, с подробным описанием каждой из задач.
Деньги списываются сразу после ассайна на сотрудника
Analysis:
Actor - "Task.Assigned" event
Command - Add audit log to user account
Data - User ID, Amount, AuditLog Type, MetaInfo (task description)
Event - Accounting.LogEntryAdded

3) Requirement: деньги начисляются после выполнения задачи.
Analysis:
Actor - "Task.Completed" event
Command - Add audit log to user account
Data - User ID, Amount, AuditLog Type, MetaInfo (task description)
Event - Accounting.LogEntryAdded

4) Requirement: считать текущий баланс
Analysis:
Actor - "Accounting.LogEntryAdded" event
Command - Update earned by user amount
Data - UserID, Amount
Event - Accounting.UserEarningsUpdated

5) Requirement: выплата баланса на счет пользователя
Analysis:
Actor - Cronjob
Command - PayDay
Data - [UserID, TotalUserBalance] - array
Event - Accounting.PayDay

6) Requirement: отправлять на почту сумму выплаты.
Analysis:
Actor - "Accounting.PayDay" event
Command - Notify
Data - Email, Amount
Event - Emailer.PayNotification

7) Requirement: После выплаты баланса (в конце дня) он должен обнуляться, и в аудитлоге всех операций аккаунтинга должно быть отображено, что была выплачена сумма.
Analysis:
Actor - "Accounting.PayDay" event
Command - Add audit log to user account
Data - UserID, Amount
Event - Accounting.LogEntryAdded

8) Requirement: Дешборд должен выводить количество заработанных топ-менеджментом за сегодня денег.
Analysis:
Actor - "Accounting.LogEntryAdded" event
Command - Recalculate top earnings
Data - Amount, AuditLogType
Event - Accounting.TopManagerEarnings

9) Requirement: Дашборд должен выводить информацию по дням, а не за весь период сразу.
Analysis:
Actor - Any user
Command - Get audit log
Data - UserID
Event - Accounting.AuditLogRequested

# Analytic requirements

1) Requirement: Аналитика — это отдельный дашборд, доступный только админам.
Analysis:
Actor - Admin user
Command - Login to analytic dashboard
Data - UserAuthInfo
Event - Analytic.Logined

2) Requirement: Нужно указывать, сколько заработал топ-менеджмент за сегодня
Analysis:
Actor - "Accounting.TopManagerEarnings" event
Command - Get top manager earnings
Data - Amount
Event - Analytic.TopManagerEarningsUpdated

3) Requirement: Нужно указывать, сколько попугов ушло в минус.
Analysis:
Actor - "Accounting.LogEntryAdded" event
Command - UpdatePopugEarnings
Data - UserID, Amount
Event - Analytic.PopugEarningsUpdated

4) Requirement: Нужно показывать самую дорогую задачу за день, неделю или месяц.
Analysis:
Actor - "Task.Appraised" event
Command - UpdatePopugEarnings
Data - Amount
Event - Analytic.TaskCostListUpdated