
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| cluster_name | The name that will be assigned to the cluster | string | `crud-cluster` | no |
| json_cred_file | The location of the credentials file for GCP | string | - | yes |
| master_node_password | A secure password for the cluster master | string | - | yes |
| master_node_username | A username for the cluster master | string | - | yes |
| node_pool_name | A name to set to the node pool | string | - | yes |
| project_id | The project id in GCP to use | string | - | yes |
| zone | The zone where the cluster will be deployed | string | `us-central1-a` | no |

