# SIGCPC Backend Documentation + To Do

## Run locally

use Go 1.26.5 or newer, as specified in `go.mod`.

from the repo root:

```sh
cd backend
go run .
```

the server listens at `http://localhost:8080`. no database setup or environment variables are currently required!

## backend structure
```text
backend/
├── main.go
├── go.mod
└── internal/
    ├── db/
    ├── handlers/
    └── middleware/
```

## what each folder does

- db
    - handles connection to db and models for data
    - currently holds dummy data
- handlers
    - api endpoints for each data type
    - eg: members.go has the GET and POST end point for members table only
- middleware
    - contains reusable request handling logic, currently rate limiting
- main.go
    - registers routes, wraps the router with middleware, and starts the server.

## to-do
- [x] create project with basic endpoints and dummy data
- [ ] create docs(possibly with swagger)
- [ ] create connection to db
- [ ] add more endpoints for more data
- [ ] add auth
- [ ] make deployment settings configurable
- [ ] add automated tests in main_test.go and other *_test.go files, runnable with go test ./...

Follow the repository's [contribution guidelines](../CONTRIBUTING.md).
Mark completed tasks before merging!