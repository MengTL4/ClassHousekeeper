# AutoDaXueXi
自动化实现：江苏共青团青年大学习+班级小管家截图提交
由于江苏共青团cookie过期时间过短，项目停止维护
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
![](https://github.com/MengTL4/AutoDaXueXi/blob/main/image/1.png)
Imprint字段为班级小管家凭证
![](https://github.com/MengTL4/AutoDaXueXi/blob/main/image/2.png)
identity字段为区分管理员身份，默认是学委，不用改，如果为老师，请改为1

------------


### 编译命令
```go
go build main.go
```
