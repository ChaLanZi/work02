package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GoodsSerializer struct {
	Name   string
	Price  float64
	Number uint32
}

func GoodsCreate(ctx *gin.Context) {
	var service GoodsService
	err := ctx.ShouldBind(&service)
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusBadRequest, "客户端参数出错")
		return
	}
	err = service.Create()
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusInternalServerError, "服务端处理出错")
		return
	}

	ctx.String(http.StatusOK, "创建成功")

}

func GoodsFindOne(ctx *gin.Context) {
	var service GoodsService
	var goodsSeria GoodsSerializer
	err := ctx.ShouldBind(&service)
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码
		ctx.JSON(http.StatusBadRequest, goodsSeria)
		return
	}
	goods, err := service.FindOne()
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusOK, goodsSeria)
		return
	}

	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码
		ctx.JSON(http.StatusInternalServerError, goodsSeria)
		return
	}
	goodsSeria.Number = goods.Number
	goodsSeria.Price = goods.Price
	goodsSeria.Name = goods.Name

	ctx.JSON(http.StatusOK, goodsSeria)
}

func GoodsFindAll(ctx *gin.Context) {
	var service GoodsService
	var goodsSerias []*GoodsSerializer
	err := ctx.ShouldBind(&service)
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码
		ctx.JSON(http.StatusBadRequest, goodsSerias)
		return
	}
	goods, err := service.FindAll()
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusOK, goodsSerias)
		return
	}

	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码
		ctx.JSON(http.StatusInternalServerError, goodsSerias)
		return
	}

	for i := 0; i < len(goods); i++ {
		gs := &GoodsSerializer{
			Name:   goods[i].Name,
			Price:  goods[i].Price,
			Number: goods[i].Number,
		}
		goodsSerias = append(goodsSerias, gs)
	}
	ctx.JSON(http.StatusOK, goodsSerias)
}

func GoodsUpdate(ctx *gin.Context) {
	var service GoodsService
	err := ctx.ShouldBind(&service)
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusBadRequest, "客户端参数出错")
		return
	}
	err = service.Update()
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusInternalServerError, "服务端处理出错")
		return
	}

	ctx.String(http.StatusOK, "更新成功")

}

func GoodsDelete(ctx *gin.Context) {
	var service GoodsService
	err := ctx.ShouldBind(&service)
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusBadRequest, "客户端参数出错")
		return
	}
	err = service.Delete()
	if err != nil {
		// 输出日志
		logrus.Error(err)
		//将错误转为 状态码返回
		ctx.String(http.StatusInternalServerError, "服务端处理出错")
		return
	}

	ctx.String(http.StatusOK, "删除成功")

}
