package routes

import (
	"productMs/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	groupProduct := r.Group("v1")
	{
		groupProduct.GET("product", controllers.GetProducts)
		/*groupProduct.POST("product", Controllers.CreateProduct)
		groupProduct.GET("product/:id", Controllers.GetProductByID)
		groupProduct.PUT("product/:id", Controllers.UpdateProduct)
		groupProduct.DELETE("product/:id", Controllers.DeleteProduct)
		*/
	}
	return r
}
