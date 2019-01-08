variable "cluster_name" {
  default = "crud-cluster"
  description = "The name that will be assigned to the cluster"
  type        = "string"
}

variable "json_cred_file" {
  description = "The location of the credentials file for GCP"
  type        = "string"
}

variable "master_node_password" {
  description = "A secure password for the cluster master"
  type        = "string"
}

variable "master_node_username" {
  description = "A username for the cluster master"
  type        = "string"
}

variable "project_id" {
  description = "The project id in GCP to use"
  type        = "string"
}

variable "node_pool_name" {
  description = "A name to set to the node pool"
  type        = "string"
}

variable "zone" {
  description = "The zone where the cluster will be deployed"
  default     = "us-central1-a"
  type        = "string"
}
