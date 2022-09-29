package input

import (
	"context"
	"fmt"

	task "challenge/task"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

/*
	provide access to the necessary funtions provided by the TaskManager
*/
type TaskManager interface {
	PushToQueue(task.TaskTemplate)
}

type InputHandler struct {
	tskManager TaskManager
	InputChan  chan []byte
	ErrorChan  chan error
}

/*
	initialises InputHandler
*/
func InitInputHandler(taskManager TaskManager) *InputHandler {
	return &InputHandler{
		tskManager: taskManager,
		InputChan:  make(chan []byte),
		ErrorChan:  make(chan error),
	}
}

/*
	parseInput:
		--> listens to InputChan
		--> identifies task ID, which is the 1st byte of the input.
		--> submits task (match by taskID) to the task queue
		--> detects if the channel has been closed, if so pass error "{parseinput} inputChan has been closed", and return
*/
func (h *InputHandler) ParseInput() {
	/*
		Task ID to task Funcs:
			- 0x0a --> push taskA to queue
			- Ox0b --> push taskB to queue
	*/

	for {
		input, ok := <-h.InputChan
		if !ok {
			h.ErrorChan <- fmt.Errorf("inputChan has been closed")
			return
		}

		switch hexutil.Encode(input[:1]) {
		case "0x0a":
			go h.tskManager.PushToQueue(func(ctx context.Context) (string, []byte) { return task.TaskA(ctx, input) })
		case "0x0b":
			go h.tskManager.PushToQueue(func(ctx context.Context) (string, []byte) { return task.TaskB(ctx, input) })
		}
	}
}
