package controllers

import (
	"github.com/revel/revel"
	"fmt"

//	"mime/multipart"
//	"strconv"
//	"time"
//	"io"
//	"os"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// file upload
func (c App) UpLoad() revel.Result {
	req := c.Request
	fmt.Printf("requset params =  %+v\n" , req)

	//file, handle , err:= req.FormFile("file")
	//if err!=nil {
	//	panic(err.Error())
	//}
	//
	//f, err := os.OpenFile("./upload/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	//io.Copy(f, file)
	//
	//if err!= nil {
	//	panic(err.Error())
	//}
	//
	//defer f.Close()
	//defer file.Close()
	return c.Render()
}
