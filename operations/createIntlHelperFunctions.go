package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

type JsFunction struct {
	Filename string
	content  string
}

var JS_FUNCTIONS = []JsFunction{
	{"download.mjs", DOWNLOAD_FUNCTION},
	{"upload.mjs", UPLOAD_FUNCTION},
	{"format.js", FORMAT_FUNCTION},
	{"compile.js", COMPILE_FUNCTION},
}

func CreateIntlHelperFunctions(wg *sync.WaitGroup, srcFolderPath string) {
	defer wg.Done()
	folderPath := utils.BuildIntlPath(srcFolderPath) + "/utils"
	utils.CreateFolder(folderPath)
	for _, function := range JS_FUNCTIONS {
		utils.WriteToFile(folderPath+"/"+function.Filename, function.content)
	}
	utils.Success("✔ Created utility functions to interact with POEditor")
}

const UPLOAD_FUNCTION = `
import fetch, { FormData, fileFromSync } from 'node-fetch';
import dotenv from 'dotenv';
import localesData from '../supportedLocales.js';

const { defaultLocale } = localesData;

dotenv.config();

if (!process.env.POEDITOR_TOKEN && !process.env.POEDITOR_PROJECT_ID) {
	console.error('Error: You need to specity POEDITOR_TOKEN and POEDITOR_PROJECT_ID environment variables in .env file!');
} else {
	console.log("Uploading default language " + defaultLocale + "...");
	const form = new FormData();
	const blob = fileFromSync("./src/i18n/messages/extracted/" + defaultLocale + ".json");

	form.append('api_token', process.env.POEDITOR_TOKEN);
	form.append('id', process.env.POEDITOR_PROJECT_ID);
	form.append('updating', 'terms_translations');
	form.append('language', defaultLocale);
	form.append('sync_terms', '1');
	form.append('file', blob);

	fetch('https://api.poeditor.com/v2/projects/upload', {
		method: 'post',
		body: form
	})
		.then((res) => res.json())
		.then((data) => console.log(data))
		.catch((error) => console.error(error));
}
`

const DOWNLOAD_FUNCTION = `
import fetch, { FormData } from 'node-fetch';
import { createWriteStream } from 'fs';
import { pipeline } from 'stream';
import { promisify } from 'util';
import dotenv from 'dotenv';
import localesData from '../supportedLocales.js';

const { Locale: localesEnum } = localesData;
const locales = Object.values(localesEnum);

const streamPipeline = promisify(pipeline);
dotenv.config();

if (!process.env.POEDITOR_TOKEN && !process.env.POEDITOR_PROJECT_ID) {
	console.error('Error: You need to specity POEDITOR_TOKEN and POEDITOR_PROJECT_ID environment variables in .env file!');
} else {
	locales.forEach(async (locale) => {
		console.log('Downloading language ' + locale + '...');
		const form = new FormData();
		form.append('api_token', process.env.POEDITOR_TOKEN);
		form.append('id', process.env.POEDITOR_PROJECT_ID);
		form.append('type', 'json');
		form.append('language', locale);
		form.append('filters', ['translated']);

		const res = await fetch('https://api.poeditor.com/v2/projects/export', {
			method: 'post',
			body: form
		});

		const json = await res.json();
		if (json.result && json.result.url) {
			const response = await fetch(json.result.url);
			if (!response.ok) throw new Error('Unexpected response: ' + response.statusText);
			await streamPipeline(response.body, createWriteStream('./src/i18n/messages/translated/' + locale + '.json'));

			console.log('Downloading language ' + locale + '...done');
		} else {
			console.log(json);
		}
	});
}`

const FORMAT_FUNCTION = `
exports.format = function (msgs) {
	return Object.entries(msgs).map(([id, msg]) => {
		return {
			term: id,
			definition: msg.defaultMessage,
			context: msg.description,
		};
	});
};
`

const COMPILE_FUNCTION = `
exports.compile = function (msgs) {
	return msgs.reduce((res, curr) => {
		res[curr.term] = curr.definition;
		return res;
	}, {});
};
`
