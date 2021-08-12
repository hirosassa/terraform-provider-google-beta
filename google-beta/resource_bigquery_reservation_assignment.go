// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceBigqueryReservationReservationAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigqueryReservationReservationAssignmentCreate,
		Read:   resourceBigqueryReservationReservationAssignmentRead,
		Update: resourceBigqueryReservationReservationAssignmentUpdate,
		Delete: resourceBigqueryReservationReservationAssignmentDelete,

		Importer: &schema.ResourceImporter{
			State: resourceBigqueryReservationReservationAssignmentImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assignee": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The resource which will use the reservation.`,
			},
			"job_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PIPELINE", "QUERY", "ML_EXTERNAL"}, false),
				Description: `Which type of jobs will use the reservation. The following values are supported:
PIPELINE: Pipeline (load/export) jobs from the project will use the reservation.
QUERY: Query jobs from the project will use the reservation.
ML_EXTERNAL: BigQuery ML jobs that use services external to BigQuery for model training. These jobs will not utilize idle slots from other reservations. Possible values: ["PIPELINE", "QUERY", "ML_EXTERNAL"]`,
			},
			"reservation_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The parent resource name of the assignment.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceBigqueryReservationReservationAssignmentCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	assigneeProp, err := expandBigqueryReservationReservationAssignmentAssignee(d.Get("assignee"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("assignee"); !isEmptyValue(reflect.ValueOf(assigneeProp)) && (ok || !reflect.DeepEqual(v, assigneeProp)) {
		obj["assignee"] = assigneeProp
	}
	jobTypeProp, err := expandBigqueryReservationReservationAssignmentJobType(d.Get("job_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("job_type"); !isEmptyValue(reflect.ValueOf(jobTypeProp)) && (ok || !reflect.DeepEqual(v, jobTypeProp)) {
		obj["jobType"] = jobTypeProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments?assignmentId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ReservationAssignment: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ReservationAssignment: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ReservationAssignment %q: %#v", d.Id(), res)

	return resourceBigqueryReservationReservationAssignmentRead(d, meta)
}

func resourceBigqueryReservationReservationAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("BigqueryReservationReservationAssignment %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}

	if err := d.Set("assignee", flattenBigqueryReservationReservationAssignmentAssignee(res["assignee"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}
	if err := d.Set("job_type", flattenBigqueryReservationReservationAssignmentJobType(res["jobType"], d, config)); err != nil {
		return fmt.Errorf("Error reading ReservationAssignment: %s", err)
	}

	return nil
}

func resourceBigqueryReservationReservationAssignmentUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	assigneeProp, err := expandBigqueryReservationReservationAssignmentAssignee(d.Get("assignee"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("assignee"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, assigneeProp)) {
		obj["assignee"] = assigneeProp
	}
	jobTypeProp, err := expandBigqueryReservationReservationAssignmentJobType(d.Get("job_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("job_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, jobTypeProp)) {
		obj["jobType"] = jobTypeProp
	}

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ReservationAssignment %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ReservationAssignment %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ReservationAssignment %q: %#v", d.Id(), res)
	}

	return resourceBigqueryReservationReservationAssignmentRead(d, meta)
}

func resourceBigqueryReservationReservationAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ReservationAssignment: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{BigqueryReservationBasePath}}projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ReservationAssignment %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "ReservationAssignment")
	}

	log.Printf("[DEBUG] Finished deleting ReservationAssignment %q: %#v", d.Id(), res)
	return nil
}

func resourceBigqueryReservationReservationAssignmentImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/reservations/(?P<reservation_id>[^/]+)/assignments/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<reservation_id>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<reservation_id>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/reservations/{{reservationId}}/assignments/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenBigqueryReservationReservationAssignmentAssignee(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenBigqueryReservationReservationAssignmentJobType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandBigqueryReservationReservationAssignmentAssignee(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryReservationReservationAssignmentJobType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}