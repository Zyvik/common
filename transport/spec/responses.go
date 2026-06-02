// Contains common responses
package spec

// ErrorResponse is a standard response struct for error responses.
type ErrorResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"error"`
}

type Pagination struct {
	PageNumber int         `json:"pageNumber" query:"pageNumber" validate:"min=0"` // indexing from 0
	PageSize   int         `json:"pageSize" query:"pageSize" validate:"min=1,max=50"`
	OrderBy    string      `json:"orderBy" query:"orderBy"`
	SortOder   string      `json:"sortOrder" query:"sortOrder" validate:"oneof=ASC asc DESC desc"`
	Filters    interface{} `json:"filters"` // (response only) used for conveying what filters were applied
}

func (p *Pagination) GetLimitAndOffset() (limit, offset int) {
	return p.PageSize, p.PageNumber * p.PageSize
}

// Healthcheck
const (
	StatusHealthy  = "HEALTHY"  // Everything works as expected.
	StatusSick     = "SICK"     // Able to connect, but there are some other issues.
	StatusDeceased = "DECEASED" // Unable to connect/ping. The service migt not be actually dead but it is for me.
)

// HealthCheckResp is the response struct for the healthcheck endpoint.
type HealthCheckResp struct {
	Status       string             `json:"status"`
	Dependencies []DependencyStatus `json:"dependencies"`
}

// SetOverallStatus sets the status to HEALTHY unless one or more dependencies are not healthy.
func (h *HealthCheckResp) SetOverallStatus() {
	h.Status = StatusHealthy
	for _, dependency := range h.Dependencies {
		if dependency.Status != StatusHealthy {
			h.Status = StatusSick
			return
		}
	}
}

// DependencyStatus holds information about services which the app uses (databases, etc.)
type DependencyStatus struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}
