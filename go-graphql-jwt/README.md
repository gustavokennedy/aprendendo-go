# Project initialization with gqlgen

```bash
$ mkdir PROJECT_NAME
$ cd PROJECT_NAME
$ go mod init github.com/[username]/PROJECT_NAME
$ go get github.com/99designs/gqlgen
```

## Create the project skeleton

```bash
$ go run github.com/99designs/gqlgen init
```


## At the top of our resolver.go, between package and import, add the following comment:

```golang

package graph

//go:generate go run github.com/99designs/gqlgen

import "gitlab.com/pragmaticreviews/graphql-go/graph/model" 

```

## Run this every time you change the GraphQL Schema

```bash
go generate
```