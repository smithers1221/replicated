/*
 * Vendor API V1
 *
 * Apps documentation
 *
 * API version: 1.0.0
 * Contact: info@replicated.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type AppCollector struct {
	Config    string    `json:"Config,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	Name      string    `json:"Name,omitempty"`
	SpecID    string    `json:"ID,omitempty"`
}
