# go-nostr-event

A Go library that helps generating Event of [go-nostr](https://github.com/nbd-wtf/go-nostr)

<a href="https://godoc.org/github.com/foxytanuki/go-nostr-event"><img src="https://img.shields.io/badge/api-reference-blue.svg?style=flat-square" alt="GoDoc"></a>

## Overview
You can create and sign nostr events with less code

### with go-nostr-event (nostrevent)

```go
sk := nostr.GeneratePrivateKey()
cev := nostrevent.NewNote("Hello World!")
cev.SignPk(sk)
```

### original

```go
sk := nostr.GeneratePrivateKey()
pub, _ := nostr.GetPublicKey(sk)

ev := nostr.Event{
	PubKey:    pub,
	CreatedAt: time.Now(),
	Kind:      1,
	Tags:      nil,
	Content:   "Hello World!",
}

ev.Sign(sk)
```

## Getting started

```go
import nostrevent "github.com/foxytanuki/go-nostr-event"
```

or

```sh
go get -u github.com/foxytanuki/go-nostr-event
```

## Example script
[example/post_note.go](https://github.com/foxytanuki/go-nostr-event/blob/main/example/post_note.go
)

```sh
go run example/post_note.go
```


