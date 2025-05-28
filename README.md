# Go RESTful API Boilerplate

A RESTful API built using Go (Gin Framework) with PostgreSQL, designed for selecting and managing your favorite Pokémon. The project follows clean and scalable architecture with authentication, role-based access control (RBAC), dynamic filtering, and pagination.

## 📁 Project Structure

```
.
├── cmd/                   # Database and environment configuration
│   └── main.go            # Application entry point
├── config/                # Database and environment configuration
├── internal/
│   ├── auth               # OAuth
│   ├── handlers/          # Route handlers (e.g., GetUsers, CreatePokemon)
│   ├── middleware/        # Middleware (Auth, Role, CORS, Logger)
|   ├── models/            # GORM models
|   ├── utils/             # Helper utilities (pagination, response formatting)
│   └── routes/            # All route registrations
├── docs/                  # Swagger docs
```

## 🔐 Features

- JWT-based Authentication
- Role-Based Access Control (Admin, Manager, User)
- CORS & Logging Middleware
- Swagger API Documentation
- Modular Clean Code Architecture
- Dynamic Filtering & Pagination

## 📦 Requirements

- Go 1.24+
- PostgreSQL
- Gin Framework
- GORM
- JWT library (`github.com/golang-jwt/jwt/v5`)
- swag CLI (`github.com/swaggo/swag/cmd/swag`)

## ▶️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-rest-api-boilerplate.git
cd go-rest-api-boilerplate
```

### 2. Setup Environment Variables

Create a `.env` file in the root directory:

```env
PORT=your_main_port
PGHOST=your_db_host
PGUSER=your_db_user
PGPASSWORD=your_db_password
PGDATABASE=your_db_name
PGPORT=your_db_port
JWT_SECRET=your_jwt_secret
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the App

```bash
go run cmd/main.go
```

## 📂 API Endpoints

### Auth (Public)
- `POST /api/v1/login`
- `POST /api/v1/register`

### Users (Admin Only)
- `GET /api/v1/users`
- `GET /api/v1/user?id=1`
- `POST /api/v1/user/create`
- `PUT /api/v1/user/update`
- `DELETE /api/v1/user/delete`

### Pokémons (Authenticated Users)
- `GET /api/v1/pokemons`
- `GET /api/v1/pokemon?id=1`
- `POST /api/v1/pokemon/create` _(Authenticated)_
- `PUT /api/v1/pokemon/update` _(Authenticated)_
- `DELETE /api/v1/pokemon/delete` _(Authenticated)_

## 📄 Filtering & Pagination Example

```http
GET /api/v1/users?name=ash&type=grass&limit=5&page=1&sort=id desc
```

## 👮 Role-Based Access Middleware

The `RoleMiddleware` uses a rank-based map:

```go
var roleRank = map[string]int{
  "user": 1,
  "manager": 2,
  "admin": 3,
}
```

- `manager` can access all `user` routes  
- `admin` can access everything

## 📦 Pagination Helper Usage

```go
db, pagination := utils.ApplyPagination(c, db, &models.User{})
```

Returns structured pagination metadata alongside the result.

## 📚 API Documentation (Swagger)

Swagger docs are available at:

```
GET /swagger/index.html
```

To generate Swagger docs:

```bash
swag init -g main.go --output ./docs
```

## 🧠 Credits

Created with ❤️ by KomangArmawan
