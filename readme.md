# Account and Mutation Services

This repository contains two GoLang services: Account Service and Mutation Service. These services are designed to run in Docker containers and utilize PostgreSQL for data storage, RabbitMQ for event messaging, and Swagger for API documentation.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Setting Up Environment Variables](#setting-up-environment-variables)
  - [Running the Services](#running-the-services)
- [Docker Compose](#docker-compose)
- [Database Migration](#database-migration)
- [Swagger Documentation](#swagger-documentation)
- [RabbitMQ Setup](#rabbitmq-setup)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

Make sure you have the following installed on your system:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [GoLang](https://golang.org/)

## Getting Started

### Setting Up Environment Variables

Before running the services, you'll need to set up environment variables. Copy the provided `.env.example` file and rename it to `.env`. Modify the values as needed.

```bash
cp .env.example .env
cp account-service/.env.example account-service/.env
cp mutation-service/.env.example mutation-service/.env
```

### Running the Services
To start the services, use the following commands:

```
make run
```

### Stop the Services
To start the services, use the following commands:

```
make stop
```

## Swagger Documentation
Swagger documentation for both services can be accessed at the following URLs:

- Account Service: http://localhost:5100/swagger/index.html
- Mutation Service: http://localhost:5200/swagger/index.html

## RabbitMQ Setup
Make sure you have RabbitMQ running. The default settings in .env should work for a local setup. If you need to change the RabbitMQ connection details, update the environment variables accordingly on every services.




