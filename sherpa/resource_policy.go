package sherpa

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/jrasell/sherpa/pkg/api"
)

func policyResource() *schema.Resource {
	return &schema.Resource{
		Create: policyResourceCreate,
		Read:   policyResourceRead,
		Update: policyResourceCreate,
		Delete: policyResourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"job_id": {
				Description: "The public key material.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"group_name": {
				Description: "The job group to interact with.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"policy_document": {
				Description: "JSON representation of the scaling policies. Use file() to specify a file as input.",
				Required:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func policyResourceCreate(d *schema.ResourceData, meta interface{}) error {
	providerConfig := meta.(ProviderConfig)
	client := providerConfig.client

	// Pull the job name, group and policy document from the parameters.
	job := d.Get("job_id").(string)
	group := d.Get("group_name").(string)
	policy := d.Get("policy_document").(string)

	var (
		err            error
		jobPolicy      map[string]*api.JobGroupPolicy
		jobGroupPolicy *api.JobGroupPolicy
	)

	if group == "" {
		if err := json.Unmarshal([]byte(policy), &jobPolicy); err != nil {
			return fmt.Errorf("error parsing job scaling policy: %s", err)
		}
		err = client.Policies().WriteJobPolicy(job, &jobPolicy)
	} else {
		if err := json.Unmarshal([]byte(policy), jobGroupPolicy); err != nil {
			return fmt.Errorf("error parsing job group scaling policy: %s", err)
		}
		err = client.Policies().WriteJobGroupPolicy(job, group, jobGroupPolicy)
	}

	if err != nil {
		return fmt.Errorf("error writing scaling document: %v", err)
	}
	d.SetId(job)

	return nil
}

func policyResourceRead(d *schema.ResourceData, meta interface{}) error {
	providerConfig := meta.(ProviderConfig)
	client := providerConfig.client

	job := d.Id()
	group := d.Get("group_name").(string)

	if group == "" {
		policies, err := client.Policies().ReadJobPolicy(job)
		if err != nil {
			return fmt.Errorf("error reading job scaling policy: %v", err)
		}
		_ = d.Set("policy_document", policies)
	} else {
		policy, err := client.Policies().ReadJobGroupPolicy(job, group)
		if err != nil {
			return fmt.Errorf("error deleting job group scaling policy: %v", err)
		}
		_ = d.Set("policy_document", policy)
	}
	return nil
}

func policyResourceDelete(d *schema.ResourceData, meta interface{}) error {
	providerConfig := meta.(ProviderConfig)
	client := providerConfig.client

	job := d.Id()
	group := d.Get("group_name").(string)

	if group == "" {
		if err := client.Policies().DeleteJobPolicy(job); err != nil {
			return fmt.Errorf("error deleting job scaling policy: %v", err)
		}
	} else {
		if err := client.Policies().DeleteJobGroupPolicy(job, group); err != nil {
			return fmt.Errorf("error deleting job group scaling policy: %v", err)
		}
	}
	return nil
}
