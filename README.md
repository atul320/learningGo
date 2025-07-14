# Go Projects Collection

This repository contains two Go projects:
1. URL Shortener API (REST API with MongoDB)
2. Banking System (CLI application)

## Project 1: URL Shortener API

A RESTful API for creating and managing short URLs with user authentication.

### Features
- User registration and login (JWT authentication)
- Create short URLs that expire after 15 days
- Track click counts
- View/manage your URLs
- MongoDB backend

### Installation
```bash
cd url-shortener
go mod download
```

### Configuration
1. Copy  `.env`
2. Set your MongoDB URI and JWT secret

### Running
```bash
go run main.go
```

### API Endpoints
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login and get JWT token
- `POST /url/create` - Create short URL (authenticated)
- `GET /:shortCode` - Redirect to original URL

---

## Project 2: Banking System CLI

A command-line banking application with basic account operations.

### Features
- Create bank accounts
- Deposit/withdraw funds
- View account balances
- Simple console interface

### Installation
```bash
cd banking-app
go mod download
```

### Running
```bash
go run main.go
```

### Menu Options
```
1. Create Account
2. Deposit 
3. Withdraw
4. Show Accounts
5. Exit
```

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/atul320/learningGo.git
```

2. Navigate to either project folder and follow its specific instructions.

