package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func CreateDummyLangPicker(wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()

	utils.WriteToFile(srcFolderPath+"/components/LanguagePicker.tsx", LANG_PICKER)
	utils.Success("✔ Created dummy language picker")
}

const LANG_PICKER = `import Link from 'next/link';
import { useRouter } from 'next/router';
import { FC } from 'react';
import { localeLanguageMap, supportedLocales } from '../i18n/supportedLocales';

export const LanguagePicker: FC = () => {
	const { route, asPath, locale: currentLocale } = useRouter();

	const dropdownItems = supportedLocales.map((locale) => ({
		href: route,
		as: asPath,
		locale: locale,
		label: localeLanguageMap[locale as keyof typeof localeLanguageMap],
	}));

	return (
		<details style={{ zIndex: 10 }}>
			<summary>{localeLanguageMap[currentLocale as keyof typeof localeLanguageMap]}</summary>
			<ul style={{ backgroundColor: 'lightgrey', listStyle: 'none', padding: 0 }}>
				{dropdownItems
					.filter((item) => item.locale !== currentLocale)
					.map((item, idx) => {
						return (
							<li key={idx} style={{ padding: '0 10px' }}>
								<Link {...item}>
									<a>{item.label}</a>
								</Link>
							</li>
						);
					})}
			</ul>
		</details>
	);
};`
