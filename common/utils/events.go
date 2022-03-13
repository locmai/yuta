package utils

type KubeopsActionData struct {
	Action    string  `json:"action"`
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	Value     float64 `json:"value"`
}
