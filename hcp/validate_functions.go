package hcp

import (
	"fmt"
	"regexp"
)

func validateHardQuota(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	re := regexp.MustCompile("^\\d+ (GB|TB)$")
	if !re.MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q uses invalid format: '%s'. Valid examples are: '10 GB' and '1 TB'", k, value))
	}
	return

}
