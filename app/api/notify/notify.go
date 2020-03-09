package notify

import (
	"fmt"
	"gf-app/app/cron/notify"
	di "gf-app/library/Di"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
)

// 查询队列剩余数量，要重启服务之前应该先关闭添加入口，然后查询队列剩余数量 为0 后 间隔3s（防止有任务消费一半） 再stop服务
func QueryResidualLength(r *ghttp.Request) {
	queue := di.Get("queue_normal").(*gqueue.Queue)
	queueResend1 := di.Get("queue_resend1").(*gqueue.Queue)
	queueResend2 := di.Get("queue_resend2").(*gqueue.Queue)
	queueResend3 := di.Get("queue_resend3").(*gqueue.Queue)

	text := `当前剩余
normal -> %d
resend1 -> %d
resend2 -> %d
resend3 -> %d
`

	r.Response.Writeln(fmt.Sprintf(text, queue.Len(), queueResend1.Len(), queueResend2.Len(), queueResend3.Len()))
}

// 关闭添加入口
func RefuseQueue(r *ghttp.Request) {
	di.Set("queue_status", false)
	r.Response.Writeln("ok")
}

// 打开添加入口
func OpenQueue(r *ghttp.Request) {
	di.Set("queue_status", true)
	r.Response.Writeln("ok")
}

// 推送进队列
func PushQueue(r *ghttp.Request) {
	queueStatus := di.Get("queue_status")

	if queueStatus == false {
		r.Response.Writeln("queue_status is false")
		return
	}

	// 校验参数
	url := r.Get("url")
	data := r.Get("data")

	if e := gvalid.Check(url, "url", nil); e != nil {
		r.Response.Writeln(e.String())
		return
	}

	queue := di.Get("queue_normal").(*gqueue.Queue)

	taskData := notify.TaskData{
		Url:        url.(string),
		Data:       data,
		TryTime:    0,
		NextDoTime: gtime.Timestamp(),
	}
	queue.Push(taskData)

	r.Response.Writeln("ok")
}
