package api

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	//code := e.SUCCESS
	//data := make(map[string]string)
	//file, image, err := c.Request.FormFile("image")
	//if err != nil {
	//	logging.Warn(err)
	//	code = e.ERROR
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": code,
	//		"msg":  e.GetMsg(code),
	//		"data": data,
	//	})
	//}
	//
	//if image == nil {
	//	code = e.INVALID_PARAMS
	//} else {
	//	imageName := upload.GetImageName(image.Filename)
	//	fullPath := upload.GetImageFullPath()
	//	savePath := upload.GetImagePath()
	//
	//	src := fullPath + imageName
	//
	//	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
	//		code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
	//	} else {
	//		err = upload.CheckImage(fullPath)
	//		if err != nil {
	//			logging.Warn(err)
	//			code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
	//		} else if err = c.SaveUploadedFile(image, src); err != nil {
	//			logging.Warn(err)
	//			code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
	//		} else {
	//			data["image_url"] = upload.GetImageFullUrl(imageName)
	//			data["image_save_url"] = savePath + imageName
	//		}
	//	}
	//}
	//c.JSON(http.StatusOK,gin.H{
	//	"code":code,
	//	"msg":e.GetMsg(code),
	//	"data":data,
	//})
	appG :=app.Gin{C:c}
	file,image,err := c.Request.FormFile("image")
	if err!= nil{
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError,e.ERROR,nil)
		return
	}

	if image==nil{
		appG.Response(http.StatusBadRequest,e.INVALID_PARAMS,nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath :=upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath+imageName

	if ! upload.CheckImageExt(imageName) || !upload.CheckImageSize(file){
		appG.Response(http.StatusBadRequest,e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT,nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err !=nil{
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError,e.ERROR_UPLOAD_SAVE_IMAGE_FAIL,nil)
		return
	}

	if err:=c.SaveUploadedFile(image,src);err!=nil{
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError,e.ERROR_UPLOAD_SAVE_IMAGE_FAIL,nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS,map[string]string{
		"image_url":upload.GetImageFullUrl(imageName),
		"image_save_url": savePath+imageName,
	})
}
