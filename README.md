# pipeserve

Instantly serve any piped input over HTTP right from your terminal.

## Install

```bash
go install github.com/djosh34/pipeserve@latest
```

## Usage

```bash
# Serve on default port 8080
cat file.html | pipeserve

# Serve on specific port
cat file.json | pipeserve 9222
```

Any path on `localhost:PORT` will serve the piped file.
