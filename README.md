# Keepalive

A small, text-backed reminder service to remind you to eat and sleep.

## Tech Stack

### Frontend

React? i guess

### Backend

#### Database

We're going to use [github.com/peterbourgon/diskv/v3](https://github.com/peterbourgon/diskv) for it's good performance and plaintext persistance. This is a little unconventional but we'll also be using the well known [todo.txt](http://todotxt.org/) format (via [github.com/1set/todotxt](https://github.com/1set/todotxt) for storing todo lists.

To see how the data breaks down refer to `keepalive-api/db/user.go` and [the Task type from `todotxt`](https://pkg.go.dev/github.com/1set/todotxt#Task). There is also a separate type for auth (currently it just wraps a token string as a placeholder).

I'm not sure how much I like the current layout where the main user type live in the `db` package but it's helping write now while I flesh things out.

#### API

We'll be building the API in [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin). I've got an `api` package going that wraps the gin Engine (along with it's precious RouterGroups) and provides configuration options for CORS and auth with the [functional configuration](https://sagikazarmark.hu/blog/functional-options-on-steroids/) pattern.

Auth is provided by `github.com/gin-contrib/authz`, which consumes an `Enforcer` from the `github.com/casbin/casbin/v2` package for configuration. The CORS option consumes a `cors.Config` and attaches it to the provided group.

All configuration functions take a string as their first argument that identifies the subgroup you're trying to configure. If the string is "engine", the engine itself is configured. After that initial run of NewEngine(), further configuration can be done via the map of RouterGroup references returned from NewEngine().
