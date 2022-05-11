FROM node:16-alpine AS dev
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn
COPY . .
EXPOSE 3000

FROM develop AS builder
RUN yarn generate --fail-on-error
