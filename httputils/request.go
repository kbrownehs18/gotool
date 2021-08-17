package httputils

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Method http request method
type Method int8

const (
	// POST post http request
	POST Method = iota
	// GET http request get method
	GET
)

func (m Method) String() string {
	var name string
	switch m {
	case POST:
		name = "POST"
	case GET:
		name = "GET"
	default:
		name = "UNKNOWN"
	}

	return name
}

// HTTPRequest request
// url request url
// method request method post or get
// args[0] type is map[string]string or string, request paramaters, \x00@ if upload file
// args[1] type is map[string]string, request headers
// args[2] type is bool, whether to return the result
// args[3] type is *http.Client, custom client
func HTTPRequest(url string, method Method, args ...interface{}) (string, error) {
	paramsMap := make(map[string]string) // request parameters
	var paramsStr string
	var paramsIsStr bool
	headers := make(map[string]string) // request headers
	rtn := true
	var client *http.Client
	var ok bool
	argsLen := len(args)
	if argsLen > 0 {
		paramsMap, ok = args[0].(map[string]string)
		if !ok {
			paramsStr, ok = args[0].(string)
			if !ok {
				return "", errors.New("Params error")
			}
			paramsIsStr = true
		}
	}
	if argsLen > 1 {
		headers, ok = args[1].(map[string]string)
		if !ok {
			return "", errors.New("Headers error")
		}
	}
	if argsLen > 2 {
		rtn, ok = args[2].(bool)
		if !ok {
			return "", errors.New("Return bool error")
		}
	}
	if argsLen > 3 {
		client, ok = args[3].(*http.Client)
		if !ok {
			return "", errors.New("Http client error")
		}
	} else {
		client = http.DefaultClient
	}

	var req *http.Request
	var err error
	contentType := "" // default content-type
	var queryString string
	if paramsIsStr {
		queryString = URLEncode(paramsStr)
	} else {
		queryString = URLEncode(paramsMap)
	}
	if method == GET {
		// GET
		if queryString != "" {
			if strings.Index(url, "?") != -1 {
				// has params
				url += "&" + queryString
			} else {
				// no params
				url += "?" + queryString
			}
		}

		req, err = http.NewRequest("GET", url, nil)
	} else {
		// POST
		// whether there is upload file
		var isFile bool
		if !paramsIsStr {
			for _, v := range paramsMap {
				if strings.Index(v, "\x00@") == 0 {
					// there is upload file
					isFile = true
					break
				}
			}
		}
		if isFile {
			bodyBuf := new(bytes.Buffer)
			bodyWriter := multipart.NewWriter(bodyBuf)

			for key, value := range paramsMap {
				if strings.Index(value, "\x00@") == 0 {
					value = strings.Replace(value, "\x00@", "", -1)
					fileWriter, err := bodyWriter.CreateFormFile(key, filepath.Base(value))
					if err != nil {
						return "", err
					}
					fh, err := os.Open(value)
					if err != nil {
						return "", err
					}
					defer fh.Close()

					// iocopy
					_, err = io.Copy(fileWriter, fh)
					if err != nil {
						return "", err
					}
				} else {
					bodyWriter.WriteField(key, value)
				}
			}

			// Important if you do not close the multipart writer you will not have a terminating boundry
			bodyWriter.Close()
			contentType = bodyWriter.FormDataContentType()
			req, err = http.NewRequest("POST", url, bodyBuf)
		} else {
			if paramsIsStr {
				contentType = "application/json; charset-utf-8"
				req, err = http.NewRequest("POST", url, strings.NewReader(paramsStr))
			} else {
				contentType = "application/x-www-form-urlencoded; charset=utf-8"
				req, err = http.NewRequest("POST", url, strings.NewReader(queryString))
			}

		}
	}

	if err != nil {
		return "", err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// add headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if rtn {
		// need return
		bData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}

		return string(bData), nil
	}

	return "", nil
}
