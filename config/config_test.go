package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	ci, err := Read([]byte(profile))
	if err != nil {
		t.Error("failed reading yaml", err)
		t.Fail()
	}
	if len(ci.Packages) != 4 {
		t.Error("Expected 4 packages, got ", len(ci.Packages))
		t.Fail()
	}
	if len(ci.RunCommands) < 1 {
		t.Error("Expected some run commands")
		t.Fail()
	}
	if len(ci.RunCommands) != 1 {
		t.Error("Expected 1 run command, got ", len(ci.RunCommands))
		t.Fail()
	}
	if len(ci.RunCommands[0]) != 5 {
		t.Error("Expected 5 run command segments, got ", len(ci.RunCommands[0]))
		t.Fail()
	}
}

func TestMerge(t *testing.T) {
	ci, err := Read([]byte(profile))
	if err != nil {
		t.Error("failed reading yaml", err)
		t.Fail()
	}

	other, err := Read([]byte(profile2))
	if err != nil {
		t.Error("failed reading yaml", err)
		t.Fail()
	}

	ci.Append(other)

	if len(ci.Packages) != 8 {
		t.Error("Expected 8 packages, got ", len(ci.Packages))
		t.Fail()
	}
	if len(ci.RunCommands) != 9 {
		t.Error("Expected 0 run command, got ", len(ci.RunCommands))
		t.Fail()
	}
	if len(ci.RunCommands[0]) != 5 {
		t.Error("Expected 5 run command segments, got ", len(ci.RunCommands[0]))
		t.Fail()
	}
}

const profile = `runcmd:
  - [ systemctl, disable, --now, snapd, snapd.socket ]
packages:
  - git
  - mercurial
  - yadm
  - build-essential`

const profile2 = `runcmd:
  - 'wget -O code.deb https://go.microsoft.com/fwlink/?LinkID=760865'
  - 'dpkg -i code.deb'
  - 'apt-get install --yes -f'
  - [ wget, 'https://dl.google.com/go/go1.12beta2.linux-amd64.tar.gz' ]
  - [ mkdir, -p, /usr/local/go ]
  - [ tar, xvzf, go1.13beta2.linux-amd64.tar.gz, -C, /usr/local/go, --strip-components=1 ]
  - [ rm, -f, go1.12beta2.linux-amd64.tar.gz ]
  - 'echo export PATH=/usr/local/go/bin:$PATH | tee --append /home/ubuntu/.bashrc'
packages:
  - git
  - mercurial
  - yadm
  - build-essential`
