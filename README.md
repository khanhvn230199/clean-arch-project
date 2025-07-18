# Clean Architecture Go Project

Dự án Go được xây dựng theo kiến trúc Clean Architecture với cấu trúc thư mục rõ ràng và tách biệt các layer.

## Cấu trúc thư mục

```
├── cmd/
│   └── server/
│       └── main.go                 # Entry point của ứng dụng
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   │   └── user.go            # Domain entities
│   │   ├── repository/
│   │   │   └── user_repository.go # Repository interfaces
│   │   └── service/
│   │       └── user_service.go    # Domain services
│   └── usecase/
│       └── user_usecase.go        # Application logic
├── infrastructure/
│   ├── repository/
│   │   └── user_repository_impl.go # Repository implementations
│   └── database/
│       └── connection.go          # Database connection
├── api/
│   ├── handler/
│   │   └── user_handler.go        # HTTP handlers
│   └── route/
│       └── routes.go              # Route configuration
├── config/
│   └── config.go                  # Configuration management
├── migrations/
│   └── 001_create_users_table.sql # Database migrations
├── .env                           # Environment variables
├── go.mod                         # Go module file
└── README.md                      # Tài liệu hướng dẫn
```

## Công nghệ sử dụng

- **Go 1.21+**: Ngôn ngữ lập trình
- **Gin**: Web framework cho HTTP routing
- **PostgreSQL**: Database
- **UUID**: Tạo unique identifiers
- **Godotenv**: Quản lý environment variables

## Cài đặt và chạy

### 1. Cài đặt dependencies

```bash
go mod download
```

### 2. Cấu hình database

Tạo database PostgreSQL:

```bash
# Tạo database
createdb your_database_name

# Chạy migration
psql -d your_database_name -f migrations/001_create_users_table.sql
```

### 3. Cấu hình environment

Tạo file `.env` hoặc set environment variables:

```bash
export DATABASE_URL="postgres://username:password@localhost/your_database_name?sslmode=disable"
export PORT=8080
export ENVIRONMENT=development
```

### 4. Chạy ứng dụng

```bash
go run cmd/server/main.go
```

Server sẽ chạy tại `http://localhost:8080`

## API Endpoints

### User Management

- `POST /api/v1/users` - Tạo user mới
- `GET /api/v1/users/:id` - Lấy user theo ID
- `PUT /api/v1/users/:id` - Cập nhật user
- `DELETE /api/v1/users/:id` - Xóa user
- `GET /api/v1/users` - Lấy tất cả users

### Health Check

- `GET /health` - Kiểm tra trạng thái server

## Ví dụ sử dụng API

### Tạo user mới

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "John Doe"
  }'
```

### Lấy user theo ID

```bash
curl http://localhost:8080/api/v1/users/{user-id}
```

### Cập nhật user

```bash
curl -X PUT http://localhost:8080/api/v1/users/{user-id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe"
  }'
```

## Kiến trúc Clean Architecture

### Domain Layer (`internal/domain/`)
- **Entities**: Chứa business objects chính
- **Repository Interfaces**: Định nghĩa contracts cho data access
- **Services**: Business logic thuần túy

### Use Case Layer (`internal/usecase/`)
- **Application Logic**: Orchestration và workflow
- **Business Rules**: Validation và processing logic

### Infrastructure Layer (`infrastructure/`)
- **Repository Implementations**: Concrete implementations
- **Database**: Connection và configuration
- **External Services**: Third-party integrations

### Presentation Layer (`api/`)
- **Handlers**: HTTP request/response handling
- **Routes**: API endpoint configuration
- **Middleware**: Cross-cutting concerns

## Lợi ích của kiến trúc này

1. **Tách biệt rõ ràng**: Mỗi layer có trách nhiệm riêng biệt
2. **Dễ test**: Có thể mock dependencies dễ dàng
3. **Maintainable**: Code dễ đọc và bảo trì
4. **Scalable**: Dễ dàng thêm features mới
5. **Independent**: Database và framework có thể thay đổi mà không ảnh hưởng business logic

## Development

### Thêm entity mới

1. Tạo entity trong `internal/domain/entity/`
2. Tạo repository interface trong `internal/domain/repository/`
3. Implement repository trong `infrastructure/repository/`
4. Tạo use case trong `internal/usecase/`
5. Tạo handler trong `api/handler/`
6. Cấu hình routes trong `api/route/`

### Testing

```bash
# Chạy tests
go test ./...

# Test với coverage
go test -cover ./...
```

### Build

```bash
# Build binary
go build -o bin/server cmd/server/main.go

# Cross compile
GOOS=linux GOARCH=amd64 go build -o bin/server-linux cmd/server/main.go
```

## Đóng góp

1. Fork repository
2. Tạo feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## License

MIT License


###  🔧 Commands hữu ích:

make run         # Chạy app
make dev         # Hot reload development
make test        # Chạy tests
make build       # Build binary
make docker-run  # Chạy với Docker