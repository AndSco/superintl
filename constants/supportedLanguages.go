package constants

func IsSupportedLanguage(lang string) bool {
	_, ok := ISO_LANGS[lang]
	return ok
}

type Iso_Lang struct {
	Name       string
	NativeName string
	Code       string
}

var ISO_LANGS map[string]Iso_Lang = map[string]Iso_Lang{"ab": {Name: "Abkhazian", Code: "ab", NativeName: "Аҧсуа"}, "aa": {Name: "Afar", Code: "aa", NativeName: "Afaraf"}, "af": {Name: "Afrikaans", Code: "af", NativeName: "Afrikaans"}, "ak": {Name: "Akan", Code: "ak", NativeName: "Akan"}, "sq": {Name: "Albanian", Code: "sq", NativeName: "Shqip"}, "am": {Name: "Amharic", Code: "am", NativeName: "አማርኛ"}, "ar": {Name: "Arabic", Code: "ar", NativeName: "العربية"}, "ar-ae": {Name: "Arabic (AE)", Code: "ar-ae", NativeName: "Arabic (AE)"}, "ar-bh": {Name: "Arabic (BH)", Code: "ar-bh", NativeName: "Arabic (BH)"}, "ar-dz": {Name: "Arabic (DZ)", Code: "ar-dz", NativeName: "Arabic (DZ)"}, "ar-eg": {Name: "Arabic (EG)", Code: "ar-eg", NativeName: "Arabic (EG)"}, "ar-iq": {Name: "Arabic (IQ)", Code: "ar-iq", NativeName: "Arabic (IQ)"}, "ar-jo": {Name: "Arabic (JO)", Code: "ar-jo", NativeName: "Arabic (JO)"}, "ar-kw": {Name: "Arabic (KW)", Code: "ar-kw", NativeName: "Arabic (KW)"}, "ar-lb": {Name: "Arabic (LB)", Code: "ar-lb", NativeName: "Arabic (LB)"}, "ar-ly": {Name: "Arabic (LY)", Code: "ar-ly", NativeName: "Arabic (LY)"}, "ar-ma": {Name: "Arabic (MA)", Code: "ar-ma", NativeName: "Arabic (MA)"}, "ar-om": {Name: "Arabic (OM)", Code: "ar-om", NativeName: "Arabic (OM)"}, "ar-qa": {Name: "Arabic (QA)", Code: "ar-qa", NativeName: "Arabic (QA)"}, "ar-sa": {Name: "Arabic (SA)", Code: "ar-sa", NativeName: "Arabic (SA)"}, "ar-sy": {Name: "Arabic (SY)", Code: "ar-sy", NativeName: "Arabic (SY)"}, "ar-tn": {Name: "Arabic (TN)", Code: "ar-tn", NativeName: "Arabic (TN)"}, "ar-ye": {Name: "Arabic (YE)", Code: "ar-ye", NativeName: "Arabic (YE)"}, "an": {Name: "Aragonese", Code: "an", NativeName: "Aragonés"}, "hy": {Name: "Armenian", Code: "hy", NativeName: "Հայերեն"}, "as": {Name: "Assamese", Code: "as", NativeName: "অসমীয়া"}, "ast": {Name: "Asturian", Code: "ast", NativeName: "Asturian"}, "av": {Name: "Avaric", Code: "av", NativeName: "Авар мацӀ"}, "ae": {Name: "Avestan", Code: "ae", NativeName: "Avesta"}, "ay": {Name: "Aymara", Code: "ay", NativeName: "Aymar aru"}, "az": {Name: "Azerbaijani", Code: "az", NativeName: "Azərbaycan dili"}, "az-cyrl": {Name: "Azerbaijani (Cyrillic)", Code: "az-cyrl", NativeName: "Azerbaijani (Cyrillic)"}, "bm": {Name: "Bambara", Code: "bm", NativeName: "Bamanankan"}, "ba": {Name: "Bashkir", Code: "ba", NativeName: "Башҡорт теле"}, "eu": {Name: "Basque", Code: "eu", NativeName: "Euskara"}, "be": {Name: "Belarusian", Code: "be", NativeName: "Беларуская"}, "bn": {Name: "Bengali", Code: "bn", NativeName: "বাংলা"}, "bh": {Name: "Bihari languages", Code: "bh", NativeName: "भोजपुरी"}, "bi": {Name: "Bislama", Code: "bi", NativeName: "Bislama"}, "bs": {Name: "Bosnian", Code: "bs", NativeName: "Bosanski jezik"}, "br": {Name: "Breton", Code: "br", NativeName: "Brezhoneg"}, "bg": {Name: "Bulgarian", Code: "bg", NativeName: "Български език"}, "my": {Name: "Burmese", Code: "my", NativeName: "ဗမာစာ"}, "ca": {Name: "Catalan", Code: "ca", NativeName: "Català"}, "ceb": {Name: "Cebuano", Code: "ceb", NativeName: "Cebuano"}, "ch": {Name: "Chamorro", Code: "ch", NativeName: "Chamoru"}, "ce": {Name: "Chechen", Code: "ce", NativeName: "Нохчийн мотт"}, "ny": {Name: "Chichewa", Code: "ny", NativeName: "ChiCheŵa"}, "zh-cn": {Name: "Chinese", Code: "zh-cn", NativeName: "Chinese"}, "zh-hk": {Name: "Chinese (HK)", Code: "zh-hk", NativeName: "Chinese (HK)"}, "zh-mo": {Name: "Chinese (MO)", Code: "zh-mo", NativeName: "Chinese (MO)"}, "zh-sg": {Name: "Chinese (SG)", Code: "zh-sg", NativeName: "Chinese (SG)"}, "zh-hans": {Name: "Chinese (simplified)", Code: "zh-hans", NativeName: "Chinese (simplified)"}, "zh-hant": {Name: "Chinese (traditional)", Code: "zh-hant", NativeName: "Chinese (traditional)"}, "zh-tw": {Name: "Chinese (TW)", Code: "zh-tw", NativeName: "Chinese (TW)"}, "cu": {Name: "Church Slavic", Code: "cu", NativeName: "Ѩзыкъ словѣньскъ"}, "cv": {Name: "Chuvash", Code: "cv", NativeName: "Чӑваш чӗлхи"}, "kw": {Name: "Cornish", Code: "kw", NativeName: "Kernewek"}, "co": {Name: "Corsican", Code: "co", NativeName: "Corsu"}, "cr": {Name: "Cree", Code: "cr", NativeName: "ᓀᐦᐃᔭᐍᐏᐣ"}, "hr": {Name: "Croatian", Code: "hr", NativeName: "Hrvatski"}, "cs": {Name: "Czech", Code: "cs", NativeName: "Česky"}, "da": {Name: "Danish", Code: "da", NativeName: "Dansk"}, "prs": {Name: "Dari", Code: "prs", NativeName: "Dari"}, "dv": {Name: "Divehi", Code: "dv", NativeName: "ދިވެހި"}, "nl": {Name: "Dutch", Code: "nl", NativeName: "Nederlands"}, "dz": {Name: "Dzongkha", Code: "dz", NativeName: "Dzongkha"}, "en": {Name: "English", Code: "en", NativeName: "English"}, "en-au": {Name: "English (AU)", Code: "en-au", NativeName: "English (AU)"}, "en-bz": {Name: "English (BZ)", Code: "en-bz", NativeName: "English (BZ)"}, "en-ca": {Name: "English (CA)", Code: "en-ca", NativeName: "English (CA)"}, "en-gh": {Name: "English (GH)", Code: "en-gh", NativeName: "English (GH)"}, "en-hk": {Name: "English (HK)", Code: "en-hk", NativeName: "English (HK)"}, "en-ie": {Name: "English (IE)", Code: "en-ie", NativeName: "English (IE)"}, "en-in": {Name: "English (IN)", Code: "en-in", NativeName: "English (IN)"}, "en-jm": {Name: "English (JM)", Code: "en-jm", NativeName: "English (JM)"}, "en-ke": {Name: "English (KE)", Code: "en-ke", NativeName: "English (KE)"}, "en-mu": {Name: "English (MU)", Code: "en-mu", NativeName: "English (MU)"}, "en-ng": {Name: "English (NG)", Code: "en-ng", NativeName: "English (NG)"}, "en-nz": {Name: "English (NZ)", Code: "en-nz", NativeName: "English (NZ)"}, "en-ph": {Name: "English (PH)", Code: "en-ph", NativeName: "English (PH)"}, "en-sg": {Name: "English (SG)", Code: "en-sg", NativeName: "English (SG)"}, "en-tt": {Name: "English (TT)", Code: "en-tt", NativeName: "English (TT)"}, "en-gb": {Name: "English (UK)", Code: "en-gb", NativeName: "English (UK)"}, "en-us": {Name: "English (US)", Code: "en-us", NativeName: "English (US)"}, "en-za": {Name: "English (ZA)", Code: "en-za", NativeName: "English (ZA)"}, "en-zw": {Name: "English (ZW)", Code: "en-zw", NativeName: "English (ZW)"}, "eo": {Name: "Esperanto", Code: "eo", NativeName: "Esperanto"}, "et": {Name: "Estonian", Code: "et", NativeName: "Eesti"}, "ee": {Name: "Ewe", Code: "ee", NativeName: "Eʋegbe"}, "fo": {Name: "Faroese", Code: "fo", NativeName: "Føroyskt"}, "fj": {Name: "Fijian", Code: "fj", NativeName: "Vosa Vakaviti"}, "fil": {Name: "Filipino", Code: "fil", NativeName: "Filipino"}, "fi": {Name: "Finnish", Code: "fi", NativeName: "Suomi"}, "nl-be": {Name: "Flemish", Code: "nl-be", NativeName: "Flemish"}, "fr": {Name: "French", Code: "fr", NativeName: "Français"}, "fr-be": {Name: "French (BE)", Code: "fr-be", NativeName: "French (BE)"}, "fr-ca": {Name: "French (CA)", Code: "fr-ca", NativeName: "French (CA)"}, "fr-ch": {Name: "French (CH)", Code: "fr-ch", NativeName: "French (CH)"}, "fr-lu": {Name: "French (LU)", Code: "fr-lu", NativeName: "French (LU)"}, "fr-mc": {Name: "French (MC)", Code: "fr-mc", NativeName: "French (MC)"}, "ff": {Name: "Fulah", Code: "ff", NativeName: "Fulfulde"}, "gl": {Name: "Galician", Code: "gl", NativeName: "Galego"}, "lg": {Name: "Ganda", Code: "lg", NativeName: "Luganda"}, "ka": {Name: "Georgian", Code: "ka", NativeName: "Ქართული"}, "de": {Name: "German", Code: "de", NativeName: "Deutsch"}, "de-at": {Name: "German (AT)", Code: "de-at", NativeName: "German (AT)"}, "de-be": {Name: "German (BE)", Code: "de-be", NativeName: "German (BE)"}, "de-ch": {Name: "German (CH)", Code: "de-ch", NativeName: "German (CH)"}, "de-li": {Name: "German (LI)", Code: "de-li", NativeName: "German (LI)"}, "de-lu": {Name: "German (LU)", Code: "de-lu", NativeName: "German (LU)"}, "el": {Name: "Greek", Code: "el", NativeName: "Ελληνικά"}, "gn": {Name: "Guarani", Code: "gn", NativeName: "Avañeẽ"}, "gu": {Name: "Gujarati", Code: "gu", NativeName: "ગુજરાતી"}, "ht": {Name: "Haitian Creole", Code: "ht", NativeName: "Kreyòl ayisyen"}, "ha": {Name: "Hausa", Code: "ha", NativeName: "Hausa"}, "haw": {Name: "Hawaiian", Code: "haw", NativeName: "Hawaiian"}, "he": {Name: "Hebrew", Code: "he", NativeName: "עברית"}, "hz": {Name: "Herero", Code: "hz", NativeName: "Otjiherero"}, "hi": {Name: "Hindi", Code: "hi", NativeName: "हिन्दी"}, "ho": {Name: "Hiri Motu", Code: "ho", NativeName: "Hiri Motu"}, "hu": {Name: "Hungarian", Code: "hu", NativeName: "Magyar"}, "is": {Name: "Icelandic", Code: "is", NativeName: "Íslenska"}, "io": {Name: "Ido", Code: "io", NativeName: "Ido"}, "ig": {Name: "Igbo", Code: "ig", NativeName: "Asụsụ Igbo"}, "id": {Name: "Indonesian", Code: "id", NativeName: "Bahasa Indonesia"}, "ia": {Name: "Interlingua", Code: "ia", NativeName: "Interlingua"}, "ie": {Name: "Interlingue", Code: "ie", NativeName: "Originally called Occidental; then Interlingue after WWII"}, "iu": {Name: "Inuktitut", Code: "iu", NativeName: "ᐃᓄᒃᑎᑐᑦ"}, "ik": {Name: "Inupiaq", Code: "ik", NativeName: "Iñupiaq"}, "ga": {Name: "Irish", Code: "ga", NativeName: "Gaeilge"}, "it": {Name: "Italian", Code: "it", NativeName: "Italiano"}, "it-ch": {Name: "Italian (CH)", Code: "it-ch", NativeName: "Italian (CH)"}, "jam": {Name: "Jamaican Patois", Code: "jam", NativeName: "Jamaican Patois"}, "ja": {Name: "Japanese", Code: "ja", NativeName: "日本語 (にほんご／にっぽんご)"}, "jv": {Name: "Javanese", Code: "jv", NativeName: "Basa Jawa"}, "kab": {Name: "Kabyle", Code: "kab", NativeName: "Kabyle"}, "kl": {Name: "Kalaallisut", Code: "kl", NativeName: "Kalaallisut"}, "kn": {Name: "Kannada", Code: "kn", NativeName: "ಕನ್ನಡ"}, "kr": {Name: "Kanuri", Code: "kr", NativeName: "Kanuri"}, "ks": {Name: "Kashmiri", Code: "ks", NativeName: "कश्मीरी"}, "kk": {Name: "Kazakh", Code: "kk", NativeName: "Қазақ тілі"}, "km": {Name: "Khmer", Code: "km", NativeName: "ភាសាខ្មែរ"}, "ki": {Name: "Kikuyu; Gikuyu", Code: "ki", NativeName: "Gĩkũyũ"}, "rw": {Name: "Kinyarwanda", Code: "rw", NativeName: "Ikinyarwanda"}, "ky": {Name: "Kirghiz", Code: "ky", NativeName: "Кыргыз тили"}, "kv": {Name: "Komi", Code: "kv", NativeName: "Коми кыв"}, "kg": {Name: "Kongo", Code: "kg", NativeName: "KiKongo"}, "ko": {Name: "Korean", Code: "ko", NativeName: "한국어 (韓國語)"}, "kj": {Name: "Kuanyama; Kwanyama", Code: "kj", NativeName: "Kuanyama"}, "ku": {Name: "Kurdish", Code: "ku", NativeName: "Kurdî"}, "lo": {Name: "Lao", Code: "lo", NativeName: "ພາສາລາວ"}, "la": {Name: "Latin", Code: "la", NativeName: "Latine"}, "lv": {Name: "Latvian", Code: "lv", NativeName: "Latviešu valoda"}, "li": {Name: "Limburgish", Code: "li", NativeName: "Limburgs"}, "ln": {Name: "Lingala", Code: "ln", NativeName: "Lingála"}, "lt": {Name: "Lithuanian", Code: "lt", NativeName: "Lietuvių kalba"}, "jbo": {Name: "Lojban", Code: "jbo", NativeName: "Lojban"}, "lu": {Name: "Luba-Katanga", Code: "lu", NativeName: ""}, "lb": {Name: "Luxembourgish", Code: "lb", NativeName: "Lëtzebuergesch"}, "mk": {Name: "Macedonian", Code: "mk", NativeName: "Македонски јазик"}, "mg": {Name: "Malagasy", Code: "mg", NativeName: "Malagasy fiteny"}, "ms": {Name: "Malay", Code: "ms", NativeName: "Bahasa Melayu"}, "ms-bn": {Name: "Malay (BN)", Code: "ms-bn", NativeName: "Malay (BN)"}, "ml": {Name: "Malayalam", Code: "ml", NativeName: "മലയാളം"}, "mt": {Name: "Maltese", Code: "mt", NativeName: "Malti"}, "gv": {Name: "Manx", Code: "gv", NativeName: "Gaelg"}, "mi": {Name: "Maori", Code: "mi", NativeName: "Te reo Māori"}, "mr": {Name: "Marathi", Code: "mr", NativeName: "मराठी"}, "mh": {Name: "Marshallese", Code: "mh", NativeName: "Kajin M̧ajeļ"}, "mo": {Name: "Moldavian; Moldovan", Code: "mo", NativeName: "Moldavian; Moldovan"}, "mn": {Name: "Mongolian", Code: "mn", NativeName: "Монгол"}, "me-me": {Name: "Montenegrin", Code: "me-me", NativeName: "Montenegrin"}, "me-cyrl": {Name: "Montenegrin (Cyrillic)", Code: "me-cyrl", NativeName: "Montenegrin (Cyrillic)"}, "na": {Name: "Nauru", Code: "na", NativeName: "Ekakairũ Naoero"}, "nv": {Name: "Navajo; Navaho", Code: "nv", NativeName: "Diné bizaad"}, "ng": {Name: "Ndonga", Code: "ng", NativeName: "Owambo"}, "ne": {Name: "Nepali", Code: "ne", NativeName: "नेपाली"}, "nd": {Name: "North Ndebele", Code: "nd", NativeName: "IsiNdebele"}, "se": {Name: "Northern Sami", Code: "se", NativeName: "Davvisámegiella"}, "no": {Name: "Norwegian", Code: "no", NativeName: "Norsk"}, "nb": {Name: "Norwegian Bokmål", Code: "nb", NativeName: "Norsk bokmål"}, "nn": {Name: "Norwegian Nynorsk", Code: "nn", NativeName: "Norsk nynorsk"}, "oc": {Name: "Occitan", Code: "oc", NativeName: "Occitan"}, "oj": {Name: "Ojibwa", Code: "oj", NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ"}, "or": {Name: "Oriya", Code: "or", NativeName: "ଓଡ଼ିଆ"}, "om": {Name: "Oromo", Code: "om", NativeName: "Afaan Oromoo"}, "os": {Name: "Ossetian; Ossetic", Code: "os", NativeName: "Ирон æвзаг"}, "pi": {Name: "Pali", Code: "pi", NativeName: "पाऴि"}, "pa": {Name: "Panjabi; Punjabi", Code: "pa", NativeName: "ਪੰਜਾਬੀ"}, "fa": {Name: "Persian", Code: "fa", NativeName: "فارسی"}, "pl": {Name: "Polish", Code: "pl", NativeName: "Polski"}, "pt": {Name: "Portuguese", Code: "pt", NativeName: "Português"}, "pt-br": {Name: "Portuguese (BR)", Code: "pt-br", NativeName: "Portuguese (BR)"}, "ps": {Name: "Pushto; Pashto", Code: "ps", NativeName: "پښتو"}, "qu": {Name: "Quechua", Code: "qu", NativeName: "Runa Simi"}, "rom": {Name: "Romani", Code: "rom", NativeName: "Romani"}, "ro": {Name: "Romanian", Code: "ro", NativeName: "Română"}, "rm": {Name: "Romansh", Code: "rm", NativeName: "Rumantsch grischun"}, "rn": {Name: "Rundi", Code: "rn", NativeName: "KiRundi"}, "ru": {Name: "Russian", Code: "ru", NativeName: "Русский язык"}, "ry": {Name: "Rusyn", Code: "ry", NativeName: "Rusyn"}, "sm": {Name: "Samoan", Code: "sm", NativeName: "Gagana faa Samoa"}, "sg": {Name: "Sango", Code: "sg", NativeName: "Yângâ tî sängö"}, "sa": {Name: "Sanskrit", Code: "sa", NativeName: "संस्कृतम्"}, "sat": {Name: "Santali", Code: "sat", NativeName: "Santali"}, "sc": {Name: "Sardinian", Code: "sc", NativeName: "Sardu"}, "gd": {Name: "Scottish Gaelic", Code: "gd", NativeName: "Gàidhlig"}, "sr": {Name: "Serbian", Code: "sr", NativeName: "Српски језик"}, "sr-cyrl": {Name: "Serbian (Cyrillic)", Code: "sr-cyrl", NativeName: "Serbian (Cyrillic)"}, "sh": {Name: "Serbo-Croatian", Code: "sh", NativeName: "Serbo-Croatian"}, "sn": {Name: "Shona", Code: "sn", NativeName: "ChiShona"}, "ii": {Name: "Sichuan Yi", Code: "ii", NativeName: "ꆈꌠ꒿ Nuosuhxop"}, "scn": {Name: "Sicilian", Code: "scn", NativeName: "Sicilian"}, "sd": {Name: "Sindhi", Code: "sd", NativeName: "सिन्धी"}, "si": {Name: "Sinhalese", Code: "si", NativeName: "සිංහල"}, "sk": {Name: "Slovak", Code: "sk", NativeName: "Slovenčina"}, "sl": {Name: "Slovenian", Code: "sl", NativeName: "Slovenščina"}, "so": {Name: "Somali", Code: "so", NativeName: "Soomaaliga"}, "st": {Name: "Sotho, Southern", Code: "st", NativeName: "Sesotho"}, "nr": {Name: "South Ndebele", Code: "nr", NativeName: "IsiNdebele"}, "es": {Name: "Spanish", Code: "es", NativeName: "Español"}, "es-ar": {Name: "Spanish (AR)", Code: "es-ar", NativeName: "Spanish (AR)"}, "es-bo": {Name: "Spanish (BO)", Code: "es-bo", NativeName: "Spanish (BO)"}, "es-cl": {Name: "Spanish (CL)", Code: "es-cl", NativeName: "Spanish (CL)"}, "es-co": {Name: "Spanish (CO)", Code: "es-co", NativeName: "Spanish (CO)"}, "es-cr": {Name: "Spanish (CR)", Code: "es-cr", NativeName: "Spanish (CR)"}, "es-do": {Name: "Spanish (DO)", Code: "es-do", NativeName: "Spanish (DO)"}, "es-ec": {Name: "Spanish (EC)", Code: "es-ec", NativeName: "Spanish (EC)"}, "es-gt": {Name: "Spanish (GT)", Code: "es-gt", NativeName: "Spanish (GT)"}, "es-hn": {Name: "Spanish (HN)", Code: "es-hn", NativeName: "Spanish (HN)"}, "es-419": {Name: "Spanish (LA & C)", Code: "es-419", NativeName: "Spanish (LA & C)"}, "es-mx": {Name: "Spanish (MX)", Code: "es-mx", NativeName: "Spanish (MX)"}, "es-ni": {Name: "Spanish (NI)", Code: "es-ni", NativeName: "Spanish (NI)"}, "es-pa": {Name: "Spanish (PA)", Code: "es-pa", NativeName: "Spanish (PA)"}, "es-pe": {Name: "Spanish (PE)", Code: "es-pe", NativeName: "Spanish (PE)"}, "es-pr": {Name: "Spanish (PR)", Code: "es-pr", NativeName: "Spanish (PR)"}, "es-py": {Name: "Spanish (PY)", Code: "es-py", NativeName: "Spanish (PY)"}, "es-sv": {Name: "Spanish (SV)", Code: "es-sv", NativeName: "Spanish (SV)"}, "es-uy": {Name: "Spanish (UY)", Code: "es-uy", NativeName: "Spanish (UY)"}, "es-ve": {Name: "Spanish (VE)", Code: "es-ve", NativeName: "Spanish (VE)"}, "su": {Name: "Sundanese", Code: "su", NativeName: "Basa Sunda"}, "sw": {Name: "Swahili", Code: "sw", NativeName: "Kiswahili"}, "ss": {Name: "Swati", Code: "ss", NativeName: "SiSwati"}, "sv": {Name: "Swedish", Code: "sv", NativeName: "Svenska"}, "sv-fi": {Name: "Swedish (FI)", Code: "sv-fi", NativeName: "Swedish (FI)"}, "tl": {Name: "Tagalog", Code: "tl", NativeName: "Wikang Tagalog"}, "ty": {Name: "Tahitian", Code: "ty", NativeName: "Reo Tahiti"}, "tg": {Name: "Tajik", Code: "tg", NativeName: "Тоҷикӣ"}, "ta": {Name: "Tamil", Code: "ta", NativeName: "தமிழ்"}, "tt": {Name: "Tatar", Code: "tt", NativeName: "Татарча"}, "te": {Name: "Telugu", Code: "te", NativeName: "తెలుగు"}, "th": {Name: "Thai", Code: "th", NativeName: "ไทย"}, "bo": {Name: "Tibetan", Code: "bo", NativeName: "བོད་ཡིག"}, "ti": {Name: "Tigrinya", Code: "ti", NativeName: "ትግርኛ"}, "to": {Name: "Tonga", Code: "to", NativeName: "Faka Tonga"}, "ts": {Name: "Tsonga", Code: "ts", NativeName: "Xitsonga"}, "tn": {Name: "Tswana", Code: "tn", NativeName: "Setswana"}, "tr": {Name: "Turkish", Code: "tr", NativeName: "Türkçe"}, "tk": {Name: "Turkmen", Code: "tk", NativeName: "Türkmen"}, "tw": {Name: "Twi", Code: "tw", NativeName: "Twi"}, "uk": {Name: "Ukrainian", Code: "uk", NativeName: "Українська"}, "ur": {Name: "Urdu", Code: "ur", NativeName: "اردو"}, "ug": {Name: "Uyghur", Code: "ug", NativeName: "Uyƣurqə"}, "uz": {Name: "Uzbek", Code: "uz", NativeName: "Zbek"}, "uz-cyrl": {Name: "Uzbek (Cyrillic)", Code: "uz-cyrl", NativeName: "Uzbek (Cyrillic)"}, "ve": {Name: "Venda", Code: "ve", NativeName: "Tshivenḓa"}, "vi": {Name: "Vietnamese", Code: "vi", NativeName: "Tiếng Việt"}, "vo": {Name: "Volapük", Code: "vo", NativeName: "Volapük"}, "wa": {Name: "Walloon", Code: "wa", NativeName: "Walon"}, "cy": {Name: "Welsh", Code: "cy", NativeName: "Cymraeg"}, "fy": {Name: "Western Frisian", Code: "fy", NativeName: "Frysk"}, "wo": {Name: "Wolof", Code: "wo", NativeName: "Wollof"}, "xh": {Name: "Xhosa", Code: "xh", NativeName: "IsiXhosa"}, "yi": {Name: "Yiddish", Code: "yi", NativeName: "ייִדיש"}, "yo": {Name: "Yoruba", Code: "yo", NativeName: "Yorùbá"}, "za": {Name: "Zhuang; Chuang", Code: "za", NativeName: "Saɯ cueŋƅ"}, "zu": {Name: "Zulu", Code: "zu", NativeName: "Zulu"}}