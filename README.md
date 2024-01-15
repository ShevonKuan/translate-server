# Translate-Server

![translate-server](https://socialify.git.ci/ShevonKuan/translate-server/image?description=1&descriptionEditable=%E4%B8%80%E4%B8%AA%E5%9F%BA%E4%BA%8EGolang%E6%9E%84%E5%BB%BA%E7%9A%84%E9%9B%86%E6%88%90%E7%BF%BB%E8%AF%91api%E6%9C%8D%E5%8A%A1%E7%AB%AF%E3%80%82%20An%20integrated%20translation%20api%20server%20built%20on%20Golang.&font=Jost&forks=1&issues=1&language=1&name=1&owner=1&pattern=Solid&pulls=1&stargazers=1&theme=Light)

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
  <a href="https://github.com/ShevonKuan/translate-server/releases">
    <img src="https://img.shields.io/github/downloads/ShevonKuan/translate-server/total?color=%239F7AEA&logo=github" alt="Downloads" />
  </a>
    <a href="https://github.com/ShevonKuan/translate-server/releases">
    <img src="https://img.shields.io/github/v/release/shevonkuan/translate-server?include_prereleases&label=pre-release" alt="Downloads" />
  </a>
  <!-- <a href="https://hub.docker.com/r/xhofe/alist">
    <img src="https://img.shields.io/docker/pulls/xhofe/alist?color=%2348BB78&logo=docker&label=pulls" alt="Downloads" />
  </a> -->

<img src="https://img.shields.io/github/go-mod/go-version/Shevonkuan/translate-server">
</div>

## Description

### Support Translate Engine

-   [x] [DeepL](https://www.deepl.com/translator)
-   [x] [Google](https://translate.google.com)
-   [ ] [tencent](https://fanyi.qq.com/)
-   [ ] [baidu](https://fanyi.baidu.com/)
-   [ ] [youdao](https://fanyi.youdao.com/)
-   [ ] others..

## Usage

### Translate API

#### Request

-   method: `POST`
-   url: `/translate` or `https://translate-server-five.vercel.app/api/translate`
-   params: engine: `google` or `deepl`
-   body application/json:

```json
{
    "text": "Hello world",
    "source_lang": "en",
    "target_lang": "zh"
}
```

##### Response

```json
{
    "alternatives": ["你好世界", "世界你好"],
    "code": 200,
    "data": "你好世界"
}
```

#### Request(JSONP)

You can use JSONP to request the api, just add `callback` param to the url, like this:

-   method: `POST`
-   url: `/translate` or `https://translate-server-five.vercel.app/api/translate?callback=tr`
-   params:
    -   engine: `google` or `deepl`
-   body application/json:

```json
{
    "text": "Hello world",
    "source_lang": "en",
    "target_lang": "zh"
}
```

##### Response

```javascript
tr({
    alternatives: ["你好世界", "世界你好"],
    code: 200,
    data: "你好世界",
});
```

### Translate XML

Always used to translate RSS feed:

1. Add prefix `https://translate-server-five.vercel.app/api/rss?url=` to the original rss feed url. e.g.
   `http://export.arxiv.org/rss/cs.DC`
   ->
   `https://translate-server-five.vercel.app/api/rss?url=http://export.arxiv.org/rss/cs.DC`
2. Add `&engine=deepl` or `&engine=google` to the end of the url. The default engine is google. if you want to specify the engine, you should run your own instance of the server instead of using the vercel one, like this:
   `http://127.0.0.1:1188/rss?url=http://export.arxiv.org/rss/cs.DC&engine=deepl`

### Run with Docker

// TODO

## Author

**Shevon Kwan** © [translate-server Contributors](https://github.com/ShevonKuan/translate-server/contributors)
