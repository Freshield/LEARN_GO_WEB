/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: server_suite_test.py
@Time: 2021-11-17 14:38
@Last_update: 2021-11-17 14:38
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package main

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostTestSuite1 struct {
	mux *http.ServeMux
	post *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite1{})
}

func TestSuite(t *testing.T) {
	TestingT(t)
}

func (s *PostTestSuite1) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/post/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite1) TestGetPost(c *C) {
	request, _ := http.NewRequest("GET", "/post/1", nil)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	var post FakePost
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(post.Id, Equals, 1)
}

func (s *PostTestSuite1) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite1) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuite1) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}