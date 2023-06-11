package api

import "github.com/julienschmidt/httprouter"

func (app application) router() *httprouter.Router {
	router := httprouter.New()

	//get
	router.GET("/url/:id", app.ForwardURL)

	//post
	router.POST("/url", app.CreateURL)

	//put
	router.PUT("/url/:id", app.UpdateURL)

	//delete
	router.DELETE("/url/:id", app.DeleteURL)

	return router
}
