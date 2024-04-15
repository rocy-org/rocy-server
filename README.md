# rocy-server

Server for the rocy project.

Rocy project aims to build a B/S arch music playback and management solution including frontend music player, backend management server and music source adapter.

This repo contains the source code of backend management server that provides:

* API and music data for frontend player.
* Manage music related user data: music/artist/album/playlists.
* Tool to manage music source adapter.

## Build from source

### Generate swagger doc

```
# Install swag cli.
go install github.com/swaggo/swag/cmd/swag@latest

# Generate doc.
swag init -g server/api/api.go -o docs/swagger/
```
