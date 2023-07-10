# Parsing custom configs

This small example aims to show how we can use Golang text/template to easily create
different config files (in this case, for BIRD) by providing a JSON which contains
relevant fields that will be applied into the "bird.conf" template.  

As for now, this example prints the result in stdout, but it can be changed to generate
a new file or even send the result as an HTTP request's response.

## How to run

### Without Golang installed

A binary was provided in this directory's `bin/` folder. Running it:

```sh
// Only need to run the following command once, to grant permission
chmod +x ./bin/fileparser

// Running the binary
./bin/fileparser
```

### Installing Golang

Check the [official documentation](https://go.dev/doc/install) on how to install Go.  
Once installed, simply run:  

```sh
go run .
```

