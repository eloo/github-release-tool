package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/eloo/github-release-tool/models"
	"github.com/urfave/cli"
	"gopkg.in/resty.v0"
	"os"
	"sort"
	"strings"
	"github.com/Unknwon/log"
)

const GITHUB_RELEASE_URL_TEMPLATE = "https://api.github.com/repos/%s/releases"

var (
	Download = cli.Command{
		Name:        "download",
		Description: "Download a release file",
		ShortName:   "d",
		Usage:       "Download a release file",
		ArgsUsage:   "<github_repository>",
		Flags: []cli.Flag{
			cli.StringFlag{
				Value: "",
				Name:  "search, s",
				Usage: "search string for filename matching",
			}},
		Action: func(c *cli.Context) error {
			downloadRelease(c.Args().First(), c.String("search"))
			return nil
		},
	}
)

func downloadRelease(repository string, pattern string) {
	releases := getReleases(repository)
	sort.Slice(releases, func(i, j int) bool {
		return releases[i].PublishedAt.Before(releases[i].PublishedAt)
	})
	for _, release := range releases {
		fmt.Println(release.TagName)
	}

	var filtered_assets []models.Asset
	if len(pattern) > 0 {
		log.Info("Pattern found: %s",pattern)
		for _, asset := range releases[0].Assets {
			if strings.Contains(asset.Name, pattern) {
				filtered_assets = append(filtered_assets, asset)
			}
		}
	}
	if len(filtered_assets) > 2 {
		log.Error("found multiple candidate files.")
	}
	for _, file := range filtered_assets {
		fmt.Println(file.Name)
	}
	fmt.Println(filtered_assets[0].DownloadUrl)
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	response, err := resty.R().
		SetOutput(filtered_assets[0].Name).
		SetAuthToken(os.Getenv("GITHUB_TOKEN")).
		Get(filtered_assets[0].DownloadUrl)
	fmt.Println(response.RawResponse)
	if err != nil {
		fmt.Errorf("download failed: %s", err)
	}
}

func getReleases(repository string) []models.Release {
	url := fmt.Sprintf(GITHUB_RELEASE_URL_TEMPLATE, repository)
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
