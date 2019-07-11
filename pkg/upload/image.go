package upload

import (
	"gin-blog/pkg/file"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"log"
	"mime/multipart"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string  {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName :=strings.TrimSuffix(name,ext)
	fileName = util.EncodeMD5(fileName)
	
	return fileName + ext
}

func GetImagePath() string  {
	return setting.AppSetting.ImageSaverPath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath +GetImagePath()
}

func CheckImageExt(fileName string) bool  {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext){
			return true
		}
	}
	return false
}

func CheckImageSize(f multipart.File) bool  {
	size,err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}




