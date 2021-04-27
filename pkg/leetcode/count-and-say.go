package leetcode

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSayStr(countAndSay(n - 1))
	return str
}

type kmap struct {
	v     string
	count int
}

func countAndSayStr(n string) string {
	var rs string
	kmaps := []kmap{}
	for i, j := 0, 0; i < len(n); i = j {
		km := kmap{
			v:     string(n[i]),
			count: 1,
		}
		for j = i + 1; j < len(n); j++ {
			if n[i] == n[j] {
				// fmt.Printf("%v %v %v\n", i, j, km)
				km.count = km.count + 1
			} else {
				break
			}
		}
		kmaps = append(kmaps, km)
	}
	// fmt.Printf("%+v\n", kmaps)
	for i := 0; i < len(kmaps); i++ {
		count := byte(kmaps[i].count) + '0'
		say := kmaps[i].v
		rs = rs + string([]byte{count}) + say
	}

	return rs
}
