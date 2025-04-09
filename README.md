# Medpoint

A backend app for medpoint system.
Medpoint system is a platform that allows users to make various types of online medical reservations,
including doctor consultations, laboratory tests, medical procedures, and vaccinations.
This system facilitates users with a more efficient reservation process.

### Technology used:
* Language: Go
* Framework: Raiden
* Backend-As-A-Service: Supabase (PostgreSQL, Auth, Storage)

### Project Structure:
```
├── configs                     # Configuration file
├── cmd
│   └── project-name
│       └── project_name.go     # Main project
│   └── apply/main.go
│   └── import/main.go
├── internal                    # To save assets file (font, images, etc)
│   ├── bootstrap
│   │   ├── route.go
│   │   ├── rpc.go
│   │   ├── roles.go
│   │   ├── models.go
│   │   └── storages.go
│   ├── controllers
│   │   ├── auth.go             # Example controller
│   │   └── ...
│   ├── roles                   # ACL/RLS definition
│   │   ├── members.go
│   │   └── ...
│   ├── models                  # All model will auto-sync
│   │   ├── users.go
│   │   └── ...
│   └── storages
│       └── avatar.go
├── build
│   └── state                   # Bytecode containing raiden state
├── go.sum
└── go.mod
```

### How to Run
1. Make sure go and raiden is installed in your machine
2. Run `go mod tidy`
3. Run `raiden configure`
4. Run `raiden run`
