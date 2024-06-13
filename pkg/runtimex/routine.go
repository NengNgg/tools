package runtimex

import (
	"log"
	"runtime/debug"
	"tools/pkg/bytesconv"
)

func GoRoutine(method func()) {
	if method == nil {
		return
	}

	go func() {
		defer Guard()
		method()
	}()
}

func Guard() {
	if r := recover(); r != nil {
		log.Printf("recover panic: %v,  stack: %v",
			r, bytesconv.BytesToString(debug.Stack()))
	}
}
