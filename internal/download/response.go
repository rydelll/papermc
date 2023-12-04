package download

type baseResponse struct {
	ProjectId   string `json:"project_id"`
	ProjectName string `json:"project_name"`
}

type versionResponse struct {
	baseResponse
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

type buildResponse struct {
	baseResponse
	Version string  `json:"version"`
	Builds  []build `json:"builds"`
}

type build struct {
	Build    int      `json:"build"`
	Time     string   `json:"time"`
	Channel  string   `json:"channel"`
	Promoted bool     `json:"promoted"`
	Changes  []change `json:"changes"`
	Download download `json:"downloads"`
}

type change struct {
	Commit  string `json:"commit"`
	Summary string `json:"summary"`
	Message string `json:"message"`
}

type download struct {
	Application application `json:"application"`
}

type application struct {
	Name   string `json:"name"`
	SHA256 string `json:"sha256"`
}
