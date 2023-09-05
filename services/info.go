package services

import (
	"gin_test/config"
	"gin_test/model"
)

type InfoService struct{}

var (
	InfoServiceDao *InfoService
)

func init() {
	InfoServiceDao = &InfoService{}
}
func (c *InfoService) Add(v *model.Info) error {
	o := config.GetMySql()
	info := model.Info{}
	return info.Add(v, o)
}
