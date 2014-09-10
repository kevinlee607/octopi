/******************************************************
Authors: Kevin Lee(kevin@yunify.com)
License: Apache License, Version 2.0
Copyright:
******************************************************/
package qingcloud

type qc_in_type struct {
	cpu    int
	memory int
}

type instance struct {
	inst_type, action, login_mod             string
	resource                                 qc_in_type
	inst_name, login_keypair, login_password string
	zone, userdata_path, userdata_value      string
	need_userdata                            int
	vxnet, image_id, security_group          string
}

func (qc_instance *instance, qc_signature string) Run_instance() {
	
}