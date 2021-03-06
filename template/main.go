package main

import (
	"fmt"
	"html/template"

	//"net"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Machine struct {
	Name string
}

type Processe struct {
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(0)
	}
}

func index(response http.ResponseWriter, request *http.Request) {
	machine := Machine{}
	temp, _ := template.ParseFiles("template/index.html")
	hostname, err := os.Hostname()
	handleErr(err)
	machine.Name = hostname //get host name
	temp.Execute(response, machine)
}

func GetInterfaceInfos(w http.ResponseWriter, r *http.Request) {

	// get interfaces MAC/hardware address

	interfStat, err := net.Interfaces()
	handleErr(err)

	html := "<html>Interfaces MAC and Hardware address infos " + "<br>"
	html = html + "<br>"

	for _, interf := range interfStat {
		html = html + "------------------------------------------------------<br>"
		html = html + "Interface Name: " + interf.Name + "<br>"

		if interf.HardwareAddr != "" {
			html = html + "Hardware(MAC) Address: " + interf.HardwareAddr + "<br>"
		}

		for _, flag := range interf.Flags {
			html = html + "Interface behavior or flags: " + flag + "<br>"
		}

		for _, addr := range interf.Addrs {
			html = html + "IPv6 or IPv4 addresses: " + addr.String() + "<br>"

		}

	}

	html = html + "</html>"

	w.Write([]byte(html))

}

func GetCPUData(w http.ResponseWriter, r *http.Request) {

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	handleErr(err)
	percentage, err := cpu.Percent(0, true)
	handleErr(err)

	html := "<html>CPU infos " + "<br>"

	html = html + "<br>"
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"

	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}

	html = html + "</html>"

	w.Write([]byte(html))

}

func GetDiskData(w http.ResponseWriter, r *http.Request) {

	diskStat, err := disk.Usage("/")
	handleErr(err)

	html := "<html> Disk infos " + "<br>"
	html = html + "<br>"

	html = html + "Disk Path: " + diskStat.Path + "  <br>"
	html = html + "Disk File system type: " + diskStat.Fstype + "  <br>"
	html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>"
	html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>"
	html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>"
	html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	html = html + "</html>"

	w.Write([]byte(html))

}

func GetHostInfos(w http.ResponseWriter, r *http.Request) {
	runtimeOS := runtime.GOOS

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	handleErr(err)

	//html := "<html>OS : " + runtimeOS + "<br>"
	html := "<html> Host infos " + "<br>"
	html = html + "<br>"
	html = html + " runtime OS : " + runtimeOS + "<br>"

	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"

	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"

	// the unique hardware id for this machine
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"
	html = html + "</html>"

	w.Write([]byte(html))

}

func GetProcInfos(w http.ResponseWriter, r *http.Request) {

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	handleErr(err)
	//get running proccesses
	miscStat, err := load.Misc()
	handleErr(err)

	//html := "<html>OS : " + runtimeOS + "<br>"
	html := "<html>Processes infos " + "<br>"
	html = html + "<br>"

	html = html + "total Number of processes: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"
	html = html + "Number of processes running: " + strconv.FormatInt(int64(miscStat.ProcsRunning), 10) + "<br>"
	html = html + "Number of blocked  prossesses: " + strconv.FormatInt(int64(miscStat.ProcsBlocked), 10) + "<br>"

	html = html + "</html>"

	w.Write([]byte(html))

}

func GetMemoryInfos(w http.ResponseWriter, r *http.Request) {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	handleErr(err)

	html := "<html>Memory infos " + "<br>"
	html = html + "<br>"
	html = html + "OS : " + runtimeOS + "<br>"
	html = html + "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>"
	html = html + "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>"
	html = html + "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"
	html = html + "</html>"

	w.Write([]byte(html))

}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	handleErr(err)

	// disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	handleErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	handleErr(err)
	percentage, err := cpu.Percent(0, true)
	handleErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	handleErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	handleErr(err)

	//get running proccesses
	miscStat, err := load.Misc()
	handleErr(err)

	//html := "<html>OS : " + runtimeOS + "<br>"
	html := "<html>System and memory infos " + "<br>"
	html = html + "<br>"
	html = html + "OS : " + runtimeOS + "<br>"
	html = html + "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>"
	html = html + "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>"
	html = html + "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// get disk serial number.... strange... not available from disk package at compile time
	// undefined: disk.GetDiskSerialNumber
	//serial := disk.GetDiskSerialNumber("/dev/sda")

	//html = html + "Disk serial number: " + serial + "<br>"

	html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>"
	html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>"
	html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>"
	html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	html = html + "<br>"
	html = html + "<br>"
	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	html = html + "CPU infos " + "<br>"
	html = html + "<br>"
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"

	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}

	html = html + "<br>"
	html = html + "<br>"
	html = html + "Prossesses infos " + "<br>"
	html = html + "<br>"
	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	html = html + "total Number of processes: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"
	html = html + "Number of processes running: " + strconv.FormatInt(int64(miscStat.ProcsRunning), 10) + "<br>"
	html = html + "Number of blocked  prossesses: " + strconv.FormatInt(int64(miscStat.ProcsBlocked), 10) + "<br>"

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	html = html + "<br>"
	html = html + "<br>"
	html = html + "Platform,interface and IP @ infos " + "<br>"
	html = html + "<br>"
	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"

	// the unique hardware id for this machine
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"

	for _, interf := range interfStat {
		html = html + "------------------------------------------------------<br>"
		html = html + "Interface Name: " + interf.Name + "<br>"

		if interf.HardwareAddr != "" {
			html = html + "Hardware(MAC) Address: " + interf.HardwareAddr + "<br>"
		}

		for _, flag := range interf.Flags {
			html = html + "Interface behavior or flags: " + flag + "<br>"
		}

		for _, addr := range interf.Addrs {
			html = html + "IPv6 or IPv4 addresses: " + addr.String() + "<br>"

		}

	}

	html = html + "</html>"

	w.Write([]byte(html))

}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/gethwdata", GetHardwareData)
	http.HandleFunc("/getCPUdata", GetCPUData)
	http.HandleFunc("/getDiskdata", GetDiskData)
	http.HandleFunc("/getHostInfos", GetHostInfos)
	http.HandleFunc("/getProcInfos", GetProcInfos)
	http.HandleFunc("/getMemoryInfos", GetMemoryInfos)
	http.HandleFunc("/getInterfaceInfos", GetInterfaceInfos)

	http.ListenAndServe(":7000", nil)

}
