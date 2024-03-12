package dataCalculations

import (
	"math"
	"time"

	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
)

func CalculateVolumePerDay(m *dataTransforms.MarketOrder) (float64, error) {
	downloadedAtTime, err := time.Parse("2006-01-02T15:04:05Z", m.DownloadedAt)
	if err != nil {
		return 0, err
	}
	issuedAtTime, err := time.Parse("2006-01-02T15:04:05Z", m.Issued)
	if err != nil {
		return 0, err
	}
	downloadedAtTs := downloadedAtTime.Unix()
	issuedAtTimeTs := issuedAtTime.Unix()
	timePassedInSeconds := downloadedAtTs - issuedAtTimeTs
	numOfDays := math.Ceil((float64(timePassedInSeconds) / 86400))
	volumeSold := m.VolumeTotal - m.VolumeRemain
	volumeSoldPerDay := float64(volumeSold) / numOfDays
	return volumeSoldPerDay, nil
}
