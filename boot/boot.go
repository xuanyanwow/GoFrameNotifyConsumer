package boot

import (
	"gf-app/app/cron/notify"
	di "gf-app/library/Di"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"time"
)


func init() {
	// 初始化队列
	di.Set("queue_normal", gqueue.New())
	di.Set("queue_resend", gqueue.New())

	// 定时生产者
	gtimer.AddSingleton(500 * time.Millisecond, func(){
		queue := di.Get("queue_normal").(*gqueue.Queue)
		taskData := notify.TaskData{
			Url:     "http://www.baidu.com",
			Data:    gtime.Datetime(),
			TryTime: 1,
		}
		queue.Push(taskData)
	})

	// 启动监听任务定时器
	queue := di.Get("queue_normal").(*gqueue.Queue)
	notify.Consumer(queue)


}

