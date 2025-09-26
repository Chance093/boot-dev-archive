# Gator

Gator is a CLI application that aggregates RSS feeds. You can register 
different users, subscribe to feeds, and aggregate their posts in the 
background. This CLI is very simple to use, and was built for the purpose of me 
learning Go.

## Features

- Register multiple users
- Each user can create and follow feeds
- Aggregate posts in the background
- Browse through posts for feeds you follow

## Installation

### Using `go install`
To install gator, use the following command:
```bash
go install github.com/Chance093/gator@latest
```
Ensure that `$GOBIN` is in your system `PATH` so you can run `gator` from anywhere.

### Building from source
Alternatively, you can build from source:
```bash
git clone https://github.com/Chance093/gator.git
cd gator
make build # or go build -o gator
mv gator /usr/local/bin/ # Optional: Move to a directory in your PATH
```

## Setup

### Postgres Installation
You will first need to locally install postgresql on your machine.

**macOS with brew**
```bash
brew install postgresql@15
```

**Linux / WSL (Debian). Here are the docs from Microsoft, but simply:**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Ensure the installation worked:
```bash
psql --version
```

(Linux Only) Update postgres password:
```bash
sudo passwd postgres
```

Start your Postgres server in the background:
- Mac: `brew services start postgresql`
- Linux: `sudo service postgresql start`

### Gator config file
You will next need to set up a gator config in your root dir:
```bash
touch ~/.gatorconfig.json
```

Your json file should look like this:
```json
{
  "db_url": "postgres://example"
}
```
Obviously change the db url.

## Usage

Once installed you can use `gator` as follows:
```bash
gator [command] <args...>
```

### Commands

- `register <username>` - registers user and sets as the active user
```bash
gator register chanceyboy
```

- `login <username>` - sets user as the active user (only if already registered)
```bash
gator login chanceyboy
```

- `users` - lists all users in registry
```bash
gator users
```

- `addfeed <feed_name url>` - adds a feed and associates with that user
```bash
gator addfeed TechCrunch https://techcrunch.com/feed/
```

- `feeds` - lists all feeds for the active user
```bash
gator feeds
```

- `follow <url>` - Follows a feed made by another user
```bash
gator follow https://techcrunch.com/feed/
```

- `unfollow <url>` - Unfollows a feed for the active user
```bash
gator unfollow https://techcrunch.com/feed/
```

- `following` - Lists all the feeds the active user is following
```bash
gator following
```

- `agg <time_duration>` - Runs in the background and fetches all posts for the users followed feeds
```bash
gator agg 30s
```

- `browse <limit>` - Displays all posts that the user has aggregated
```bash
gator browse 20
```

- `reset` - Resets all data in your database
```bash
gator reset
```

## Contributing

1. Fork the repo

2. Create a new branch (git checkout -b feature-branch)

3. Make your changes

4. Submit a pull request
