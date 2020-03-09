package notify

type TaskData struct {
	Url        string
	Data       interface{}
	TryTime    int
	NextDoTime int64
}
