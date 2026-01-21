# SubMinder API

SubMinder is a high-performance RESTful backend service designed to track recurring subscriptions and calculate renewal dates automatically.

This project demonstrates **Production-Grade Go development** practices, including **Clean Architecture**, **Concurrency**, **Containerization**, and **Automated Documentation**.

## Key Features

* **RESTful API:** Built with Gin Framework for high performance.
* **Clean Architecture:** Strict separation of concerns (Handler, Service, Repository layers).
* **Concurrency:** A concurrent background worker that checks for expiring subscriptions daily without blocking the main thread.
* **Database Management:** Uses PostgreSQL with GORM for object-relational mapping and auto-migrations.
* **Auto Seeding:** Automatically populates the database with mock data for testing purposes upon startup.
* **Containerization:** Fully dockerized application with Docker Compose for consistent deployment.
* **API Documentation:** Interactive API documentation using Swagger (OpenAPI).
* **Input Validation:** Robust request validation to ensure data integrity using DTOs.

## Tech Stack

* **Language:** Go (Golang 1.25)
* **Framework:** Gin Web Framework
* **Database:** PostgreSQL 15
* **ORM:** GORM
* **Configuration:** Viper
* **Documentation:** Swaggo (Swagger)
* **Deployment:** Docker & Docker Compose

## Project Structure

The project follows the Layered Architecture pattern:

* **cmd/api:** Application entry point.
* **internal/domain:** Entity definitions and Data Transfer Objects (DTOs).
* **internal/repository:** Database access layer (GORM implementation).
* **internal/service:** Business logic layer.
* **internal/transport:** HTTP handlers and router configuration.
* **internal/worker:** Background jobs (Renewal Checker).
* **pkg/database:** Database connection and seeding logic.

## Getting Started

### Prerequisites

* Docker Desktop installed on your machine.

### Installation & Running

1.  Clone the repository:
    ```bash
    git clone https://github.com/IlhanKazan/subminder.git
    cd subminder
    ```

2.  Run with Docker Compose:
    This command builds the Go application and sets up the PostgreSQL database.
    ```bash
    docker-compose up --build
    ```

3.  Access the Application:
    * **API Documentation (Swagger):** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
    * **API Base URL:** http://localhost:8080
    * **Database:** localhost:5432

### Verifying the Background Worker

The project includes an automatic **Seeder** that inserts mock data. To see the concurrency in action:

1.  Start the app and wait for the logs.
2.  Look for the following log entry in your terminal:
    > `[INFO] Checking renewals...`
3.  Since the mock data includes a subscription expiring "tomorrow", the worker will trigger an alert without blocking API requests.

### Stopping the Application

* **Stop only:** Press `Ctrl + C` in the terminal.
* **Stop and Remove Containers:** Run the following command:
    ```bash
    docker-compose down
    ```

> **Note:** These commands work on Windows, macOS, and Linux.

### Cleaning Up

To remove the database volumes and start fresh (this will delete all data):
```bash
docker-compose down -v