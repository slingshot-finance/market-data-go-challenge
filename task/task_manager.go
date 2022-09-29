package tasks

import (
	"context"
)

/*
	Pushes task to the TaskQueue to be executed by the Executioner
		- Don't push task if TaskQueue is maxed out.
*/
func (t *TaskManager) PushToQueue(task TaskTemplate) {
}

/*
	executes tasks in the TaskQueue as go-routine which pushes results into the resultMap
		--> resultMap a mapping of source ID to output
		--> detects if the channel has been closed, if so pass error "{executioner} inputChan has been closed", and return
*/
func (t *TaskManager) Executioner(ctx context.Context) {
}

/*
	applyFilter applies the filter func on all values in Storage.
		--> if filter returns true on value, drop value from Storage
*/
func (t *TaskManager) applyFilter(ctx context.Context, filter FilterTemplate) {
}

/*
	FaultDetector kicks off applyFilter as a go-routine to filter Storage based on the received filter function.
		--> When a new filter is received, cancel previous applyFilter go-routine and kiff off the new applyFilter go-routine.
*/
func (t *TaskManager) FaultDetector(filter chan FilterTemplate) {
}
