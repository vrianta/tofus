package browser

type Charset string

const (
	Def     Charset = ""
	UTF8    Charset = "utf-8"
	UTF16   Charset = "utf-16"
	ISO8859 Charset = "iso-8859-1"
)

// var Charset = UTF8

func (Charset) ToUTF8() Charset {
	return UTF8
}

func (Charset) ToUTF16() Charset {
	return UTF16
}

func (Charset) ToISO88591() Charset {
	return ISO8859
}

func SetChatset() Charset {
	return Def
}
