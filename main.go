package main

import (
	"coco/cmd"
	"coco/common/setting"
	"fmt"
)

func main() {
	app := cmd.App()
	app.Run(fmt.Sprintf(":%s", setting.Config.Server.Port))
}
