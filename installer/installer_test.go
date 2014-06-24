package installer

import (
	"testing"
)

var noDependenciesPkgInstalled bool

type NoDependenciesPkg struct {
	Package
}

func (pkg NoDependenciesPkg) IsInstalled() bool {
	return noDependenciesPkgInstalled
}

func (pkg NoDependenciesPkg) Dependencies() []Package {
	return nil
}

func (pkg NoDependenciesPkg) Install() error {
	noDependenciesPkgInstalled = true
	return nil
}

func TestInstallNoDependenciesPkg(t *testing.T) {
	pkg := NoDependenciesPkg{}

	// preconditions
	noDependenciesPkgInstalled = false
	if pkg.IsInstalled() {
		t.Errorf("%T is unexpectedly installed already", pkg)
	}

	// installation
	if err := installPackage(pkg); err != nil {
		t.Errorf("Unexpectedly encountered error '%v' while installing %T", err, pkg)
	}

	// postconditions
	if !pkg.IsInstalled() {
		t.Errorf("%T did not install successfully even though the installer reported no errors", pkg)
	}
}


var parentPkgInstalled bool
var childPkgInstalled bool

type ParentPkg struct {
	Package
}

func (pkg ParentPkg) IsInstalled() bool {
	return parentPkgInstalled
}

func (pkg ParentPkg) Dependencies() []Package {
	return []Package {ChildPkg{}}
}

func (pkg ParentPkg) Install() error {
	parentPkgInstalled = true
	return nil
}


type ChildPkg struct {
	Package
}

func (pkg ChildPkg) IsInstalled() bool {
	return childPkgInstalled
}

func (pkg ChildPkg) Dependencies() []Package {
	return nil
}

func (pkg ChildPkg) Install() error {
	childPkgInstalled = true
	return nil
}

func TestInstallPackageThatHasDependency(t *testing.T) {
	parentPkg := ParentPkg{}
	
	// preconditions
	parentPkgInstalled = false
	childPkgInstalled = false
	if parentPkg.IsInstalled() {
		t.Errorf("%T is unexpectedly installed already", parentPkg)
	}
	if parentPkg.Dependencies()[0].IsInstalled() {
		t.Errorf("%T is unexpectedly installed already", parentPkg.Dependencies()[0])
	}

	// installation
	if err := installPackage(parentPkg); err != nil {
		t.Errorf("Error while installing %T: %v", parentPkg, err)
	}

	// postconditions
	if !parentPkg.IsInstalled() {
		t.Errorf("%T did not install successfully even though the installer reported no errors", parentPkg)
	}
	if !parentPkg.Dependencies()[0].IsInstalled() {
		t.Errorf("%T did not install successfully even though the installer reported no errors", parentPkg.Dependencies()[0])
	}
}

func TestInstallPackageThatHasDependencyAlreadyInstalled(t *testing.T) {
	parentPkg := ParentPkg{}
	
	// preconditions
	parentPkgInstalled = false
	childPkgInstalled = true
	if parentPkg.IsInstalled() {
		t.Errorf("%T is unexpectedly installed already", parentPkg)
	}
	if !parentPkg.Dependencies()[0].IsInstalled() {
		t.Errorf("Expected %T to already be installed", parentPkg.Dependencies()[0])
	}

	// installation
	if err := installPackage(parentPkg); err != nil {
		t.Errorf("Error while installing %T: %v", parentPkg, err)
	}

	// postconditions
	if !parentPkg.IsInstalled() {
		t.Errorf("%T did not install successfully even though the installer reported no errors", parentPkg)
	}
	if !parentPkg.Dependencies()[0].IsInstalled() {
		t.Errorf("%T was already installed but now it's not", parentPkg.Dependencies()[0])
	}
}

