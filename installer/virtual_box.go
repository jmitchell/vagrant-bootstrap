package installer

type VirtualBoxPackage struct {
	Package
}

func (pkg VirtualBoxPackage) IsInstalled() bool {
	return false
}

func (pkg VirtualBoxPackage) Dependencies() []Package {
	return nil
}

func (pkg VirtualBoxPackage) Install() error {
	return nil
}
