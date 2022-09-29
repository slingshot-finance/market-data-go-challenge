package input

type InputHandler struct {
	InputChan chan []byte
	ErrorChan chan error
}

/*
	initialises InputHandler
*/
func InitInputHandler() *InputHandler {
	return &InputHandler{
		InputChan: make(chan []byte),
		ErrorChan: make(chan error),
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

}
