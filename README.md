# go-pushover
Simple Pushover CLI client, written in Go.

# Installation
Binaries are available to download on the [releases page](https://github.com/luketainton/go-pushover/releases). Alternatively, you can install by running `go install github.com/luketainton/pushover`.

# Usage
| Flag      | Description         |
| --------- | ------------------- |
| `-t`      | Message title       |
| `-m`      | Message body        |
| `-u`      | Hyperlink URL       |
| `-r`      | Hyperlink text      |
| `-w`      | Enable HTML parsing |

Your Pushover app and user tokens are retrieved from your environment variables - you must set `PUSHOVER_USER_TOKEN` and `PUSHOVER_APP_TOKEN`.

# Examples
- Title and message only:
  - `pushover -t "Hello world" -m "Message body"`
- Title and message with hyperlink:
  - `pushover -t "Hello world" -m "Message body" -u "https://google.com" -r "Google"`
- Title and HTML message:
  - `pushover -t "Hello world" -m "<h1>Hello</h1>" -w`

# Issues or questions?
Please open an issue!