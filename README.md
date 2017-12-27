    go的分布式异步任务调度

## 目录结构

~~~

├─lib/              
│  ├─redis.go           redis操作类
│  ├─server_task.go     脚本运行文件
│  └─server_http.go     http服务文件
└─main.go               入口文件

~~~


现在有三种goroutine

    
    http    

        任务接收http server

    
    task        
        
        队列监控 任务消费获取 （按队列数量 多个goroutine）
        监测新增队列并创建task

    
    sentinel

        任务进程状态扫描检查 并启动




上传任务脚本并 配置 对应队列等参数，你所有消费机自动获取部署

conf.json
{
    "http_port":"8888",

    "task_process":1,

    "debug":true,

    "_comment" : "监测频率",
    "sentinel_time":3,

    "_comment" : "监测频率",
    // 任务脚本路径
    "task_script_file":"",

    "_comment" : "监测频率",
    // MQ类型 redis/rabbitmq
    "mq":"redis",

    "_comment" : "监测频率",
    // redis
    "redis_host":"localhost",
    "redis_port":6379,
    "redis_db":0,

    "_comment" : "mysql",
    "mysql_host":"localhost",
    "mysql_port":3306,
    "mysql_user":"",
    "mysql_pass":"",

    "shell_cli":    "/bin/sh",
    "php_cli":      "/usr/bin/php7.0",
    "java_cli":     "/usr/bin/java",
    "python2_cli":  "/usr/bin/python2",
    "python3_cli":  "/usr/bin/python3",

}

开发中...

