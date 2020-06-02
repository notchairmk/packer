// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package ncloud

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName                   *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType                 *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug                       *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                       *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError                     *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars                    map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars               []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AccessKey                         *string           `mapstructure:"access_key" cty:"access_key" hcl:"access_key"`
	SecretKey                         *string           `mapstructure:"secret_key" cty:"secret_key" hcl:"secret_key"`
	ServerImageProductCode            *string           `mapstructure:"server_image_product_code" required:"true" cty:"server_image_product_code" hcl:"server_image_product_code"`
	ServerProductCode                 *string           `mapstructure:"server_product_code" required:"true" cty:"server_product_code" hcl:"server_product_code"`
	MemberServerImageNo               *string           `mapstructure:"member_server_image_no" required:"false" cty:"member_server_image_no" hcl:"member_server_image_no"`
	ServerImageName                   *string           `mapstructure:"server_image_name" required:"false" cty:"server_image_name" hcl:"server_image_name"`
	ServerImageDescription            *string           `mapstructure:"server_image_description" required:"false" cty:"server_image_description" hcl:"server_image_description"`
	UserData                          *string           `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
	UserDataFile                      *string           `mapstructure:"user_data_file" required:"false" cty:"user_data_file" hcl:"user_data_file"`
	BlockStorageSize                  *int              `mapstructure:"block_storage_size" required:"false" cty:"block_storage_size" hcl:"block_storage_size"`
	Region                            *string           `mapstructure:"region" required:"false" cty:"region" hcl:"region"`
	AccessControlGroupConfigurationNo *string           `mapstructure:"access_control_group_configuration_no" required:"false" cty:"access_control_group_configuration_no" hcl:"access_control_group_configuration_no"`
	Type                              *string           `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect                *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                           *string           `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                           *int              `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername                       *string           `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword                       *string           `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName                    *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName           *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys            *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile                 *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHPty                            *bool             `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                        *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout                    *string           `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth                      *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding         *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts              *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost                    *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort                    *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth               *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername                *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword                *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive             *bool             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile          *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod             *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost                      *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort                      *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername                  *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword                  *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval              *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout               *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels                  []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels                   []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey                      []byte            `mapstructure:"ssh_public_key" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey                     []byte            `mapstructure:"ssh_private_key" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                         *string           `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword                     *string           `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                         *string           `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMPort                         *int              `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout                      *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL                       *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure                     *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM                      *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
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
		"packer_build_name":                     &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":                   &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                          &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                          &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                       &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":                 &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":            &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                            &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"secret_key":                            &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"server_image_product_code":             &hcldec.AttrSpec{Name: "server_image_product_code", Type: cty.String, Required: false},
		"server_product_code":                   &hcldec.AttrSpec{Name: "server_product_code", Type: cty.String, Required: false},
		"member_server_image_no":                &hcldec.AttrSpec{Name: "member_server_image_no", Type: cty.String, Required: false},
		"server_image_name":                     &hcldec.AttrSpec{Name: "server_image_name", Type: cty.String, Required: false},
		"server_image_description":              &hcldec.AttrSpec{Name: "server_image_description", Type: cty.String, Required: false},
		"user_data":                             &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":                        &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"block_storage_size":                    &hcldec.AttrSpec{Name: "block_storage_size", Type: cty.Number, Required: false},
		"region":                                &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"access_control_group_configuration_no": &hcldec.AttrSpec{Name: "access_control_group_configuration_no", Type: cty.String, Required: false},
		"communicator":                          &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":               &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                              &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                              &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                          &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                          &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":                      &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":               &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":             &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":                  &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                               &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                           &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":                      &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":                        &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding":          &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":                &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":                      &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":                      &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":                &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":                  &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":                  &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":               &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file":          &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":              &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":                        &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":                        &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":                    &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":                    &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":               &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":                &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":                    &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":                     &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":                        &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":                       &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":                        &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":                        &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                            &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                            &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                         &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                         &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":                        &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":                        &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
	}
	return s
}
