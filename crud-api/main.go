package main

import (
	"zkfmapf123/crud-api/src/handler"
	"zkfmapf123/crud-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	middleware(r)
	router(r)
	r.Run(":8080")

}

// middleware
func middleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery()) // panic occured response to 500
}

// customer middleware
func getMiddleware() map[string]gin.HandlerFunc{
	mw := make(map[string]gin.HandlerFunc)
	mw["auth"] = middlewares.AuthMiddleware()
	mw["valid"] = middlewares.ValidMiddleware()
	mw["serialize"] = middlewares.SerializeMiddleware()

	return mw
}

// router
func router(r *gin.Engine){
	
	mw := getMiddleware()
	v1 := r.Group("/v1", mw["auth"],mw["valid"],mw["serialize"])
	{
		v1.GET("/movies",handler.HandleGetAllMovie)
		v1.GET("/movies/:id",handler.HandleGetMovie)
		v1.POST("/movies",handler.HandleCreateMovie)
		v1.PUT("/movies/:id",handler.HandleUpdateMovie)
		v1.DELETE("/movies/:id",handler.HandleDeleteMovie)
	}

}