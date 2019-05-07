package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	VIDEO_PATH ="./videos/"
)

//预定义数据模型
type controlChan chan string
type dataChan chan interface{} //下发数据
type fn func(dc dataChan) error
