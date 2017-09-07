package hcp

import (
	"crypto/sha512"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
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
