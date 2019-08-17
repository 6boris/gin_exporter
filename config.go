package gin_exporter

type Config struct {
	Group     string `json:"group"`
	Type      string `json:"type"`
	System    string `json:"system"`
	Instance  string `json:"instance"`
	Version   string `json:"version"`
	CommitId  string `json:"commit_id"`
	StartedAt string `json:"started_at"`
	Platform  string `json:"platform"`
}
