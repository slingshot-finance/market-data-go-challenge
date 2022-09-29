package challenge

import (
	input "challenge/input"
	task "challenge/task"
	"context"
	"math/big"
)

type AppCtx struct {
	Ctx          context.Context
	TaskManager  *task.TaskManager
	InputHandler *input.InputHandler

	Storage        map[string][]byte
	MaxStorageSize int64
}

/*
	Pushes result to Storage
		--> Drop the oldest value out of Storage if it's size reaches MaxStorageCap
*/
func (app *AppCtx) PushResult(sourceID string, result []byte) {
	app.Storage[sourceID] = result
}

/*
	Fetches Result from Storage
*/
func (app *AppCtx) FetchResult(sourceID string) []byte {
	return app.Storage[sourceID]
}

/*
	Check if value exist in Storage by SourceID
		- If Exist Drops value from Storage by sourceID and Update StorageKeys
*/
func (app *AppCtx) DropResult(sourceID string) {
}

/*
	Return all Storage Keys
*/
func (app *AppCtx) GetSourceIDs() (sourceIDs []string) {
	return sourceIDs
}

func FilterA(in []byte) bool {
	return big.NewInt(0).SetBytes(in).Cmp(big.NewInt(10)) >= 0
}

func FilterB(in []byte) bool {
	return big.NewInt(0).SetBytes(in).Cmp(big.NewInt(50)) <= 0
}

func InitApp(maxPendingTasks int64, maxStorageSize int64) *AppCtx {
	app := &AppCtx{
		Ctx:            context.Background(),
		Storage:        make(map[string][]byte),
		MaxStorageSize: maxStorageSize,
	}
	return app
}

func (app *AppCtx) Run() {
	go app.TaskManager.Executioner(app.Ctx) // run task executioner in the background
	go app.InputHandler.ParseInput()        // start listening to input
}
