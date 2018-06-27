package httpreq

import (
	"bytes"
	"fmt"
	"infolog"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"regexp"
)

const APN_URL = "https://api.xyz.com"

type HttpReq struct {
	Url     string
	Method  string
	Reqdata []byte
	AddAuth bool
	Token   string
	Success string
}

type HttpResp struct {
	Status        string
	Header        map[string][]string
	Hresp         *http.Response
	Body          []byte
	Contentlength int64
	Location      string
}

func MakeUrl(apiService string, apiQuery string) (finalUrl string) {
	finalUrl = fmt.Sprintf("%s/%s?%s", APN_URL, apiService, apiQuery)
	return finalUrl
}

func (reqstr *HttpReq) SendHttpRequest() (httpresPtr *HttpResp, err error) {
	httpres := new(HttpResp)
	req, err := http.NewRequest(reqstr.Method, reqstr.Url, bytes.NewBuffer(reqstr.Reqdata))
	if reqstr.AddAuth == true {
		infolog.Log(infolog.DEBUG, "Adding auth token :: Token=%s", reqstr.Token)
		req.Header.Add("Authorization", reqstr.Token)
	}

	infolog.Log(infolog.DEBUG, "Sending HTTP request :: Method = %s URL = %s", reqstr.Method, reqstr.Url)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		infolog.Log(infolog.ERROR, "Error in sending HTTP request :: Error=%s", err)
		return httpresPtr, err
	}
	defer resp.Body.Close()
	httpres.Status = resp.Status
	httpres.Contentlength = resp.ContentLength
	httpres.Header = resp.Header

	httpres.Body, _ = ioutil.ReadAll(resp.Body)
	httpres.Hresp = resp

	infolog.Log(infolog.DEBUG2, "HTTP  Response  :: Body=%s", httpres.Body)

	httpresPtr = httpres
	return httpresPtr, err
}

func (reqstr *HttpReq) SendHttpRoundtrip() (httpresPtr *HttpResp, err error) {
	var DefaultTransport http.RoundTripper = &http.Transport{}
	httpres := new(HttpResp)
	req, err := http.NewRequest("GET", reqstr.Url, bytes.NewBuffer(reqstr.Reqdata))
	if reqstr.AddAuth == true {
		infolog.Log(infolog.DEBUG, "Adding auth token :: Token=%s", reqstr.Token)
		req.Header.Add("Authorization", reqstr.Token)
	}
	dump, err := httputil.DumpRequestOut(req, true)
	infolog.Log(infolog.DEBUG, "SendHttpRoundtrip: DumpReq=%s", dump)

	infolog.Log(infolog.DEBUG, "Sending HTTP Roundtrip :: Method = %s URL = %s", reqstr.Method, reqstr.Url)

	resp, err := DefaultTransport.RoundTrip(req)
	if err != nil {
		infolog.Log(infolog.ERROR, "SendHttpRoundtrip: Error RoundTrip :: Error=%s", err)
		return httpresPtr, err
	}

	defer resp.Body.Close()

	loc, err := resp.Location()
	if err != nil {
		infolog.Log(infolog.ERROR, "SendHttpRoundtrip: HTTP  Roundtrip Error :: %s", err)
		return httpresPtr, err
	} else {
		infolog.Log(infolog.DEBUG, "SendHttpRoundtrip: HTTP  Roundtrip  :: Location=%s", httpres.Location)
	}
	re := regexp.MustCompile("http://")
	httpsloc := re.ReplaceAllString(loc.String(), "https://")

	httpres.Location = httpsloc
	httpres.Status = resp.Status

	httpresPtr = httpres
	return httpresPtr, err
}
