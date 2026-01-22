# Gator - RSS aggregator
Gator is a toy command line RSS aggre-gator application written in GO to practice interacting with a Postgres database. It stores most of it's data

## Requirements:
1. Golang 1.25+
2. A fresh Postgres DB for storing application data.
3. Goose for applying DB migrations

## Setup:
1. Determine a connection string for the Postgres DB. This will be platform dependent. On Linux: the format is of the form `postgres://<user>:<password>@<domain>:<port>/<dbname>`
2. Apply all of the DB migrations in sql/schema by running the command `goose postrgres <connection-string> up` which if successfull will display that migration number that corresponds to the migration files.
3. Create a JSON config file in your home directory with the name `.gatorconfig.json`
4. The content of the file should be `{"db_url":"<connection-string>?sslmode=disable"}` where `<connection-string>` is the connection string from step 1.
5. Run `go mod download` to fetch go dependencies.
6. Run `go build` to build the application.

## User instructions:
1. The basic usage is `<excutable> <command> <args>` where `<executable>` is the executable (e.g. `./gator` on unix like systems), `<command>` is one of the commands described below (required) and `<args>` are any arguments that apply to the command (varies by command).
2. The first time you run gator you must register: `<execuatable> register <name>` where `<name>` must be a single word.
3. Then you must login: `<executable> login <name>` where `<name>` must be the same as in step 2.
4. Add feeds: `<executable> addfeed <name> <url>` where `<name>` is any single word name you want to give the feed and `<url>` is a valid URL of an RSS feed. Note the command doesn't validate the URL. This will automatically cause your user to follow the feed.
5. View listing of feeds that have been added: `<executable> feeds`. Note: this command ignores args.
6. If a feed was already added, you can follow it: `<executable> follow <url>` Note: the url must already have been added by another user (if you added it, you're alreadying following unless you unfollowed).
7. You can unfollow a feed: `<executable> unfollow <url>` You must already be following the feed. Note this does not remove the feed from the DB.
8. You can review what fees you're following: `<executable> following` Note: args are ingored.
9. You can trigger the application to aggregate the RSS feeds: `<executable> agg <fetch-interval>` where `<fetch-interval>` is the time the application waits before fetching the next feed. For example `30s` is 30 seconds, `2m` is 2 minutes. The application will loop through all feeds in an infinite loop, so press <ctrl>+C (or <cmd>+C on MacOS) to kill the application. Use a sufficiently long interval to prevent spaming the feed providers. With a long enough interval, you could theoretically run the agg command in a background process and interact with other commands in a seperate shell. 
10. You can view the messages in the feeds that you are following: `<executable> browse <limit>` where <limit> is optional (default is 2), specified max number of messages to display. Messages are displayed from recent to longer ago.
11. You can force the DB to be cleared: `<executable> reset` This command clears all tables in the db. Note: args are ignored.

