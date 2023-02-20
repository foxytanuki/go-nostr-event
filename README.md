# go-nostr-event

A Go library that helps generating Event of [go-nostr](https://github.com/nbd-wtf/go-nostr)

### Original

```
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


### With go-nostr-event

```
sk := nostr.GeneratePrivateKey()
cev := nostrevent.NewNote("Hello World!")
cev.SignPk(sk)
```

## Example script

```
go run example/post_note.go
```
