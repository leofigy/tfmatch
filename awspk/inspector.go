package awspk

import (
	"fmt"

	pkill "github.com/leofigy/aws_resources"
)

func InspectResources(stack pkill.FactoryData) {
	info, err := stack.CallAWS()
	if err != nil {
		fmt.Printf("Following errors found\n%s\n", err)
	}
	fmt.Println(info)

}
