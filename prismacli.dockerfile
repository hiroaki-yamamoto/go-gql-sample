FROM node:current-alpine
RUN npm i -g prisma
RUN mkdir /app
WORKDIR /app
