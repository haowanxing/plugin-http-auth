# HTTP鉴权插件

用于对推流和拉流客户端进行鉴权

## 插件地址

https://github.com/haowanxing/plugin-http-auth

## 插件引入

```go
import (
    _ "github.com/haowanxing/plugin-http-auth"
)
```

## 配置

```yaml
httpauth:
  onsubaddr: http://localhost:9091/subauth
  onpubaddr: http://localhost:9091/pubauth
```

## API回调请求

```json
{
    "action": "subscribe/publish",
    "app": "live",
    "stream": "test",
    "param": "field1=value1&field2=value2",
    "client_id": "127.0.0.1:53779"
}
```