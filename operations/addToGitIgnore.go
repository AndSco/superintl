package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func AddToGitIgnore(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	utils.AppendToFile(rootPath+"/"+".gitignore", LINES_TO_ADD)
	utils.Success("✔ Updated .gitignore")
}

const LINES_TO_ADD = `
#Intl folders
src/i18n/messages/extracted/
src/i18n/messages/compiled/
src/i18n/supportedLocales.js`
