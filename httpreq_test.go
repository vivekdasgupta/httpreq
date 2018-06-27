package httpreq

import "testing"

func TestMakeUrl(t *testing.T) {
	finalUrl := MakeUrl("advertiser", "country=australia")
	if finalUrl != "https://api.appnexus.com/advertiser?country=australia" {
		t.Errorf("FinalURL was incorrect, got: %s, want: %s.", finalUrl, "https://api.appnexus.com/advertiser?country=australia")
	}
}
