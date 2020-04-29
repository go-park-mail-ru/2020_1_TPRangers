package csrf


import(
	"main/internal/tools/errors"
	"testing"
)


func TestHashToken_Check(t *testing.T) {


	hTKN , _ := NewHMACHashToken("secret")

	cookie := "123"
	str , _ := hTKN.Create(cookie, int64(2000000000000))

	if flag , _ := hTKN.Check(cookie,str); !flag {
		t.Error("error check")
	}

	if _ , err := hTKN.Check(cookie, "1:1"); err == nil {
		t.Error("error check")
	}

	if _ , err := hTKN.Check(cookie, "1:c1"); err == nil {
		t.Error("error check")
	}


	if _ , err := hTKN.Check(cookie, "1:-2"); err != errors.CookieExpired {
		t.Error("error check")
	}

	if _ , err := hTKN.Check(cookie, "asdpoi d98n12:220000000000"); err == nil {
		t.Error("error check")
	}
}