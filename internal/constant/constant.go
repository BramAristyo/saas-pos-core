package constant

type ContextKey string

const (
	AuthHeader string = "Authorization"

	ClaimUserID string = "userID"
	ClaimRole   string = "role"

	CtxUserID ContextKey = "userID"
	CtxRole   ContextKey = "role"
)
