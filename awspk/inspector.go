package awspk

import (
	"fmt"

	pkill "github.com/leofigy/aws_resources"
)

func InspectResources(stack pkill.Factory) {
	fmt.Println(stack)
}
