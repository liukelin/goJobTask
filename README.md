    go的分布式任务

## 目录结构

~~~

├─lib/              处理方法
│  ├─redis.go           redis操作类
│  ├─server_task.go     脚本运行文件
│  └─server_http.go     http服务文件
└─main.go               入口文件

~~~


http        任务接收http server
task        任务处理主进程
sentinel    任务进程扫描检查进程

开发中...

