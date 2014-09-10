/******************************************************
Authors: Kevin Lee(kevin@yunify.com)
License: Apache License, Version 2.0
Copyright:
******************************************************/
package octopi

import (
	"fmt"
	"octopi/signature"
)

func main() {
	var sigstr string = "GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1"
	var sigtype string = "sha256"
	var sigkey string = "aaaaaaa"

	//siged_str := signature.Encrypt_hash(sigtype, sigkey, sigstr)
	siged_str := signature.Qingcloud_sig(sigtype, sigkey, sigstr)
	fmt.Println(siged_str)
}
