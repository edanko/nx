package tenant

import "context"

// ctxKeyTenantID represents the type of value for the context key.
type ctxKeyTenantID int

// tenantIDKey is how tenant id value is stored/retrieved.
const tenantIDKey ctxKeyTenantID = 0

// ContextWithTenantID returns a new Context that carries tenant id.
func ContextWithTenantID(ctx context.Context, tenantID string) context.Context {
	if ctx == nil {
		return nil
	}
	if tenantID == "" {
		return ctx
	}

	return context.WithValue(ctx, tenantIDKey, tenantID)
}

// FromContext returns the tenant id value stored in ctx, if any.
func FromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}

	t, ok := ctx.Value(tenantIDKey).(string)

	return t, ok
}
