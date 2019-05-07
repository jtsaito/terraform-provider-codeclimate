package codeclimate

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRepositoryRead,

		Schema: map[string]*schema.Schema{
			"repository_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"test_reporter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRepositoryRead(d *schema.ResourceData, client interface{}) error {
	repositoryId := d.Get("repository_id").(string)
	repositoryData, err := getRepository(client.(Client), repositoryId)
	if err != nil {
		return err
	}

	d.SetId(repositoryData.(ReadRepositoryResponse).Data.ID)
	// TODO: Check that repositoryData.Data.Attributes.TestReporterID exists
	d.Set("test_reporter_id", repositoryData.(ReadRepositoryResponse).Data.Attributes.TestReporterID)

	return err
}
