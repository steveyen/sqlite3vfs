package sqlite3vfs

import (
	"crypto/rand"
	"time"
)

type DefaultVFSv1 struct {
	VFS
}

func (vfs *DefaultVFSv1) Randomness(n []byte) int {
	i, err := rand.Read(n)
	if err != nil {
		panic(err)
	}
	return i
}

func (vfs *DefaultVFSv1) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (vfs *DefaultVFSv1) CurrentTime() time.Time {
	return time.Now()
}
