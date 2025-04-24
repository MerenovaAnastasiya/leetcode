package tasks

func FindAnagrams(s string, p string) []int {

	var pCount [26]int
	var windowCount [26]int
	var result []int

	if len(p) > len(s) {
		return []int{}
	}

	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}

	for r := 0; r < len(s); r++ {
		windowCount[s[r]-'a']++
		if r > len(p)-1 {
			windowCount[s[r-len(p)]-'a']--
		}
		flag := true
		for i := 0; i < 26; i++ {
			if pCount[i] != windowCount[i] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, r-len(p)+1)
		}

	}
	return result
}
