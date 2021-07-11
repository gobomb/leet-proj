package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := range ransomNote {
		m[ransomNote[i]] = m[ransomNote[i]] + 1
	}
	for j := 0; j < len(magazine); j++ {
		if v, ok := m[magazine[j]]; ok {
			m[magazine[j]] = v - 1
		}
	}
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
