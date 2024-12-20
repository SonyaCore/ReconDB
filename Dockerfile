FROM golang:latest

WORKDIR /app

# Create a non-root user
RUN useradd -u 10001 user \
    && mkdir /home/user \
    && chown -R user:user /home/user \
    && chown -R user:user /app \
    && chmod 755 /home/user

USER user

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ReconDB -ldflags "-s -w -buildid=" -buildvcs=false .

EXPOSE 8080

CMD ["./ReconDB"]
