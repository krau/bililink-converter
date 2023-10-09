# Bililink Converter

B 站链接转换

## Features

- [x] av , bv 互转
- [x] b23.tv 短链接转换为 bv 链接, av 链接, iframe 链接 (并去除跟踪参数) 

## Usage

config.toml
```toml
Address = "127.0.0.1:39080" # 监听地址
```

Route

- GET /av2bv/av

- GET /bv2av/bv

- GET /b23/b23 (b23.tv/后的代码)

Docs

- GET /swagger/index.html
