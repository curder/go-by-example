- [并发启动单个goroutine](/src/concurrences/signal_go_routine_test.go)
- [并发启动多个goroutine](/src/concurrences/multi_go_routine_test.go)
- [并发回调函数方式goroutine](/src/concurrences/callback_go_routine_test.go)
- [GOMAXPROCS](/src/concurrences/go_max_procs_test.go)
- [通道](/src/concurrences/channels/README.md)
  - [通道channel](/src/concurrences/channels/channel_test.go)
  - [通道小练习](/src/concurrences/channels/worker_pool_demo_test.go)
  - [通道select](/src/concurrences/channels/channel_select_test.go)
- [并发操作全局变量问题](/src/concurrences/syncs/multi_go_routine_modify_global_variable_test.go)
- [并发安全和互斥锁](/src/concurrences/syncs/multi_go_routine_modify_global_variable_by_sync_mutex_test.go)
- [并发安全和读写互斥锁](/src/concurrences/syncs/multi_go_routine_modify_global_variable_by_sync_rw_mutex_test.go)
- [sync.WaitGroup](/src/concurrences/syncs/go_routine_sync_wait_group_test.go)
- [sync.Once](/src/concurrences/syncs/go_routine_sync_once_test.go)


## channel 常见的异常总结
| channel | nil | 非空 | 空 | 满了 | 没满 |
| :---- | :---- | :---- | :---- | :---- | :---- |
| **接收** | 阻塞 | 接收值 | 阻塞 | 接收值| 接收值 |
| **发送** | 阻塞 | 发送值 | 发送值 | 阻塞 | 发送值 |
| **关闭** | panic | 关闭成功，读完数据后返回零值| 关闭成功，返回零值| 关闭成功读取数据和返回零值| 关闭成功读取数据和返回零值 |
