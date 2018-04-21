package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eloo/github-release-tool/log"
	"github.com/eloo/github-release-tool/models"
	"gopkg.in/resty.v0"
	"os"

	"github.com/spf13/cobra"
	"path/filepath"
	"robpike.io/filter"
	"strings"
)

const githubReleaseURLTemplate = "https://api.github.com/repos/%s/releases"
const githubLatestReleaseURLTemplate = "https://api.github.com/repos/%s/releases/latest"

var outputDirectory string
var searchPattern string

func init() {
	RootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringVarP(&outputDirectory, "output", "o", ".", "Output directory")
	downloadCmd.Flags().StringVarP(&searchPattern, "pattern", "p", "", "Pattern for download candidate name")
	downloadCmd.SetUsageTemplate("github-release-tool download <:owner/:repo> [flags]")
}

var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"d"},
	Short:   "Download a release file",
	Long: "Download a release file for the passed repository. Per default the latest release will be " +
		"downloaded to the current folder.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		downloadRelease(args[0], searchPattern, outputDirectory)
	},
}

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
	log.Debug("Download a release of repo %s", repository)
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
