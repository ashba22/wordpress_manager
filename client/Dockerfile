# client/Dockerfile

# Stage 1: Build
FROM node:latest AS build
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Stage 2: Serve with `serve`
FROM node:latest
WORKDIR /app
RUN npm install -g serve
COPY --from=build /app/dist /app
EXPOSE 5000
CMD ["serve", "-s", ".", "-l", "5000"]
