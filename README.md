# gator_rss

# RSS Feed Aggregator

A command-line RSS feed aggregator written in Go.

## Prerequisites

Before running this program, ensure you have the following installed:

- **Go** : [https://golang.org/dl/](https://golang.org/dl/)
- **PostgreSQL** : [https://www.postgresql.org/download/](https://www.postgresql.org/download/)

You must also have a running PostgreSQL server and a database set up.

## 🔧 Installation

You can install the `gator_rss` CLI using Go:

```bash
go install github.com/yourusername/gator_rss@latest


📘 Commands


🔐 gator login

Sets the current user in the local config file.

login <username>

This command will set your username and store it locally so future commands know which user is active.

📝 gator register

Registers a new user in the database.

register <username>

This command creates a new user entry in your Postgres database.


👥 gator users

Lists all registered users from the database.

gator users

Useful for viewing all available users. This can help you confirm a registration worked or find a user ID for testing or development.
