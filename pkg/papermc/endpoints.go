package papermc

type endpoint string

const (
	// URL format: /projects/{project}
	endpointProjectVersion endpoint = "/projects/%s"
	// URL format: /projects/{project}/versions/{version}/builds
	endpointProjectBuild endpoint = "/projects/%s/versions/%s/builds"
	// URL format: /projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
	endpointProjectDownload endpoint = "/projects/%s/versions/%s/builds/%d/downloads/%s"
)
