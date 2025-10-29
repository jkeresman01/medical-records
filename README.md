<div align="center">

  <h1>medical-records</h1>
  <h6>  Medical Records Management System </h4>

[![Go](https://img.shields.io/badge/Go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
![Status](https://img.shields.io/badge/status-work--in--progress-yellow?style=for-the-badge)

A web application for managing medical records.
Main purpose is to demonstrate moving from gorm to apolon.

## Objectives ##

Respect general principles such as:

SOC - separation of concerns 
SRP - single responsibility principle, 
loose coupling, high cohesion,
prefer composition over inheritance
DRY - don’t repeat yourself...

`Repository Factory` design pattern is used.

## Project Structure

```
.
├── config/              # Configuration management
├── db/                  # Database connection
├── handlers/            # HTTP request handlers
├── models/              # Data models
├── repository/          # Repos
│   └── factory/         # Repository factory
├── routes/              # Route definitions
├── docker/              # Docker componse
├── viewmodels/          # View models
├── static/              # Static nonsense
│   ├── css/
│   └── templates/       # HTML templates
│       ├── patients/
│       ├── medications/
│       ├── prescriptions/
│       ├── exam_types/
│       └── exams/
└── main.go              # Application entry point
```

## Technologies

- [Fiber](https://gofiber.io/) - Cool Web framework
- [GORM](https://gorm.io/) - Cool ORM library
- [htmx](https://htmx.org/) - So random
- [PostgreSQL](https://www.postgresql.org/) - basedata
