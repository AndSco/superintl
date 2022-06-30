package operations

import (
	"fmt"
	"os"
	"os/exec"

	"superkoders.com/intl/utils"
)

func CloneNextStarterKitRepo(projectName string) {
	cloneRepoCommand, err := exec.Command("git", "clone", "git@bitbucket.org:superkoders/nextjs_stack.git", projectName).Output()
	if err != nil {
		utils.Alert(fmt.Sprintf("Error cloning the repo. Check if a folder with that name already exists: %s", err) )
		os.Exit(0)
	}
	utils.Success(string(cloneRepoCommand))
	utils.Success("\n✔ Cloned SuperKoders' base Nextjs stack")
}
