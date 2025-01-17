package constant

type ContextKey string

const (
	FiberContext  ContextKey = "fiberCtx"
	HeaderContext ContextKey = "headerCtx"
)

const (
	Success = "success"
	Error   = "error"
)
