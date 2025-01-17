package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"github.com/mat/besticon/besticon"
	"io/ioutil"
)

type CreateSiteData struct {
	CategoryId  int32  `json:"category_id"`
	Thumb       string `json:"thumb"`
	Title       string `json:"title"`
	Url         string `json:"Url"`
	Description string `json:"description"`
}

// 获取网站 logo
func getWebLogoIconUrl(siteData *CreateSiteData) string {
	if siteData.Thumb == "" {
		b := besticon.New(besticon.WithLogger(besticon.NewDefaultLogger(ioutil.Discard))) // Disable verbose logging
		finder := b.NewIconFinder()
		icons, err := finder.FetchIcons(siteData.Url)
		if err != nil || len(icons) == 0 {
			return ""
		}
		return icons[0].URL
	}
	return siteData.Url

}

func (s *service) Create(ctx core.Context, siteData *CreateSiteData) (id int32, err error) {

	model := site.NewModel()
	model.CategoryId = siteData.CategoryId
	model.Thumb = getWebLogoIconUrl(siteData)
	model.Title = siteData.Title
	model.Url = siteData.Url
	model.Description = siteData.Description
	model.IsUsed = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
