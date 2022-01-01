package api

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/authz"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Config is used internally before the API is returned to New's caller
//		This enables chaining together configuration functions for things
//		like multiple middlewares on subroutes without *needing* access
//		to the RouterGroups. You do get the map used to pull this off
//		back as a second return value. This is useful for trying out new
//		configurations outside the module before needing to try to write
//		a generally useful higher order configuration function.
type config struct {
	// TODO: enforce overwriting keys after it is returned
	groups map[string]*gin.RouterGroup
	engine *gin.Engine
}

// . makes sure that the RouterGroup you're trying to configure exists
func (c *config) groupMustExist(group string) {
	if _, exists := c.groups[group]; !exists {
		if group == "engine" || group == "" {
			c.groups[group] = c.engine.Group(group)
		} else {
			c.groups[group] = c.groups[group].Group(group)
		}
	} else {
		// Configuration time is a better time to panic than most, I find.
		panic("You may not reuse group names.")
	}
}

// NewEngine implements the functional config of a router via a wrapper
//		All RouterGroups are created via config functions. All config
//		functions take a group basePath as a first parameter.
//		Also, it returns a pointer to the map of router groups
//			for further tweaking, for now.
func NewEngine(opts ...func(*config)) (*gin.Engine, map[string]*gin.RouterGroup) {
	c := config{
		groups: make(map[string]*gin.RouterGroup),
		engine: gin.New(),
	}

	// TODO: logger

	// enable panic recovery
	c.engine.Use(gin.Recovery())

	// add handlers from config options
	for _, opt := range opts {
		opt(&c)
	}

	return c.engine, c.groups
}

// Group creates a new RouterGroup from another group. It is the simplest
//		example of a configuration function.
func Group(group string, newGroup string) func(*config) {
	return func(c *config) {
		c.groupMustExist(group)
	}
}

// AuthzEnforce attaches pointer to a casbin Enforcer to a RouterGroup
func AuthzEnforce(group string, enforcer *casbin.Enforcer) func(*config) {
	return func(c *config) {
		if group == "engine" || group == "" {
			c.engine.Use(authz.NewAuthorizer(enforcer))
		} else {
			c.groupMustExist(group)
			c.groups[group].Use(authz.NewAuthorizer(enforcer))
		}
	}
}

// CORS attaches a cors configuration struct to a RouterGroup
func CORS(group string, corsConfig cors.Config) func(*config) {
	return func(c *config) {
		if group == "engine" || group == "" {
			c.engine.Use(cors.New(corsConfig))
		} else {
			c.groupMustExist(group)
			c.groups[group].Use(cors.New(corsConfig))
		}
	}
}
