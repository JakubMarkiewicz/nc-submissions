type TimestampValue struct {
	timestamp int
	value string
}

type TimeMap struct {
	data map[string][]TimestampValue
}

func Constructor() TimeMap {
	return TimeMap{data: map[string][]TimestampValue{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	_, ok := this.data[key]
	if !ok {
		this.data[key] = []TimestampValue{}
	}
	this.data[key] = append(this.data[key], TimestampValue{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps, ok := this.data[key]

	if !ok {
		return ""
	}

	closest := -1

	l, r := 0, len(timestamps) - 1

	for l <= r {
		mid := l + (r-l)/2
		midVal := timestamps[mid]

		if midVal.timestamp > timestamp {
			r = mid - 1
		} else {
			closest = mid
			l = mid + 1
		}
	}

	if closest == -1 {
		return ""
	}

	return timestamps[closest].value
}
