package graphql

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserField,
	NewPostField,
	NewRootSchema,
	NewUserResolver,
	NewPostResolver,
)
