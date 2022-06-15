# Application for monitoring the availability and immutability of sites over Telegram bot

Go (golang) application for checking and monitoring nodes through HTTP and HTTPS requests to nodes with a response code check other than '200'. There is also an additional check of nodes via PING (ICMP).

## Install and config
Download the required repository;

```bash
$ git clone https://github.com/leech001/site-monitor.git
```

Change dir to 'app' and edit config.yaml file;

```yaml
app:
  update: 30   #time to rechecking hosts (sec)

telegram:
  token: 244516775:AAGZп55654ASsFFpbjyNA9su6gQU-Qs  #Token for you Telegram BOT
  group: 123456     # Telegram you ID or group ID (use command for BOT /start

http:
  repeat: 5   # number of rechecks
  timeout: 1  # HTTP(s) timeout (sec) 
  sites:
    - url: http://yandex.ru/   # hosts for monitoring over HTTP or HTTPS with basic auth
      elements:
        - YandexZen   # content(element) in the source code of the site page
    - url: https://example.com/ # hosts for monitoring over HTTP or HTTPS

icmp:
  count: 10         # number of ICMP requests in a ping packet
  timeout: 4        # time out for ping error status (sec)
  timedelay: 100    # time delay for ping warning status (ms) 
  hosts:
    - google.com    # hosts for monitoring over ICMP
```

## Run
Change dir to root project and run docker container;
```bash
$ sudo docker-compose up -d --build
```

## Telegram BOT command
```
/start  # Print you ID or group ID need you for config
/list   # Print monitoring sites and hosts
```
