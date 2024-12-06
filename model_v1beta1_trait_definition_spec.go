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

type V1beta1TraitDefinitionSpec struct {
	AppliesToWorkloads []string `json:"appliesToWorkloads,omitempty"`
	ConflictsWith []string `json:"conflictsWith,omitempty"`
	ControlPlaneOnly bool `json:"controlPlaneOnly,omitempty"`
	DefinitionRef *CommonDefinitionReference `json:"definitionRef,omitempty"`
	Extension string `json:"extension,omitempty"`
	ManageWorkload bool `json:"manageWorkload,omitempty"`
	PodDisruptive bool `json:"podDisruptive,omitempty"`
	RevisionEnabled bool `json:"revisionEnabled,omitempty"`
	Schematic *CommonSchematic `json:"schematic,omitempty"`
	Stage string `json:"stage,omitempty"`
	Status *CommonStatus `json:"status,omitempty"`
	WorkloadRefPath string `json:"workloadRefPath,omitempty"`
}