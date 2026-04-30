# BionicPro SSO: Microservices Auth & RBAC PoC

A professional reference implementation of a Single Sign-On (SSO) system featuring **Keycloak** as the Identity Provider, a **Golang** backend, and a **React** frontend. This project demonstrates modern authentication patterns and secure communication in a containerized environment.

## 🚀 Project Overview

The goal of this PoC is to showcase a robust **OpenID Connect (OIDC)** integration. It focuses on:
- Centralized authentication via Keycloak.
- **Role-Based Access Control (RBAC)** enforced at both Frontend and Backend levels.
- Secure JWT validation and claim verification.
- Automated infrastructure setup using Docker Compose.

## 🛠 Tech Stack

- **Backend:** Go 1.23 (Chi Router, Gocloak)
- **Frontend:** React, TypeScript, Tailwind CSS
- **IAM:** Keycloak 25.0
- **Infrastructure:** Docker, Docker Compose, Nginx

## ✨ Key Features

- **Full OIDC Flow:** Seamless login/logout experience using `keycloak-js`.
- **JWT Middleware:** Custom Golang middleware for validating tokens and enforcing realm-level permissions.
- **Pre-configured Environment:** Includes a ready-to-use Keycloak realm export (`realm-export.json`) for instant testing.
- **Stateless Security:** Backend validates signatures without frequent round-trips to the IAM server.

## 📂 Architecture & Structure

```txt
├── backend/            # Go microservice with OIDC validation logic
├── frontend/           # React SPA with Keycloak integration
├── keycloak/           # Realm configurations and custom themes
└── docker-compose.yaml # Orchestration for all services
```

## 🚦 Quick Start

Ensure you have **Docker** and **Docker Compose** installed.

1. **Clone the repository:**
   ```bash
   git clone https://github.com/nickzhog/bionicpro-sso-keycloak.git
   cd bionicpro-sso-keycloak
   ```

2. **Launch the stack:**
   ```bash
   docker-compose up --build
   ```

3. **Access the services:**
   - **Frontend:** [http://localhost:3000](http://localhost:3000)
   - **Keycloak Admin:** [http://localhost:8080](http://localhost:8080) (Admin/Admin)

## 🔐 Implementation Details

- **Token Validation:** The backend utilizes public keys from Keycloak's JWKS endpoint to verify token integrity.
- **Service Discovery:** Uses Docker's internal DNS for secure service-to-service communication.
- **Clean Architecture:** Separation of concerns between authentication logic and business handlers.

