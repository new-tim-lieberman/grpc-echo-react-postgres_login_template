A starter project with ready login.

TODO:  The one thing that is missing from here is Enviornment variables.  I will update this in the future to use them I honestly just forgot.   So as it stands everything is using localhost and salts are hardcoded.

Stack:
**Docker, GRPC, Echo Api, Redis, Postgres, React Native, React, Sqlc.**


Usage.

Backend:
* `make proto`
* `docker-compose up --build`
* `go test ./...`

Add user:
```
curl -X POST http://localhost:8080/api/register \                                                 
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com","password":"password" }'
```



React Native:
* `npm install`
* `rm -rf .expo`
* `npx expo start -c`



React:
* `npm install`
* `npm run dev`
