package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func WrapAppInLocaleContext(wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()

	utils.WriteToFile(srcFolderPath+"/pages/_app.tsx", APP_FILE_CONTENT)
	utils.Success("✔ Added locale provider to _app file")
}

const APP_FILE_CONTENT = `import '../styles/globals.css';
import type { AppPropsWithLayout } from '../types/next';
import { Layout } from '../components/Layout/Layout';
import { LocaleProvider } from '../components/LocaleProvider';

function MyApp({ Component, pageProps }: AppPropsWithLayout) {
	const getLayout = Component.getLayout || ((page) => <Layout>{page}</Layout>);

	return getLayout(
		<LocaleProvider>
			<Component {...pageProps} />
		</LocaleProvider>,
	);
}

export default MyApp;
`
