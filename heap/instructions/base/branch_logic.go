package base

import "JVM-Go/rtda/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.GetThread().GetPC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
