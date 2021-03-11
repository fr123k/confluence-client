[![Build Status](https://travis-ci.com/fr123k/confluence-client.svg?branch=main)](https://travis-ci.com/fr123k/confluence-client)

# confluence-client

## Targets

### Build

The following command will build the golang binary and run the unit tests.
The result of this build step is the standalone binary in the `./build/` folder. 

```
make build
```

Example Output:
```
  go build -o build/confluence-cli cmd/confluence-cli.go
  go test -v --cover ./...
  ?   	github.com/fr123k/confluence-client/cmd	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/cmd	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/config	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/confluence	[no test files]
```

### Run

The following make target will first build and then execute the golang binary.
```
  make run
```

Example Output:
```
  go build -o build/confluence-cli cmd/confluence-cli.go
  go test -v --cover ./...
  ?   	github.com/fr123k/confluence-client/cmd	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/cmd	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/config	[no test files]
  ?   	github.com/fr123k/confluence-client/pkg/confluence	[no test files]
  ./build/confluence-cli
  NAME:
    consluence-cli - A new cli application

  USAGE:
    confluence-cli [global options] command [command options] [arguments...]

  VERSION:
    0.0.1

  DESCRIPTION:
    A conflucnce cli application.

  AUTHOR:
    fr123k <fr123k@yahoo.de>

  COMMANDS:
    page     Create/Read/Update/Delete conflunece pages.
    label    Create/Read/Update/Delete conflunece labels.
    help, h  Shows a list of commands or help for one command

  GLOBAL OPTIONS:
    --confluence.url value, --url value            The root url of the confluence instance like ('https://examplecompany.atlassian.net/wiki') [$CONFLUENCE_URL]
    --confluence.username value, --username value  The confluence username for authentication. [$CONFLUENCE_USERNAME]
    --confluence.token value, --password value     The confluence user api-token or password for authentication. [$CONFLUENCE_PASSWORD]
    --debug, -d                                    This enables verbose logging. (default: false) [$CONFLUENCE_DEBUG, $DEBUG]
    --config-file FILE, -c FILE                    Load configuration from FILE (default: "config.yaml") [$CONFLUENCE_CLI_CONFIGFILE, $CONFIGFILE]
    --help, -h                                     show help (default: false)
    --version, -v                                  print the version (default: false)
```

### Clean

The following make target will remove the `./build/` folder.
**No confirmation needed**
```
  make clean
```

Example Output:
```
  rm -rfv ./build
  ./build/main
  ./build
```

# Changelog

* setup travis build
* implement following commands page get, page search and label get


# Todos

* adding additional operations for page like update, delete, add label, remove label...
* adding json export capabilities
