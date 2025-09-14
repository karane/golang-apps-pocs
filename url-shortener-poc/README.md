# URL Shortener in Go

## How to Run
```bash
docker compose up --build -d

curl -X POST http://localhost:8080/shorten \
     -H "Content-Type: application/json" \
     -d '{"original_url":"https://example.com"}'
```

You should expect an output similar to this:

```
{"ID":0,"ShortCode":"251da339","OriginalURL":"https://example.com","CreatedAt":"0001-01-01T00:00:00Z"}

```