# 本工具基于原项目v1.0.26版本改造

### 增加了请求类型:option

```
service sdkapi-api {
	@handler OptionsHandler
	options /js-sdk();
}
```

### 配合修改过后的工程，增加路径通配

```
type AssetsRequest struct {
	FileName string `path:"name,optional"`
}

type AssetsResponse struct {
	Content string            `json:"-"`
	Headers map[string]string `json:"-"`
}

service api.dbmis-api {
  @handler GreetHandler
  get /greet/from/:name(Request) returns (Response);
  @handler AssetsHandler
  get /:name*(AssetsRequest) returns(AssetsResponse);
  @handler AssetsHandler
  get /(AssetsRequest) returns(AssetsResponse);
}
```

上面的例子中 / 和 /:name* 都可以被识别到，*号为通配符，如果不满足上面的绝对路径，则按照顺序匹配下面的路径，注意顺序