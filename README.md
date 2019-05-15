# 点播测试程序

点播白金版压测程序。

```sh
platinumc
    [-x ClientIdentifier]
    [-t ProtocolType] (tcp, udp, wss)
    [-a ServerAddress] (tcp, wss, udp or mona address)
    [-p PeerId] (optional, connect through mona if specified)
    [-f FileIndex] (must not be empty)
    [-b BlockIndex] (default = 0)
    [-o SavePath]
    [-s StartPieceIndex]
    [-n FetchPieceCount]
    [-v (VerboseLog)]
    [-c (CheckPiece)]
```

本工具支持使用三种协议从服务器拉取数据：

1. tcp - TCP
2. udp - RTMFP
3. wss - Secure WebSocket

选项说明：

| 选项 | 说明                          |
|-----|-------------------------------|
|`-x` | 客户端标识符，便于在插件上定位该客户端连接    |
|`-t` | 协议类型，必须为`tcp`，`udp`和`wss`三者之一 |
|`-a` | 服务器地址，格式为`IP:Port` |
|`-p` | RTMFP对应的`PeerID` |
|`-f` | 请求的文件唯一ID |
|`-b` | 请求的文件块号 |
|`-o` | 保存路径，不提供此值时不保存 |
|`-s` | 请求的起始`Piece`编号 |
|`-n` | 从起始编号起取几个`Piece` |
|`-v` | 输出详细日志 |
|`-c` | 校验`Piece` |

**注意：**当`-t`为`tcp`或`wss`时，`-a`指的是插件的对应的服务地址；
当`-t`为`udp`时，`-a`的意义需要通过`PeerID`来确定：

1. `-p`不为空，则`-a`表示的是Mona的地址；
2. `-p`为空，则`-a`表示的是插件的地址；

示例：

```sh
# 1. 使用TCP拉取数据
platinumc -x tcp001 -t tcp -a '192.168.200.28:59606' -f 'scdn.00cdn.com/p2p/test/1.mp4' -b 0

# 2. 使用WSS拉取数据
platinumc -x wss001 -t wss -a '192.168.200.28:59843' -f 'scdn.00cdn.com/p2p/test/1.mp4' -b 0

# 3. 使用RTMFP直接连接拉取数据
platinumc -x udp001 -t udp -a '192.168.200.28:59608' -f 'scdn.00cdn.com/p2p/test/1.mp4' -b 0

# 4. 使用RTMFP打洞连接拉取数据
platinumc -x udp002 -t udp -a '61.155.182.194:1942' -p '38cbe49e1c9166f4bb4311204f4a284e7b1e1f43f8cb263ce15ddabde463c36b' -f 'scdn.00cdn.com/p2p/test/1.mp4' -b 0

# 5. 开启校验，打开详细日志，同时保存文件
platinumc -x tcp002 -t tcp -a '192.168.200.28:59606' -f 'scdn.00cdn.com/p2p/test/1.mp4' -b 0 -cv -o /tmp/1.mp4
```