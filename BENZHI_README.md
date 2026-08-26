# AI Course Certification

这是一个面向高校外语专业的课程与能力认证服务，提供课程窗口、成果提交、AI 使用声明、证据版本、教师与复核委员评审、退回补交、归档、审计及后台任务能力。

## 构建与运行

```sh
go build ./cmd/server
ADDR=:8080 DB_PATH=course.db go run ./cmd/server
```

## 测试

```sh
go test ./... -count=1
go test -race ./... -count=1
go vet ./...
```

## Docker

```sh
./build_benzhi_docker.sh
docker run --rm -p 8080:8080 ai-course-cert-benzhi
```

存活检查地址为 `/healthz`，就绪检查地址为 `/readyz`。环境变量见 `.env.example`。
