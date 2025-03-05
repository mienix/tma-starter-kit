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
- LocalTunnel (for exposing the app to the internet)


### Installation
Clone the repository:
```sh
 git clone https://github.com/your-repo/tma-starter-kit.git
 cd tma-starter-kit
```

#### Run with Docker Compose
```sh
docker compose -f devops/docker-compose.dev.yml up -d
```
#### Exposing the Application to the Internet

Install LocalTunnel:
```sh
npm install -g localtunnel
```

Run the following command to expose your app:
```sh
lt --port 9000 --subdomain <subdomain>
```

Your application will now be accessible online at https://<subdomain>.loca.lt/.

# Setting Up a Telegram Web App

To integrate your application as a **Telegram Web App**, follow these steps:

## 1. Create a Telegram Bot
- Open Telegram and search for [@BotFather](https://t.me/BotFather).
- Start a chat and send the command:
  ```
  /newbot
  ```
- Follow the instructions to provide a name and username for your bot.
- After completion, BotFather will provide a **bot token**. Save this token for later use.

## 2. Register a Telegram Web App
- In BotFather, send:
  ```
  /newapp
  ```
- Provide a name for your Web App when prompted.
- Enter a short description of your application.
- Provide the **URL** where your Web App is hosted (e.g., `https://<subdomain>.loca.lt/`).
- After registration, you will receive confirmation that the Web App has been created.

Now your Telegram Web App is fully configured and ready to use! 🚀

## Contribution
Contributions are welcome! Feel free to submit a PR or open an issue.

## License
MIT License.

