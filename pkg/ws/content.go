package ws

// Content represents a content user bookmarked.
// NOTE: 最低限のフィールドしか定義していません。
type Content struct {
	ID         string `json:"int"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	PDFPageNum int    `json:"pdf_page_num"`
}

type ContentResponse struct {
	Content Content `json:"data"`
}
