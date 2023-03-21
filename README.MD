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
#global:
#  enableauth: true # 开启鉴权（默认开启）
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

开发者可以根据鉴权回调请求到API接口上的参数进行鉴权，响应HTTP状态码为`200`即为放行，其他状态均为拒绝。