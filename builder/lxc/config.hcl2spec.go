// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package lxc

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	ConfigFile          *string           `mapstructure:"config_file" required:"true" cty:"config_file" hcl:"config_file"`
	OutputDir           *string           `mapstructure:"output_directory" required:"false" cty:"output_directory" hcl:"output_directory"`
	ContainerName       *string           `mapstructure:"container_name" required:"false" cty:"container_name" hcl:"container_name"`
	CommandWrapper      *string           `mapstructure:"command_wrapper" required:"false" cty:"command_wrapper" hcl:"command_wrapper"`
	InitTimeout         *string           `mapstructure:"init_timeout" required:"false" cty:"init_timeout" hcl:"init_timeout"`
	CreateOptions       []string          `mapstructure:"create_options" required:"false" cty:"create_options" hcl:"create_options"`
	StartOptions        []string          `mapstructure:"start_options" required:"false" cty:"start_options" hcl:"start_options"`
	AttachOptions       []string          `mapstructure:"attach_options" required:"false" cty:"attach_options" hcl:"attach_options"`
	Name                *string           `mapstructure:"template_name" required:"true" cty:"template_name" hcl:"template_name"`
	Parameters          []string          `mapstructure:"template_parameters" required:"false" cty:"template_parameters" hcl:"template_parameters"`
	EnvVars             []string          `mapstructure:"template_environment_vars" required:"true" cty:"template_environment_vars" hcl:"template_environment_vars"`
	TargetRunlevel      *int              `mapstructure:"target_runlevel" required:"false" cty:"target_runlevel" hcl:"target_runlevel"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"config_file":                &hcldec.AttrSpec{Name: "config_file", Type: cty.String, Required: false},
		"output_directory":           &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"container_name":             &hcldec.AttrSpec{Name: "container_name", Type: cty.String, Required: false},
		"command_wrapper":            &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"init_timeout":               &hcldec.AttrSpec{Name: "init_timeout", Type: cty.String, Required: false},
		"create_options":             &hcldec.AttrSpec{Name: "create_options", Type: cty.List(cty.String), Required: false},
		"start_options":              &hcldec.AttrSpec{Name: "start_options", Type: cty.List(cty.String), Required: false},
		"attach_options":             &hcldec.AttrSpec{Name: "attach_options", Type: cty.List(cty.String), Required: false},
		"template_name":              &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"template_parameters":        &hcldec.AttrSpec{Name: "template_parameters", Type: cty.List(cty.String), Required: false},
		"template_environment_vars":  &hcldec.AttrSpec{Name: "template_environment_vars", Type: cty.List(cty.String), Required: false},
		"target_runlevel":            &hcldec.AttrSpec{Name: "target_runlevel", Type: cty.Number, Required: false},
	}
	return s
}
