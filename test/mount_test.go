package test

import (
	"testing"

	"fmt"
	"os/exec"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestFindMount(t *testing.T) {

	path := "/var/lib/kubelet/plugins/kubernetes.io/flexvolume/lcj/lvm/mounts/test"

	t.Logf("find %s is mounted", path)
	{
		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf(`findmnt -n %s 2>/dev/null | cut -d' ' -f1`, path))
		if r, err := cmd.Output(); err != nil {
			t.Fatal("\t\tShould be able exec findmnt command", ballotX, err)
		} else if string(r) == path {
			t.Log("\t\t mounted", checkMark)
		} else {
			t.Log("\t\t unmounted", checkMark)
		}
	}

}
