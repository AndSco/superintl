package operations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"superkoders.com/intl/utils"
)

type ResponseInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ProjectInfo struct {
	Id int `json:"id"`
}

type ProjectData struct {
	Project ProjectInfo `json:"project"`
}

type POEditorResponse struct {
	Response ResponseInfo `json:"response"`
	Result   ProjectData  `json:"result"`
}

func CreatePOEditorProject(projectName, POEditorToken string) {
	data := strings.NewReader(`api_token=` + POEditorToken + `&name=` + projectName)
	projectCreationResponse := sendRequestToPOEditor("https://api.poeditor.com/v2/projects/add", data)

	if requestWasSuccesful(projectCreationResponse) {
		addProjectLanguages(projectCreationResponse.Result.Project.Id, POEditorToken, []string{"it", "de", "fr"})
	}
}

func addProjectLanguages(projectId int, POEditorToken string, langCodes []string) {
	resultsChannel := make(chan POEditorResponse)
	var wg sync.WaitGroup
	wg.Add(len(langCodes))

	for _, code := range langCodes {
		go addLanguage(projectId, POEditorToken, code, resultsChannel, &wg)
	}

	go communicateResults(resultsChannel) 

	wg.Wait()
	close(resultsChannel)
	specifyDefaultLang(projectId, POEditorToken, "it")
}

func addLanguage(projectId int, POEditorToken string, languageCode string, resultsChannel chan POEditorResponse, wg *sync.WaitGroup) {
	defer wg.Done()

	stringifiedId := strconv.Itoa(projectId)
	data := strings.NewReader(`api_token=` + POEditorToken + `&id=` + stringifiedId + `&language=` + languageCode)

	languageAddedResponse := sendRequestToPOEditor("https://api.poeditor.com/v2/languages/add", data)
	resultsChannel <- languageAddedResponse
}

func communicateResults(ch chan POEditorResponse) {
	for result := range ch {
		fmt.Println(result.Response.Message)
	}
}

func specifyDefaultLang(projectId int, POEditorToken string, defaultLangCode string) {
	stringifiedId := strconv.Itoa(projectId)
	data := strings.NewReader(`api_token=` + POEditorToken + `&id=` + stringifiedId + `&reference_language=` + defaultLangCode)
	sendRequestToPOEditor("https://api.poeditor.com/v2/projects/update", data)

}

func sendRequestToPOEditor(url string, data *strings.Reader) POEditorResponse {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, data)
	utils.Check(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	utils.Check(err)
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)
	var parsedReponse POEditorResponse
	json.Unmarshal(bodyText, &parsedReponse)
	wasSuccessful := parsedReponse.Response.Code == "200"
	if !wasSuccessful {
		alertInsuccess()
	}
	return parsedReponse
}

func requestWasSuccesful(res POEditorResponse) bool {
	return res.Response.Code == "200"
}

func alertInsuccess() {
	utils.Alert("Problems encontered when creating a POEditor project! Create one yourself on https://poeditor.com")
}
