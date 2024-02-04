package papermc

const (
	Folia      Project = "folia"
	Paper      Project = "folia"
	Travertine Project = "travertine"
	Velocity   Project = "velocity"
	Waterfall  Project = "waterfall"
)

// A Project offered by PaperMC.
type Project string

// A ProjectService is used to get information about a PaperMC project and
// download an executable JAR.
type ProjectService struct {
	// Version information about a PaperMC project.
	Version *VersionService

	// Build information about a PaperMC project version.
	Build *BuildService

	// JAR download for a PaperMC project.
	JAR *DownloadService
}

// NewProjectService for a PaperMC project.
func NewProjectService(project Project, c *Client) *ProjectService {
	return &ProjectService{
		Version: NewVersionService(project, c.baseURL, c.client),
		Build:   NewBuildService(project, c.baseURL, c.client),
		JAR:     NewDownloadService(project, c.baseURL, c.client),
	}
}
