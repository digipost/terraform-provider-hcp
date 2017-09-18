package hcp

import (
	"crypto/sha512"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func suppressPasswordDiffs(k, old, new string, d *schema.ResourceData) bool {
	suppress := old == sha512sum(new)
	log.Printf("[DEBUG] suppressPasswordDiffs old: %s new: %s suppress: %t", old, new, suppress)
	return suppress
}

func sha512sum(new string) string {
	sum := sha512.Sum512([]byte(new))
	shasum := fmt.Sprintf("%x", sum)
	return shasum
}

func suppressHardQuotaDiffs(k, old, new string, d *schema.ResourceData) bool {
	suppress := normalizeHardQuota(old) == normalizeHardQuota(new)
	log.Printf("[DEBUG] suppressHardQuotaDiffs old: %s new: %s suppress: %t", old, new, suppress)
	return suppress
}

func normalizeHardQuota(hardQuota string) string {
	split := strings.Split(hardQuota, " ")
	number := split[0]
	unit := split[1]
	floatNumber, _ := strconv.ParseFloat(number, 64)
	return fmt.Sprintf("%0.2f %s", floatNumber, unit)
}
