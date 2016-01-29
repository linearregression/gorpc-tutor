package sum

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Caption() string {
	return `"Summarize numbers`
}

func (h *Handler) Description() string {
	return `Returns sum of numbers`
}
