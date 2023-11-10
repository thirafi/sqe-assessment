FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
COPY .env ./
RUN go build -o /docker-alta
EXPOSE  8080
CMD ["/docker-alta"]