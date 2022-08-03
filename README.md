###################
Little Things
###################

This is for shorting url.
It's running on port 80

*******************
Requirement
*******************

Go v.1.11+
Redis Server

**************************
How to use ?
**************************

go run main.go

*******************
Structure
*******************
```
├── handler
│   └── handlers.go
├── shortener
│   ├── shortener_generate_test.go
│   └── shortener_generate.go
├── store
│   ├── store_service_test.go
│   └── store_service.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```