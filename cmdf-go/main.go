package main

import (
	"fmt"

	"actiontech.cloud/universe/ucommon/v4/log"
	"actiontech.cloud/universe/ucommon/v4/os"
)

func main() {
	stage := log.NewStage().Enter("begin cmf2")
	defer stage.Exit()
	output, err := os.Cmdf2(stage, "ls")
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("ls: %v", output)
}
