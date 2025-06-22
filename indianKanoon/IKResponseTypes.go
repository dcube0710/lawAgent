package indianKanoon

type IKSearchResponse struct {
	Docs []IKSearchDocumentType `json:"docs"`
}

type IKSearchDocumentType struct {
	Tid         int    `json:"tid"`
	DocType     int    `json:"doctype"`
	PublishDate string `json:"publishdate"`
	DocSize     int    `json:"docsize"`
	Headline    string `json:"headline"`
}
