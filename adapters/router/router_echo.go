package router

//
//import (
//	"github.com/labstack/echo/v4"
//	"github.com/omid-h70/bookstore/users-api/adapters/controller"
//	"github.com/omid-h70/bookstore/users-api/controllers"
//	"io"
//	"net/http"
//)
//
//type EchoRouter struct {
//	echoMux *echo.Echo
//}
//
//func (e *EchoRouter) NewRouter() (Router, error) {
//	router := &EchoRouter{
//		echoMux: echo.New(),
//	}
//	return router, nil
//}
//
//func (e *EchoRouter) SetAppHandlers() {
//	e.echoMux.GET("/ping", buildEchoHandlerFunc(e.echoMux, controllers.Ping))
//}
//
//func buildEchoHandlerFunc(e *echo.Echo, h controller.RequestHandler) func(c echo.Context) error {
//	return func(c echo.Context) error {
//		bytes, err := io.ReadAll(c.Request().Body)
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, err.Error())
//		}
//		//Unmarshal Data to Desired Struct
//
//		var req controller.Request
//		req.SetData(bytes)
//
//		resp := h(&req)
//		return c.JSON(http.StatusOK, resp.Data())
//	}
//}
//
//func (e *EchoRouter) Send(status int, data string) {
//	//return e.echoMux.cont
//}
