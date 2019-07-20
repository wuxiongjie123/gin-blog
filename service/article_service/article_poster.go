package article_service

import (
	"gin-blog/pkg/file"
	"gin-blog/pkg/qrcode"
)

type ArticlePoster struct {
	PosterName string
	*Article
	Qr *qrcode.QrCode
}

func NewArticlePoster(posterName string,article *Article,qr *qrcode.QrCode) *ArticlePoster {
	return &ArticlePoster{
		PosterName:posterName,
		Article:article,
		Qr:qr,
	}
}

func GetPosterFlag() string {
	return "poster"
}

func (a *ArticlePoster) CheckMergeImage(path string) bool {
	if file.CheckExist(path+a.PosterName) ==true{
		return false
	}
	return true
}