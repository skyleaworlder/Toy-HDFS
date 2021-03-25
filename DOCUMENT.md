# DOCUMENT

## 一、各服务端口

```yaml
Client:
    rcv-DataNode: 12000
    rcv-NameNode: 12001
DataNode:
    rcv-Client: 12002
    rcv-DataNode-ACK: 12003
    rcv-NameNode-Beg-End: 12004
NameNode:
    rcv-Client-File-Query: 12005
    rcv-DataNode-Heart-Beat: 12006
```

## 二、chunk & packet & block

仍旧以这三者为基本的写入写出单位。

```yaml
buf:
    totalSize: 512*9
chunk:
    content: 512B
    checksum: 4B
    totalSize: 516B
packet:
    chunkNum: 126
    totalSize: 65016B
```
