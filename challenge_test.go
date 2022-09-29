package challenge

import (
	"bytes"
	task "challenge/task"
	"fmt"
	"testing"
	"time"
)

func verifyVal(sourceID string, output []byte) ([]byte, bool) {
	var expectedVal []byte
	switch sourceID {
	case "0x0c":
		expectedVal = []byte{6}
	case "0x17":
		expectedVal = []byte{15}
	case "0x22":
		expectedVal = []byte{90}
	case "0x15":
		expectedVal = []byte{156}
	case "0x04":
		expectedVal = []byte{240}
	case "0x0e":
		expectedVal = []byte{30}
	}

	if !bytes.Equal(output, expectedVal) {
		return expectedVal, false
	}
	return expectedVal, true

}

// go test -run=TestApp
func TestApp(t *testing.T) {
	_max_storage_cap := int64(5)
	app := InitApp(2, _max_storage_cap) // 1 task at a time.
	app.Run()

	testCases := [][]byte{[]byte{10, 1, 2, 3, 12}, []byte{11, 2, 3, 3, 23}, []byte{10, 5, 6, 3, 34}}
	for _, c := range testCases {
		app.InputHandler.InputChan <- c
	}

	time.Sleep(1 * time.Second)

	fmt.Print("\n.....")
	for i, v := range app.GetSourceIDs() {
		_res := app.FetchResult(v)
		expected, valid := verifyVal(v, _res)
		fmt.Printf("\n\t SourceID: %s Got: %+v Expected; %+v", v, _res, expected)
		if !valid || int64(i) > _max_storage_cap {
			t.FailNow()
		}
	}

	filterChan := make(chan task.FilterTemplate)

	go app.TaskManager.FaultDetector(filterChan)
	filterChan <- FilterA

	testCases = [][]byte{[]byte{11, 5, 5, 3, 14}, []byte{11, 4, 8, 13, 21}, []byte{10, 5, 16, 3, 4}}
	for _, c := range testCases {
		app.InputHandler.InputChan <- c
	}

	time.Sleep(1 * time.Second)

	fmt.Print("\n.....")
	for i, v := range app.GetSourceIDs() {
		_res := app.FetchResult(v)
		expected, valid := verifyVal(v, _res)
		fmt.Printf("\n\t SourceID: %s Got: %+v Expected: %+v", v, _res, expected)
		if !valid {
			t.Error("nvalid Value !")
			t.FailNow()
		}

		if int64(i) > _max_storage_cap {
			t.Error("Storage size exceeded _max_storage_cap ! ")
			t.FailNow()

		}
	}

	filterChan <- FilterB

	time.Sleep(1 * time.Second)

	fmt.Print("\n.....")
	for i, v := range app.GetSourceIDs() {
		_res := app.FetchResult(v)
		expected, valid := verifyVal(v, _res)
		fmt.Printf("\n\t SourceID: %s Got: %+v Expected: %+v", v, _res, expected)
		if !valid {
			t.Error("nvalid Value !")
			t.FailNow()
		}

		if int64(i) > _max_storage_cap {
			t.Error("Storage size exceeded _max_storage_cap ! ")
			t.FailNow()

		}

		if FilterB(_res) {
			t.Error("Value was not filtered out... ")
			t.FailNow()
		}
	}
}
