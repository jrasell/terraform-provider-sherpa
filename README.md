# Sherpa Terraform Provider

The Terraform Sherpa provider is used to interact with Sherpa scaling policy documents.

## Download & Install

* The provider binary can be downloaded from the [GitHub releases page](https://github.com/jrasell/terraform-provider-sherpa/releases) using `curl -L https://github.com/jrasell/terraform-provider-sherpa/releases/download/v0.0.1/terraform-provider-sherpa_0.0.1_linux_amd64 -o terraform-provider-sherpa`
* Follow the HashiCorp instructions on [installing plugins](https://www.terraform.io/docs/plugins/basics.html#installing-plugins) 

## Provider Usage
```hcl
provider "sherpa" {
  addr = "http://sherpa.mycompany.io:8000"
}
```

#### Provider Argument Reference

* `addr` (Optional) The HTTP(S) address of the sherpa server. Defaults to `http://localhost:8000`.
* `ca_file` (Optional) The PEM encoded CA cert file to use to verify the Sherpa server SSL certificate.
* `cert_file` (Optional) The PEM encoded client certificate for TLS authentication to the Sherpa server.
* `key_file` (Optional) The unencrypted PEM encoded private key matching the client certificate.

## Policy Resource Usage
```hcl
resource "sherpa_policy" "job_example" {
  job_id          = "example"
  policy_document = <<EOF
{
  "cache": {
    "Enabled": true,
    "MinCount": 2,
    "MaxCount": 10,
    "ScaleOutCount": 1,
    "ScaleInCount": 1,
    "ScaleOutCPUPercentageThreshold": 75,
    "ScaleOutMemoryPercentageThreshold": 75,
    "ScaleInCPUPercentageThreshold": 30,
    "ScaleInMemoryPercentageThreshold": 30
  }
}
EOF 
}
```

#### Policy Resource Argument Reference

* `job_id` (Required) The Nomad job which the policy will be used for.
* `group_name` (Optional) The group name within the job which the scaling policy is for.
* `policy_document` (Required) The scaling policy to be written to Sherpa.
