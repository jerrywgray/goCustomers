package routes

import (
	"github.com/palantir/stacktrace"
)

type sumResults struct {
	Total int64
}

const (
	eCodeTimestampFormat = stacktrace.ErrorCode(iota + 600)
	eCodeRangeFormat
)
