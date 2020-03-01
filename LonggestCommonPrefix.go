func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		firststr := strs[0]
		potision := 0
		tmpstring := ""

		for k := 1; k < len(firststr)+1; k++ {
			for i := 1; i < len(strs); i++ {
				tmpstr := strs[i]
				if len(tmpstr) > k {
					tmpstring = tmpstr[:k]
				} else {
					tmpstring = tmpstr
				}
				if firststr[:k] != tmpstring {
					break
				} else {
					if i == len(strs)-1 {
						potision = k
					}
				}
			}
		}
		if potision == 0 {
			return ""
		} else {
			return firststr[:potision]
		}
	}
}
