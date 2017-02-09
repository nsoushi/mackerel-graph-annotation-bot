# mackerel Graph Annotation tool

This repository contains the commnd line tool for the [macarerel Graph Annotations](https://mackerel.io/api-docs/entry/graph-annotations).
See also [release entry](https://mackerel.io/blog/entry/weekly/20170127).

## Usage
```
Tne post command creates graph annotations via graph-annotations API.

Usage:
  graph-annotations post [flags]

Examples:
graph-annotations post -s 'deploy application' ExampleRole1 ExampleRole2
graph-annotations post --service ExampleService -s 'deploy application' ExampleRole1 ExampleRole2

Flags:
      --description string   [optional] annotation details
      --service string       required: service name, when it's empty value then will use enviroment variable of 'MACKEREL_SERVICE_NAME'' (default "local")
  -s, --title string         required: annotation title
```

## How to install

```
$ go get github.com/nsoushi/mackerel-graph-annotation-bot/mackerel/cmd/graph-annotations
$ source $GOPATH/src/github.com/nsoushi/mackerel-graph-annotation-bot/setup.sh <MACKEREL_API_KEY> <MACKEREL_SERVICE_NAME>
```

then you can use following.
```
$ graph-annotations post -s 'Annotation title' ExampleRole1 ExampleRole2
```

## Post annotation via hubot

See also [graph-annotations.coffee](https://github.com/nsoushi/mackerel-graph-annotation-tool/blob/master/hubot/scripts/graph-annotations.coffee) file.
This repository contains bot script. 
If you include script to your bot, then you can post via bot that using following message.

```
@yourbot note -s 'sent push notifications.' web
```

