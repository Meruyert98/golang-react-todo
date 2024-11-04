# Go To Do App

Полноценное приложение для управления задачами (To-Do) с использованием Golang для бэкенда, React с Semantic UI для фронтенда и MongoDB в качестве базы данных. Этот проект позволяет пользователям создавать, обновлять и удалять задачи, предлагая удобный интерфейс для управления делами.

## Оглавление

- [Функции](#функции)
- [Стек технологий](#стек-технологий)
- [Структура проекта](#структура-проекта)
- [Установка](#установка)
- [Конфигурация](#конфигурация)
- [Запуск](#запуск)
- [API Эндпоинты](#api-эндпоинты)
- [Использование Docker](#использование-docker)
- [Вклад в проект](#вклад-в-проект)
- [Лицензия](#лицензия)

## Функции

- Создание, чтение, обновление и удаление задач (CRUD).
- Постоянное хранение данных в MongoDB.
- Адаптивный интерфейс на базе React и Semantic UI.

## Стек технологий

- **Бэкенд**: Golang
- **Фронтенд**: React, Semantic UI
- **База данных**: MongoDB (локальная)

## Структура проекта

```plaintext
Go-To-Do-App/
├── client/                 # Фронтенд на React
│   ├── public/
│   └── src/
│       ├── components/     # Повторно используемые компоненты интерфейса
│       ├── pages/          # Основные страницы
│       ├── To-Do-List.js   # API-сервисы
│       └── App.js          # Основной файл приложения
├── server/                 # Бэкенд на Golang
│   ├── handlers/           # Обработчики API маршрутов
│   ├── models/             # Модели данных
│   ├── main.go             # Основной файл сервера
│   └── router              # Настройка роутера
└── README.md
└── docker-compose.yml
└── .env
```

## Установка

**1. Клонирование репозитория**

```bash
git clone https://github.com/Meruyert98/golang-react-todo.git
cd golang-react-todo
```

**2. Настройка бэкенда**

- Перейдите в директорию server:
```bash
cd server
```
- Установите зависимости для Go (если есть) и соберите сервер:
```bash
go mod download
go build
```

**3. Настройка фронтенда**

- Перейдите в директорию client:
```bash
cd ../client
```
- Установите зависимости для Node:
```bash
npm install
```
**4. Настройка MongoDB**
Убедитесь, что MongoDB установлена и запущена локально на порту 27017.

## Конфигурация
**Конфигурация бэкенда**
- В директории `server` создайте файл `.env` со следующими переменными окружения:
```dotenv
DB_URI=mongodb://localhost:27017
DB_NAME=todo
DB_COLLECTION_NAME=tasks
```
**Конфигурация фронтенда**
Настройте конечную точку API на фронтенде в файле `client/src/services/api.js`, указав базовый URL для сервера.


## Запуск

**1. Запуск бэкенда**
- Из директории `server` запустите Go сервер:
```bash
go run main.go
```
Сервер будет запущен на http://localhost:8080.

**2. Запуск фронтенда**
- Из директории client запустите React приложение:
```bash
npm start
```
Приложение будет доступно по адресу http://localhost:3000.

**3. Открытие приложения**
Откройте браузер и перейдите по адресу http://localhost:3000 для использования приложения.

## API Эндпоинты

| Метод	 |Эндпоинт	|Описание   |
|--------|----------|-----------|
|GET	 |/api/todos |	Получить все задачи
|POST	 |/api/todos |	Создать новую задачу
|PUT	 |/api/todos/:id |	Обновить задачу
|DELETE	 |/api/todos/:id |	Удалить задачу


## Использование Docker
Для удобства развертывания создадим Dockerfile и Docker Compose для всего приложения.

**Dockerfile для бэкенда**
Создайте файл Dockerfile в директории server с таким содержанием:
```dockerfile
# Dockerfile для Go сервера
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
```

**Dockerfile для фронтенда**
Создайте файл Dockerfile в директории client с таким содержанием:

```dockerfile
# Dockerfile для React клиента
FROM node:18-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
```

**Docker Compose**
Создайте файл docker-compose.yml в корневой директории проекта с таким содержанием:

```yaml
version: '3.8'

services:
  backend:
    build: ./server
    ports:
      - "8080:8080"
    environment:
      - DB_URI=mongodb://mongodb:27017
      - DB_NAME=todo
      - DB_COLLECTION_NAME=tasks
    depends_on:
      - mongodb

  frontend:
    build: ./client
    ports:
      - "3000:3000"
    environment:
      - REACT_APP_API_URL=http://localhost:8080

  mongodb:
    image: mongo:latest
    ports:
      - "27017:27017"
    volumes:
      - mongo-data:/data/db

volumes:
  mongo-data:
```

**Запуск с Docker Compose**
- Убедитесь, что Docker установлен и запущен.
- Выполните команду для запуска всех сервисов:
```bash
docker-compose up --build
```
Откройте http://localhost:3000 в браузере, чтобы начать работу с приложением.