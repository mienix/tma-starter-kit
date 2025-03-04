FROM node:18-bullseye

WORKDIR /app

#COPY package.json package-lock.json ./
COPY . .
RUN npm install
RUN npm install -g @quasar/cli


EXPOSE 9000

CMD ["npx", "quasar", "dev"]
