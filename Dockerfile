FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

ARG SERVICE_PATH

RUN go build -o /main ./services/${SERVICE_PATH}/cmd/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder /main /main

USER nonroot:nonroot

ENTRYPOINT ["/main"]
