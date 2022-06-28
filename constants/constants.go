package constants

import "os"

// PATHS
const INTL_FOLDER_PATH = "i18n"
const NEW_PROJECT_FOLDER_NAME = "my-superintl-app"

var PACKAGE_JSON_PATH = "/package.json"
var currentDir, _ = os.Getwd()
var ROOT_PATH = currentDir

// USER-INTENTIONS
const FROM_SCRATCH = "Set up a brand-new NextJS project with multiple languages"
const MAKE_PROJECT_MULTILANGUAGE = "Configure an existing NextJS project to support multiple languages (still to be implemented)"
const ADD_LANGUAGES = "Add extra-languages to an already-configured NextJS project (still to be implemented)"
const NO_NEXT = "Set up a non-Next multi-language React project (still to be implemented)"

// OTHER
const EMPTY_JSON = "{}"
const EMPTY_ARRAY = "[]"
