package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestKubernetes(t *testing.T) {
	t.Parallel()
	terraformOptions := &terraform.Options{
		TerraformDir: "../gcp",
		Vars: map[string]interface{}{
			"cluster_name":         fmt.Sprintf("default-%s", random.UniqueId()),
			"json_cred_file":       os.Getenv("CRED_FILE"),
			"master_node_password": os.Getenv("MASTER_PASS"),
			"master_node_username": "admin",
			"project_id":           os.Getenv("PROJ_ID"),
			"node_pool_name":       "default",
			"zone":                 "us-central1-a",
		},
	}
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)
}
