package controller

import (
	"encoding/json"
	_const "gin_test/const"
	"gin_test/model"
	"gin_test/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type InfoController struct{}

var (
	InfoControllerInstance *InfoController
)

func init() {
	InfoControllerInstance = &InfoController{}
}

func (u *InfoController) Add(g *gin.Context) {
	body, _ := ioutil.ReadAll(g.Request.Body)
	msg := &_const.ResultInfo{}
	if string(body) == "" {
		msg.Code = "-1"
		msg.Message = "body 参数为空"
		g.JSONP(200, msg)
		g.Abort()
		return
	}
	var v model.Info
	json.Unmarshal(body, &v)

	info := model.Info{
		Name:      v.Name,
		Address:   v.Address,
		Phone:     v.Phone,
		Sex:       v.Sex,
		CreatedAt: time.Now(),
	}
	err := services.InfoServiceDao.Add(&info)
	if err != nil {
		msg.Code = "-1"
		msg.Message = "新增失败"
		g.JSONP(200, msg)
		g.Abort()
		return
	}
	msg.Code = "0"
	msg.Message = "新增成功"
	g.JSONP(200, msg)
	g.Abort()
}
