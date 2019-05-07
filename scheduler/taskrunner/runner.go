package taskrunner

// runnner ->startDispatcher->control channer ->data channer

type Runner struct {
	Controller  controlChan
	Error       controlChan
	Data        dataChan
	dataSize    int
	longLived   bool
	Dispathcher fn
	Executor    fn
}

func NewRunnner(size int, longlived bool, d fn, e fn) *Runner {
	//预定义
	return &Runner{
		Controller:  make(chan string, 1),
		Error:       make(chan string, 1),
		Data:        make(chan interface{}, size),
		longLived:   longlived,
		dataSize:    size,
		Dispathcher: d,
		Executor:    e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispathcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:


		}
	}
}

func (r *Runner) startAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
