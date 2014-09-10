/******************************************************
Authors: Kevin Lee(kevin@yunify.com)
License: Apache License, Version 2.0
Copyright:
******************************************************/
package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	//"net/url"
)

func Qingcloud_sig(sigtype string, sigkey string, sigstr string) string {
	var string_siged string = ""
	sig_key := []byte(sigkey)
	sig_str := []byte(sigstr)
	switch sigtype {
	case "sha256":
		hmac_sig := hmac.New(sha256.New, sig_key)
		hmac_sig.Write(sig_str)
		hmac_dig := []byte(hmac_sig.Sum(nil))
		string_siged = base64.StdEncoding.EncodeToString(hmac_dig)
		fmt.Println(hmac_sig)
	case "sha128":
		fmt.Println("Ongoing")
	}
	return string_siged
}

func encrypt_string() {

}
