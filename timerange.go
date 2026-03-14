package timerange

import "time"

type TimeRange struct {
	Start         float64 `json:"start" xml:"start" bson:"start" yaml:"start"`
	StartIncluded bool    `json:"startIncluded" xml:"startIncluded" bson:"startIncluded" yaml:"startIncluded"`
	StartInfinite bool    `json:"startInfinite" xml:"startInfinite" bson:"startInfinite" yaml:"startInfinite"`
	End           float64 `json:"end" xml:"end" bson:"end" yaml:"end"`
	EndIncluded   bool    `json:"endIncluded" xml:"endIncluded" bson:"endIncluded" yaml:"endIncluded"`
	EndInfinite   bool    `json:"endInfinite" xml:"endInfinite" bson:"endInfinite" yaml:"endInfinite"`
}

func (tr *TimeRange) Valid() (ok bool) {
	if tr.StartInfinite && tr.EndInfinite {
		return true
	}

	return tr.Contains(float64(time.Now().UnixNano()) / 1e9)
}

func (tr *TimeRange) Before(timeStamp float64) (ok bool) {
	return (timeStamp >= tr.End && !tr.EndIncluded) || (timeStamp > tr.End && tr.EndIncluded)
}

func (tr *TimeRange) BeforeNow() (ok bool) {
	return tr.Before(float64(time.Now().UnixNano()) / 1e9)
}

func (tr *TimeRange) After(timeStamp float64) (ok bool) {
	return (timeStamp <= tr.Start && !tr.StartIncluded) || (timeStamp < tr.Start && tr.StartIncluded)
}

func (tr *TimeRange) AfterNow() (ok bool) {
	return tr.After(float64(time.Now().UnixNano()) / 1e9)
}

func (tr *TimeRange) Contains(timeStamp float64) (ok bool) {
	if tr.StartInfinite && tr.EndInfinite {
		return true
	}

	if tr.StartInfinite && !tr.EndInfinite {
		return timeStamp < tr.End || (timeStamp <= tr.End && tr.EndIncluded)
	} else if tr.EndInfinite && !tr.StartInfinite {
		return timeStamp > tr.Start || (timeStamp >= tr.Start && tr.StartIncluded)
	}

	return (timeStamp < tr.End || (timeStamp <= tr.End && tr.EndIncluded)) &&
		(timeStamp > tr.Start || (timeStamp >= tr.Start && tr.StartIncluded))
}

func (tr *TimeRange) StartToTime() (start time.Time, include bool, infinite bool) {
	if tr.StartInfinite {
		return time.Time{}, tr.StartIncluded, true
	}

	return time.Unix(0, int64(tr.Start*1e9)), tr.StartIncluded, false
}

func (tr *TimeRange) EndToTime() (end time.Time, include bool, infinite bool) {
	if tr.EndInfinite {
		return time.Time{}, tr.EndIncluded, true
	}

	return time.Unix(0, int64(tr.End*1e9)), tr.EndIncluded, false
}
