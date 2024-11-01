package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nsiebenaller/go-api-example/controllers"
	books_controllers "github.com/nsiebenaller/go-api-example/controllers/books"
	members_controllers "github.com/nsiebenaller/go-api-example/controllers/members"
)

func ConfigureRouter(server *gin.Engine) {
	// Health route
	server.GET("/health", controllers.GetHealth)

	// Member routes
	server.GET("/members", members_controllers.GetMembers)
	server.POST("/members", members_controllers.CreateMember)
	server.GET("/members/:id", members_controllers.FindMember)
	server.DELETE("/members/:id", members_controllers.DeleteMember)
	server.POST("/members/checkout-book", members_controllers.CheckoutBook)

	// Book routes
	server.GET("/books", books_controllers.GetBooks)
	server.POST("/books", books_controllers.CreateBook)
	server.DELETE("/books/:id", books_controllers.DeleteBook)
}
