// Package util provides utility functionality
package util

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

// GenerateBlockDevices generates block devices from /dev, not including any of the following: loop, dm, nbd, sr
func GenerateBlockDevices() ([]string, error) {
	log := zap.L().Sugar()

	files, err := os.ReadDir("/sys/block")
	if err != nil {
		return nil, err
	}

	var blockDevices []string
	for i := range files {
		if strings.HasPrefix(files[i].Name(), "dm") || strings.HasPrefix(files[i].Name(), "loop") || strings.HasPrefix(files[i].Name(), "nbd") || strings.HasPrefix(files[i].Name(), "sr") || strings.HasPrefix(files[i].Name(), "vd") {
			continue
		}

		absPath := fmt.Sprintf("/dev/%s", files[i].Name())

		if _, err := os.Stat(absPath); errors.Is(err, os.ErrNotExist) {
			log.Warnf("block device %s does not exist", absPath)

			continue
		}

		blockDevices = append(blockDevices, absPath)
	}

	return blockDevices, nil
}
