# gator

A command-line RSS reader written in Go.

## Requirements

You will need:

- **Go** installed to build and install the CLI
- **PostgreSQL** installed and running

## Install

If you want to install the CLI globally, use `go install`:

```bash
go install github.com/shafayetsadi/gator@latest
```

After that, make sure your Go bin directory is on your `PATH` so you can run:

```bash
gator
```

> Go programs are statically compiled binaries. After running `go build`
> or `go install`, you should be able to run the program without needing
> the Go toolchain.
>
> `go run .` is for development only.

## Configuration

The app reads its config from:

```text
~/.gatorconfig.json
```

Example config:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

### Set up PostgreSQL

Create a database for the project, then put its connection string in `db_url`.

For example:

```bash
createdb gator
```

Then update `~/.gatorconfig.json` with your database URL.

## Running locally

For development, you can run the app directly from the repo:

```bash
go run . <command> [args...]
```

For example:

```bash
go run . register sadi
go run . login sadi
go run . addfeed "TechCrunch" https://techcrunch.com/feed/
```

## Commands

Here are some of the available commands:

- `register <name>` - create a new user and set it as the current user
- `login <name>` - switch the current user
- `users` - list all users
- `reset` - clear users and feeds from the database
- `addfeed <name> <url>` - add a feed and follow it as the current user
- `feeds` - list all feeds
- `follow <url>` - follow an existing feed by URL
- `unfollow <url>` - stop following a feed by URL
- `following` - list feeds the current user follows
- `agg <duration>` - scrape feeds on a timer, for example `10s` or `1m`
- `browse [limit]` - show recent posts for the current user

### Example workflow

```bash
gator register sadi
gator addfeed "TechCrunch" https://techcrunch.com/feed/
gator feeds
gator following
gator agg 30s
gator browse 5
```
