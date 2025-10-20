# User Service
User Service is part of a microservices system that is responsible for managing user data, authentication, and relationships between other services.

---

## Futures
- Auth with JWT
- User management

---

## Tech Stack
- **Programming Language:** Go
- **Database:** PostgreSQL
- **Logging:** Logrus
- **Validation:** Go Validator
- **API Documentation:** Swagger (via swaggo)
- **Containerization:** Docker
- **Deployment:** Kubernetes

---

## Setup
```bash
# clone repository
https://github.com/nurmanhadi/microservices-user-service.git

# go to directory
cd user-service

# install dependency
go mod tidy
```

---

## Environment Variable
Create file `.env`
```bash
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=user_service
DB_USERNAME=user_service
DB_PASSWORD=user_service
DB_MAX_POOL_CONNS=10
DB_MAX_IDLE_CONNS=5
DB_CONN_MAX_LIFETIME=300
```

---

## Usage

### Run Localy
```bash
go run main.go
```

---

## API Documentation
Access url Endpoint `/swagger/index.html`

---

## License
This project is licensed under the MIT License.

---

## Author
**Nurman Hadi**  
Backend Developer (Golang, Microservices)  
GitHub: [nurmanhadi](https://github.com/nurmanhadi)