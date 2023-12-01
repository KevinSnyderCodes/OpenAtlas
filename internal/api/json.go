package api

import "time"

const (
	RFC3339Plus = "2006-01-02T15:04:05-07:00"
)

type TimeRFC3339Plus time.Time

func (o TimeRFC3339Plus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(o).Format(RFC3339Plus) + `"`), nil
}
