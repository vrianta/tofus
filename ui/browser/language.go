package browser

type Language string

const (
	Ldef     Language = ""
	English  Language = "en"
	Spanish  Language = "es"
	French   Language = "fr"
	German   Language = "de"
	Hindi    Language = "hi"
	Japanese Language = "ja"
)

func (Language) ToEnglish() Language {
	return English
}

func (Language) ToSpanish() Language {
	return Spanish
}

func (Language) ToFrench() Language {
	return French
}

func (Language) ToGerman() Language {
	return German
}

func (Language) ToHindi() Language {
	return Hindi
}

func (Language) ToJapanese() Language {
	return Japanese
}

func SetLanguage() Language {
	return Ldef
}
