package usage

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func getUsageUint(path, param string) (uint64, error) {

	contents, err := ioutil.ReadFile(filepath.Join(path, param))

	if err != nil {
		return 0, err
	}

	return ParseUint(strings.TrimSpace(string(contents)), 10, 64)

}

func ParseUint(s string, base, bitSize int) (uint64, error) {
	value, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		intValue, intErr := strconv.ParseInt(s, base, bitSize)

		if intErr == nil && intValue < 0 {
			return 0, nil

		} else if intErr != nil && intErr.(*strconv.NumError).Err == strconv.ErrRange && intValue < 0 {
			return 0, nil
		}
		return value, err
	}
	return value, err
}
