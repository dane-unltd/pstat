package pstat

type System struct {
	oldStat Stat
	Stat    Stat

	CpuTime CPUStat

	NumCpus int
}

func NewSystem() *System {
	return nil
}
