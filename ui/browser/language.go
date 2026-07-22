package browser

type Language string

const (
	ldef     Language = ""
	english  Language = "en"
	spanish  Language = "es"
	french   Language = "fr"
	german   Language = "de"
	hindi    Language = "hi"
	japanese Language = "ja"
)

func (Language) ToEnglish() Language {
	return english
}

func (Language) ToSpanish() Language {
	return spanish
}

func (Language) ToFrench() Language {
	return french
}

func (Language) ToGerman() Language {
	return german
}

func (Language) ToHindi() Language {
	return hindi
}

func (Language) ToJapanese() Language {
	return japanese
}

func SetLanguage() Language {
	return ldef
}
