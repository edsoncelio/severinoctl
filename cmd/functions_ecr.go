package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type Image struct {
	Images   []ImageOpts   `json:"images,omitempty"`
	Failures []FailureOpts `json:"failures,omitempty"`
}

type ImageOpts struct {
	RegistryId             string      `json:"registryId"`
	RepositoryName         string      `json:"repositoryName"`
	ImageId                ImageIDOpts `json:"imageId"`
	ImageManifest          string      `json:"imageManifest"`
	ImageManifestMediaType string      `json:"imageManifestMediaType"`
}

type FailureOpts struct {
	ImageId       ImageIDOpts `json:"imageId"`
	FailureCode   string      `json:"failureCode"`
	FailureReason string      `json:"failureReason"`
}

type ImageIDOpts struct {
	ImageDigest string `json:"imageDigest,omitempty"`
	ImageTag    string `json:"imageTag"`
}

func checkImage(repository string, tag string) {
	var resultJson Image

	_, lookErr := exec.LookPath("aws")
	if lookErr != nil {
		panic(lookErr)
	}

	out, err := exec.Command("aws", "ecr", "batch-get-image", "--repository-name", repository, "--image-ids", "imageTag="+tag).Output()

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(out), &resultJson)

	if len(resultJson.Images) == 1 {
		fmt.Printf("✅ Tag '%s' found! digest: %s and repository: %s\n", resultJson.Images[0].ImageId.ImageTag, resultJson.Images[0].ImageId.ImageDigest, resultJson.Images[0].RepositoryName)
	} else if len(resultJson.Failures) == 1 {
		fmt.Printf("❌ Tag '%s' not found, reason: '%s' - try again!\n", resultJson.Failures[0].ImageId.ImageTag, resultJson.Failures[0].FailureReason)
	}

}
