package gitconnector

import (
	"fmt"
	"strings"

	"git.rickiekarp.net/rickie/tree2yaml/executor"
)

func GetRepoTopLevelDir() string {
	stdout, _, _ := executor.ExecuteCmd("git", "rev-parse", "--show-toplevel")
	return stdout
}

func GetGitLogHashes(depth int) []string {
	stdout, _, _ := executor.ExecuteCmd("git", "log", "--format=format:%H", "-"+fmt.Sprint(depth))
	hashes := strings.Split(stdout, "\n")
	return hashes
}

func ShowFileAtRevision(file, revision string) (string, int) {
	fileContent, _, exitCode := executor.ExecuteCmd("git", "show", revision+":"+file)
	return fileContent, exitCode
}
