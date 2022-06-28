package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func CreateLocalesTsConfig(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	utils.CreateRootFile("tsconfig-locales.json", TS_CONFIG_CONTENT, rootPath)
	utils.Success("✔ Created tsconfig file for project's locales")
}

const TS_CONFIG_CONTENT = `{
	"extends": "./tsconfig.json",
	"compilerOptions": {
		"noEmit": false,
		"sourceMap": false,
		"module": "commonjs",
		"incremental": false
	},
	"include": ["src/i18n/supportedLocales.ts"],
	"outdir": "./src/i18n"
}
`
