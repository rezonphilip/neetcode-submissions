type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 { return "" }
	meta := make([]string, len(strs))
	stringsT := ""
	for i, str := range strs {
		meta[i] = strconv.Itoa(len(str))
		stringsT = stringsT + str
	}
	metaT := strings.Join(meta, "_")
	stringsT = metaT+":"+stringsT
	return stringsT
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 { return []string{} }
	decodedStrs := []string{}
	encodedSplit := strings.SplitN(encoded, ":", 2)
	meta := encodedSplit[0]
	metaIndex := strings.Split(meta, "_")
	indexStart := len(meta)+1
	for _, val := range metaIndex {
		strLen, _ := strconv.Atoi(val)
		decodedStrs = append(decodedStrs, encoded[indexStart:indexStart+strLen])
		indexStart = indexStart + strLen
	}
	return decodedStrs
}
