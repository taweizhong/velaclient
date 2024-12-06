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

type V1AddonRegistry struct {
	Git    *AddonGitAddonSource    `json:"git,omitempty"`
	Gitee  *AddonGiteeAddonSource  `json:"gitee,omitempty"`
	Gitlab *AddonGitlabAddonSource `json:"gitlab,omitempty"`
	Helm   *AddonHelmSource        `json:"helm,omitempty"`
	Name   string                  `json:"name"`
	Oss    *AddonOssAddonSource    `json:"oss,omitempty"`
}
