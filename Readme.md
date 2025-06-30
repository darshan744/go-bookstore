# 📚 Go Bookstore API

A RESTful Bookstore API built with **Go**, using the **Gin** web framework and **GORM** ORM to interact with a **PostgreSQL** database. This project is structured for scalability, readability, and ease of use.

---

## 🚀 Features

- CRUD operations for managing books
- PostgreSQL database connection using GORM
- Environment-based configuration (`.env`)
- Clean, modular structure (controllers, models, routes, configs)
- Input validation with Gin
- Auto-migration for database schema
- Docker-compatible for easy setup

---

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/your-username/go-bookstore.git
cd go-bookstore
```

### 2. Create a `.env` File

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASS=your_password
DB_NAME=go_bookstore
PORT=3000
```

### 3.Install Dependencies

```
go mod tidy
```

### 4. Run the Server

```
go run main.go
```

### 5. Tech stack
- Go (Golang)
- Gin (Web Framework)
- GORM (ORM)
- PostgreSQL

