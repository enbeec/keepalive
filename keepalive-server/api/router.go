package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
	Groups map[string]*gin.RouterGroup
}

// NewAPI implements the functional config of a router via a wrapper
//		All RouterGroups are created via config functions. All config
//		functions take a group basePath as a first parameter.
func NewAPI(opts ...func(*API)) *API {
	api := API{
		Router: gin.New(),
		Groups: make(map[string]*gin.RouterGroup),
	}

	// TODO: logger

	// enable panic recovery
	api.Router.Use(gin.Recovery())

	// add handlers from config options
	for _, opt := range opts {
		opt(&api)
	}

	return &api
}

func (api *API) groupMustExist(group string) {
	if _, exists := api.Groups[group]; !exists {
		api.Groups[group] = api.Router.Group(group)
	}
}

// AuthzEnforce attaches a casbin Enforcer (via a reference) to a RouterGroup
//		specified by it's basePath/name
func AuthzEnforce(group string, enforcer *casbin.Enforcer) func(*API) {
	return func(api *API) {
		api.groupMustExist(group)
		api.Groups[group].Use(authz.NewAuthorizer(enforcer))
	}
}

// CORS attaches a cors configuration struct to a RouterGroup
//		specified by it's basePath/name
func CORS(group string, config cors.Config) func(*API) {
	return func(api *API) {
		api.groupMustExist(group)
		api.Groups[group].Use(cors.New(config))
	}
}
