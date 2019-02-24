# Experimental CMS with Buffalo!

CMS project. Administrator/content/user management tool. Experimental project.

## System setup

- go, npm, buffalo, dep, etc.

## Database Setup

- postgres db, user, password with the necessary grants. (database.yml)

### Create Your Databases

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



[Powered by Buffalo](http://gobuffalo.io)
