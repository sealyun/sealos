package v1alpha

type ClusterResource struct {
	Node int64 `json:"node"`
	CPU  int64 `json:"cpu"`
	Mem  int64 `json:"mem"`
}
