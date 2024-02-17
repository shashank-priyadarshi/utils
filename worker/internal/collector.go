package internal

var workqueue = make(chan *Work, 100)

func Collector(w *Work) {
	workqueue <- w
}
