# Hotel reservation backend

---

## Project outline
- users -> book room from an hotel (Книжный зал из отеля)
- admins -> going to check reservation/bookings (проверка бронирования и заказов)
- authentication and authorization -> JWT tokens (Аутентификация и авторизаци)
- hotels -> CRUD API -> JSON
- rooms -> CRUD API -> JSON
- scripts -> database management -> seeding, migration (Сценарии -> Управление базами данных -> Посев, миграция)

---

## Resources

### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```
Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

---

### Fiber
Documentation
```
https://gofiber.io
```
Installing mongodb client
```
go get github.com/gofiber/fiber/v2
```

### Bcrypt
Documentation
```
```
Installing bcrypt
```
go get golang.org/x/crypto/bcrypt
```

---
## Docker

### Installing mongodb as a Docker container
```
docker run --name mongodb-hotel -p 27017:27017 -d mongo:latest
```