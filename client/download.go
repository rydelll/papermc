package client

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// URL format: /projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
const downloadEndpoint string = "/projects/%s/versions/%s/builds/%d/downloads/%s"

func (s *ProjectService) Download(build Build) error {
	out, err := os.Create(fmt.Sprintf("./%s.jar", s.project))
	if err != nil {
		return fmt.Errorf("could not create %s.jar file: %w", s.project, err)
	}
	defer out.Close()

	url := fmt.Sprintf(s.baseURL+downloadEndpoint, s.project, build.Version, build.Build, build.Jar)
	resp, err := s.client.Get(url)
	if err != nil {
		return fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save request body to %s.jar file: %w", s.project, err)
	}

	err = validateChecksum(s.project, build.Checksum)
	if err != nil {
		return err
	}

	return nil
}

func validateChecksum(project Project, checksum string) error {
	file, err := os.Open(fmt.Sprintf("./%s.jar", project))
	if err != nil {
		return fmt.Errorf("could not read %s.jar file: %w", project, err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return fmt.Errorf("could not create SHA256 hash of %s.jar: %w", project, err)
	}

	if hex.EncodeToString(hasher.Sum(nil)) != checksum {
		return fmt.Errorf("invalid SHA256 hash, expected %s but got %x", checksum, hasher.Sum(nil))
	}

	return nil
}
