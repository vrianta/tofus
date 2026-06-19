package browser

type Charset string

const (
	def     Charset = ""
	uTF8    Charset = "utf-8"
	uTF16   Charset = "utf-16"
	iSO8859 Charset = "iso-8859-1"
)

// var Charset = uTF8

func (c *Charset) ToUTF8() Charset {
	return uTF8
}

func (c *Charset) ToUTF16() Charset {
	return uTF16
}

func (c *Charset) ToISO88591() Charset {
	return iSO8859
}

func SetChatset() Charset {
	return def
}
