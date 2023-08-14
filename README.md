### Pulsar
API Backend for IGAG

##### Build Pulsar
```shell
go build -o pulsar.exe -ldflags '-s -w'
```

Create Post
```shell
curl --request POST \
     --url 'http://localhost:8080/posts'\
     --data '{"title": "Title", "body": "body"}'
```