# 回调任务消费

模拟微信支付回调，多条推送队列

## 相关api url

- 打开推送入口 /notify/OpenQueue （默认打开）
- 关闭推送入口 /notify/RefuseQueue
- 推送数据进队列 /notify/PushQueue post两个字段
- - url
- - data 数组

## 相关代码所在文件

boot/boot.go 开启监听的队列

app/cron/notify/consumer.go 设置每次队列失败后增加的间隔

## 默认配置

默认失败4次就丢弃，每次间隔为：{0, 5, 10, 30}
默认推送的url返回"success"即为成功，若要修改，在app/cron/notify/consumer.go的`doJob`方法中找到相关逻辑修改 `if responseContext != "success" {`
`

## 如何新增次数

- app/cron/notify/consumer.go 新增delayTimeArray数组的元素
- boot/boot.go 增加初始化队列、消费者开启（传入队列，最后一个队列传递自身就好了 反正不会用到）


