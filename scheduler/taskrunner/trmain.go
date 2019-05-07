package taskrunner

//timer setup  start{}  start{trigger->task->runner}
//timer ,task,runner,longlived.
import (
	"time"
)

type Worker struct {
	ticker *time.Ticker //初始化一个ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.startAll()
		}
	}
}
func Start() {
	// Start video file cleaning
	r := NewRunnner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
