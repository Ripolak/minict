package commands

import (
	"fmt"

	"github.com/Ripolak/minict/run"

	"github.com/apex/log"
)

func Start(containerDir string, name string) {
	if !doesContainerExist(name, containerDir) {
		log.Fatal(fmt.Sprintf("Container with name %s does not exist. Use the 'run' option to start a new container", name))
	}
	err := run.RunContainer(containerDir, name)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to run image. Error received: %s", err.Error()))
	}
}
