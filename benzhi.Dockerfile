FROM golang:1.26
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/course-cert ./cmd/server
EXPOSE 8080
ENTRYPOINT ["/bin/course-cert"]
