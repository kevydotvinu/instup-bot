package command

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"

	hbot "github.com/whyrusleeping/hellabot"
)

// GetTaskGraph prints the taskgraph
func GetTaskGraph(release string) (string, error) {
	workDir, _ := filepath.Abs("./cluster-version-util")
	filename := "cluster-version-util.sh"
	link, err := exec.Command("/bin/bash", path.Join(workDir, filename), release).Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(link), nil
}

// Manifestgraph of OpenShift release image
func (core Core) Manifestgraph(m *hbot.Message, args []string) {
	if len(args) < 1 {
		core.Bot.Reply(m, "Please tell me the release. For example: 4.10.10.")
		return
	}
	release := args[0]
	core.Bot.Reply(m, fmt.Sprintf("Processing OpenShift %s release manifests. Hold on!", release))
	link, err := GetTaskGraph(release)
	if err != nil {
		core.Bot.Reply(m, fmt.Sprintf("Image not found!: %s", err))
	} else {
		core.Bot.Reply(m, fmt.Sprintf("The graph is ordered by the number and component of the manifest file. Here is the OpenShift %s manifest graph. \n%s", release, link))
	}
}
