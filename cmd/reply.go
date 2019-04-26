package cmd

import (
	"fmt"
)

const (
	InitResponse      = `{"status": "Success","capabilities":{"attach":True}}`
	AttachResponse    = `{status: Success, device:"%s"}`
	DetachResponse    = `{"status": "Success"}`
	IsAttachResponse  = `{"status": "Success", "attached":true}`
	FailResponse      = `{"status":"Failure","message":"%s"}`
	SuccessResponse   = `{"status": "Success"}`
	UnsupportResponse = `{"status": "Not supported"}`
)

func ReplyStr(response string) error {
	fmt.Println(response)
	return nil
}

func ReplyError(err error) error {
	fmt.Printf(FailResponse, err.Error())
	return nil
}
