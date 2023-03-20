# Translate-Server
一个基于`Golang`构建的集成翻译api服务端。
An integrated translation api server built on `Golang`.
<div>
  <a href="https://goreportcard.com/report/github.com/shevonkuan/translate-server">
    <img src="https://goreportcard.com/badge/github.com/ShevonKuan/translate-server" alt="latest version" />
  </a>
  <a href="https://github.com/ShevonKuan/translate-server/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/ShevonKuan/translate-server" alt="License" />
  </a>
  <a href="https://github.com/ShevonKuan/translate-server/actions?query=workflow%3ABuild">
    <img src="https://img.shields.io/github/actions/workflow/status/ShevonKuan/translate-server/build.yml" alt="Build status" />
  </a>
  <a href="https://github.com/ShevonKuan/translate-server/releases">
    <img src="https://img.shields.io/github/release/ShevonKuan/translate-server" alt="latest version" />
  </a>
</div>
<div>
  <a href="https://github.com/ShevonKuan/translate-server/discussions">
    <img src="https://img.shields.io/github/discussions/ShevonKuan/translate-server?color=%23ED8936" alt="discussions" />
  </a>
  <a href="https://github.com/Xhofe/alist/releases">
    <img src="https://img.shields.io/github/downloads/ShevonKuan/translate-server/total?color=%239F7AEA&logo=github" alt="Downloads" />
  </a>
  <!-- <a href="https://hub.docker.com/r/xhofe/alist">
    <img src="https://img.shields.io/docker/pulls/xhofe/alist?color=%2348BB78&logo=docker&label=pulls" alt="Downloads" />
  </a> -->

<img src="https://img.shields.io/github/go-mod/go-version/Shevonkuan/translate-server">
</div>

## Description
### Support Translate Engine
- [x] [DeepL](https://www.deepl.com/translator)
- [x] [Google](https://translate.google.com)
- [ ] [tencent](https://fanyi.qq.com/)
- [ ] [baidu](https://fanyi.baidu.com/)
- [ ] [youdao](https://fanyi.youdao.com/)
- [ ] others.. 

## Usage
### Translate API
#### Request
- method: `POST`
- url: `/translate` or `https://translate-server-five.vercel.app/api/translate`
- params: engine: `google` or `deepl`(`engine` param is only avaliable in the local server instead of the vercel one)
- body application/json: 
```json
{
  "text": "Hello world",
  "source_lang": "en",
  "target_lang": "zh"
}
```

#### Response
```json
{
  "alternatives": [
    "你好世界",
    "世界你好"
  ],
  "code": 200,
  "data": "你好世界",
}

```
### Translate XML
always used to translate RSS feed.

origin rss feed url: `http://export.arxiv.org/rss/cs.DC`
translated rss feed url: `https://translate-server-five.vercel.app/api/rss?url=http://export.arxiv.org/rss/cs.DC`

the default engine is google. if you want to specify the engine, you should run your own instance of the server instead of using the vercel one, like this:
`http://127.0.0.1:1188/rss?url=http://export.arxiv.org/rss/cs.DC&engine=deepl`


### Run with Docker

// TODO


## Author
**Shevon Kwan** © [translate-server Contributors](https://github.com/ShevonKuan/translate-server/contributors)