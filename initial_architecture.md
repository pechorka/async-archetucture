# Initial architecture

High level overview of service relationships here <link to miro missing for now>

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

1) Create task
2) Re-assign tasks
3) Complete task
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
1) Balance for given day for all users (for cronjob)
2) Add log (for cronjob and task service)
3) Balance for regular user for today (to render dashboard for regular user)
4) Balance for top manager for today (to render dashboard for regular user, for admin and for analytics)
5) Audit log for regular user with infinite scroll (to render dashboard for regular user)
6) Audit log for top manager with infinite scroll (to render dashboard for admin user)
Analytics api:
1) Count users with balance below zero
2) Most expensive task for given period: day, week, month

### Cron job

Will start at the end of each day, receive balance for all users.
If balance is positive, add log with reason "payday" and positive balance sum

## Auth service

Takes care of authorization and access level to services.

### Database structure

Role table:
1) Name string
2) Access map[string]bool // string - is a service name and bool is access level to a service. true - can read and write, false - can only read and null - no access
3) Comment string

Users table:
1) Name string
2) RoleID uuid
3) Password string

### API
