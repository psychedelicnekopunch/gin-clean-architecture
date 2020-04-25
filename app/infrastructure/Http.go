
package infrastructure


import (
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)


type Http struct {}


func NewHttp() *Http {
	return new(Http)
}


func (h *Http) GetResponse(request *http.Request) (response *http.Response, err error) {
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}


func (h *Http) GetBody(response *http.Response) (body string, err error) {
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}


func (h *Http) setParams(values url.Values, params map[string]string) url.Values {
	for key, param := range params {
		values.Add(key, param)
	}
	return values
}


func (h *Http) Get(URL string, params map[string]string) (request *http.Request, err error) {

	request, err = http.NewRequest("GET", URL, nil)

	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = h.setParams(request.URL.Query(), params).Encode()
	return request, nil
}


func (h *Http) Post(URL string, params map[string]string) (request *http.Request, err error) {

	request, err = http.NewRequest("POST", URL, strings.NewReader(h.setParams(url.Values{}, params).Encode()))

	if err != nil {
		return nil, err
	}

	return request, nil
}


func (h *Http) Delete(URL string, params map[string]string) (request *http.Request, err error) {

	request, err = http.NewRequest("DELETE", URL, strings.NewReader(h.setParams(url.Values{}, params).Encode()))

	if err != nil {
		return nil, err
	}

	return request, nil
}
