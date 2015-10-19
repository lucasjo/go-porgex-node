package usage

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/lucasjo/go-porgex-node/models"
)

var AppMemCgroupPath = "/cgroup/memory/openshift/"

const (
	memstat  = "memory"
	megaByte = 1024 * 1024
)

func SetMemoryStats(uuid string, v *models.MemStats) error {

	cgroupPath := filepath.Join(AppMemCgroupPath, uuid)

	usagefile := strings.Join([]string{memstat, "usage_in_bytes"}, ".")
	maxUsagefile := strings.Join([]string{memstat, "max_usage_in_bytes"}, ".")
	limitfile := strings.Join([]string{memstat, "limit_in_bytes"}, ".")

	usageValue, err := getUsageUint(cgroupPath, usagefile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", usagefile, err)
		return err
	}

	maxUsageValue, err := getUsageUint(cgroupPath, maxUsagefile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", maxUsagefile, err)
		return err
	}

	limitValue, err := getUsageUint(cgroupPath, limitfile)

	if err != nil {
		fmt.Errorf("failed to parse %s - %v\n", limitfile, err)
		return err
	}

	v.Current_usage = usageValue
	v.Max_usage = maxUsageValue
	v.Limit_usage = limitValue
	v.Create_at = time.Now()
	v.AppId = uuid

	return nil

}
