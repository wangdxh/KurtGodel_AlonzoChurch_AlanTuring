package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Do(req *http.Request) ([]byte, error) {
	return DoWithClient(req, http.DefaultClient)
}

// DoWithClient executes an HTTP request and returns the response body.
// Any errors or non-200 status code result in an error.
func DoWithClient(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	return body, nil
}

// https://blog.csdn.net/liyunlong41/article/details/87932953
func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.37.133:6831",
		},
		ServiceName: service,
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func foo1(req string, ctx context.Context) (reply string) {
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo1")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("requestfoo1", req)
		span.SetTag("replyfoo1", reply)
		span.Finish()
	}()
	span.LogKV("sble", "hhhh", "status", 2000)
	println(req)
	//2.模拟处理耗时
	time.Sleep(time.Second / 2)
	//3.返回reply
	reply = "foo1Reply"
	return
}

func foo2(req string, ctx context.Context) (reply string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo2")
	defer func() {
		span.SetTag("requestfoo2", req)
		span.SetTag("replyfoo2", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second / 2)
	reply = "foo2Reply"
	return
}

func testdefault() {
	tracer, closer := InitJaeger("jaeger-default")
	defer closer.Close()
	//
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("span_root")
	defer span.Finish()
	span.LogKV("you sb", "good", "status", 200)

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo1("Hello foo1", ctx)
	r2 := foo2("Hello foo2", ctx)
	fmt.Println(r1, r2)
	span.LogKV("sble", "hehehe jiushi sb")
}

func systeminfo1(req string, parentSpan opentracing.Span) (reply string) {
	//span, _ := opentracing.StartSpanFromContext(ctx, "span_foo4")
	fmt.Println("systeminfo cxt :", parentSpan.Context())
	fmt.Println("son", opentracing.ChildOf(parentSpan.Context()))
	fmt.Println("son2", g_systrace)

	span := g_systrace.StartSpan("systeminfo1", opentracing.ChildOf(parentSpan.Context()))
	defer func() {
		span.SetTag("system start", req)
		span.Finish()
	}()

	println(req)
	span.LogKV("systeminfo1", "i am")
	time.Sleep(time.Second / 2)
	reply = "system start"
	return
}

func systeminfo2(req string, parentSpan opentracing.Span) (reply string) {
	fmt.Println("systeminfo cxt :", parentSpan.Context())
	fmt.Println("son", opentracing.ChildOf(parentSpan.Context()))
	fmt.Println("son2", g_systrace)

	span := g_systrace.StartSpan("systeminfo2", opentracing.ChildOf(parentSpan.Context()))
	defer func() {
		span.SetTag("system down", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second / 2)
	span.LogKV("systeminfo2", "i am")
	reply = "system down"
	return
}

var g_systrace opentracing.Tracer
var g_closer io.Closer

func testsystemlog() {
	fmt.Println("systemtracer", g_systrace)
	span := g_systrace.StartSpan("mysystem")
	defer span.Finish()

	span.LogKV("you sb", "good", "status", 200)
	//ctx2 := opentracing.ContextWithSpan(context.Background(), span2)
	fmt.Println(g_systrace)
	systeminfo1("systeminfo1", span)
	systeminfo2("systeminfo2", span)

	span.LogKV("root", "i am")
}

func testfromstringspan() {
	span := g_systrace.StartSpan("local")

	span.SetBaggageItem("aaa", "bbb")

	span.LogKV("test from string", span.Context().(jaeger.SpanContext).String())
	ctx, err := jaeger.ContextFromString(span.Context().(jaeger.SpanContext).String())
	if err != nil {
		fmt.Println(err)
		return
	}

	spanstr := g_systrace.StartSpan("local2", opentracing.FollowsFrom(ctx))
	fmt.Println(spanstr)

	spanstr.LogKV("thisisson", "test")
	spanstr.SetTag("good", "ok")
	spanstr.Finish()

	spanfollow := g_systrace.StartSpan("local3", opentracing.FollowsFrom(ctx))
	spanfollow.LogKV("thisisson3", "test")
	spanfollow.SetTag("good23", "ok")
	spanfollow.Finish()

	spanstr4 := g_systrace.StartSpan("local4", opentracing.ChildOf(spanstr.Context()))
	fmt.Println(spanstr4)

	spanstr4.LogKV("thisisson", "test4")
	spanstr4.SetTag("good", "ok")
	spanstr4.Finish()

	span.Finish()
}

func testhttp() {
	http.HandleFunc("/getperson/", handleGetPerson)
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()
	time.Sleep(3 * time.Second)

	starthttprequest()
}
func starthttprequest() {
	operationName := "getperson"
	url := "http://127.0.0.1:8081/getperson/aaawxh"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	span := g_systrace.StartSpan(operationName)
	defer span.Finish()

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
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

func handleGetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)

	spanCtx, _ := g_systrace.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := g_systrace.StartSpan(
		"/getperson",
		opentracing.ChildOf(spanCtx),
	)
	ext.HTTPStatusCode.Set(span, 200)
	defer span.Finish()

	//ctx := opentracing.ContextWithSpan(r.Context(), span)

	name := strings.TrimPrefix(r.URL.Path, "/getperson/")
	/*person, err := repo.GetPerson(ctx, name)
	if err != nil {
		span.SetTag("error", true)
		span.LogFields(otlog.Error(err))
		http.Error(w, "error ", http.StatusInternalServerError)
		return
	}*/

	span.LogKV("name", name,
		"title", name,
		"description", name)
	bytes, _ := json.Marshal("great this is good")
	w.Write(bytes)
}

func main() {
	testdefault()

	g_systrace, g_closer = InitJaeger("systeminfo")
	testsystemlog()
	testfromstringspan()
	testhttp()

	g_closer.Close()
}

// should get information from
// https://github.com/opentracing/opentracing-go
// https://www.lizenghai.com/archives/6130.html
/*
main info
https://yq.aliyun.com/articles/514488?utm_content=m_43347
开放分布式追踪（OpenTracing）入门与 Jaeger 实现

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

http://192.168.37.133:16686


https://github.com/PacktPublishing/Mastering-Distributed-Tracing/
*/
