package world

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Caption() string {
	return `"Hello world" handler`
}

func (h *Handler) Description() string {
	return `Returns string with "Hello world" message`
}

