package router

type (
	Router interface {
		NewRouter() (Router, error)
		SetAppHandlers()
		//Send(status int, data string)
	}
)
