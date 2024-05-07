## Migoro

- [x] Enums for all env vars
- [x] Adapter to expose Validate ENV function
- [x] Automatically Create a database on init if not existing
- [x] How do I define which schema to use fo db migrations
- [x] Implement Adapter pattern for Postgres and SQLite
- [x] Move Queries in to the adapter functions
- [] Put it in the container and see what happens, docker log collector output
- [] Possibly use `log` from standard library and set output to stdout?? or else
- [] Have a way of outputting queries in the logs
- [] Weird behavior when having commented out migration code in a file, executes a comment and inserts migration log ?? 🤔
- [] Migration File creation have separate file for up and down - ?? Not sure ?? 🤔
- [] What about empty rollback SQLx snippets ??? 🤔 I think this is ok...
