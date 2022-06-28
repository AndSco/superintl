package operations

import (
	"io/ioutil"
	"sync"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func AddNpmScripts(defaultLanguage string, wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	unmarshaledPackageJson := utils.UnmarshalPackageJson(rootPath)

	if packageScripts, ok := unmarshaledPackageJson["scripts"].(map[string]interface{}); ok {
		scriptsToAdd := createNewScripts(defaultLanguage)
		for key, val := range scriptsToAdd {

			packageScripts[key] = val
		}
	}

	dataBytes, err := utils.JSONMarshal(unmarshaledPackageJson, true)
	utils.Check(err)
	err = ioutil.WriteFile(rootPath+"/"+constants.PACKAGE_JSON_PATH, dataBytes, 0644)
	utils.Check(err)
	utils.Success("✔ Added i18n npm scripts to package.json")
}

func createNewScripts(defaultLanguage string) map[string]string {
	return map[string]string{
		"predev":                         "npm run i18n:prepare",
		"prebuild":                       "npm run i18n:prepare",
		"i18n:prepare":                   "npm run i18n:compile-messages && npm run i18n:transpile-locales",
		"i18n:transpile-locales":         "tsc --p tsconfig-locales.json",
		"i18n:extract-messages":          "formatjs extract \"src/**/*.ts*\" --ignore=\"**/*.d.ts\" --id-interpolation-pattern \"[sha512:contenthash:base64:6]\" --out-file src/i18n/messages/extracted/" + defaultLanguage + ".json --format src/i18n/utils/format.js",
		"i18n:compile-messages":          "formatjs compile-folder --ast src/i18n/messages/translated src/i18n/messages/compiled --format src/i18n/utils/compile.js",
		"prei18n:upload-translations":    "npm run i18n:extract-messages",
		"i18n:upload-translations":       "node ./src/i18n/utils/upload.mjs",
		"i18n:download-translations":     "node ./src/i18n/utils/download.mjs",
		"posti18n:download-translations": "npm run i18n:compile-messages",
	}
}
