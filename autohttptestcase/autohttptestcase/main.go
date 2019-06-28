package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	f, err := os.OpenFile(`D:\github\kill-that-programmer\autohttptestcase\tcpstreamdata\172.16.64.92_3555_1561704446.5528667.txt`,
		os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	for {
		fmt.Println("----------------")
		err := readrequest(bufReader)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			continue
		}
		err = readresponse(bufReader)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
	}
}

func readrequest(reader *bufio.Reader) error {
	req, err := http.ReadRequest(reader)
	if err != nil {
		fmt.Println(err)
		return err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer req.Body.Close()

	s := len(body)
	fmt.Println("req body len is ", s)

	if req.ContentLength != int64(s) {
		fmt.Println("sble ,contentlent is not ", req.ContentLength, s)
	}
	return nil
}

func readresponse(reader *bufio.Reader) error {
	rsp, err := http.ReadResponse(reader, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rsp.Body.Close()

	bodyrsp, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	srsp := len(bodyrsp)
	fmt.Println("rsp body len is ", srsp)
	if rsp.ContentLength != int64(srsp) {
		fmt.Println("sble ,contentlent is not ", rsp.ContentLength, srsp)
	}
	return nil
}
