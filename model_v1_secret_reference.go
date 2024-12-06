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

// SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
type V1SecretReference struct {
	// name is unique within a namespace to reference a secret resource.
	Name string `json:"name,omitempty"`
	// namespace defines the space within which the secret name must be unique.
	Namespace string `json:"namespace,omitempty"`
}
