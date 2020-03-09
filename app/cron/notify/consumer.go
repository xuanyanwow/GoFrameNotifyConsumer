package notify

import (
	"fmt"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"time"
)

var (
	// 第一次失败马上重发
	// 2失败延迟5s
	// 3失败延迟10s
	// 4失败延迟30s
	delayTimeArray = [4]int{0, 5, 10, 30}
)

/**
 * 队列监听入口
 */
func Consumer(queue *gqueue.Queue, name string, nextQueue *gqueue.Queue) {
	for {
		v := <-queue.C
		if v != nil {
			data, ok := v.(TaskData)
			if ok == false {
				// log
				fmt.Println("error")
				continue
			}

			go doJob(queue, name, nextQueue, data)
		}
		time.Sleep(time.Millisecond)
	}
}

/**
 * 实际任务内容
 */
func doJob(queue *gqueue.Queue, name string, nextQueue *gqueue.Queue, data TaskData) {
	// 队列最前面的还不可以消费，整个协程堵塞 等待
	if data.NextDoTime > gtime.Timestamp() {
		sleepTime := data.NextDoTime - gtime.Timestamp()
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	isSuccess := false
	responseContext := ""

	if response, err := ghttp.Post(data.Url, data.Data); err != nil {
		afterJob(name, data, false, err.Error())
		return
	} else {
		defer response.Close()
		responseContext = response.ReadAllString()
		if responseContext != "success" {
			isSuccess = false
		} else {
			isSuccess = true
		}
	}

	if isSuccess {
		afterJob(name, data, true)
		return
	}

	// 超出设置次数了
	if data.TryTime >= len(delayTimeArray) {
		fmt.Printf("超出次数 %d 结束 \n", data.TryTime)
		return
	}

	// 如果失败 投递到下一个梯队队列中
	delayTime := delayTimeArray[data.TryTime]

	// 第一次失败投递自身
	data.TryTime += 1
	data.NextDoTime = int64(int(gtime.Timestamp()) + delayTime)

	if data.TryTime == 1 {
		queue.Push(data)
	} else {
		nextQueue.Push(data)
	}

	afterJob(name, data, false, responseContext)
}

// 后置操作，比如汇总结果到日志平台
func afterJob(name string, data TaskData, result bool, content ...string) {
	fmt.Println(gtime.Datetime())
	fmt.Printf("%s -> result %t content : %s \n", name, result, content)
	fmt.Println(data)
}
