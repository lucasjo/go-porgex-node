package usage

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lucasjo/go-porgex-node/models"
	"github.com/lucasjo/go-porgex-node/system"
)

var appCpuAcctPath = "/cgroup/cpuacct/openshift/"

const (
	cpuacctStat         = "cpuacct.stat"
	nanosecondsInSecond = 1000000000
)

var clockTicks = uint64(system.GetClockTicks())

func GetCpuUsageStat(path string) (uint64, uint64, error) {
	userModeUsage := uint64(0)
	systemModeUsage := uint64(0)

	const (
		userField   = "user"
		systemField = "system"
	)

	// user <usage in ticks>
	// system <usage in ticks>

	data, err := ioutil.ReadFile(filepath.Join(path, cpuacctStat))

	if err != nil {
		return 0, 0, err
	}
	fields := strings.Fields(string(data))

	fmt.Printf("fields : %v\n", fields)

	if len(fields) != 4 {
		return 0, 0, fmt.Errorf("failure - %s is expected to have 4 fields", filepath.Join(path, cpuacctStat))
	}

	if fields[0] != userField {
		return 0, 0, fmt.Errorf("unexpected field %q in %q, expected %q", fields[0], cpuacctStat, userField)
	}
	if fields[2] != systemField {
		return 0, 0, fmt.Errorf("unexpected field %q in %q, expected %q", fields[2], cpuacctStat, systemField)
	}
	if userModeUsage, err = strconv.ParseUint(fields[1], 10, 64); err != nil {
		return 0, 0, err
	}

	if systemModeUsage, err = strconv.ParseUint(fields[3], 10, 64); err != nil {
		return 0, 0, err
	}

	return (userModeUsage * nanosecondsInSecond) / clockTicks, (systemModeUsage * nanosecondsInSecond) / clockTicks, nil

}

func SetCpuUsage(uuid string, stats *models.AppCpuStats) error {

	appCgroupPath := filepath.Join(appCpuAcctPath, uuid)

	userModeUsage, systemModeUsage, err := GetCpuUsageStat(appCgroupPath)

	if err != nil {
		return err
	}

	totalUsage, err := getUsageUint(appCgroupPath, "cpuacct.usage")

	if err != nil {
		fmt.Errorf("Error : ", err)
		return err
	}

	perCpuUsage, err := getPercpuUsage(appCgroupPath)

	if err != nil {
		return err
	}

	stats.CPUStats.CPUUsage.TotalUsage = totalUsage
	stats.CPUStats.CPUUsage.UsageInSytemmode = systemModeUsage
	stats.CPUStats.CPUUsage.UsageInUsermode = userModeUsage
	stats.CPUStats.CPUUsage.PercpuUsage = perCpuUsage

	return nil
}

func getPercpuUsage(path string) ([]uint64, error) {
	percpuUsage := []uint64{}
	data, err := ioutil.ReadFile(filepath.Join(path, "cpuacct.usage_percpu"))
	if err != nil {
		return percpuUsage, err
	}
	for _, value := range strings.Fields(string(data)) {
		value, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return percpuUsage, fmt.Errorf("Unable to convert param value to uint64: %s", err)
		}
		percpuUsage = append(percpuUsage, value)
	}
	return percpuUsage, nil
}

func calculateCPUPercent(previousCpu uint64, v *models.AppCpuStats) float64 {

	var cpuPercnt = 0.0

	fmt.Printf("previous : %v , current : %v \n", previousCpu, v.CPUStats.CPUUsage.TotalUsage)

	cpudeta := float64((v.CPUStats.CPUUsage.TotalUsage - previousCpu)) / float64(1000000000)

	cpuPercnt = (cpudeta / float64(len(v.CPUStats.CPUUsage.PercpuUsage))) * 100.0

	return cpuPercnt

}
