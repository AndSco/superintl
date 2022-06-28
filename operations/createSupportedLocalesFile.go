package operations

import (
	"fmt"
	"sync"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func CreateSupportedLocalesFile(defaultLocale constants.Iso_Lang, otherLocales []constants.Iso_Lang, wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()
	allLocalesStructs := append(otherLocales, defaultLocale)
	imports := "import { MessageFormatElement } from 'react-intl';\n\n"
	locales := []string{defaultLocale.Code}
	for _, loc := range otherLocales {
		locales = append(locales, loc.Code)
	}

	enumStart := "export enum Locale {\n"
	enumEnd := "};\n\n"
	enumBody := ""
	for _, locale := range locales {
		enumBody += "\t" + "'" + locale + "'" + " = " + "'" + locale + "',\n"
	}

	localesEnum := fmt.Sprintf(enumStart + enumBody + enumEnd)

	defaultLocaleDefinition := "export const defaultLocale = Locale['" + defaultLocale.Code + "'];\n"

	getLocaleMessagesFunctionStart := `export const getLocaleMessages = (shortLocale: Locale): IntlMessages => {
	switch (shortLocale) {` + "\n"
	getLocaleMessagesFunctionEnd := "\t}\n}"

	switchBody := ""
	for _, otherLocale := range otherLocales {
		switchBody += "\t\t case " + "'" + otherLocale.Code + "':\n" + specifyFileImport(otherLocale.Code)
	}

	switchBody += "\t\t default:\n" + specifyFileImport(defaultLocale.Code)

	getLocaleMessagesFunction := fmt.Sprintf(getLocaleMessagesFunctionStart + switchBody + getLocaleMessagesFunctionEnd)

	languageMapStart := "\n" + `type LocaleLanguageMap = Record<SupportedLocale, string>;` + "\n\n" + `export const localeLanguageMap: LocaleLanguageMap = {` + "\n"

	languageMapBody := ""
	for idx, loc := range allLocalesStructs {
		languageMapBody += "\t'" + loc.Code + "': '" + loc.NativeName + "',"
		if idx < len(allLocalesStructs)-1 {
			languageMapBody += "\n"
		}
	}

	languageMapEnd := "\n" + "};\n\n"

	languageMap := languageMapStart + languageMapBody + languageMapEnd

	restOfFile := `export const supportedLocales = (Object.keys(Locale) as Array<keyof typeof Locale>).map((k) => Locale[k]);

export type SupportedLocale = typeof supportedLocales[number];

export const isValidLocale = (locale: unknown): locale is SupportedLocale => {
	return supportedLocales.includes(locale as SupportedLocale);
};

type IntlMessages = Promise<Record<string, MessageFormatElement[]>>;

`
	wholeFile := fmt.Sprintf(imports + localesEnum + defaultLocaleDefinition + languageMap + restOfFile + getLocaleMessagesFunction)
	utils.WriteToFile(utils.BuildIntlPath(srcFolderPath)+"/"+SUPPORTED_LOCALES_FILE, wholeFile)

	utils.Success("✔ Created supportedLocales.ts file, source of truth for project's locales")
}

func specifyFileImport(locale string) string {
	return "\t\t\t return import('./messages/compiled/" + locale + ".json') as unknown as IntlMessages;\n"
}

const SUPPORTED_LOCALES_FILE = "supportedLocales.ts"
