# Citassauro 🦖

A simple and robust API for creating and managing quotes, written in Golang.

## 📦 About

Citassaurus is an open-source project designed to deliver a lightweight and efficient API for quote management — perfect for testing, learning, or as a base for larger applications.  
Built with educational and portfolio purposes in mind.

## 🚀 Technologies

- [Go (Golang)](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) — HTTP framework
- [GORM](https://gorm.io/) — ORM for database
- SQLite (default, but easily swappable)
- Docker (optional for deployment)

## 🔧 Features

- 🔹 Create new quotes
- 🔹 List all quotes
- 🔹 Get quote by ID
- 🔹 Delete quotes
- 🔹 (Coming soon) Update quotes

## ⚙️ How to run

```bash
# Clone the repository
git clone https://github.com/walacysilvam/citassauro.git
cd citassauro

# Install dependencies
go mod tidy

# Run the application
go run main.go