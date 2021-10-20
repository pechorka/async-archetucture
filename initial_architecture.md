# Initial architecture

High level overview of service relationships here https://miro.com/app/board/o9J_lpuD3tE=/

## When something goes wrong
1) Database failed: return 500 error and try to reconnect
2) Message broker failed: 
    send missed messages after connection re-established based on rows in databases with
    technical flags like: isSendToMessageBroker=false
3) Can't return a response to the user, because network connection is down: don't know what to do

## Task service

CRUD for tasks

### Database structure

Tasks:
1) Name string
2) Description string
3) Status enum // enum values are: open, done
4) Deadline datetime
5) CreatedBy uuid
6) AssignedTo uuid
7) AssignCost int // how much to deduct from popug account for task assignment
8) Reward int // how much to add to popug account for task completion

### API

1) Create task. params:
    a) name
    b) description
    c) assignedTo
    d) deadline
2) Re-assign tasks
3) Complete task. params:
    a) taskID
4) List tasks

## Accounting service

### Database structure

AccountLog:
1) UserID uuid
2) Reason enum // enum values are: task assignment, task completion, payday
3) Sum int // can be positive or negative
3) When datetime

### API

Accounting api:
1) Pay to all users (for cronjob)
2) Add log (for task service). params - all fields in AccountLog
3) Balance for regular user for today (to render dashboard for regular user)
4) Balance for top manager for today (to render dashboard for regular user, for admin and for analytics)
5) Audit log for regular user with infinite scroll (to render dashboard for regular user)
6) Audit log for top manager with infinite scroll (to render dashboard for admin user)
Analytics api:
1) Count users with balance below zero
2) Most expensive task for given period. params:
    a) period: day, week, month

### Cron job

Will start at the end of each day and call "Pay to all users" api handler

## Auth service

Takes care of authorization and access level to services.

### Database structure
Users table:
1) Name string
2) RoleName enum // enum values are: admin, manager, developer, auditor
3) Password string

### API
1) Authorization. Returns access and refresh jwt tokens. params:
    a) name
    b) password
2) Refresh tokens. params:
    a) refresh token
3) User names. params:
    a) user ids
