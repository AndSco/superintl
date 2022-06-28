package operations

import (
	"os"
	"os/exec"

	"superkoders.com/intl/utils"
)

func CloneNextStarterKitRepo(projectName string) {
	cloneRepoCommand := exec.Command("git", "clone", "git@bitbucket.org:superkoders/nextjs_stack.git", projectName)
	cloneRepoCommand.Stdout = os.Stdout
	cloneRepoCommand.Stderr = os.Stderr
	cloneRepoCommand.Run()
	utils.Success("\n✔ Cloned SuperKoders' base Nextjs stack")
}
