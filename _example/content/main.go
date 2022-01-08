package main

import (
	"context"
	"flag"
	"fmt"
	"ws/pkg/ws"
)

func main() {
	cid := flag.String("content-id", "", "content id")
	email := flag.String("email", "", "email")
	password := flag.String("password", "", "password")
	flag.Parse()

	c := ws.NewClient()
	auth := ws.AuthParams{
		EMail:    *email,
		PassWord: *password,
	}
	c.AuthService.SignIn(context.Background(), &auth)

	content := ws.Content{
		ID:         *cid,
		PDFPageNum: 5,
	}
	contentResp, err := c.ContentService.Update(context.Background(), &content)
	if err != nil {
		panic(err)
	}

	fmt.Printf("pdf_page_num: %d\n", contentResp.Content.PDFPageNum)
}
