package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eloo/github-release-tool/src/log"
	"github.com/eloo/github-release-tool/src/models"
	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
	"os"

	"path/filepath"
	"robpike.io/filter"
	"strings"
)

const githubReleaseURLTemplate = "https://api.github.com/repos/%s/releases"
const githubLatestReleaseURLTemplate = "https://api.github.com/repos/%s/releases/latest"

var (
	// Download command used from the cli
	Download = cli.Command{
		Name: "download",
		Description: "Download a release file for the passed repository. Per default the latest release will be " +
			"downloaded to the current folder",
		ShortName: "d",
		Usage:     "Download a release file",
		ArgsUsage: "<:owner/:repo>",
		Flags: []cli.Flag{
			searchFlag,
			outputFlag,
		},
		Action: func(c *cli.Context) error {
			downloadRelease(c.Args().First(), c.String("search"), c.String("output"))
			return nil
		},
	}
	outputFlag = cli.StringFlag{
		Value: ".",
		Name:  "output, o",
		Usage: "sets an output directory for the directory",
	}
	searchFlag = cli.StringFlag{
		Value: "",
		Name:  "search, s",
		Usage: "search string for filename matching",
	}
)

func download(candidate models.DownloadCandidate, outputDirectory string) {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	destinationPath := filepath.Join(outputDirectory, candidate.File.Name)
	response, err := resty.R().
		SetOutput(destinationPath).
		SetAuthToken(os.Getenv("GITHUB_TOKEN")).
		Get(candidate.File.DownloadURL)
	fmt.Println(response.RawResponse)
	if err != nil {
		fmt.Errorf("download failed: %s", err)
	}
}

func downloadRelease(repository string, pattern string, outputDirectory string) {
	latestRelease := getLatestRelease(repository)

	var assets []models.Asset
	if len(pattern) > 0 {
		log.Info("Pattern found: %s", pattern)
		for _, asset := range latestRelease.Assets {
			if strings.Contains(asset.Name, pattern) {
				assets = append(assets, asset)
			}
		}
	}

	binaryAssets := getBinaries(assets)
	checksumAssets := getChecksums(assets)
	candidates := createDownloadCandidates(binaryAssets, checksumAssets)

	switch len(candidates) {
	case 0:
		log.Error("No possible download candidate found.")
	case 1:
		download(candidates[0], outputDirectory)
	default:
		log.Error("Found %d possible download candidates. Please use an more accurate search string", len(candidates))

	}

}
func createDownloadCandidates(binaries []models.Asset, checksums []models.Asset) []models.DownloadCandidate {
	var downloadCandidates []models.DownloadCandidate
	for i := 0; i < len(binaries); i++ {
		downloadCandidates = append(downloadCandidates, models.DownloadCandidate{File: binaries[i]})
	}
	return downloadCandidates
}

func getBinaries(assets []models.Asset) []models.Asset {
	files := filter.Choose(assets, func(v models.Asset) bool {
		return !strings.HasSuffix(v.Name, ".sha256")
	})
	binaries, ok := files.([]models.Asset)
	if !ok {
		return make([]models.Asset, 0)
	}
	return binaries
}

func getChecksums(assets []models.Asset) []models.Asset {
	files := filter.Choose(assets, func(v models.Asset) bool {
		return strings.HasSuffix(v.Name, ".sha256")
	})
	checksums, ok := files.([]models.Asset)
	if !ok {
		return make([]models.Asset, 0)
	}
	return checksums
}

func getLatestRelease(repository string) models.Release {
	url := fmt.Sprintf(githubLatestReleaseURLTemplate, repository)
	resp, err := resty.R().SetAuthToken(os.Getenv("GITHUB_TOKEN")).Get(url)
	if err != nil {
		log.Fatal("Error", err)
	}
	var release models.Release
	err = json.Unmarshal(resp.Body(), &release)
	if err != nil {
		log.Fatal("Error", err)
	}
	log.Debug("%s: %s", release.TagName, release.PublishedAt.Local().String())
	return release
}

func getReleases(repository string) []models.Release {
	url := fmt.Sprintf(githubReleaseURLTemplate, repository)
	resp, err := resty.R().SetAuthToken(os.Getenv("GITHUB_TOKEN")).Get(url)
	if err != nil {
		log.Fatal("Error", err)
	}
	var releases []models.Release
	err = json.Unmarshal(resp.Body(), &releases)
	if err != nil {
		log.Fatal("Error", err)
	}
	for _, release := range releases {
		fmt.Println(release.TagName + " " + release.PublishedAt.Local().String())
	}
	return releases
}
