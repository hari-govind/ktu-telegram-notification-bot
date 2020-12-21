# KTU Telegram Notification Bot

Relay notifications from the official KTU website to a telegram channel

## Usage
- Clone this repository and run `go build` to compile the binary or get the prebuilt binary from releases.
- Rename `conf.sample.json` to `conf.json` and fill in the details(instructions below).
- Run the binary. Currently the bot is not daemonized and hence you might need to run it using [screen](https://www.howtogeek.com/662422/how-to-use-linuxs-screen-command/).

## Configuration File Options
- `token`: Token of your telegram bot. If you don't have one, follow [this guide](https://core.telegram.org/bots#6-botfather).
- `interval`: Time delay in minutes before each poll. An interval of 5 means that a new notification will be relayed within five minutes. Lower values will cause unnecessary load on both servers, a value above 15 or 30 is recommend.
- `channel`: Channel name to relay notifications. E.g, `@your_channel_name`.

## To Be Implemented
- [ ] Daemonize the bot
- [ ] Send error message to admin
- [ ] Make the bot easier to deploy using Heroku
## Contributing
PRs are always welcome!
