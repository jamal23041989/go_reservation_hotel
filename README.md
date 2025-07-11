# Hotel reservation backend

---

## Project outline

- users -> book room from an hotel (Книжный зал из отеля)
- admins -> going to check reservation/bookings (проверка бронирования и заказов)
- authentication and authorization -> JWT tokens (Аутентификация и авторизаци)
- hotels -> CRUD API -> JSON
- rooms -> CRUD API -> JSON
- scripts -> database management -> seeding, migration (Сценарии -> Управление базами данных -> Посев, миграция)

```
go_reservation_hotel/
/cmd                 → входная точка (main.go)
/internal
    /domain          → сущности + интерфейсы (без зависимостей)
    /usecase         → бизнес-логика (исп. интерфейсы domain)
    /repository
        /mongo       → конкретная реализация для MongoDB
    /handler         → HTTP/REST обработчики
/pkg                 → вспомогательные пакеты (ошибки, конфиг и пр.)
/configs             → .env, конфиги
/scripts             → сиды, миграции
```



### В чистой архитектуре, слой domain:
- 🔒 Не зависит ни от БД, ни от веб-сервера, ни от логирования.
- 💡 Содержит ядро бизнес-логики — всё, что определяет что делает приложение, а не как оно это делает.
- 👥 Сущности (User, Hotel, Booking, и т.д.) описывают реальные объекты, с которыми работает бизнес, независимо от хранения.
- 🔌 Интерфейсы (UserRepository, HotelRepository) нужны, чтобы можно было внедрить (inject) зависимости извне.
- types.Hotel, types.User, types.Room, types.Booking — это сущности (entities) — они существуют независимо от БД, фреймворков или веб-интерфейса.
- HotelStore, UserStore, и другие интерфейсы — это абстракции (порт), через которые бизнес-логика общается с инфраструктурой.

###  В чистой архитектуре, cлой usecase:
- 💼 Отвечает за бизнес-сценарии — то, как мы используем сущности из domain;
- 🔁 Работает только через интерфейсы из domain (например, UserRepository);
- ❌ Не зависит от web-фреймворков, баз данных, логгирования и т.д.;
- 📋 Здесь определяются все правила — что должно произойти, если пользователь хочет забронировать номер, создать отель и т.д.
- 🔒 Бизнес-логика (ограничить limit, валидация, условия) хранится централизованно;
- 👨‍🔧 Контроллеру (или handler'у) не нужно знать, как работает БД;
- 🔁 Легко писать тесты — ты просто мокаешь repo.



---

# Project environment variable

```
HTTP_LISTEN_ADDRESS=:3000
JWT_SECRET=somethingsupersecret
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
```
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