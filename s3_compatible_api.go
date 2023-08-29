package config

type S3CompatibleAPI struct {
	Enable      bool
	Address     string
	HTTPHeaders map[string][]string // Leave nil for default headers
}
