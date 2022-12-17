package v1

import (
	"DineEasy/utils/errmsg"
	"DineEasy/utils/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"time"
)

func Upload(c *gin.Context) (int, any) {
	file, fileHeader, err := c.Request.FormFile("file")
	class := c.Request.FormValue("class")
	if err != nil {
		code = errmsg.ERROR_FILE_WRONG
		return code, err
	}

	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	if class == "category" || class == "item" || class == "combo" {
		now := time.Now()
		uploadName := fmt.Sprintf("%s/%d%s", class, now.UnixNano(), path.Ext(fileName))
		url, code := upload.QiniuUpload(uploadName, file, fileSize)
		return code, url
	} else {
		code = errmsg.ERROR_CLASS_WRONG
		return code, nil
	}
}
