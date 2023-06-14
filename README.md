# assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is my implementation for backend assignment of 2023 TikTok Tech Immersion.

Requirements: https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle

## How to run with Docker

Make sure you have Docker installed. Run following command to start the app:

```
docker-compose up -d
```

## API Documentation

### Ping:

> Check if the server is running

```bash
curl -X GET http://localhost:8080/ping
```

Expected response: status 200

```json
{
  "message": "pong"
}
```

### Send message:

> Send message in a chat

```bash
curl -X POST \
  http://localhost:8080/api/send \
  -H 'Content-Type: application/json' \
  -d '{
    "Chat": "a:b",
    "Text": "Hello World",
    "Sender": "a"
}'
```

Expected response: status 200

### Pull messages:

> Retrieve messages in a chat from Cursor with Limit and sorting order Reverse (default: False)

```bash
curl -X GET \
  http://localhost:8080/api/pull \
  -H 'Content-Type: application/json' \
  -d '{
    "Chat": "a:b",
    "Cursor": 0,
    "Limit": 20,
    "Reverse": false
}'
```

Expected response: status 200

```json
    {
      "messages": [
        {
          "chat": "a:b",
          "text": "Hello World",
          "sender": "a",
          "send_time": 1684744610
        }, ...
      ]
    }
```

## Tech stack

- Go
- Redis
- Kitex
- Docker
