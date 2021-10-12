package state

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	defaultProject := &Project{
		Name:               "default",
		LXDName:            "default",
		ContainerMountName: "projects",
		Profiles:           []string{"default"},
	}

	Projects["default"] = defaultProject

	servicesProject := &Project{
		Name:               "services",
		LXDName:            "services",
		ContainerMountName: "services",
		Profiles:           []string{"default"},
	}

	Projects["services"] = servicesProject
}

var Projects = make(map[string]*Project)

const configDirName = "dlxpersist"

type Project struct {
	Name               string
	LXDName            string
	ContainerMountName string
	Profiles           []string
}

func GetProject(name string) *Project {
	if name == "" {
		name = "default"
	}
	if project, ok := Projects[name]; ok {
		return project
	}

	return nil
}

func (p *Project) MountPath() string {
	return filepath.Join(GetHomePath(), configDirName, p.Name)
}

func (p *Project) ContainerMountPath() string {
	return filepath.Join(GetHomePath(), p.ContainerMountName)
}

func (p *Project) CreateMountPath() error {
	if _, err := os.Stat(p.MountPath()); os.IsNotExist(err) {
		err := os.MkdirAll(p.MountPath(), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// CommonMountPath is a project specific directory that is mounted
// into every container, to provide shared storage for provisioning
// scripts or configuration files.
func (p *Project) CommonMountPath() string {
	return filepath.Join(GetHomePath(), configDirName, p.Name, "common")
}

// ContainerCommonMountPath determines the mount path for the CommonMountPath
// inside the container.
func (p *Project) ContainerCommonMountPath() string {
	return filepath.Join(GetHomePath(), "common")
}

// CreateCommonMountPath creates the CommonMountPath on the LXD host
func (p *Project) CreateCommonMountPath() error {
	if _, err := os.Stat(p.CommonMountPath()); os.IsNotExist(err) {
		err := os.MkdirAll(p.CommonMountPath(), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) InstanceMountPath(instanceName string) string {
	return filepath.Join(p.MountPath(), instanceName)
}

func (p *Project) CreateInstanceMountPath(instanceName string) error {
	if _, err := os.Stat(p.InstanceMountPath(instanceName)); os.IsNotExist(err) {
		err := os.MkdirAll(p.InstanceMountPath(instanceName), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetHomePath() string {
	// Find home directory form env
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Get Home Dir: " + err.Error())
	}
	return home
}
