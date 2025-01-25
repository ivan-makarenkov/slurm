# gometr

Простой http сервер, который отдает go метрики

## Зависимости

Для установки зависимостей перейдите в корень проекта и выполните следующую команду

```bash
make install
```

## Запуск

Чтобы запустить сервис, из директории Service выполните команду  

```bash
make run
```

## Линтер

Для запуска линтера, перейдите в директорию Service проекта и выполните следующую команду

```bash
make lint
```

## Структура папок

```bash
.
├── Makefile
├── build
│   └── gometr                   # Dockerfile приложения gometr
│       └── Dockerfile
├── cmd
│   └── gometr
│       ├── README.md
│       └── main.go
├── configs                      # папка с конфигами приложения
│   └── gometr.yaml              # конфиг gometr
├── internal
│   └── gometr
│       ├── app                  # папка с описанием и стартом сервиса
│       │   ├── bootstrap.go
│       │   ├── main.go
│       │   └── run.go
│       ├── handlers             # пакет с api хэндлерами нашего сервиса
│       │   ├── handler.go
│       │   ├── health.go
│       │   ├── models
│       │   │   └── checks.go
│       │   └── router.go
│       └── infrastructure
│           └── config
│               └── config.go
├── main.go
└── pkg
    └── graceful
        ├── graceful.go
        └── handler.go
```

Для более подробной информации о том, как строить структуру файлов и папок для проекта на GO читайте по [ссылке](https://github.com/golang-standards/project-layout)

## http Методы

Приложение стартует на порту 8000. После запуска будет доступен по url localhost:8000

* /metrics - отдает go метрики (согласно протоколу prometheus)
* /health - сокращенный формат ответа стандарта хелсчека

```json
{
     "status": "pass",
     "service_id": "gometr",
     "checks": {
         "ping_mysql": {
             "component_id": "mysql",
             "component_type": "db",
             "status": "pass"
         }
     }
 }
```
