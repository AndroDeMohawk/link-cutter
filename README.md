# Link Cutter

A fast, lightweight, and production-ready **URL Shortener service** built with **Go (Golang)**. The project is fully containerized using **Docker** and includes automatic database migrations.

## 🚀 Features
* **URL Shortening**: Convert long, bulky URLs into short, manageable links.
* **Redirection**: Fast HTTP redirection from short aliases to the original URLs.
* **Database Migrations**: Automated database schema management.
* **Dockerized**: Easy setup and deployment via Docker and Docker Compose.
* **Clean Architecture**: Structured following Go best practices (`cmd/`, `internal/`, `pkg/`).

## 📁 Project Structure
* `cmd/` — Entry points for the application applications (e.g., main API server).
* `configs/` — Configuration files and environment setups.
* `internal/` — Private application and business logic (API handlers, repositories, services).
* `migrations/` — SQL migration files for setting up the database schema.
* `pkg/` — Publicly importable utility packages and helper functions.

## 🛠️ Prerequisites
Before running the project, make sure you have the following installed:
* [Go](https://go.dev) (1.21+ recommended)
* [Docker](https://docker.com) & [Docker Compose](https://docker.com)

## ⚡ Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com
cd link-cutter
```

### 2. Configuration
Create a configuration file or environment variables based on the templates provided in the `configs/` directory.

### 3. Run with Docker Compose (Recommended)
To spin up the entire stack (including the application and database), simply run:
```bash
docker-compose up --build
```

### 4. Run Locally
If you prefer to run the Go application directly on your machine:
```bash
# Install dependencies
go mod download

# Run the application
go run cmd/main.go
```

## 🔌 API Endpoints

All endpoints (except redirection) require authentication via middleware.

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :---: |
| `POST` | `/link` | Create a new short link | Yes |
| `PATCH` | `/link/{id}` | Update an existing link by ID | Yes |
| `DELETE` | `/link/{id}` | Delete a link by ID | Yes |
| `GET` | `/link` | Get all links for the authenticated user | Yes |
| `GET` | `/{hash}` | Redirect to the original long URL | No |

## 📝 License
This project is licensed under the MIT License — see the LICENSE file for details.
