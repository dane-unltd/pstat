package pstat

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type Process struct {
	User    user.User
	oldStat ProcessStat
	Stat    ProcessStat
	CpuTime ProcessLoad
}

type ProcessStatus struct{}
type ProcessIO struct{}
type ProcessStatm struct{}

type ProcessLoad struct {
	Utime int
	Stime int
}

type ProcessStat struct {
	Pid         int
	Comm        string
	State       string
	Ppid        int
	Pgrp        int
	Session     int
	Tty_nr      int
	Tpgid       int
	Flags       int
	Minflt      int
	Cminflt     int
	Majflt      int
	Cmajflt     int
	Utime       int
	Stime       int
	Cutime      int
	Cstime      int
	Priority    int
	Nice        int
	Itrealvalue int
	Starttime   int
	Vsize       int
	Rss         int
	Rlim        int
	Startcode   int
	Endcode     int
	Startstack  int
	Kstkesp     int
	Kstkeip     int
	Signal      int
	Blocked     int
	sigignore   int
	sigcatch    int
	wchan       int
	nswap       int
	cnswap      int
	Exit_signal int
	Processor   int
}

func NewProcess(pid int) (*Process, error) {
	p := Process{}
	p.Stat.Pid = pid
	err := UpdateProcessStat(&p.Stat, &p.User)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Process) Update() error {
	p.oldStat = p.Stat
	err := UpdateProcessStat(&p.Stat, &p.User)
	if err != nil {
		return err
	}
	p.CpuTime.Utime = p.Stat.Utime - p.oldStat.Utime
	p.CpuTime.Stime = p.Stat.Stime - p.oldStat.Stime
	return nil
}

func UpdateProcessStat(ps *ProcessStat, u *user.User) error {
	baseDir := "/proc/"
	pidDir := filepath.Join(baseDir, strconv.Itoa(ps.Pid))
	fi, err := os.Stat(pidDir)
	if err != nil {
		return err
	}
	if u != nil {
		ut, err := user.LookupId(strconv.Itoa(int(fi.Sys().(*syscall.Stat_t).Uid)))
		if err != nil {
			return err
		}
		*u = *ut
	}

	statFile := filepath.Join(pidDir, "stat")

	b, err := ioutil.ReadFile(statFile)
	if err != nil {
		return err
	}

	content := string(b)

	fields := strings.Fields(content)

	ps.Comm = fields[1]
	ps.State = fields[2]
	ps.Ppid, _ = strconv.Atoi(fields[3])
	ps.Pgrp, _ = strconv.Atoi(fields[4])
	ps.Session, _ = strconv.Atoi(fields[5])
	ps.Tty_nr, _ = strconv.Atoi(fields[6])
	ps.Tpgid, _ = strconv.Atoi(fields[7])
	ps.Flags, _ = strconv.Atoi(fields[8])
	ps.Minflt, _ = strconv.Atoi(fields[9])
	ps.Cminflt, _ = strconv.Atoi(fields[10])
	ps.Majflt, _ = strconv.Atoi(fields[11])
	ps.Cmajflt, _ = strconv.Atoi(fields[12])
	ps.Utime, _ = strconv.Atoi(fields[13])
	ps.Stime, _ = strconv.Atoi(fields[14])
	ps.Cutime, _ = strconv.Atoi(fields[15])
	ps.Cstime, _ = strconv.Atoi(fields[16])
	ps.Priority, _ = strconv.Atoi(fields[17])
	ps.Nice, _ = strconv.Atoi(fields[18])
	ps.Itrealvalue, _ = strconv.Atoi(fields[19])
	ps.Starttime, _ = strconv.Atoi(fields[20])
	ps.Vsize, _ = strconv.Atoi(fields[21])
	ps.Rss, _ = strconv.Atoi(fields[22])
	ps.Rlim, _ = strconv.Atoi(fields[23])
	ps.Startcode, _ = strconv.Atoi(fields[24])
	ps.Endcode, _ = strconv.Atoi(fields[25])
	ps.Startstack, _ = strconv.Atoi(fields[26])
	ps.Kstkesp, _ = strconv.Atoi(fields[27])
	ps.Kstkeip, _ = strconv.Atoi(fields[28])
	ps.Signal, _ = strconv.Atoi(fields[29])
	ps.Blocked, _ = strconv.Atoi(fields[30])
	ps.sigignore, _ = strconv.Atoi(fields[31])
	ps.sigcatch, _ = strconv.Atoi(fields[32])
	ps.wchan, _ = strconv.Atoi(fields[33])
	ps.nswap, _ = strconv.Atoi(fields[34])
	ps.cnswap, _ = strconv.Atoi(fields[35])
	ps.Exit_signal, _ = strconv.Atoi(fields[36])
	ps.Processor, _ = strconv.Atoi(fields[37])

	return nil
}
