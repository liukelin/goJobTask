    go的分布式异步任务调度

## 目录结构

~~~

├─lib/              
│  ├─redis.go           redis操作类
│  ├─server_task.go     脚本运行文件
│  └─server_http.go     http服务文件
└─main.go               入口文件

~~~


现在有三种:goroutine
http        任务接收http server
task        任务处理
sentinel    任务进程状态扫描检查

开发中...

