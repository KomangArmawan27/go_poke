# Go RESTful API Boilerplate

A simple RESTful API built using Go (Gin Framework) with PostgreSQL, following clean and scalable project architecture, authentication, role-based access control (RBAC), dynamic filtering, and pagination.

## 📁 Project Structure

```
.
├── config/
│   └── db.go                # Database connection
├── controllers/handlers/   # Route handlers (e.g., GetUsers, CreateProduct)
├── middleware/             # Middleware (Auth, Role)
├── models/                 # GORM models
├── routes/                 # All route registrations
├── utils/                  # Helper utilities (pagination, response formatting)
├── main.go                 # Application entry point
```

## 🔐 Features

- JWT-based Authentication
- Role-based Access Control (Admin, Manager, User)
- Dynamic Filtering & Pagination
- Modular Clean Code Architecture

## 📦 Requirements

- Go 1.21+
- PostgreSQL
- Gin Framework
- GORM
- JWT library (e.g., github.com/golang-jwt/jwt/v5)

## ▶️ Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/go-rest-api-boilerplate.git
cd go-rest-api-boilerplate
```

### 2. Setup Environment Variables
Create a `.env` file in the root directory:
```
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
JWT_SECRET=your_jwt_secret
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the App
```bash
go run main.go
```

## 📂 API Endpoints

### Auth
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Users (Admin Only)
- `GET /api/v1/users`
- `GET /api/v1/user?id=1`
- `POST /api/v1/user/create`
- `PUT /api/v1/user/update`
- `DELETE /api/v1/user/delete`

### Products
- `GET /api/v1/products`
- `GET /api/v1/product?id=1`
- `POST /api/v1/product/create` _(Manager and Admin)_
- `PUT /api/v1/product/update` _(Manager and Admin)_
- `DELETE /api/v1/product/delete` _(Manager and Admin)_

## 📄 Example Query Parameters for Filtering & Pagination

```http
GET /api/v1/users?name=john&email=gmail.com&limit=5&page=2&sort=id desc
```

## 👮 Role-Based Access Middleware

The `RoleMiddleware` uses a rank map:
```go
var roleRank = map[string]int{
  "user": 1,
  "manager": 2,
  "admin": 3,
}
```
So a `manager` can access all `user` routes, and `admin` can access all.

## 📦 Pagination Helper Usage

```go
db, pagination := utils.ApplyPagination(c, db, &models.User{})
```

Returns structured pagination info in response.

## 🧠 Credits
Created with ❤️ by KomangArmawan

---

# Swagger initiation command
```bash
swag init -g cmd/main.go --output ./docs
```
