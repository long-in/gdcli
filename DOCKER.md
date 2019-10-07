## Docker

### Dockerイメージをビルド

```shell
$ docker build -t centos7/go .
```

### Dockerコンテナを作成

```shell
$ docker run -i --name gdcli -v `pwd`:/root/go/src/github.com/gdcli -t centos7/go
```

### Dockerコンテナを起動

```shell
$ docker start gdcli
```

### Dockerコンテナに接続

```shell
$ docker attach gdcli
```
