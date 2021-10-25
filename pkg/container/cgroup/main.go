package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()

		cmd := exec.Command("sh", "-c", `stress --vm-bytes 2000m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		fmt.Printf("%v", cmd.Process.Pid)
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)

		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit",
			"tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)

		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit",
			"memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait()
}
