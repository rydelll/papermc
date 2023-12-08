package papermc

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type projectDownloadService service

func (srv *projectDownloadService) Download(project Project, info BuildInfo) error {
	out, err := os.Create(fmt.Sprintf("./%s.jar", project))
	if err != nil {
		return fmt.Errorf("could not create %s.jar file: %w", project, err)
	}
	defer out.Close()

	url := fmt.Sprintf(srv.client.baseURL+string(srv.endpoint), project, info.version, info.build, info.jar)
	resp, err := srv.client.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("could not reach %s: %w", url, err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save request body to %s.jar file: %w", project, err)
	}

	err = validateChecksum(project, info.checksum)
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
