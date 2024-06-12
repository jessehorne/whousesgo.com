package util

import "github.com/gosimple/slug"

func GetSlugFromName(name string) string {
	return slug.Make(name)
}
