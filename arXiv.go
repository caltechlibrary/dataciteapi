package dataciteapi

import (
	"fmt"
	"strings"
)

// Checks if a string is an ArXiv id.
func IsArXiv(s string) bool {
	if strings.HasPrefix(s, "https://arxiv.org/abs/") {
		return true
	}
	if strings.HasPrefix(s, "arXiv:") || strings.HasPrefix(s, "https://10.48550/arxiv.") {
		return true
	}
	return false
}

// ArXivURLtoArXiv will convert an absolute arxiv.org URL to an arXiv id format.
func ArXivURLtoArXiv(s string) string {
	if strings.HasPrefix(s, "https://arxiv.org/abs/") {
		return fmt.Sprintf("arXiv:%s", strings.TrimPrefix(s, "https://arxiv.org/abs/"))
	}
	return s
}

// ArXivToDOI converts an arXiv id to a DOI (without the "https://doi.org/" prefix) formatted id per
// Instructions in the announcement that all arXiv gets DOI at
// https://blog.arxiv.org/2022/02/17/new-arxiv-articles-are-now-automatically-assigned-dois/
func ArXivToDOI(s string) string {
	if strings.Contains(s, "arxiv.org/abs") {
		s = ArXivURLtoArXiv(s)
	}
	if strings.HasPrefix(s, "arXiv:") {
		return fmt.Sprintf("10.48550/%s", strings.Replace(strings.TrimSpace(s), ":", ".", 1))
	}
	return s
}
