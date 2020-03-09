package boot

import (
	"gf-app/app/cron/notify"
	di "gf-app/library/Di"
	"github.com/gogf/gf/container/gqueue"
)

func init() {
	// ================= 方案1 go 内存队列 需要自己完成数据落地备份 以免进程挂掉或者重启的时候数据丢失 ===============
	// 初始化队列
	di.Set("queue_status", true)
	di.Set("queue_normal", gqueue.New())
	di.Set("queue_resend1", gqueue.New())
	di.Set("queue_resend2", gqueue.New())
	di.Set("queue_resend3", gqueue.New())

	// 启动时数据恢复
	// 定时数据落地

	// ================= 方案2 其他队列如redis等 go服务挂了不会影响数据储存  ===============

	// 定时生产者
	//gtimer.AddSingleton(10000*time.Millisecond, func() {
	//	queue := di.Get("queue_normal").(*gqueue.Queue)
	//	taskData := notify.TaskData{
	//		Url:        "http://www.baidu.com",
	//		Data:       gtime.Datetime(),
	//		TryTime:    0,
	//		NextDoTime: gtime.Timestamp(),
	//	}
	//	queue.Push(taskData)
	//})

	// 启动监听任务定时器
	queue := di.Get("queue_normal").(*gqueue.Queue)
	queueResend := di.Get("queue_resend1").(*gqueue.Queue)
	queueResend2 := di.Get("queue_resend2").(*gqueue.Queue)
	queueResend3 := di.Get("queue_resend3").(*gqueue.Queue)

	go notify.Consumer(queue, "normal", queueResend)
	go notify.Consumer(queueResend, "resend1", queueResend2)
	go notify.Consumer(queueResend2, "resend2", queueResend3)
	// 最后一个队列了，nextQueue 传递自身就可以了
	go notify.Consumer(queueResend3, "resend3", queueResend3)
}
