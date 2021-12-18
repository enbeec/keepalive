# Keepalive

A small, text-backed reminder service to remind you to eat and sleep.

## Tech Stack

### Frontend

React?

### Backend

#### Gateway

I'm going to cheat on AA (as usual for side projects at this scale) and use a combination of Caddy (a single static webpage and basic auth) for authentication and the simplest possible authorization scheme: non-rotating tokens and no shared data. The API deployment script will expect `caddy` in your path but all the features I'm using are native and stable in the current major version (2) so no need to vendor it in or anything.

#### Database

We're gonna use [tiedot](https://github.com/HouzuoGuo/tiedot) which is a document database written in Go. All the data is persisted as plain text on the filesystem. It also supports an HTTP(S) API for querying and mutating data. 

#### API

Golang, maybe Gin. Using tiedot as the database means when you're running the API and database locally you can directly query it with `fetch` as an escape hatch.
