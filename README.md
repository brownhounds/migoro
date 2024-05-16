## Migoro

- [x] Enums for all env vars
- [x] Adapter to expose Validate ENV function
- [x] Automatically Create a database on init if not existing
- [x] How do I define which schema to use fo db migrations
- [x] Implement Adapter pattern for Postgres and SQLite
- [x] Move Queries in to the adapter functions
- [x] Put it in the container and see what happens, docker log collector output
- [x] What about empty rollback SQLx snippets ??? 🤔 I think this is ok...
- [x] Migration File creation have separate file for up and down
- [x] Weird behavior when having commented out migration code in a file, executes a comment and inserts migration log ?? 🤔
- [x] Do I need an injection of ENV Vars in to migrations ?? 🤔
- [x] Add version Command
- [] Do I want to use configurable logger ?? 🤔
  - [] Possibly use `log` from standard library and set output to stdout?? or else

## Testing

- [] Write a simple snapshot library - something like jest has 👍
  - I want an interface to `toMatchSnapshot`
  - give name of the file
  - pass a pointer to t testing library
  - make automatic assertion
  - detect if snapshot needs updating before asserting shit!
- [] Do I want to have single snap per file, I say yes
- [] Make it configurable, give it a dir via ENV var, fuck windows don't worry about!!

## Learning

- [x] Pointers
- [x] Pointer to interfaces instead of concrete types 🤔
- [x] Structures Comparison
- [x] File Embed - https://stackoverflow.com/questions/13904441/how-can-i-bundle-static-resources-in-a-go-program
- [] Error Handling
- [] Go publishing packages not a seamless as others

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
