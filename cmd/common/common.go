package common

import (
	"fmt"
)

const (
	DefaultHost string = "127.0.0.1"
	DefaultPort string = "52500"
)

const (
	Succeed uint = 0
	Failed  uint = 1
	Error   uint = 2
	Unauth  uint = 3
)

func codeMessage(code uint) string {
	switch code {
	case 0:
		return "Succeed"
	case 1:
		return "Failed"
	case 2:
		return "Error"
	case 3:
		return "Unauthenticated"
	default:
		return "None"
	}
}

func GenMessage(msgType string, msgCode uint, msgDesc string) {
	fmt.Printf("[%s] %s - %s\n", codeMessage(msgCode), msgType, msgDesc)
}
