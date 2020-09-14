# BLOGM

Blogm is a tiny microblog engine.



## Prerequisites

Tested with `go1.15` on `windows` and `linux`. There is no reason it doesn't work on `MacOS`.



## Installation

Download the default binary, or build one from source. Building from source let you include your own templates and style.

To build from source, you just have to run the `build.bat` or `build.sh` script located inside the project.

This result in a binary. I advise you to put that binary in a folder where the `PATH` environment variable points to, to just have to run `blogm` command to use it.



## Setting up your own blog

You know have to go to an empty directory. This is where your blog is going to be set up.

Use :

```shell
blogm init
```

to create all the files and folder needed.



Now run :

```shell
blogm server start
```

To start the http server.



## Configure your blog

The `config.json` file contains all the configurations.

You can edit the differents messages, menu links, blog name, etc...

You just have to stop and restart the http server to make the updates effectives.



## Publish posts and pages

To publish a post or a page, you just have to put a `.md` markdown file into the `posts` or `pages` directory.



The `pages` are accessible at the `/pages/...` URL.

The `posts` are accessible at the `/posts/...` URL.

You can list the differents posts sorted by modification date at the `/posts` or `/posts/` URL.



Note that the `...` in the URL is simply the markdown file name, without the terminating `.md`.



Posts title (displayed on the list posts page, for instance), is simply the file name with spaces replacing underscores.



## Notes

- I know the README isn't that good. I'm working on improving it.

- At the moment, BLOGM is meant to run behind a reverse proxy such as `nginx`, as it doesn't support SSL/TLS (yet).
- You can use `BLOGM` alongside `fs-server`  which can enable you to login, upload, edit or remove posts and pages online.