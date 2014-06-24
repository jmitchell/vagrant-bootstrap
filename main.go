package main

import (
	"fmt"

	"github.com/jmitchell/vagrant-bootstrap/installer"
)

func main() {
	err := installer.InstallVagrant()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
