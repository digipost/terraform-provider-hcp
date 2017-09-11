package hcp

import (
	"fmt"
	"github.com/digipost/hcp"
	"regexp"
	"strings"
)

func validateHardQuota(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	re := regexp.MustCompile("^\\d+(\\.\\d{1,2})? (GB|TB)$")
	if !re.MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"Namespace hard quota uses invalid format: '%s'. Valid examples are: '10 GB' and '1.05 TB'", value))
	}
	return

}

func validateSoftQuota(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 0 || value > 100 {
		errors = append(errors, fmt.Errorf(
			"Namespace soft quota has to be an integer between 0 and 100: '%s'", value))
	}
	return
}

/**
* In English, the name you specify for a namespace must be from one through 63 characters long and can contain only
* alphanumeric characters and hyphens (-) but cannot start or end with a hyphen. In other languages, because the
* derived hostname cannot be more than 63 characters long, the name you specify may be limited to fewer than 63 characters.
* Namespace names cannot contain special characters other than hyphens and are not case sensitive. White space is not allowed.
* Namespace names cannot start with xn-- (that is, the characters x and n followed by two hyphens).
 */
func validateNamespaceName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 63 {
		errors = append(errors, fmt.Errorf(
			"Namespace name must be from one through 63 characters long: '%s'", value))
	}
	re := regexp.MustCompile("^[a-zA-Z0-9\\-]+$")
	if !re.MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"Namespace name can contain only alphanumeric characters and hyphens (-): '%s'", value))
	}
	if strings.HasPrefix(value, "-") || strings.HasSuffix(value, "-") {
		errors = append(errors, fmt.Errorf(
			"Namespace name cannot start or end with a hyphen (-): '%s'", value))
	}
	if strings.HasPrefix(value, "xn--") {
		errors = append(errors, fmt.Errorf(
			"Namespace name cannot start with xn-- (that is, the characters x and n followed by two hyphens): '%s'", value))
	}
	return
}

func validateHashScheme(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	schemes := map[string]bool{
		hcp.SHA_1:     true,
		hcp.SHA_256:   true,
		hcp.SHA_384:   true,
		hcp.SHA_512:   true,
		hcp.MD5:       true,
		hcp.RIPEMD160: true,
	}

	if !schemes[value] {
		errors = append(errors, fmt.Errorf("Namespace hash scheme uses an illegal value: '%s'", value))
	}
	return
}

func validateOptimizedFor(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	schemes := map[string]bool{
		hcp.CLOUD: true,
		hcp.ALL:   true,
	}

	if !schemes[value] {
		errors = append(errors, fmt.Errorf("Namespace hash scheme uses an illegal value: '%s'", value))
	}
	return
}

func validateAclsUsage(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	schemes := map[string]bool{
		hcp.NOT_ENABLED:  true,
		hcp.NOT_ENFORCED: true,
		hcp.ENFORCED:     true,
	}

	if !schemes[value] {
		errors = append(errors, fmt.Errorf("Namespace acls usage uses an illegal value: '%s'", value))
	}
	return
}
