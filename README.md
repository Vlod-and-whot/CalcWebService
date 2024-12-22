# CalcWebService

# Простой веб-сервис для подсчета арифметических уравнений(калькулятор)

## Описание
Здесь представлен простой веб-сервис на языке golang, способный выполнять функции калькулятора, а именно считать выражения, которые ему передал пользователь через HTTP-запрос.

## Структура проекта

- `cmd/` — точка входа приложения.
- `Packages/` — внутренняя логика и модули приложения.
- `Tests/` — директория тестов к кальулятору и веб-сервису.
- `custom_errors` - спецальные ситуативные случаи ошибок(деление на 0, некорректный ввод и тд).

## Запуск сервиса

1. Установите [Go](https://go.dev/dl/).
2. Склонируйте проект с GitHub:
    ```bash
    git clone https://github.com/Vlod-and-Whot/CalcWebService.git
    ```
3. Перейдите в папку проекта и запустите сервер:
    ```bash
    go run ./main.go
    ```
4. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

### Альтернативный запуск
Вы можете использовать скрипты для сборки и запуска:
- **Для Linux/MacOS:**
    ```bash
    ./build/build.sh
    ```
- **Для Windows:**
    ```powershell
    .\build\build.bat
    ```

## Эндпоинты

### `POST /api/v1/calculate`

#### Описание
Эндпоинт принимает JSON с математическим выражением.

#### Пример запроса с использованием PowerShell

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
Пример успешного ответа
json
Копировать код
{
  "result": "6.000000"
}
Пример ошибки 422
Если выражение содержит некорректный символ $, сервер вернёт ошибку 422:

Пример запроса
powershell
Копировать код
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "1+2"}'
Пример ответа
json
Копировать код
{
  "result": "3"
}
Тестирование
Для запуска тестов выполните:

bash
Копировать код
go test ./...

cURLs:
curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "error"}'

curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d ''

curl -X POST http://localhost:8080/api/v1/calculate/ \
-d '{"expression": "1 + 2"}'

curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "1 + 2"}'

curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{"expression": "invalid"}'

curl -X POST http://localhost:8080/api/v1/calculate/ \
-H "Content-Type: application/json" \
-d '{invalid}'

## Структура проекта




Примечания
Для работы API требуется установленный Go (версии 1.18 и выше).
Все зависимости проекта управляются через go mod. Убедитесь, что в корне проекта находится go.mod.
