# **UserGate**

UserGate is a RESTful API service that provides basic user management capabilities, including creating and retrieving user information. The service is built using the [**Echo**](https://github.com/labstack/echo) framework in Go, MySQL for storage, and [**goose**](https://github.com/pressly/goose) for database migrations. This project includes a CI/CD pipeline integrated with GitHub Actions and the GitHub Container Registry (GHCR) for streamlined testing, building, and deployment.

---

## **Features**
- **User Management**: Create and retrieve user information via RESTful APIs.
- **Database Migrations**: Manage schema changes with `goose`.
- **Echo Framework**: Provides a robust and efficient web framework for API development using `Echo`.
- **Dockerized Deployment**: Build and run the application in a Dockerized environment.
- **CI/CD Integration**: Automated pipeline for testing, building, and deploying Docker images.

---

## **Table of Contents**
1. [Prerequisites](#prerequisites)
2. [Installation](#installation)
3. [Usage](#usage)
4. [API Endpoints](#api-endpoints)
5. [Database Migrations](#database-migrations)
6. [Development](#development)
7. [Testing](#testing)
8. [Deployment](#deployment)
9. [Contributing](#contributing)
10. [License](#license)

---

## **Prerequisites**
- Go `1.21` or later
- Docker and Docker Compose
- MySQL `8.0` or later
- GitHub Container Registry (GHCR) setup (optional for deployment)

---

## **Installation**

1. Clone the repository:
   ```bash
   git clone https://github.com/AshkanShakiba/UserGate.git
   cd UserGate
   ```

2. Set up environment variables in a `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=password
   DB_NAME=usergate
   ```

3. Build the application:
   ```bash
   make build
   ```

4. Run the Dockerized environment:
   ```bash
   make docker-run
   ```

---

## **Usage**

The service runs on `http://localhost:8080`. You can use tools like `curl` or Postman to interact with the API.

Example:
```bash
# Create a user
curl --location 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data '{
    "Name": "Anakin Skywalker"
}'

# Get a user
curl --location 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data '{
    "ID": 1
}'
```

---

## **API Endpoints**

### **POST /user**
- **Description**: Create a new user.
- **Request Body**:
  ```json
  {
    "Name": "Anakin Skywalker"
  }
  ```
- **Response**:
  ```json
  ```

### **GET /user**
- **Description**: Retrieve a user by ID.
- **Request Body**:
  ```json
  {
    "ID": 1
  }
  ```
- **Response**:
  ```json
  {
    "ID": 1,
    "Name": "Anakin Skywalker"
  }
  ```

---

## **Database Migrations**

This project uses `goose` for managing database schema. Run migrations with the following commands:

```bash
# Run migrations
make migrate-up

# Rollback migrations
make migrate-down
```

---

## **Development**

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the application locally:
   ```bash
   go run cmd/webserver/main.go
   ```

---

## **Testing**

Run unit tests:
```bash
make test
```

The project includes tests for key functionalities using `sqlmock` and `testify`.

---

## **Deployment**

This project includes a GitHub Actions CI/CD pipeline for automated testing, building, and deploying Docker images to GHCR.

1. **CI/CD Workflow**:
    - Push changes to the `main` branch to trigger the pipeline.
    - Monitor the pipeline in the **Actions** tab on GitHub.

2. **Docker Image**:
    - Images are pushed to GHCR:
      ```
      ghcr.io/<your-username-or-org>/usergate:latest
      ```

3. **Run the Docker Image**:
   ```bash
   docker pull ghcr.io/<your-username-or-org>/usergate:latest
   docker run -p 8080:8080 ghcr.io/<your-username-or-org>/usergate:latest
   ```

---

## **Contributing**

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit changes:
   ```bash
   git commit -m "Add a feature"
   ```
4. Push and create a pull request.

---

## **License**

This project is licensed under the [MIT License](LICENSE).

---

Let me know if you'd like further refinements! 🚀