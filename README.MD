# Application for monitoring the availability and immutability of sites over Telegram bot
![Made](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/grfc-ru/telegram-site-monitor/docker-image.yml?label=BUILD%20%26%20PUBLISH%20APPLICATION&logo=github) [![GitHub](https://img.shields.io/badge/Git-Hub-purple.svg)](https://github.com/leech001/telegram-site-monitor) [![Docker](https://img.shields.io/badge/Docker-hub-2496ed.svg)](https://hub.docker.com/r/leech001/telegram-site-monitor)

Go (golang) application for checking and monitoring nodes through HTTP and HTTPS requests to nodes with a response code check other than '200'.

## Simple use
Download the required repository;

```bash
$ git clone https://github.com/leech001/site-monitor.git
```

Change `conf/config.yaml`

Run `sudo docker-compose up -d`

## Configure

```yaml
app:
  update: 30   #time to rechecking hosts (sec)

telegram:
  token: 244516775:AAGZп55654ASsFFpbjyNA9su6gQU-Qs  #Token for you Telegram BOT
  group: 123456     # Telegram you ID or group ID (use command for BOT /start

http:
  repeat: 5   # number of rechecks
  timeout: 30  # HTTP(s) timeout (sec)
  delay: 1.5 # HTTP(s) time delay (sec)
  sites:
    - url: http://yandex.ru/   # hosts for monitoring over HTTP or HTTPS with basic auth
      elements:
        - YandexZen   # content(element) in the source code of the site page
    - url: https://example.com/ # hosts for monitoring over HTTP or HTTPS
```

## Telegram BOT command
```
/start  # Print you ID or group ID need you for config
/list   # Print monitoring sites and hosts
```
