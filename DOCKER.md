## Docker

### Dockerイメージをビルド

```shell
$ docker build -t ubuntu2004/go116x .
```

### Dockerコンテナを作成

```shell
$ docker run -i --name gdcli -v `pwd`:/root/go/src/github.com/gdcli -t ubuntu2004/go116x
```

### Dockerコンテナを起動

```shell
$ docker start gdcli
```

### Dockerコンテナに接続

```shell
$ docker attach gdcli
```
