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

type IKFetchDocumentType struct {
	Tid         int    `json:"tid"`
	PublishDate string `json:"publishdate"`
	Title       string `json:"title"`
	Doc         string `json:"doc"`
	DocSource   string `json:"docsource"`
	DivType     string `json:"divtype"`
	CourtCopy   bool   `json:"courtcopy"`
}
