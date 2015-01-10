package api

import "github.com/jmcvetta/napping"
import "github.com/apellizzn/lelylan-go/lelylan/errors"
import "github.com/apellizzn/lelylan-go/lelylan/devices"
import "net/http"

func prepareRequest(headers map[string]string) napping.Session{
	s := napping.Session{ Log: false }
	h := &http.Header{}
	for header := range headers {
		h.Set(header,headers[header])
	}
	s.Client = &http.Client{}
	s.Header = h
	return s
	}

func Get(url string, headers map[string]string) (*napping.Response, errors.LelylanHttpFail){
	s := prepareRequest(headers)
	errMsg := errors.LelylanHttpFail{} 
	res, err := s.Get(url, nil, nil, &errMsg)
	if err != nil {
		panic(err)
	}
	return res, errMsg
}

func Post(url string, headers map[string]string, device devices.Device) (*napping.Response, errors.LelylanHttpFail){
	s := prepareRequest(headers)
	errMsg := errors.LelylanHttpFail{} 
	res, err := s.Post(url, &device, nil, &errMsg)
	if err != nil {
		panic(err)
	}
	return res, errMsg 	
}

func Put(url string, headers map[string]string, device devices.Device) (*napping.Response, errors.LelylanHttpFail){
	s := prepareRequest(headers)
	errMsg := errors.LelylanHttpFail{} 
	res, err := s.Put(url, &device, nil, &errMsg)
	if err != nil {
		panic(err)
	}
	return res, errMsg	
}

func Delete(url string, headers map[string]string) (*napping.Response, errors.LelylanHttpFail){
	s := prepareRequest(headers)
	errMsg := errors.LelylanHttpFail{} 
	res, err := s.Delete(url, nil, &errMsg)
	if err != nil {
		panic(err)
	}
	return res, errMsg
}
