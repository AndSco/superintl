package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func EditIndexFile(wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()

	utils.WriteToFile(srcFolderPath+"/pages/index.tsx", INDEX_FILE_CONTENT)
	utils.Success("✔ Updated homepage")

}

const INDEX_FILE_CONTENT = `import type { NextPage } from 'next';
import { FormattedMessage } from 'react-intl';
import { LanguagePicker } from '../components/LanguagePicker';

const Home: NextPage = () => {
	return (
		<div style={{ padding: 20, lineHeight: 2.2 }}>
			<nav style={{ width: '95vw', height: 40, display: 'flex', justifyContent: 'flex-end' }}>
				<LanguagePicker />
			</nav>
			<main>
				<h1>
					<FormattedMessage description={'main-heading'} defaultMessage="Multi-language NextJS app" />
				</h1>
				<section>
					<p>
						To handle app locales, this project uses React-Intl integrated with the POEditor translation-management tool.
						Throughout the project, use React-Intl&apos;s recommended way of declaring intl messages (refrain from using IDs,
						use defaultMessage and description instead - see{' '}
						<a href="https://formatjs.io/docs/getting-started/message-declaration">here</a>).
					</p>
					<h3>This starter-kit includes:</h3>
					<ul>
						<li>Files to store your intl messages (compiled, extracted, translated)</li>
						<li>A LocaleProvider wrapping the _app.tsx file</li>
						<li>A number of utility functions to interact automatically with POEditor</li>
						<li>A file (supportedLocales.ts), which is the only source of truth for your project&apos;s locales</li>
						<li>
							tsconfig-locales file, emitting (on npm run dev and npm run build) a js version of the above file to be used in
							non-ts files (ex next.config and POeditor utils)
						</li>
						<li>Pre-populated .babelrc and next.config files</li>
						<li>A number of npm scripts to automate the handling of intl messages</li>
					</ul>

					<h3>Intl flow</h3>
					<ol>
						<li>Set up a project on POEditor (https://poeditor.com) with the same languages as the ones in this project</li>
						<li>Retrieve the associated POEditor token and project ID at https://poeditor.com/account/api</li>
						<li>
							Enter these values in the <code>POEDITOR_TOKEN</code> and <code>POEDITOR_PROJECT_ID</code> env variables in your
							.env file.
						</li>
						<li>
							Add your <code>FormattedMessages</code> to this project, providing the defaultMessage in the default language
						</li>
						<li>
							At any point, use the script <code>npm run i18n:upload-translations</code> to extract all react-intl messages in
							the default language and automatically upload them to POEditor.
						</li>
						<li>
							Translate the messages to the relevant target languages inside POEditor itself (https://poeditor.com - filter
							by <code>Untranslated</code> to see the messages missing translations).
						</li>
						<li>
							To download and compile all project translations, use the <code>npm run i18n:download-translations</code>{' '}
							script. The script takes the POEditor translations via API, compiles them and serve them to the app.
						</li>
					</ol>
				</section>
				<p>
					<FormattedMessage description={'test-message'} defaultMessage="Edit me" />
				</p>
			</main>
		</div>
	);
};

Home.displayName = 'Home';

export default Home;
`
