/*
 * Kubevela api doc
 *
 * Kubevela api doc
 *
 * API version: v1beta1
 * Contact: feedback@mail.kubevela.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelProviderInfo struct {
	ClusterID string `json:"clusterID"`
	ClusterName string `json:"clusterName,omitempty"`
	Labels map[string]string `json:"labels"`
	Provider string `json:"provider"`
	RegionID string `json:"regionID,omitempty"`
	VpcID string `json:"vpcID,omitempty"`
	Zone string `json:"zone,omitempty"`
	ZoneID string `json:"zoneID,omitempty"`
}