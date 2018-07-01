package httpreq

import (
	"httpreq"
	"testing"
)

const SIPHON_URL = "https://api.xyz.com/siphon"
const TEST_URL = "https://www.example.org"
const HTTP_OK = "200 OK"

func TestMakeUrl(t *testing.T) {
	finalUrl := MakeUrl("advertiser", "country=australia")
	if finalUrl != "https://api.xyz.com/advertiser?country=australia" {
		t.Errorf("FinalURL was incorrect, got: %s, want: %s.", finalUrl, "https://api.xyz.com/advertiser?country=australia")
	}
}

func TestSendHttpRequest(t *testing.T) {

	var req httpreq.HttpReq
	req.Url = TEST_URL
	req.Method = "GET"
	req.Reqdata = nil
	req.AddAuth = false
	req.Success = HTTP_OK

	apnResp, _ := req.SendHttpRequest()
	if apnResp.Status != HTTP_OK {
		t.Errorf("SendHttp response was incorrect, got: %s, want: %s.", apnResp.Status, HTTP_OK)
	}

}
