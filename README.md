# Go script template

This is a simple GoLang script template.

* CLI example
* CSV read
* CSV write

## How to run

```sh
go run main.go
# or
go run main.go --help
```

## How to create project

* Add main.go file
* write the script

### Create Go Module

```sh
# init mod and go version
go mod init go-script-template
# add module requirements and sums
go mod tidy
```

## More info

### GoLang version manager (asdf)

* https://github.com/kennyp/asdf-golang

## 25 Keywords in GoLang

> Reference: https://medium.com/wesionary-team/know-about-25-keywords-in-go-eca109855d4d

25 Keywords or Reserved words are the words in a language that are used for some internal process or represent some predefined actions.

All 25 keywords in 4 groups as follows so that we will able to know the purpose of the keyword:

| Declaration | Composite Types | Control Flow |        | Function Modifier |
|-------------|-----------------|--------------|--------|-------------------|
| const       | chain           | break        | goto   | defer
| var         | interface       | case         | if     | go
| func        | map             | continue     | range  |
| type        | struct          | default      | return |
| import      |                 | else         | select |
| package     |                 | fallthrough  | switch |
|             |                 | for          |        |

https://gobyexample.com/
