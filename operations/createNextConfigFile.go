package operations

import (
	"sync"

	"superkoders.com/intl/utils"
)

func CreateNextConfigFile(wg *sync.WaitGroup, rootPath string) {
	defer wg.Done()
	utils.CreateRootFile("next.config.js", FILE_CONTENT, rootPath)
	utils.Success("✔ Edited next.config file")
}

const FILE_CONTENT = `const withTM = require('next-transpile-modules')(['@superkoders/vercajk']);

const localesData = require('./src/i18n/supportedLocales.js');
const { Locale: localesEnum, defaultLocale } = localesData;
const locales = Object.values(localesEnum);

console.warn('HINT: Consider add security headers, see https://nextjs.org/docs/advanced-features/security-headers'); // TODO

/** @type {import('next').NextConfig} */
const nextConfig = {
	reactStrictMode: true,

	i18n: {
		locales,
		defaultLocale,
		localeDetection: false,
	},
};

module.exports = withTM(nextConfig);`
