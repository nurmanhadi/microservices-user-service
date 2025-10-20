FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/docs ./docs
ENV GIN_MODE=release
USER nonroot:nonroot
EXPOSE 4000
CMD [ "./app" ]