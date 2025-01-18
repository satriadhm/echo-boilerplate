# Echo Boilerplate with PostgreSQL and Migrations

This repository provides a boilerplate for building a web application using the [Echo](https://echo.labstack.com/) framework in Go. It supports PostgreSQL as the database and integrates with `goose` for managing database migrations.

---

## Features

- Modular project structure with `cmd/`, `internal/`, and `pkg/` directories.
- PostgreSQL as the database.
- Middleware integration for logging and error recovery.
- `goose` integration for handling database migrations.
- Configuration via YAML.
- Example APIs for authentication and managing todos.

---

## Prerequisites

- Go 1.23 or later
- PostgreSQL 13 or later
- `goose` CLI

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/satriadhm/echo-boilerplate.git
   cd echo-boilerplate
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Install `goose` CLI**:
   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

4. **Setup PostgreSQL database**:
   - Create a PostgreSQL database.
   - Update `config.yaml.example` with your database credentials:
     ```yaml
     database:
       host: "127.0.0.1"
       port: 5432
       user: "postgres"
       password: "password"
       name: "your_database_name"
     ```

---

## Usage

### 1. Run Database Migrations

Ensure you have a `migrations/` directory with migration files.

- **Apply migrations**:
  ```bash
  make migrate-up
  ```

- **Rollback migrations**:
  ```bash
  make migrate-down
  ```

- **Create a new migration**:
  ```bash
  make create-migration name=create_table_name
  ```

### 2. Run the Application

Start the application:
```bash
go run cmd/main.go
```

---

## Project Structure

```
satriadhm-echo-boilerplate/
├── cmd/
│   └── main.go             # Application entry point
├── internal/
│   ├── auth/               # Authentication module
│   ├── todo/               # Todo module
│   └── middlewares/        # Custom middleware
├── pkg/
│   ├── config/             # Configuration loader
│   ├── logger/             # Logger setup
│   └── migrations/         # Migration runner
├── migrations/             # SQL migration files
├── Makefile                # Commands for migrations
├── config.yaml.example     # Configuration template
├── go.mod                  # Dependencies
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

---

## Endpoints

### Authentication
- **Login**: `POST /login`

### Todo Management
- **Create Todo**: `POST /todo`
- **Get Todo**: `GET /todo/:id`
- **Update Todo**: `PUT /todo/:id`
- **Delete Todo**: `DELETE /todo/:id`

---

## Environment Variables

- **`JWT_SECRET`**: Secret key for JWT token generation.

---

## Development Workflow

1. **Run migrations**:
   ```bash
   make migrate-up
   ```

2. **Start the server**:
   ```bash
   go run cmd/main.go
   ```

3. **Test API endpoints**:
   Use tools like [Postman](https://www.postman.com/) or [curl](https://curl.se/).

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
