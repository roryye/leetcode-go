package main

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, s := range str {
			key[s-'a']++
		}
		m[key] = append(m[key], str)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
