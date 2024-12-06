/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package velaclient

type V1DetailAddonResponse struct {
	Annotations       map[string]string        `json:"annotations,omitempty"`
	AvailableVersions []string                 `json:"availableVersions"`
	Definitions       []V1AddonDefinition      `json:"definitions"`
	Dependencies      []AddonDependency        `json:"dependencies,omitempty"`
	DeployTo          *AddonDeployTo           `json:"deployTo,omitempty"`
	Description       string                   `json:"description"`
	Detail            string                   `json:"detail,omitempty"`
	Icon              string                   `json:"icon"`
	Invisible         bool                     `json:"invisible"`
	Name              string                   `json:"name"`
	NeedNamespace     []string                 `json:"needNamespace,omitempty"`
	RegistryName      string                   `json:"registryName,omitempty"`
	Schema            string                   `json:"schema"`
	System            *AddonSystemRequirements `json:"system,omitempty"`
	Tags              []string                 `json:"tags,omitempty"`
	UiSchema          []SchemaUiParameter      `json:"uiSchema"`
	Url               string                   `json:"url,omitempty"`
	UxPlugins         map[string]string        `json:"uxPlugins,omitempty"`
	Version           string                   `json:"version"`
}
