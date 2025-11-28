package crawler

import (
	"sync"
)

type CrawlList struct {
	visited map[string]struct{}
	lock    sync.RWMutex
}
