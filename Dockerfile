FROM golang:1.22.4-alpine AS environ

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify


FROM environ AS builder

COPY . .
RUN go build -v -o /usr/local/bin/app


FROM environ AS runner

COPY . .

CMD ["go", "run", "."]


FROM scratch AS prod

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
