package main

import "net/http"

type TestCase struct {
	Cases map[string]CaseDetail
}

func NewTestCase(casename string, detail CaseDetail) *TestCase {
	return &TestCase{
		Cases: map[string]CaseDetail{casename: detail},
	}
}

type CaseDetail struct {
	Priority       int          `json:"priority,string"`
	ParentTestCase string       `json:"parentTestCase"`
	Request        CaseRequest  `json:"request"`
	Response       CaseResponse `json:"response"`
}

type CaseRequest struct {
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Headers     http.Header       `json:"headers"`
	QueryString map[string]string `json:"querystring"`
	Payload     string            `json:"payload"`
}

type CaseResponse struct {
	Status map[string]string `json:"status"`
	Method map[string]string `json:"method"`
	Body   map[string]string `json:"body"`
}
