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
	// TODO: if group does not exist, create it
}

func BasePath(group, path string) func(*API) {
	return func(api *API) {
		api.groupMustExist(group)
		// TODO: set the base path
	}
}

func AuthzEnforce(group string, enforcer *casbin.Enforcer) func(*API) {
	return func(api *API) {
		api.groupMustExist(group)
		api.Groups[group].Use(authz.NewAuthorizer(enforcer))
	}
}

func CORS(group string, config cors.Config) func(*API) {
	return func(api *API) {
		api.groupMustExist(group)
		api.Groups[group].Use(cors.New(config))
	}
}
