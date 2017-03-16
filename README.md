### Reporting Tool
> This simple service was created to handle the reporting of daily user usage. These reports are generated on a daily basis and triggered on a given time. You can optionally filter on a front-end UI to get the data.

### Dependencies
* `"github.com/etowett/fiesta/core"` -  contains all the common methods used in `"github.com/etowett/fiesta/utils`
* `"github.com/go-sql-driver/mysql"` - Golang driver for MySQL
* `github.com/joho/godotenv` - Golang loader for environment variables from `.env`
> Most of the other dependencies are found in the import statements

### Getting Started
###### Description
* `main.go` is the main file that is run for execution i.e. `go run main.go`
* The `core` folder contains the common methods re-used throughout the project.
* The `config` folder houses configuration for use if you want to run this application as a service with `systemd`
* `utils` contains the `structs`, `interfaces` and `functions` that handle the receiving of requests to get data, parsing into required data structures and processing into the final step.

###### Setting up
* Ensure you've setup `Golang` and your `$GO_PATH` (if not, please Google :smirk: )
* Do `go get github.com/etowett/fiesta`. This will clone a copy of this repo into your `/path/to/golang/projects/src/github.com/etowett`
###### .env and environment variables
> Rename the `env.sample` file into `.env` and fill out the required variables

###### MySQL
> There are 2 scenarios for connection in MySQL. We use a remote DB so our connection looks like this in `github.com/etowett/fiesta/main.go`:
```
    common.DbCon, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":3306)/"+os.Getenv("DB_NAME")+"?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer common.DbCon.Close()
```

> If your DB is in your local server, it would be something like this:
```
    common.DbCon, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+os.Getenv("DB_NAME")+"?charset=utf8")
    if err != nil {
        panic(err.Error())
    }
    defer common.DbCon.Close()
```
Just remove the `tcp` part.

###### Logging
* Create a folder for the log file as specified in your `LOG_FILE` variable in `.env`
* Open a terminal and `tail -f /path/to/logs/fiesta.log`
> Ensure that this folder has the proper permissions to allow writing into it's files

###### Building and installing
* Run `go build` to build the app
* Run `go install`. Check the `/path/to/golang/projects/bin` folder and you should find a `fiesta` executable.
* Run `./fiesta` on one terminal.
* Open another terminal and run the python requests simulator file `client.py` i.e. `./client.py`
> You should see outputs of `Get usage from: ` and logs on the tailed logs terminal.

### TODOs
* Dockerize the app so that it runs in it's own container
* Use `systemd timers` for the daily polls to create and send reports

### Contributing
> Contributions are welcome. PRs will be accepted if they've followed atleast the minimum standards.

### Issues
> Reach out to:
* [Eutychus Towett](https://github.com/etowett)
* [Martin Musale](https://github.com/musale)
