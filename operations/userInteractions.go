package operations

import (
	"fmt"
	"strings"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func AskUserIntention() string {
	userChoice := utils.ChooseOption("What do you want to do?", []string{constants.FROM_SCRATCH, constants.MAKE_PROJECT_MULTILANGUAGE, constants.ADD_LANGUAGES, constants.NO_NEXT})
	return userChoice
}

func AskProjectName() string {
	projectName := utils.GetInputOnSameLine("❓ What's your project's name (lowercase, no spacings)? ")
	if len(projectName) < 1 || len(strings.Fields(projectName)) > 1 || utils.HasUpperCase(projectName) {
		utils.Alert("⚠️ Enter a valid project name, all lowercase and without spacings!")
		return AskProjectName()
	}
	return projectName
}

func AskDefaultLocale() constants.Iso_Lang {
	defaultLocale := utils.GetInputOnSameLine("❓ What's your project's default locale (ex en-US, fr, de, cs)? ")
	lowerCasedLocale := strings.ToLower(defaultLocale)
	if !constants.IsSupportedLanguage(lowerCasedLocale) {
		alertInvalidLanguage(lowerCasedLocale)
		return AskDefaultLocale()
	}
	return constants.ISO_LANGS[lowerCasedLocale]
}

func AskOtherLocales(defaultLocale string) []constants.Iso_Lang {
	otherSupportedLocales := utils.GetInputOnSameLine("➡️ Enter the other project's locales, separated by comma: ")
	splittedOtherLocales := strings.Split(otherSupportedLocales, ",")

	parsedOtherLocales := make([]constants.Iso_Lang, len(splittedOtherLocales))

	for idx, locale := range splittedOtherLocales {
		trimmedLocale := strings.ToLower(strings.TrimSpace(locale))
		if !constants.IsSupportedLanguage(trimmedLocale) {
			alertInvalidLanguage(trimmedLocale)
			return AskOtherLocales(defaultLocale)
		}

		if trimmedLocale == defaultLocale {
			utils.Alert(fmt.Sprintf("'%s' is already the project's default locale! Don't repeat yourself.", trimmedLocale))
			return AskOtherLocales(defaultLocale)
		}

		parsedOtherLocales[idx] = constants.ISO_LANGS[trimmedLocale]
	}

	return parsedOtherLocales
}

func HandlePOEditorProjectCreation(projectName, defaultLangCode string, otherLangCodes []string) {
	wantsToCreatePOEditorProject := utils.GetInputOnSameLine("❓ Do you want to create a project on POEditor? y/n ")
	if (sanitizeInput(wantsToCreatePOEditorProject) == "y") {
		POEditorToken := utils.GetInputOnSameLine("Enter your POEditor api token (you can find it here https://poeditor.com/account/api): ")
		CreatePOEditorProject(projectName, sanitizeInput(POEditorToken), defaultLangCode, otherLangCodes)
	}
}

func alertInvalidLanguage(input string) {
	utils.Alert(fmt.Sprintf("'%s' is not a supported locale! Have a look at https://poeditor.com/docs/languages for a comprehensive list", input))
}

func sanitizeInput(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
