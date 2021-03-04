package handlers

func NewHandlers(
	users *UserHandler,
) Handlers {
	return Handlers{
		Users: users,
	}
}

type Handlers struct {
	Users *UserHandler
}
