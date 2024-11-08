# Druni Advent Calendar Checker

This project is a Go application that checks the availability of the Advent calendar on the Druni online store. If the calendar is available, the application sends a notification to a specified Telegram chat.

---

## Requirements

- [Go](https://golang.org/doc/install) 1.19 or higher (optional if using Docker)
- [Docker](https://docs.docker.com/get-docker/) to build and run the container
- A Telegram bot and a chat or group to send notifications to (see Configuration section for details)

## Configuration

Before running the application, you need to configure some environment variables in a `.env` file in the root project directory. These variables include your Telegram bot token and the chat or group ID to which notifications will be sent.

### Step 1: Create the `.env` File

Create a file named `.env` in the root directory (`/cmd`) and add the following:

```env
TOKEN=your_telegram_bot_token
CHAT_ID=your_telegram_chat_id
```

### Step 2: Build the docker image

```bash
docker build -t druni-calendar-checker .
```

### Step 3: Run the container

```
docker run --env-file .env druni-calendar-checker
```

### Usage

The application connects to the Druni Advent calendar URL and checks if it is available. If it finds the keyword No disponible ("Not available"), it will log a notification message indicating unavailability and will exit (note: retry functionality could be added as an improvement).

If the calendar is available, the application will send a message to your Telegram chat indicating that the product is in stock.
