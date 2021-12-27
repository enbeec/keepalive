# Keepalive

A small, text-backed reminder service to remind you to eat and sleep.

## Tech Stack

### Frontend

React? i guess

### Backend

#### Gateway

I'm going to cheat on AA (as usual for side projects at this scale) and use a combination of Caddy (a single static webpage and basic auth) for authentication and the simplest possible authorization scheme: non-rotating tokens and no shared data. The API deployment script will expect `caddy` in your path but all the features I'm using are native and stable in the current major version (2) so no need to vendor it in or anything.

#### Database

We're going to use [github.com/peterbourgon/diskv/v3](https://github.com/peterbourgon/diskv) for it's good performance and plaintext persistance. This is a little unconventional but we'll also be using the well known [todo.txt](http://todotxt.org/) format (via [github.com/1set/todotxt](https://github.com/1set/todotxt) for storing todo lists.

To see how the data breaks down refer to `keepalive-api/db/user.go` and [the Task type from `todotxt`](https://pkg.go.dev/github.com/1set/todotxt#Task). There is also a separate type for auth (currently it just wraps a token string as a placeholder).

I'm not sure how much I like the current layout where the main user type live in the `db` package but it's helping write now while I flesh things out.

#### API

We'll be building the API in [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin).
