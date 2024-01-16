package client

// URL format: /projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
const downloadEndpoint string = "/projects/%s/versions/%s/builds/%d/downloads/%s"

func (s *ProjectService) Download(ver string, build int, jar string) error {
	return nil
}
