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
	for idx, v := range timestamps {
		if v.timestamp > timestamp {
			continue
		}
		if closest == -1 {
			closest = idx
			continue
		}

		currDiff := timestamp - timestamps[closest].timestamp 
		diff := v.timestamp - timestamp

		if diff < currDiff {
			closest = idx
		}
	}

	if closest == -1 {
		return ""
	}

	return timestamps[closest].value
}
