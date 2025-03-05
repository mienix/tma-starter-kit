# TMA Starter Kit

TMA Starter Kit is a developer-friendly template for building **Telegram Mini Apps**. It provides a ready-to-use foundation with a modern tech stack, allowing developers to focus on business logic rather than project setup.

## Why Use TMA Starter Kit?
- **Fast Development** – Pre-configured setup with essential tools.
- **Scalability** – Modular architecture for easy expansion.
- **Modern Tech Stack** – Uses Quasar (Vue.js), TypeScript, Go, and MongoDB.
- **Built-in DevOps** – CI/CD pipelines, Docker Compose, and deployment scripts included.

## Project Structure
```
TMA-Starter-Kit/
│── frontend/   # Quasar-based Vue.js frontend
│── backend/    # Go API backend
│── devops/     # DevOps tools, CI/CD, Docker Compose files
```

## Tech Stack
- **Frontend:** Quasar (Vue.js), TypeScript, Composition API
- **Backend:** Go
- **Database:** MongoDB
- **DevOps:** Docker, CI/CD, Deployment automation

## Getting Started

### Prerequisites
- Node.js & npm (for frontend)
- Go (for backend)
- Docker & Docker Compose (for containerized deployment)

### Installation
Clone the repository:
```sh
 git clone https://github.com/your-repo/TMA-Starter-Kit.git
 cd TMA-Starter-Kit
```

#### Frontend
```sh
cd frontend
npm install
quasar dev
```

#### Backend
```sh
cd backend
go run main.go
```

#### Run with Docker Compose
```sh
docker-compose up --build
```

## Contribution
Contributions are welcome! Feel free to submit a PR or open an issue.

## License
MIT License.

