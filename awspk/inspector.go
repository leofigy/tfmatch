package awspk

import (
	"fmt"
	pkill "github.com/pitakill/aws_resources"
)

func InspectResources(stack pkill.Factory) {
	fmt.Println(stack)
}
