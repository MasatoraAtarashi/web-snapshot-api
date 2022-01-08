package main

import (
	"context"
	"flag"
	"fmt"
	"ws/pkg/ws"
)

func main() {
	cid := flag.String("content-id", "", "content id")
	token := flag.String("access-token", "", "access-token")
	client := flag.String("client", "", "client")
	uid := flag.String("uid", "", "uid")
	flag.Parse()

	content := ws.Content{
		ID:         *cid,
		PDFPageNum: 5,
	}

	c := ws.NewClient(*token, *client, *uid)
	contentResp, err := c.ContentService.Update(context.Background(), &content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("pdf_page_num: %d", contentResp.Content.PDFPageNum)
}
