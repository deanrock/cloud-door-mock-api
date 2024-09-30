FROM golang:1-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

FROM alpine AS runtime

COPY --from=build /app/cloud-door-mock-api /cloud-door-mock-api

CMD ["/cloud-door-mock-api"]
