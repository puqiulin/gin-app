package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserHandler, NewGraphQLHandler, NewGoogleHandler, NewCacheHandler)
