package notify

import (
    "fmt"
    "github.com/gogf/gf/container/gqueue"
)


func Consumer(queue *gqueue.Queue)  {
    go func() {
        for {
            v := <-queue.C
            if v != nil {
                v = v.(TaskData)
                fmt.Println(v)
            }
        }
    }()
}
