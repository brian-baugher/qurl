# qurl

## Development

Start DB with `sudo systemctl start mysql`

Export env variables `DBUSER` and `DBPASS`

Start server with `go run cmd/qurl/main.go`

Test endpoints with `curl -H 'Content-Type: application.json' -X POST -d '{"url":"test url2"}' localhost:8000/url`