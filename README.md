# ClassHousekeeper
自动化实现：班级小管家作业(青年大学习)截图提交

- 快捷配置
- 适配班级小管家数据结构
- 可配置定时任务

------------


### 使用说明
在ClassHousekeeper.go中配置imprint的值即可使用
```Golang
const imprint = "xxxx"
```
Imprint字段为班级小管家凭证
![](https://github.com/MengTL4/AutoDaXueXi/blob/main/image/2.png)
本项目仅用于管理员系统身份为[班委]的班级，如果身份为[老师]，需要自己抓包修改数据结构

------------


### 编译命令
```go
go build main.go
```
