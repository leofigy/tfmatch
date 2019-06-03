package awspk

import (
	"fmt"
	"reflect"

	pkill "github.com/leofigy/aws_resources"
)

func InspectResources(stack pkill.FactoryData) {
	info, err := stack.CallAWS()
	if err != nil {
		fmt.Printf("[CALLING ERRORS]\n%s\n", err)
	}
	for t := range info {
		switch info[t].Kind() {
		case reflect.Slice:
			getItems(info[t])
		default:
			fmt.Println("Unexpected", info[t].Kind())
		}
	}
}

func getItems(list reflect.Value) {
	if list.Len() > 0 {
		for i := 0; i < list.Len(); i++ {
			fmt.Printf("%v\n", list.Index(i).Interface())
		}
	}
}
