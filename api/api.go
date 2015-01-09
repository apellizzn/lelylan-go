package api

import "net/http"
import "github.com/jmcvetta/napping"
import "github.com/apellizzn/lelylan-go/devices"

func prepareRequest(headers map[string]string) napping.Session{
	s := napping.Session{ Log: true }
  h := &http.Header{}
  for header := range headers {
    h.Set(header,headers[header])
	}
	s.Client = &http.Client{}
  s.Header = h
  return s
}

func Get(url string, headers map[string]string) *napping.Response{
	s := prepareRequest(headers)
  res, _ := s.Get(url, nil, nil, nil)
  return res
}

func Post(url string, headers map[string]string, device devices.Device) *napping.Response{
	s := prepareRequest(headers)
	res, _ := s.Post(url, &device, nil, nil)
  return res	
}

func Put(url string, headers map[string]string, device devices.Device) *napping.Response{
	s := prepareRequest(headers)
	res, _ := s.Put(url, &device, nil, nil)
  return res	
}

func Delete(url string, headers map[string]string) *napping.Response{
	s := prepareRequest(headers)
	res, _ := s.Delete(url, nil, nil)
	return res
}
