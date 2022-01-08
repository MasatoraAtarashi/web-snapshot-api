package ws

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type ContentService interface {
	Update(ctx context.Context, content *Content) (*ContentResponse, error)
}

type contentService struct {
	cli *Client
}

func (s *contentService) Update(ctx context.Context, content *Content) (*ContentResponse, error) {
	path := fmt.Sprintf("content/%s", content.ID)
	data := url.Values{}
	pn := strconv.Itoa(content.PDFPageNum)
	data.Set("pdf_page_num", pn)

	var c ContentResponse
	if err := s.cli.putForm(ctx, path, data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}
