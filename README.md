# AI Course Certification

面向高校外语专业的课程与能力认证后端，覆盖课程模块、学生成果提交、AI 使用声明、证据版本、教师/复核委员评审、申诉退回、审计与后台任务。

## 架构

Go HTTP 服务按 domain、service、repository、storage、middleware、worker 分层，SQLite 使用版本化 migration。提交和证据写入支持跨实体事务，状态流通过版本号进行乐观并发控制；会话支持过期与撤销。

## API

`POST /v1/login` 登录；`POST /v1/logout` 撤销会话；`GET /healthz` 存活；`GET /readyz` 就绪；认证后教师/管理员可创建课程，学生可创建并提交成果。

## 开发

```sh
go mod tidy
go test ./... -count=1
go test -race ./... -count=1
go vet ./...
go build ./...
go run ./cmd/server
```

环境变量见 `.env.example`。默认监听 `:8080`，数据库文件为 `course.db`。
