package main

import (
	"fmt"
	"os"
	"sync"

	"superkoders.com/intl/constants"
	"superkoders.com/intl/operations"
	"superkoders.com/intl/utils"
)

func main() {
	userIntention := operations.AskUserIntention()
	srcFolderPath := ""
	rootPath := constants.ROOT_PATH
	projectName := ""

	if userIntention == constants.FROM_SCRATCH {
		projectName = operations.AskProjectName()
		operations.CloneNextStarterKitRepo(projectName)
		srcFolderPath = projectName + "/src"
		rootPath = rootPath + "/" + projectName
	} else {
		utils.Alert("This functionality is not implemented yet!")
		os.Exit(1)
	}

	defaultLocale := operations.AskDefaultLocale()
	otherLocales := operations.AskOtherLocales(defaultLocale.Code)
	var otherLocalesCodes []string = []string{}
	for _, loc := range otherLocales {
		otherLocalesCodes = append(otherLocalesCodes, loc.Code)
	}
	operations.HandlePOEditorProjectCreation(projectName, defaultLocale.Code, otherLocalesCodes)
	utils.Warn("đ Creating project...")
	operations.CreateFilesForIntlMessages(defaultLocale, otherLocales, srcFolderPath)
	var wg sync.WaitGroup

	wg.Add(13)
	go operations.CreateLocalesTsConfig(&wg, rootPath)
	go operations.CreateIntlHelperFunctions(&wg, srcFolderPath)
	go operations.CreatePackageJson(&wg, projectName, defaultLocale, rootPath)
	go operations.CreateSupportedLocalesFile(defaultLocale, otherLocales, &wg, srcFolderPath)
	go operations.CreateBabelrcFile(&wg, rootPath)
	go operations.CreateNextConfigFile(&wg, rootPath)
	go operations.AddToGitIgnore(&wg, rootPath)
	go operations.AddEnvFile(&wg, rootPath)
	go operations.CreateLocaleProvider(&wg, srcFolderPath)
	go operations.WrapAppInLocaleContext(&wg, srcFolderPath)
	go operations.EditIndexFile(&wg, srcFolderPath)
	go operations.CreateDummyLangPicker(&wg, srcFolderPath)
	go operations.AppendToReadMeFile(&wg, rootPath)
	wg.Wait()

	utils.Warn2(fmt.Sprintf("\n\nđ Your project is ready! To start working on it, run: \n\nâď¸ cd %s \nâď¸ npm install \nâď¸ npm run dev\n\n", projectName))
}

