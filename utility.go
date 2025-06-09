package main

type APIResponseEnvelope struct {
	Success     bool   `json:"success"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	HTTPCode    int    `json:"http_code"`
	HTTPMessage string `json:"http_message"`
	Detailed    string `json:"detailed"`
	Count       int    `json:"count"`
	Pagination  string `json:"pagination"`
}

type Pagination struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type APIResponseEnvelopePagination struct {
	Success     bool       `json:"success"`
	Code        int        `json:"code"`
	Message     string     `json:"message"`
	HTTPCode    int        `json:"http_code"`
	HTTPMessage string     `json:"http_message"`
	Detailed    string     `json:"detailed"`
	Count       int        `json:"count"`
	Pagination  Pagination `json:"pagination"`
}
