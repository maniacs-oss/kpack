package v1alpha1

// +k8s:openapi-gen=true
type NotaryConfig struct {
	V1 *NotaryV1Config `json:"v1,omitempty"`
}

// +k8s:openapi-gen=true
type NotaryV1Config struct {
	Host             string `json:"host"`
	RootSecretName   string `json:"rootSecretName"`
	TargetSecretName string `json:"targetSecretName"`
}
