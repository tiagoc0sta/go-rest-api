# Go: Developing a REST API

## Acknowledgements

- This is a project of creating a a REST API, connect it to a PostgreSQL database using a Docker image.

## Authors

- [Alura](https://cursos.alura.com.br/formacao-go)

## Requirements

- Golang
- PostgreSQL
- React.JS
- Docker
- Postman

## Features

- Access frontend
- Create / update / edit / delete an item from a db using the frontend

## Demo

- Postgre database created via Docker
  ![image](https://github.com/tiagoc0sta/go-rest-api/assets/63982700/3f729e5f-68e8-451e-9f9b-c1125cbeb9c3)

- Frontend displaying the Database information
  ![image](https://github.com/tiagoc0sta/go-rest-api/assets/63982700/dfbeea15-a9f3-4703-89de-79f9c2f0b0ec)

- Perform CRUD tests via Postman
  ![image](https://github.com/tiagoc0sta/go-rest-api/assets/63982700/e4d26c8e-56c2-4f4c-8742-7c3becc06e78)

## Installation

1. Open Docker Desktop

2. Run Docker

```bash
  docker-compose up
```

3. Run API

```bash
  run main.go
```

3. Run Frontend

```bash
  npm start
```

4. Open Postman and perform CRUD operations on API. See the reponse on the frontend

### Pages:

- http://localhost:8000/ = homepage

- http://localhost:8000/api/personalidades - ver personalidades

- http://localhost:54321 = login to pgAdmin - PostgreSQL

# Lessons Learned

- Understood the importance of structuring the API endpoints and handling HTTP requests and responses efficiently.

Integrating with a Database

- Learned how to integrate a Go API with a database using Docker.
- Gained knowledge of using Docker for containerization and database management.
  Using GORM

- Learned how to use GORM, Go's popular ORM (Object-Relational Mapping) library, for database operations.
- Explored features such as model definitions, querying, and relationships management provided by GORM.

Creating Middleware

- Learned how to create middleware in Go to handle cross-cutting concerns such as logging, authentication, and request validation.
- Understood the benefits of using middleware to avoid code duplication and maintain clean API logic.

Integrating with a React Frontend

- Learned how to integrate a Go API with a React frontend for a full-stack application.
- Explored methods for communication between the frontend and backend, such as RESTful APIs and WebSocket.
