# 规格冻结表（foundation）

| 维度 | 冻结结论 |
|---|---|
| 业务边界 | 高校外语课程认证：课程窗口、模块、成果提交、AI 声明、证据版本、教师/复核评审、退回补交、归档。排除游戏、CRM、OA、报表等禁题材。 |
| 持久化 | SQLite 真实关系数据库；版本化 migration 001；users/courses/modules/submissions/evidences/reviews/audit_events/idempotency_records/outbox_jobs 九表及外键、唯一约束和索引。 |
| 事务 | submission 创建在事务内校验课程窗口/容量并写入；Transaction 用例支持提交与证据跨实体原子写入和回滚。 |
| 状态机 | draft→submitted→returned/approved→archived；非法转换拒绝，退回可重新提交。 |
| 并发 | submission version 条件更新、Serializable 事务、容量检查与唯一约束；race 测试命令固定。 |
| context | HTTP、service、repository、worker 全链路使用 context，取消和超时保留。 |
| worker | outbox worker 支持停止、重试策略、退避和 dead 状态。 |
| 错误传播 | apperr 稳定错误码、Unwrap 链、HTTP 状态映射和 request id。 |
| HTTP/身份 | cmd/server；login、可撤销/过期会话、logout；student/teacher/reviewer/admin 角色鉴权；healthz/readyz。 |
| Docker | Dockerfile 基于 golang:1.26，多阶段构建 cmd/server，已验证 linux/amd64 与 linux/arm64。 |
| 测试 | 领域、service、真实 SQLite migration/repository、HTTP、worker/auth、并发与恢复覆盖；测试物理行数≥1500。 |
| 规模 | foundation_profile=compact_10；目标题数=10；至少 20 个生产 Go 文件、10 个 package、2000 生产行。 |
| 出题容量 | 规划十个彼此独立运行时边界：状态转换、容量窗口、版本冲突、事务回滚、证据版本、评审退回、幂等键、context 取消、worker 重试、审计一致性；不预埋 Bug。 |
| 禁止项 | 不创建题目分支、私测、gold、答案、intake 数据或旧 seed/record/push 脚本。 |
