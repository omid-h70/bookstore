package router

type (
	Router interface {
		NewRouter() (Router, error)
		SetAppHandlers()
	}
)
