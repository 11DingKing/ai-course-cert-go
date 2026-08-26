FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/course-cert ./cmd/server
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /out/course-cert /app/course-cert
EXPOSE 8080
ENTRYPOINT ["/app/course-cert"]
