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
7) AssignCost float // how much to deduct from popug account for task assigment
8) Reward float // how much to add to popug account for task completion

### API

1) Create task
2) Re-assign tasks
3) Complete task

## Accounting service

### Database structure

AccountLog:
1) UserID uuid
2) Reason enum // enum values are: task assigment, task completion, pay day
3) Sum float // can be positive or negative
3) When datetime

### API

1) Balance for given day for all users
2) Add log
3) Balance for specified user
4) Stats for all users or for specified user, depending on request

### Cron job

Will start at the end of each day, receive balance for all users.
If balance is positive, add log with reason "pay day" and positive balance sum

