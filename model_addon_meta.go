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

type AddonMeta struct {
	Annotations   map[string]string        `json:"annotations,omitempty"`
	Dependencies  []AddonDependency        `json:"dependencies,omitempty"`
	DeployTo      *AddonDeployTo           `json:"deployTo,omitempty"`
	Description   string                   `json:"description"`
	Icon          string                   `json:"icon"`
	Invisible     bool                     `json:"invisible"`
	Name          string                   `json:"name"`
	NeedNamespace []string                 `json:"needNamespace,omitempty"`
	System        *AddonSystemRequirements `json:"system,omitempty"`
	Tags          []string                 `json:"tags,omitempty"`
	Url           string                   `json:"url,omitempty"`
	UxPlugins     map[string]string        `json:"uxPlugins,omitempty"`
	Version       string                   `json:"version"`
}
