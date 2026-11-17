package clickhouse

import (
	"strconv"
	"time"
)

func toUInt32(v any) uint32 {
	switch t := v.(type) {
	case float64:
		return uint32(t)
	case float32:
		return uint32(t)
	case int:
		return uint32(t)
	case int64:
		return uint32(t)
	case uint64:
		return uint32(t)
	case string:
		if t == "" {
			return 0
		}
		n, _ := strconv.ParseUint(t, 10, 32)
		return uint32(n)
	default:
		return 0
	}
}

func toFloat64(v any) float64 {
	switch t := v.(type) {
	case float64:
		return t
	case float32:
		return float64(t)
	case int:
		return float64(t)
	case int64:
		return float64(t)
	case uint64:
		return float64(t)
	case string:
		if t == "" {
			return 0
		}
		n, _ := strconv.ParseFloat(t, 64)
		return n
	default:
		return 0
	}
}

func parseTime(v any) time.Time {
	switch t := v.(type) {
	case time.Time:
		return t
	case string:
		if t == "" {
			return time.Now()
		}
		parsed, err := time.Parse(time.RFC3339Nano, t)
		if err == nil {
			return parsed
		}
		parsed, _ = time.Parse(time.RFC3339, t)
		return parsed
	default:
		return time.Now()
	}
}
