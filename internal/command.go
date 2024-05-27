package internal

import (
	"fmt"
	"reflect"
)

type Action struct {
	Check interface{}
	Run   []interface{}
}

func runActionItem(actionItem interface{}, cwd string) bool {
	switch v := actionItem.(type) {
	case string:
		_, err := RunCommand(v, cwd)
		if err != nil {
			fmt.Println("Error running command: ", v)
			return true
		}
	case func() bool:
		return v()
	default:
		panic(fmt.Sprintf("Invalid type: %s", reflect.TypeOf(v)))
	}
	return false
}

func RunAction(action Action, cwd string) {

	if action.Check != nil {
		if runActionItem(action.Check, cwd) {
			return
		}
	}

	for _, action := range action.Run {
		runActionItem(action, cwd)
	}
}
