# AutoDaXueXi
自动化实现：江苏共青团青年大学习+班级小管家截图提交
- 快捷配置
- 适配班级小管家数据结构
- 可配置定时任务
------------
### 使用说明
完成config.json配置即可开始使用
```json
{
  "Cookie" : "",
  "Imprint" : "",
  "identity": ""
}
```
Cookie字段为江苏共青团凭证

Imprint字段为班级小管家凭证

identity字段功能开发中，无需配置
### 编译命令
```go
go build main.go
```
