# todo-cli <div style="position: absolute; left: 112px; top: -14px">![][go-logo]</div>

![][preview-img]

## Features

* Add tasks to remember them later.
* Do or undo tasks, after it remove them.
* Count how many tasks are remaining.

## Installation

First, import this project with `git clone`:

```
git clone https://github.com/luisrojass/todo-cli.git
```

Now, install all dependencies of this project:
```
go get
```

Then, install this cli on your system:
```
go install
```

After that, you can use the cli by the command `todo`.

To show what arguments are available, just write this same command: `todo`.
```
$ todo

Usage:
  ls                    Show all tasks
  add <args> [string]   Add a task, no quotes, no limit of words
  do <args> [int]       Mark selected task as completed
  undo <args> [int]     Mark selected task as incompleted
  rm <args> [int]       Delete selected task
  help                  Get more information
```

## FAQ

After installing, if you can't use the cli or you see this error or similar:
```
Command 'todo' not found, but can be installed with:
sudo apt install devtodo
```

Then, you can add the following line to your `.bashrc` or equivalent file.
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Contributing

* You can report bugs or suggest ideas through a [issue][issue-link].
* Or clone this repository and make a [pull request][pull-request-link].

## License

Copyright (c) 2023 Luis Rojas S.

[MIT license][license-link]

[go-logo]: https://www.vectorlogo.zone/logos/golang/golang-icon.svg

[preview-img]: https://res.cloudinary.com/dda2colxy/image/upload/v1683943506/github/readmes/todo-cli/github-todo-cli-preview_likdwp.png

[issue-link]: https://github.com/luisrojass/todo-cli/issues

[pull-request-link]: https://github.com/luisrojass/todo-cli/pulls

[license-link]: https://github.com/luisrojass/todo-cli/blob/main/LICENSE
