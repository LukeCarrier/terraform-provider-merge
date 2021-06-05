package merge

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/LukeCarrier/terraform-provider-merge/merge/marshal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/imdario/mergo"
	"log"
)

func dataSourceMerge() *schema.Resource {
	return &schema.Resource{
		Description: "Deep merges a list of encoded strings into a single string.",
		ReadContext: dataSourceMergeRead,
		Schema: map[string]*schema.Schema{
			"input": {
				Description: "List of input layers to merge.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Description:      "Data format (json, yaml).",
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: dataSourceMergeFormatValidateDiagFunc(),
						},
						"data": {
							Description: "Encoded data.",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
				MinItems: 1,
				Required: true,
			},

			"merge_on_overwrite": {
				Description: "Merge values instead of overwriting?",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"overwrite_with_empty_value": {
				Description: "Overwrite existing values with empty values in later layers?",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"type_check": {
				Description: "Type check values?",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"output_format": {
				Description:      "Output format (json, yaml).",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: dataSourceMergeFormatValidateDiagFunc(),
			},

			"output": {
				Description: "Deep-merged output.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceMergeFormatValidateDiagFunc() schema.SchemaValidateDiagFunc {
	return validation.ToDiagFunc(validation.StringInSlice([]string{"json", "yaml"}, false))
}

func dataSourceMergeMakeId(output string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(output)))
}

func dataSourceMergeMakeMergoOpts(mergeOnOverwrite bool, overwriteWithEmptyValue bool, typeCheck bool) []func(*mergo.Config) {
	opts := make([]func(*mergo.Config), 0)
	if mergeOnOverwrite {
		opts = append(opts, mergo.WithOverride)
	}
	if overwriteWithEmptyValue {
		opts = append(opts, mergo.WithOverwriteWithEmptyValue)
	}
	if typeCheck {
		opts = append(opts, mergo.WithTypeCheck)
	}
	return opts
}

func dataSourceMergeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	input := d.Get("input").([]interface{})
	mergeOnOverwrite := d.Get("merge_on_overwrite").(bool)
	overwriteWithEmptyValue := d.Get("overwrite_with_empty_value").(bool)
	typeCheck := d.Get("type_check").(bool)
	outputFormat := d.Get("output_format").(string)

	var output string

	log.Println(
		"[DEBUG] received", len(input), "input layer(s) and Mergo options mergeOnOverwrite =", mergeOnOverwrite,
		"overwriteWithEmptyValue =", overwriteWithEmptyValue, "typeCheck =", typeCheck, "and output format =", outputFormat)
	log.Println("[TRACE] input layer values", input)

	var mergedLayers marshal.UnmarshalledData
	mergoOpts := dataSourceMergeMakeMergoOpts(mergeOnOverwrite, overwriteWithEmptyValue, typeCheck)
	for _, layer := range input {
		layerMap := layer.(map[string]interface{})
		data := layerMap["data"].(string)
		format := layerMap["format"].(string)

		marshaller, err := marshal.NewMarshaller(format)
		if err != nil {
			return diag.FromErr(err)
		}
		layerData := marshal.NewUnmarshalResult()
		if err := marshaller.Unmarshal(data, layerData); err != nil {
			return diag.FromErr(err)
		}
		log.Println("[TRACE] Unmarshalled layer data", layerData)

		if mergedLayers == nil {
			log.Println("[DEBUG] Assigning first layer")
			mergedLayers = layerData
		} else {
			log.Println("[DEBUG] Merging subsequent layer")
			if err = mergo.Merge(&mergedLayers, layerData, mergoOpts...); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	marshaller, err := marshal.NewMarshaller(outputFormat)
	if err != nil {
		return diag.FromErr(err)
	}
	result, err := marshaller.Marshal(mergedLayers)
	if err != nil {
		return diag.FromErr(err)
	}
	output = *result

	id := dataSourceMergeMakeId(output)
	log.Println("[DEBUG] generated resource ID", id)
	d.SetId(id)
	d.Set("output", output)

	return nil
}
