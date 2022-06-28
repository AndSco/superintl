package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func CreateBabelrcFile(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	utils.CreateRootFile(".babelrc", BABELRC_CONTENT, rootPath)
	utils.Success("✔ Created .babelrc file")
}

const BABELRC_CONTENT = `{
	"presets": ["next/babel"],
	"plugins": [
		[
			"formatjs",
			{
				"idInterpolationPattern": "[sha512:contenthash:base64:6]",
				"ast": true
			}
		]
	]
}
`
