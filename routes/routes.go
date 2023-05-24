package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodblg/Ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())

}

/*
func EcommerceRoutes(incomingRoutes *gin.Engine) {
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	incomingRoutes.GET("/addtocart", app.AddToCart())
	incomingRoutes.GET("/removeitem", app.RemoveItem())
	incomingRoutes.GET("/listcart", controllers.GetItemFromCart())
	incomingRoutes.POST("/addaddress", controllers.AddAddress())
	incomingRoutes.PUT("/edithomeaddress", controllers.EditHomeAddress())
	incomingRoutes.PUT("/editworkaddress", controllers.EditWorkAddress())
	incomingRoutes.GET("/deleteaddresses", controllers.DeleteAddress())
	incomingRoutes.GET("/cartcheckout", app.BuyFromCart())
	incomingRoutes.GET("/instantbuy", app.InstantBuy())

}


router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

*/
