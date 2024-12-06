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

import (
	"time"
)

type V1TerraformProvider struct {
	CreateTime time.Time `json:"createTime"`
	Name       string    `json:"name"`
	Provider   string    `json:"provider"`
	Region     string    `json:"region"`
}
