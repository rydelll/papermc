package papermc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
)

const downloadEndpoint = "/projects/%s/versions/%s/builds/%d/downloads/%s"

// A DownloadService is used to download a JAR for a PaperMC project.
type DownloadService struct {
	project Project
	baseURL string
	client  *http.Client
}

// NewDownloadService for a PaperMC project.
func NewDownloadService(project Project, baseURL string, client *http.Client) *DownloadService {
	return &DownloadService{
		project: project,
		baseURL: baseURL,
		client:  client,
	}
}

// Download a PaperMC executable JAR for a given version, build, and JAR name.
func (s *DownloadService) Download(ver string, build int, JAR string) error {
	out, err := os.Create(fmt.Sprintf("%s.jar", s.project))
	if err != nil {
		return fmt.Errorf("create JAR: %w", err)
	}
	defer out.Close()

	url := fmt.Sprintf(s.baseURL+downloadEndpoint, s.project, ver, build, JAR)
	resp, err := s.client.Get(url)
	if err != nil {
		return fmt.Errorf("fetch data: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("save JAR: %w", err)
	}

	return nil
}

// ValidateChecksum ensures the JAR has not been modified or corrupted.
//
// Compares the SHA256 hash of a projects JAR with the provided SHA256 hash.
func (s *DownloadService) ValidateChecksum(checksum string) error {
	file, err := os.Open(fmt.Sprintf("%s.jar", s.project))
	if err != nil {
		return fmt.Errorf("read JAR: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return fmt.Errorf("SHA256 hash: %w", err)
	}

	if hex.EncodeToString(hasher.Sum(nil)) != checksum {
		return fmt.Errorf("SHA256 hash: expected %s got %x", checksum, hasher.Sum(nil))
	}

	return nil
}
