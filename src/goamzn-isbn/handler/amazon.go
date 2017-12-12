package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"goamzn-isbn/model"

	"github.com/k0kubun/pp"
)

const (
	EC_SERVICE_ENDPOINT = "webservices.amazon.co.jp"
	EC_SERVICE_URI      = "/onca/xml"
)

func readConf() ([]byte, error) {
	f, err := os.Open("./conf/token_cred.json")
	if err != nil {
		fmt.Printf("token_cred.json open error: err; %v", err)
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("json file read error: err; %v", err)
		return nil, err
	}
	return b, nil
}

type AmazonRequest struct {
	Method string
	Url    string
	Params string
}

func SearchISBN(w http.ResponseWriter, r *http.Request) {
	b, err := readConf()
	if err != nil {
		fmt.Printf("readConf error. err: %v", err)
		return
	}

	var cred model.AmazonTokenCred
	if err := json.Unmarshal(b, &cred); err != nil {
		fmt.Printf("json unmarshal error. err: %v", err)
		return
	}

	params := url.Values{}
	params.Set("Service", "AWSECommerceService")
	params.Set("Operation", "ItemLookup")
	params.Set("ItemId", "9784774193328")
	params.Set("IdType", "ISBN")
	params.Set("SearchIndex", "Books")
	params.Set("Timestamp", time.Now().UTC().Format(time.RFC3339))
	params.Set("AWSAccessKeyId", cred.AccessKeyId)
	params.Set("AssociateTag", cred.AssociateTag)
	params.Set("ResponseGroup", "ItemAttributes, Images")

	// 署名
	canonical_params := params.Encode()
	strToSign := fmt.Sprintf("GET\n%v\n%v\n%v", EC_SERVICE_ENDPOINT, EC_SERVICE_URI, canonical_params)
	mac := hmac.New(sha256.New, []byte(cred.SecretKeyId))
	mac.Write([]byte(strToSign))
	signature := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	canonical_params = fmt.Sprintf("%v&Signature=%v", canonical_params, signature)

	// http request
	res, err := http.Get(fmt.Sprintf("http://%v%v?%v", EC_SERVICE_ENDPOINT, EC_SERVICE_URI, canonical_params))
	if err != nil {
		fmt.Printf("response error. err: %v", err)
		return
	}

	body := res.Body
	defer body.Close()
	br, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Printf("read error. err: %v", err)
		return
	}

	var result model.ItemLookupResponse
	if err := xml.Unmarshal(br, &result); err != nil {
		fmt.Printf("xml unmarshal error. err: %v", err)
	}

	fmt.Printf("%v", pp.Sprint(&result))

}
