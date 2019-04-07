# Experimental CMS with Buffalo!

[![Build Status](https://api.travis-ci.org/akosgarai/buffalo_example.svg?branch=master)](https://travis-ci.org/akosgarai/buffalo\_example)

CMS project. Administrator/content/user management tool. Experimental project.

# Start project from scratch

How to install & start the application and it's dependencies.

## Go - linux

From [https://golang.org/doc/install](https://golang.org/doc/install)

### Download & install go

- Download the go version from [here](https://golang.org/dl/). (You need at least 1.10.8)
- Extract it into `/usr/local` (root or sudo)

```bash
tar -C /usr/local -xzf go1.11.2.linux-amd64.tar.gz
```

- Add `/usr/local/go/bin` to the `PATH` environment variable. (you can paste it into your .bash\_profile/.bashrc/.zshrc/etc... file & source it.)

```bash
export PATH=$PATH:/usr/local/go/bin
```

- Try to run go version command to make sure, that the installation was successful.

```bash
$ go version
go version go1.11.2 linux/amd64
```

### Setup GOPATH environment variable

From [https://github.com/golang/go/wiki/SettingGOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
Let's assume, that you have a `go` directory in your `HOME` folder, and you want to develop your gocode under that dir. Make sure that you have a `bin` and a `src` dir. under the `go` dir.
Add the following lines in your bash/zsh rc/profile file, and then source it.

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### Setup Dep dependency manager tool

From [https://github.com/golang/dep](https://github.com/golang/dep)

- Install with apt-get

```bash
sudo apt-get install go-dep
```

- Or install with script

```bash
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

### Buffalo

From [https://gobuffalo.io/en/docs/installation](https://gobuffalo.io/en/docs/installation)

#### Frontend dependencies

- Node -gte 8 - [how to install](https://github.com/nodejs/node)
- NPM - [info](https://github.com/npm/cli)

#### Install buffalo

```bash
$ curl -OL https://github.com/gobuffalo/buffalo/releases/download/v0.14.2/buffalo_0.14.2_linux_amd64.tar.gz
$ tar -xvzf buffalo_0.14.2_linux_amd64.tar.gz
$ sudo mv buffalo /usr/local/bin/buffalo
```

### Download this application

You have a bunch of options for getting an application from github. Due to the golang's dependencies, you have to make sure, that this application is donwloaded to the right place.

- go get

```bash
$ go get github.com/akosgarai/buffalo_example
$ cd $GOPATH/src/github.com/akosgarai/buffalo_example
```

- git clone

```bash
$ mkdir -p $GOPATH/src/github.com/akosgarai
$ cd $GOPATH/src/github.com/akosgarai
$ git clone git@github.com:akosgarai/buffalo_example.git
```

## Setup postgres db

My buffalo example application uses postgres database, so that we need a working db server instance, that the app can use.
From [https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-16-04)

- Installation

```bash
$ sudo apt-get update
$ sudo apt-get install postgresql postgresql-contrib
```

- DB management - switch to postgres user

```bash
$ sudo -i -u postgres
```

- Starting psql app (postgres user)

```bash
$ psql
```

In psql shell, you are able to create users/databases, so that you can setup everything for your needs.

## Setup buffalo\_example application

Before the first start, you have to setup the database that contains the necessary tables, and the initial values (eg admins).

```bash
$ buffalo db create -a
$ soda migrate up
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```bash
$ buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should be redirected to the [http://127.0.0.1:3000/login](http://127.0.0.1:3000/login) page.

# About the tests

Good to know

- Use the built in tool: `buffalo test`
- Don't be logged in to the database when you want to run the tests

[Powered by Buffalo](http://gobuffalo.io)
