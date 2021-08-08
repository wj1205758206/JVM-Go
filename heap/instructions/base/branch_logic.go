package base

import "JVM-Go/heap/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.GetThread().GetPC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
