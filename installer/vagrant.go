package installer

type VagrantPackage struct {
	Package
}

func (pkg VagrantPackage) IsInstalled() bool {
	return false
}

func (pkg VagrantPackage) Dependencies() []Package {
	return []Package{VirtualBoxPackage{}}
}

func (pkg VagrantPackage) Install() error {
	return nil
}
