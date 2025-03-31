# Real-Time IoT Alerts System

A scalable, real-time IoT alerts management system designed to efficiently handle IoT device alerts and notifications using **gRPC**, **WebSockets**, **Kafka**, and **Gin**. The system provides real-time alerting, Kafka-based message queues, and seamless integration with a PostgreSQL database, all built in **Go**.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
  - [Docker](#docker)
  - [Kubernetes](#kubernetes)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Real-Time IoT Alerts System is designed to monitor and manage alerts for IoT devices in real time. The system supports:
- **gRPC** based communication for high-performance API interactions.
- **WebSockets** for real-time alert notifications.
- **Kafka** for message streaming and queueing.
- **Gin** for the HTTP API layer.
- Type-safe database interactions with **SQLC**.

## Features

- **Real-Time Alerts**: Provides immediate notifications on IoT device status changes via WebSocket and gRPC.
- **Scalable Architecture**: Designed with Kafka for message streaming and horizontal scaling.
- **gRPC API**: Efficient and type-safe gRPC API for interacting with the backend.
- **WebSocket Support**: Real-time alerts pushed to clients over WebSocket.
- **Database Integration**: PostgreSQL with SQLC for type-safe queries.
- **Docker/Kubernetes Support**: Ready for containerized deployments and orchestration.

## Technologies

- **Go** (Golang) for backend development.
- **gRPC** for high-performance, type-safe RPC.
- **Kafka** for message streaming and queueing.
- **Gin** for HTTP server and API.
- **WebSockets** for real-time communication.
- **SQLC** for type-safe SQL queries.
- **PostgreSQL** for relational database storage.
- **Docker** for containerization.
- **Kubernetes** for orchestration.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following tools installed:

- **Go** (v1.18 or higher) for backend development.
- **Docker** for containerization.
- **kubectl** (optional, for Kubernetes deployments).
- **Kafka** (for message streaming).
- **PostgreSQL** (for the relational database).

You can follow the instructions to install these tools from their respective documentation:

- [Go Installation](https://golang.org/doc/install)
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Kafka Setup](https://kafka.apache.org/quickstart)
- [PostgreSQL Setup](https://www.postgresql.org/download/)

### Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/real-time-iot-alerts.git
cd real-time-iot-alerts
