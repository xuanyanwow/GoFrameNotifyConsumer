# 回调任务消费

类似微信支付回调，多条推送队列

将一个推送任务推送进来，消费者将会POST任务数据到指定的url（通知客户）

客户接受接口需要响应指定内容  默认"success"字段，如果没有响应该内容，则认定回调失败，将会按一定间隔再次触发推送，直到超出次数或者推送成功

## 环境依赖

基于`goframe`框架，请先查看该框架的开发文档、gf工具使用文档

打包构建:`gf build main.go`

上传构建包和其他目录(配置目录 public目录等) 运行即可

## 默认配置

默认失败4次就丢弃，每次间隔为：{0, 5, 10, 30}

默认推送的url返回"success"即为成功，若要修改，在app/cron/notify/consumer.go的`doJob`方法中找到相关逻辑修改 `if responseContext != "success" {`


## 相关api url

- 打开推送入口 /notify/OpenQueue （默认打开）
- 关闭推送入口 /notify/RefuseQueue
- 推送数据进队列 /notify/PushQueue post两个字段
- - url
- - data 数组

可以参考api接口的逻辑，实现TCP版推送入口，增加响应速率。

## 相关代码所在文件

boot/boot.go 开启监听的队列

app/cron/notify/consumer.go 设置每次队列失败后增加的间隔

## 如何新增次数

- app/cron/notify/consumer.go 新增delayTimeArray数组的元素
- boot/boot.go 增加初始化队列、消费者开启（传入队列，最后一个队列传递自身就好了 反正不会用到）


