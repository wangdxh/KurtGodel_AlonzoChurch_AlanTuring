#### background 
开放分布式追踪（OpenTracing）入门与 Jaeger 实现  
https://yq.aliyun.com/articles/514488?utm_content=m_43347

#### install jaeger
测试使用内存jaeger all-in-one 
  
docker run -d --name jaeger \  
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \  
  -p 5775:5775/udp \  
  -p 6831:6831/udp \  
  -p 6832:6832/udp \  
  -p 5778:5778 \  
  -p 16686:16686 \  
  -p 14268:14268 \  
  -p 9411:9411 \  
  jaegertracing/all-in-one:latest  
  
#### jaeger 结果查询
通过http 16686端口进行查询   

http://172.16.64.193:16686/search
http://192.168.37.133:16686   

 
#### 参考代码 
 https://github.com/PacktPublishing/Mastering-Distributed-Tracing/

#### 事件的格式
```
traceID ：spanID ：parentID：flag
6f3813b8b0b2b976:242106bdb7462cae:6f3813b8b0b2b976:1
6f3813b8b0b2b976:6f3813b8b0b2b976:0:1
```
串行化到http头为：  
Uber-Trace-Id: 20e84f3d5867f9c6:20e84f3d5867f9c6:0:1 

#### dependencies
go get github.com/opentracing/opentracing-go  
go get github.com/uber/jaeger-client-go  

#### 创建追踪对象
```go
func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
    //采样设置
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
    //设置服务器信息
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "172.16.64.193:6831",
		},
		ServiceName: service,
	}
  //创建tracer
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))		
}

```

#### 函数间调用
```go
	//创建追踪对象，可以在查询界面根据名称进行查询
	tracer, closer := InitJaeger("jaeger-default")
	defer closer.Close()
	
	//设置为默认的追踪对象，可以不指定追踪对象进行函数调用
	opentracing.SetGlobalTracer(tracer)
  
	//创建一个span的根对象，即一次追踪事件的根节点
	span := tracer.StartSpan("span_root")
	defer span.Finish()

	//针对span进行日志添加操作，使用key value的方式
	span.LogKV("you sb", "good", "status", 200)

	//创建一个go的context对象，将span对象设置到context中，其他函数可以从context中获取对span的访问
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo1("Hello foo1", ctx)
	r2 := foo2("Hello foo2", ctx)
	fmt.Println(r1, r2)
	span.LogKV("sble", "hehehe jiushi sb")
```

```go
func foo2(req string, ctx context.Context) (reply string) {
	//从context中的span对象中，创建一个新的子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo2")
  
	defer func() {
		//每个span可以设置log，还可以设置tag，在查询界面可以根据tag的名称和值，进行查询过滤    

		span.SetTag("requestfoo2", req)
		span.SetTag("replyfoo2", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second / 2)
	reply = "foo2Reply"
	return
}

```
#### 多个tracer
一个程序中，可以创建多个追踪对象，查询时根据追踪对象的名称进行过滤，也可以一个追踪对象，span的名称不同，根据span名称进行过滤  
```go
	//创建tracer
	var g_systrace opentracing.Tracer
	g_systrace, g_closer = InitJaeger("systeminfo")

	//通过特定tracer创建根span
	span := g_systrace.StartSpan("mysystem")
	defer span.Finish()

	span.LogKV("you sb", "good", "status", 200)
	//ctx2 := opentracing.ContextWithSpan(context.Background(), span2)
	fmt.Println(g_systrace)
	systeminfo1("systeminfo1", span)
	systeminfo2("systeminfo2", span)
```

```go
func systeminfo1(req string, parentSpan opentracing.Span) (reply string) {
	//根据span创建子的span
	span := g_systrace.StartSpan("systeminfo1", opentracing.ChildOf(parentSpan.Context()))
  
	//... 操作 span ...
	return
}
```

#### span串行化到字符串和反串行化
```go
	//创建新的span
	span := g_systrace.StartSpan("local")

	// span.Context().(jaeger.SpanContext).String()
	// 将span转换成jaeger的信息，然后串行化到字符串 6f3813b8b0b2b976:6f3813b8b0b2b976:0:1
	span.LogKV("test from string", span.Context().(jaeger.SpanContext).String())

	// jaeger.ContextFromString 从字符串重新构造jaeger的spancontext
	ctx, err := jaeger.ContextFromString(span.Context().(jaeger.SpanContext).String())
	
	// 根据spancontext就可以创造新的span，这里创造新的子span
	spanstr := g_systrace.StartSpan("local2", opentracing.FollowsFrom(ctx))
	spanstr.LogKV("thisisson", "test")
	spanstr.SetTag("good", "ok")
	spanstr.Finish()
```
#### http调用
客户端发起请求
```go
func starthttprequest() {
	// 创建http请求对象
	operationName := "getperson"
	url := "http://127.0.0.1:8081/getperson/aaawxh"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 由客户端创建根span
	span := g_systrace.StartSpan(operationName)
	defer span.Finish()

	// opentracing 提供了ext包，可以针对span方便的设置tag
	// 查询界面根据tag的值进行查询，一般会把http url，方法名称，发起方为client，
	// 以及返回结果是否成功，写入到tag内
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
  
	// 将span的信息串行化到http头内部
	g_systrace.Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	ret, err := Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("http return :", string(ret))
}
```
服务器进行处理，服务器处理span时，创建服务器span以及status的tag等信息，可以放置在过滤器内

```go
func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	//从http头中间提取出客户端发起的根span的信息，
	//如果没有找到的话，一般由服务器自主创建根span，这里没有处理这个逻辑 
	spanCtx, _ := g_systrace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	//根据客户端的span信息创建新的子span
	span := g_systrace.StartSpan(
		"/getperson",
		opentracing.ChildOf(spanCtx),
	)
	// 将处理结果值写入到服务器创建的span的tag
	ext.HTTPStatusCode.Set(span, 200)
	defer span.Finish()

	name := strings.TrimPrefix(r.URL.Path, "/getperson/")
	//对服务器生成的子span，添加日志信息
	span.LogKV("name", name,
		"title", name,
    "description", name)
    
	// 返回数据到客户端  
	bytes, _ := json.Marshal("great this is good")
	w.Write(bytes)
}

```
