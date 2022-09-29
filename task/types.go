package tasks

import (
	"context"
	"sync"
	"sync/atomic"
)

type TaskTemplate func(context.Context) (string, []byte)
type FilterTemplate func([]byte) bool

type WaitGroupCount struct {
	wg    *sync.WaitGroup
	count int64
}

func NewWaitGroupCount() *WaitGroupCount {
	return &WaitGroupCount{
		wg:    new(sync.WaitGroup),
		count: 0,
	}
}

func (w *WaitGroupCount) Add(delta int) {
	atomic.AddInt64(&w.count, int64(delta))
	w.wg.Add(delta)
}

func (w *WaitGroupCount) Done() {
	atomic.AddInt64(&w.count, -1)
	w.wg.Done()
}

func (wg *WaitGroupCount) GetCount() int64 {
	return atomic.LoadInt64(&wg.count)
}

type TaskManager struct {
	TaskQueue       chan TaskTemplate
	ErrorChan       chan error
	MaxPendingTasks int64
	Wg              *WaitGroupCount
}

/*
	Initialises TaskManager
*/
func InitTaskManager(maxPendingTasks int64) *TaskManager {
	return &TaskManager{
		TaskQueue:       make(chan TaskTemplate),
		ErrorChan:       make(chan error),
		MaxPendingTasks: maxPendingTasks,
		Wg:              NewWaitGroupCount(),
	}
}

/*
	input: TaskID 1 byte, a 1 byte, b 1 byte, c 1 byte, sourceID 1 byte
	output: a * b * c
	source: source ID, last byte of the input,  returned as hex string
*/
func TaskA(ctx context.Context, input []byte) (source string, output []byte) {

	return source, output
}

/*
	input: TaskID 1 byte, a 1 byte, b 1 byte, c 1 byte, sourceID 1 byte
	output: a * b * c
	source: source ID, last byte of the input,  returned as hex string
*/
func TaskB(ctx context.Context, input []byte) (source string, output []byte) {

	return source, output
}
