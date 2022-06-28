package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func AppendToReadMeFile(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()

	utils.AppendToFile(rootPath+"/"+"README.md", INFO_TO_ADD)
	utils.Success("✔ Updated README file")
}

const INFO_TO_ADD = "\n" + `## Intl, translations and POEditor

To handle app locales, this project uses React-Intl integrated with the POEditor translation-management tool. Throughout the project, use React-Intl's recommended way of declaring intl messages (refrain from using IDs, use defaultMessage and description instead - see https://formatjs.io/docs/getting-started/message-declaration).

#### This starter-kit includes:

* Files to store your intl messages (compiled, extracted, translated) 
* A LocaleProvider wrapping the \_app file
* A number of utility functions to interact automatically with POEditor 
* A file (supportedLocales.ts), which is the only source of truth for your project's locales 
* A tsconfig-locales file, emitting (on npm run dev and npm run build) a js version of the above file to be used in non-ts files (ex next.config and POeditor utils) 
* Pre-populated .babelrc and next.config files 
* A number of npm scripts to automate the handling of intl messages

#### Intl flow
1. Set up a project on POEditor (` + "`https://poeditor.com`" + `) with the same languages as the ones in this project
2. Retrieve the associated POEditor token and project ID at ` + "`https://poeditor.com/account/api`" + `
3. Enter these values in the ` + "`POEDITOR_TOKEN`" + ` and ` + "`POEDITOR_PROJECT_ID`" + ` env variables in your .env file
4. Add your FormattedMessages to this project, providing the defaultMessage in the default language
5. At any point, use the script ` + "`npm run i18n:upload-translations`" + ` to extract all react-intl messages in the default language and automatically upload them to POEditor
6. Translate the messages to the relevant target languages inside POEditor itself (` + "`https://poeditor.com`" + ` - filter by Untranslated to see the messages missing translations)
7. To download and compile all project translations, use the ` + "`npm run i18n:download-translations`" + ` script. The script takes the POEditor translations via API, compiles them and serve them to the app
`
