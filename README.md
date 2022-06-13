# Ports

Run the project with `docker-compose up`. Default host `localhost:8000`.

API:

**POST /ports**

Example value:
```
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  }
}
```
The JSON object key will be considered as the port id.

Success: 204 No Content


**POST /ports/import**

Example value:
```
{
  "filename": "ports.json"
}
```
The file path is relative to main.go.

Success: 204 No Content

