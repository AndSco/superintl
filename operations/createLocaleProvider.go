package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func CreateLocaleProvider(wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()

	utils.WriteToFile(srcFolderPath+"/components/LocaleProvider.tsx", PROVIDER_CONTENT)
	utils.Success("✔ Created locale provider to use in your _app file")
}

const PROVIDER_CONTENT = `import { useRouter } from 'next/router';
import { FC, useEffect, useState, PropsWithChildren } from 'react';
import { IntlProvider } from 'react-intl';
import { getLocaleMessages, isValidLocale, Locale } from '../i18n/supportedLocales';

export const LocaleProvider: FC<PropsWithChildren> = ({ children }) => {
	const { locale, defaultLocale } = useRouter();
	const localeToUse = isValidLocale(locale) ? locale : (defaultLocale as Locale);
	const [messages, setMessages] = useState({});

	useEffect(() => {
		(async () => {
			const messages = await getLocaleMessages(localeToUse);
			setMessages(messages);
		})();
	}, [localeToUse]);

	return (
		<IntlProvider locale={localeToUse} messages={messages} onError={() => null}>
			{children}
		</IntlProvider>
	);
};
`
