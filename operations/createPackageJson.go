package operations

import (
	"sync"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func CreatePackageJson(wg *sync.WaitGroup, projectName string, defaultLocale constants.Iso_Lang, rootPath string) {
	defer wg.Done()

	JSON_CONTENT := FILE_TOP + projectName + "\",\n" + "\t" + FILE_MIDDLE + defaultLocale.Code + FILE_END
	utils.WriteToFile(rootPath+constants.PACKAGE_JSON_PATH, JSON_CONTENT)
	utils.Success("✔ Updated package.json with dependencies and scripts")
}

const FILE_TOP = `{
   "name": "`

const FILE_MIDDLE = `"description":"",
   "license":"UNLICENSED",
   "version":"0.3.0",
   "engines":{
      "node":"^18.4.0",
      "npm":"^8.12.0"
   },
   "scripts":{
	"build": "next build",
	"dev": "next dev",
	"prebuild": "npm run i18n:prepare",
	"predev": "npm run i18n:prepare",
	"start": "next start",
	"lint": "next lint",
	"i18n:prepare": "npm run i18n:compile-messages && npm run i18n:transpile-locales",
	"i18n:compile-messages": "formatjs compile-folder --ast src/i18n/messages/translated src/i18n/messages/compiled --format src/i18n/utils/compile.js",
	"i18n:extract-messages": "formatjs extract \"src/**/*.ts*\" --ignore=\"**/*.d.ts\" --id-interpolation-pattern \"[sha512:contenthash:base64:6]\" --out-file src/i18n/messages/extracted/`

const FILE_END = `.json --format src/i18n/utils/format.js",
	"i18n:transpile-locales": "tsc --p tsconfig-locales.json",
	"i18n:download-translations": "node ./src/i18n/utils/download.mjs",
	"posti18n:download-translations": "npm run i18n:compile-messages",
	"prei18n:upload-translations": "npm run i18n:extract-messages",
	"i18n:upload-translations": "node ./src/i18n/utils/upload.mjs"
   },
   "dependencies":{
      "@superkoders/vercajk":"^1.0.0",
      "next":"^12.1.6",
      "next-transpile-modules":"^9.0.0",
      "react":"^18.2.0",
      "react-dom":"^18.2.0",
      "react-intl":"^5.24.2"
   },
   "devDependencies":{
      "@formatjs/cli":"^4.7.1",
      "@superkoders/prettier-config":"^0.2.6",
      "@types/node":"^18.0.0",
      "@types/react":"^18.0.14",
      "@types/react-dom":"^18.0.5",
      "@typescript-eslint/eslint-plugin":"^5.28.0",
      "babel-plugin-formatjs":"^10.3.15",
      "dotenv":"^16.0.0",
      "eslint":"^8.17.0",
      "eslint-config-next":"^12.1.6",
      "eslint-config-prettier":"^8.5.0",
      "node-fetch":"^3.2.0",
      "prettier":"^2.7.1",
      "typescript":"^4.7.3"
   } 
}`
