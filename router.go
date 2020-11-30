package main

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1/goods/")
	{
		v1.GET("find_one", GoodsFindOne)
		v1.GET("find_all", GoodsFindAll)
		v1.POST("crete", GoodsCreate)
		v1.PUT("update", GoodsUpdate)
		v1.DELETE("delete", GoodsDelete)
	}
	return r
}
