package comunication

import "net/http"
import "io/ioutil"
import "bytes"

func Perform(method string, url string, headers map[string]string, body []byte) []byte{
	
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	for header := range headers {
    req.Header.Set(header,headers[header])
	}
	res, err := client.Do(req)
	if err == nil{
		defer res.Body.Close()
    data, _ := ioutil.ReadAll(res.Body)
  	return []byte(data)	
  } else {
  	panic(err)
  }
}
