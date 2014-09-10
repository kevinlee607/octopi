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

type run_instance struct {
	inst_type, action, login_mod             string
	resource                                 qc_in_type
	inst_name, login_keypair, login_password string
	zone, userdata_path, userdata_value      string
	need_userdata                            int
	vxnet, image_id, security_group          string
}

type des_instance struct {
	image_id, instacne_id, status string
}

func Run_instance(qc_instance *run_instance, qc_signature string) string {

}

func (inst_id string) Des_instance() string {

}
