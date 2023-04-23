FROM golang:latest

WORKDIR /app

# Create a non-root user
#RUN useradd -u 10001 user
#USER user


COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ReconDB .

EXPOSE 8080

CMD ["./ReconDB"]