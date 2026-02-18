FROM golang:1.25.1 AS builder

WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod 
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
GOOS=linux go build -o meuprimeirocrudgo . 

FROM golang:1.25.1-alpine AS runner

RUN adduser -D quentinhauser

COPY --from=builder /app/meuprimeirocrudgo /app/meuprimeirocrudgo

RUN chown -R quentinhauser:quentinhauser /app

RUN chmod +x /app/meuprimeirocrudgo

EXPOSE 8080

USER quentinhauser

# CMD [ "/app/meuprimeirocrudgo" ]

CMD [ "sleep", "5151851" ]

