"".main STEXT size=71 args=0x0 locals=0x20 funcid=0x0
	0x0000 00000 (main.go:5)	TEXT	"".main(SB), ABIInternal, $32-0
	0x0000 00000 (main.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:5)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:5)	PCDATA	$0, $-2
	0x000d 00013 (main.go:5)	JLS	64
	0x000f 00015 (main.go:5)	PCDATA	$0, $-1
	; 栈从高位向低位增长
	; SUBQ 对 SP 的值减32 （栈增长32）
	0x000f 00015 (main.go:5)	SUBQ	$32, SP
	; 把 BP 中的值存到 SP+24 （把 BP 存到栈上）（32-8, 8用来放 BP）
	0x0013 00019 (main.go:5)	MOVQ	BP, 24(SP)
	; 将 SP+24 的地址复制给 BP
	0x0018 00024 (main.go:5)	LEAQ	24(SP), BP
	0x001d 00029 (main.go:5)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:6)	MOVQ	$3, (SP)
	0x0025 00037 (main.go:6)	MOVQ	$5, 8(SP)
	0x002e 00046 (main.go:6)	PCDATA	$1, $0
	0x002e 00046 (main.go:6)	CALL	"".add(SB)
	; 把 SP+24 的值复制回 BP
	0x0033 00051 (main.go:7)	MOVQ	24(SP), BP
	; 栈收缩32
	0x0038 00056 (main.go:7)	ADDQ	$32, SP
	0x003c 00060 (main.go:7)	RET
	0x003d 00061 (main.go:7)	NOP
	0x003d 00061 (main.go:5)	PCDATA	$1, $-1
	0x003d 00061 (main.go:5)	PCDATA	$0, $-2
	0x003d 00061 (main.go:5)	NOP
	0x0040 00064 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (main.go:5)	PCDATA	$0, $-1
	0x0045 00069 (main.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 31 48  dH..%....H;a.v1H
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 c7 04  .. H.l$.H.l$.H..
	0x0020 24 03 00 00 00 48 c7 44 24 08 05 00 00 00 e8 00  $....H.D$.......
	0x0030 00 00 00 48 8b 6c 24 18 48 83 c4 20 c3 0f 1f 00  ...H.l$.H.. ....
	0x0040 e8 00 00 00 00 eb b9                             .......
	rel 5+4 t=17 TLS+0
	rel 47+4 t=8 "".add+0
	rel 65+4 t=8 runtime.morestack_noctxt+0
"".add STEXT nosplit size=65 args=0x18 locals=0x10 funcid=0x0
	; 即将分配栈帧为16（ BP + int 大小）; 两个 int 入参加一个返回值 == 24(调用方分配的栈大小)
	0x0000 00000 (main.go:9)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $16-24
	0x0000 00000 (main.go:9)	SUBQ	$16, SP
	0x0004 00004 (main.go:9)	MOVQ	BP, 8(SP)
	0x0009 00009 (main.go:9)	LEAQ	8(SP), BP
	0x000e 00014 (main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x000e 00014 (main.go:9)	MOVQ	$0, "".~r2+40(SP)
	0x0017 00023 (main.go:10)	MOVQ	$0, "".t(SP)
	0x001f 00031 (main.go:11)	MOVQ	$2, "".t(SP)
	0x0027 00039 (main.go:13)	MOVQ	"".a+24(SP), AX
	0x002c 00044 (main.go:13)	ADDQ	"".b+32(SP), AX
	0x0031 00049 (main.go:13)	MOVQ	AX, "".~r2+40(SP)
	0x0036 00054 (main.go:13)	MOVQ	8(SP), BP
	0x003b 00059 (main.go:13)	ADDQ	$16, SP
	0x003f 00063 (main.go:13)	NOP
	0x0040 00064 (main.go:13)	RET
	0x0000 48 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 48 c7  H...H.l$.H.l$.H.
	0x0010 44 24 28 00 00 00 00 48 c7 04 24 00 00 00 00 48  D$(....H..$....H
	0x0020 c7 04 24 02 00 00 00 48 8b 44 24 18 48 03 44 24  ..$....H.D$.H.D$
	0x0030 20 48 89 44 24 28 48 8b 6c 24 08 48 83 c4 10 90   H.D$(H.l$.H....
	0x0040 c3                                               .
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
