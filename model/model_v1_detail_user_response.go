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
import (
	"time"
)

type V1DetailUserResponse struct {
	Alias string `json:"alias,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Disabled bool `json:"disabled"`
	Email string `json:"email"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	Name string `json:"name"`
	Projects []V1UserProjectBase `json:"projects"`
	Roles []V1NameAlias `json:"roles"`
}