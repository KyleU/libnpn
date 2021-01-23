package npncore

import "time"

func PString(v string) *string {
	return &v
}

func PStringV(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func PInt(v int) *int {
	return &v
}

func PIntV(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

func PInt64(v int64) *int64 {
	return &v
}

func PInt64V(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

func PTime(v time.Time) *time.Time {
	return &v
}

func PTimeV(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}
