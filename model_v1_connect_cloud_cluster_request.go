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

type V1ConnectCloudClusterRequest struct {
	AccessKeyID     string            `json:"accessKeyID"`
	AccessKeySecret string            `json:"accessKeySecret"`
	Alias           string            `json:"alias,omitempty"`
	ClusterID       string            `json:"clusterID"`
	Description     string            `json:"description,omitempty"`
	Icon            string            `json:"icon"`
	Labels          map[string]string `json:"labels,omitempty"`
	Name            string            `json:"name"`
}
