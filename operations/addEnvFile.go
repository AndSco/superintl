package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func AddEnvFile(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	utils.WriteToFile(rootPath+"/.env", VARS_TO_ADD)

	utils.Success("✔ Added .env file with POEditor variables needed")

}

const VARS_TO_ADD = `POEDITOR_TOKEN=
POEDITOR_PROJECT_ID=`
