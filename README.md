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

## Learning

- [x] Pointers
- [x] Pointer to interfaces instead of concrete types 🤔
- [x] Structures Comparison

Printing structures:

```go
	fmt.Printf("%+v\n", <StructValue>)
```

Execute Command and wait for to finish:

```go
	command := []string{"compose", "up", "-d"}
	cmd := exec.Command("docker", command...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating stderr pipe: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Error starting Docker Compose: %v", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Println(scanner.Text())
		}
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Error waiting for Docker Compose: %v", err)
	}

	log.Println("Docker Compose started successfully")
```

Running docker container in github action
`https://www.youtube.com/watch?v=U7TY_qUD8yA`
`https://stackoverflow.com/questions/57915791/how-to-connect-to-postgres-in-github-actions`

```yml
jobs:
  dump-database:
    runs-on: ubuntu-latest
    container: ubuntu
    services:
      anonymizer:
        image: anonymizer:latest
        env:
          POSTGRES_PASSWORD: ${{ secrets.DATABASE_PASSWORD }}
        ports:
          - 5432:5432
     steps:
       - name: Dump database
         run: |
           curl anonymizer:5432/dump > dump.sql
           head -n 150 dump.sql
```
