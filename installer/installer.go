package installer

import (
	"errors"
	"fmt"
)

type Package interface {
	IsInstalled() bool
	Dependencies() []Package
	Install() error
}

func InstallVagrant() error {
	return installPackage(VagrantPackage{})
}

func installPackage(pkg Package) error {
	if pkg.IsInstalled() {
		return nil
	}

	// Assumes there are no cycles in the dependency graph. Can
	// add cycle detection later if necessary.
	for _, dependency := range pkg.Dependencies() {
		if !dependency.IsInstalled() {
			err := installPackage(dependency)
			if err != nil {
				return err
			}
		}

		if !dependency.IsInstalled() {
			msg := fmt.Sprintf("%T failed even though the installer ran without any errors.", dependency)
			return errors.New(msg)
		}
	}

	err := pkg.Install()
	if err != nil {
		return err
	}

	return nil
}
