package operations

import (
	"superkoders.com/intl/constants"
	"superkoders.com/intl/utils"
)

func CreateFilesForIntlMessages(defaultLocale constants.Iso_Lang, otherSupportedLocales []constants.Iso_Lang, srcFolderPath string) {
	foldersNeeded := []string{"extracted", "translated", "compiled"}
	allSupportedLocales := append(otherSupportedLocales, defaultLocale)
	for _, folderName := range foldersNeeded {
		fileContent := constants.EMPTY_ARRAY
		if folderName == "compiled" {
			fileContent = constants.EMPTY_JSON
		}
		folderPath := utils.BuildIntlPath(srcFolderPath) + "/messages/" + folderName
		utils.CreateFolder(folderPath)

		if folderName == "extracted" {
			utils.WriteToFile(folderPath+"/"+defaultLocale.Code+".json", fileContent)
		} else {
			for _, lang := range allSupportedLocales {
				utils.WriteToFile(folderPath+"/"+lang.Code+".json", fileContent)
			}
		}
	}
	utils.Success("✔ Created folders & files for intl messages")
}
