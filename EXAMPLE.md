## Create config file

```shell
$ gdcli config init
Create Config File: /root/.gehirun.json.sample
$ mv /root/.gehirun.json.sample /root/.gehirun.json
$ vim /root/.gehirun.json
```

## Zone

#### List

```shell
$ gdcli zone ls
  example.info
     |- ID: 00000000-0000-0000-0000-000000000000
     |- CurrentVersionID: 00000000-0000-0000-0000-000000000000
     |- CreatedAt: 2018-03-16 08:22:38 +0000 UTC
     `- Editablet: true
  test.work
     |- ID: 00000000-0000-0000-0000-000000000000
     |- CurrentVersionID: 00000000-0000-0000-0000-000000000000
     |- CreatedAt: 2018-03-26 05:53:55 +0000 UTC
     `- Editablet: true
  infraengineer.net
     |- ID: 00000000-0000-0000-0000-000000000000
     |- CurrentVersionID: 00000000-0000-0000-0000-000000000000
     |- CreatedAt: 2018-10-05 07:25:40 +0000 UTC
     `- Editablet: true
```

```shell
$ gdcli zone ls example.info
  example.info
     |- ID: 00000000-0000-0000-0000-000000000000
     |- CurrentVersionID: 00000000-0000-0000-0000-000000000000
     |- CreatedAt: 2018-03-16 08:22:38 +0000 UTC
     `- Editablet: true
```

## Record

#### Supported record types

- A
- AAAA
- MX
- TXT
- CNAME

#### List

```shell
$ gdcli record ls example.info
  example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- Nsdname: ns2.gehirndns.net.
     |- Nsdname: ns2.gehirndns.jp.
     |- Nsdname: ns2.gehirndns.com.
     |- Nsdname: ns2.gehirndns.org.
     |- Type: NS
     |- TTL: 86400
     `- EnableAlias: false
  example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- Data: gehirn-dns-verification=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
     |- Type: TXT
     |- TTL: 3600
     `- EnableAlias: false
  www.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- IPAddress: 1.1.1.1
     |- Type: A
     |- TTL: 300
     `- EnableAlias: false
```

```shell
$ gdcli record ls example.info --name www
  www.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- IPAddress: 1.1.1.1
     |- Type: A
     |- TTL: 300
     `- EnableAlias: false
```

#### Add a new **A** record

```shell
$ gdcli record add example.info --type A --ip 1.1.1.1 --name f
  f.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- IPAddress: 1.1.1.1
     |- Type: A
     |- TTL: 300
     `- EnableAlias: false
```

#### Add a new **AAAA** record

```shell
$ gdcli record add example.info --type AAAA --name aaaa --ip 0000:0000:0000:000:a150:95:151:220e
  aaaa.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- IPAddress: ::a150:95:151:220e
     |- Type: AAAA
     |- TTL: 300
     `- EnableAlias: false
```

#### Add a new **MX** record

```shell
$ gdcli record add example.info --type MX --exchange mail --name mail
  mail.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- Priority: 10
     |- Exchange: mail.example.info.
     |- Type: MX
     |- TTL: 300
     `- EnableAlias: false
```

#### Add a new **TXT** record

```shell
$ gdcli record add example.info --type TXT --name txt --data example-txt-message
  txt.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- Data: "example-txt-message"
     |- Type: TXT
     |- TTL: 300
     `- EnableAlias: false
```

#### Add a new **CNAME** record

```shell
$ gdcli record add example.info --type CNAME --name cname --cname www
  wwww.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- CNAME: cname.example.info.
     |- Type: CNAME
     |- TTL: 300
     `- EnableAlias: false
```

#### Remove a record

```shell
$ gdcli record rm example.info --type A --name www
  www.example.info.
     |- ID: 00000000-0000-0000-0000-000000000000
     |- IPAddress: 1.1.1.1
     |- Type: A
     |- TTL: 300
     `- EnableAlias: false
```
