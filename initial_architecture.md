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
