package operations

import (
	"io/ioutil"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func AddDependencies(rootPath string) {
	unmarshaledPackageJson := utils.UnmarshalPackageJson(rootPath)

	if dependencies, ok := unmarshaledPackageJson["dependencies"].(map[string]interface{}); ok {
		dependencies["react-intl"] = "^5.24.2"
	}

	if devDependencies, ok := unmarshaledPackageJson["devDependencies"].(map[string]interface{}); ok {

		for key, val := range devDepsToAdd {
			devDependencies[key] = val
		}
	}

	dataBytes, err := utils.JSONMarshal(unmarshaledPackageJson, true)
	utils.Check(err)
	err = ioutil.WriteFile(rootPath+"/"+constants.PACKAGE_JSON_PATH, dataBytes, 0644)
	utils.Check(err)
	utils.Success("✔ Added dependencies to package.json")
}

var devDepsToAdd map[string]string = map[string]string{
	"@formatjs/cli": "^4.7.1", "babel-plugin-formatjs": "^10.3.15", "dotenv": "^16.0.0", "node-fetch": "^3.2.0",
}
