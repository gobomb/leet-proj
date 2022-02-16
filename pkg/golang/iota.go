package golang

const (
	Zero   = iota // 0
	First         // 1
	Second        // 2
	Hi     = "a"  // a, 被打断后，后续值不变，直到用iota显示恢复
	Four          // a
	Five          // a
	Six    = iota // 6，显示恢复，iota接着累加，中间打断不中断累加
	Seven         // 7
)
