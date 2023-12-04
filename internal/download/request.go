package download

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Project string

type buildInfo struct {
	version  string
	build    int
	jar      string
	checksum string
}

// URL format: https://api.papermc.io/v2/projects/{project}/versions/{version}/builds/{build}/downloads/{jar}
const (
	versionURL string = "https://api.papermc.io/v2/projects/%s"
	buildURL   string = "https://api.papermc.io/v2/projects/%s/versions/%s/builds"
	jarURL     string = "https://api.papermc.io/v2/projects/%s/versions/%s/builds/%d/downloads/%s"
)

const (
	Paper     Project = "paper"
	Velocity  Project = "velocity"
	Waterfall Project = "waterfall"
	Folia     Project = "folia"
)

func getLatestVersion(project Project) (string, error) {
	resp, err := http.Get(fmt.Sprintf(versionURL, project))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	vr := versionResponse{}
	err = json.Unmarshal(body, &vr)
	if err != nil {
		return "", err
	}

	return vr.Versions[len(vr.Versions)-1], nil
}

func getLatestBuild(project Project, version string) (buildInfo, error) {
	resp, err := http.Get(fmt.Sprintf(buildURL, project, version))
	if err != nil {
		return buildInfo{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return buildInfo{}, err
	}

	br := buildResponse{}
	err = json.Unmarshal(body, &br)
	if err != nil {
		return buildInfo{}, err
	}

	info := buildInfo{
		version:  br.Version,
		build:    br.Builds[len(br.Builds)-1].Build,
		jar:      br.Builds[len(br.Builds)-1].Download.Application.Name,
		checksum: br.Builds[len(br.Builds)-1].Download.Application.SHA256,
	}
	return info, nil
}

func downloadJar(project Project, info buildInfo) error {
	out, err := os.Create(fmt.Sprintf("./%s.jar", project))
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(fmt.Sprintf(jarURL, project, info.version, info.build, info.jar))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func validateChecksum(project Project, checksum string) error {
	file, err := os.Open(fmt.Sprintf("./%s.jar", project))
	if err != nil {
		return err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return err
	}

	if hex.EncodeToString(hasher.Sum(nil)) != checksum {
		return fmt.Errorf("invalid SHA256 hash, expected %s but got %x", checksum, hasher.Sum(nil))
	}

	return nil
}

func Download(project Project, version string) error {
	var err error
	if version == "latest" {
		version, err = getLatestVersion(project)
		if err != nil {
			return err
		}
	}

	info, err := getLatestBuild(project, version)
	if err != nil {
		return err
	}

	err = downloadJar(project, info)
	if err != nil {
		return err
	}

	err = validateChecksum(project, info.checksum)
	if err != nil {
		return err
	}

	return nil
}
