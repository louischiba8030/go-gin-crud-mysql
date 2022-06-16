package util

import (
	"time"
	"math/rand"
	"github.com/oklog/ulid"
)

func GenerateULID () string {
	// Generate ULID
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return id.String()
}