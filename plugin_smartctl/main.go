package main

import (
	"bufio"
	"bytes"
	//"flag"
	"fmt"
	//"github.com/toolkits/file"
	"io"
	"os"
	//"io/ioutil"
	"net"
	"os/exec"
	"strings"
	"time"
)

//smartctl
var plu_name = "plugin_smartctl"
var id, value, thresh, raw_value, endpoint string

func main() {
	//	cfg := flag.String("c", "cfg.json", "configuration file")
	//	flag.Parse()
	//	ParseConfig(*cfg)

	//	initLog()

	//	fmt.Println("hello world")
	GetSmartInfo()
}
func Ifraid() string {
	p := bytes.NewBuffer(nil)
	command := exec.Command("/bin/sh", "-c", "lspci |grep RAID")
	command.Stdout = p
	command.Run()
	f := string(p.Bytes())
	r := strings.TrimSpace(f)
	n := len(r)
	if n != 0 {
		return "smartctl --scan |grep megaraid"
	} else {
		return "smartctl --scan |grep -v megaraid"
	}
}
func GetSmartInfo() {
	endpoint = Getip()
	command := Ifraid()
	cmd := exec.Command("/bin/sh", "-c", command)
	//fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		v := strings.Fields(line)
		var smartcmd, tags string
		if len(v) == 7 {
			smartcmd = "smartctl -a " + v[0]
			tags = v[4]
		}
		if len(v) == 8 {
			smartcmd = "smartctl -a -d " + v[2] + " " + v[0]
			tags = v[5]
		}
		command := exec.Command("/bin/sh", "-c", smartcmd)
		//fmt.Println(command.Args)
		stdout, err := command.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error=>", err.Error())
		}
		command.Start()
		reader := bufio.NewReader(stdout)
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			t := time.Now().Unix()
			timestamp := fmt.Sprintf("%d", t)
			if strings.Contains(line, "Current Drive Temperature") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				tag := "disk = " + tags
				//fmt.Println(v1, timestamp, tag, endpoint)
				pushIt(v1, timestamp, "Current Drive Temperature", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Drive Trip Temperature") {
				slice1 := strings.Fields(line)
				v1 := slice1[4]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Drive Trip Temperature", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Manufactured in week") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				v2 := slice1[6]
				date := v2 + v1
				tag := "disk = " + tags
				fmt.Println(date)
				pushIt(date, timestamp, "Manufacture date", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Specified cycle count over device lifetime") {
				slice1 := strings.Fields(line)
				v1 := slice1[6]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Specified cycle count over device lifetime", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Accumulated start-stop cycles") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Accumulated start-stop cycles", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Specified load-unload count over device lifetime") {
				slice1 := strings.Fields(line)
				v1 := slice1[6]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Specified load-unload count over device lifetime", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Accumulated load-unload cycles") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Accumulated load-unload cycles", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Elements in grown defect list") {
				slice1 := strings.Fields(line)
				v1 := slice1[5]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Elements in grown defect list", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Non-medium error count") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Non-medium error count", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Blocks sent to initiator") {
				slice1 := strings.Fields(line)
				v1 := slice1[5]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Blocks sent to initiator", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Blocks received from initiator") {
				slice1 := strings.Fields(line)
				v1 := slice1[5]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Blocks received from initiator", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Blocks read from cache and sent to initiator") {
				slice1 := strings.Fields(line)
				v1 := slice1[9]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Blocks read from cache and sent to initiator", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Number of read and write commands whose size <= segment size") {
				slice1 := strings.Fields(line)
				v1 := slice1[12]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Number of read and write commands whose size <= segment size", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Number of read and write commands whose size > segment size") {
				slice1 := strings.Fields(line)
				v1 := slice1[12]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Number of read and write commands whose size > segment size", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "number of hours powered up") {
				slice1 := strings.Fields(line)
				v1 := slice1[6]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "number of hours powered up", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "number of minutes until next internal SMART test") {
				slice1 := strings.Fields(line)
				v1 := slice1[9]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "number of minutes until next internal SMART test", tag, "", "GAUGE", endpoint)
			}

			if strings.Contains(line, "read:") {
				slice1 := strings.Fields(line)
				tag := "disk = " + tags
				pushIt(slice1[1], timestamp, "Errors Corrected  by ECC fast read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[2], timestamp, "Errors Corrected  by ECC delayed read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[3], timestamp, "Errors Corrected  by ECC rewrites read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[4], timestamp, "Errors Corrected  by ECC corrected read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[5], timestamp, "Errors Corrected  by ECC invocations read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[6], timestamp, "Errors Corrected  by ECC [10^9 bytes] read", tag, "", "GAUGE", endpoint)
				pushIt(slice1[7], timestamp, "Errors Corrected  by ECC errors read", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "write:") {
				slice1 := strings.Fields(line)
				tag := "disk = " + tags
				pushIt(slice1[1], timestamp, "Errors Corrected  by ECC fast write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[2], timestamp, "Errors Corrected  by ECC delayed write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[3], timestamp, "Errors Corrected  by ECC rewrites write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[4], timestamp, "Errors Corrected  by ECC corrected write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[5], timestamp, "Errors Corrected  by ECC invocations write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[6], timestamp, "Errors Corrected  by ECC [10^9 bytes] write", tag, "", "GAUGE", endpoint)
				pushIt(slice1[7], timestamp, "Errors Corrected  by ECC errors write", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "Non-medium error count") {
				slice1 := strings.Fields(line)
				v1 := slice1[3]
				tag := "disk = " + tags
				pushIt(v1, timestamp, "Non-medium error count", tag, "", "GAUGE", endpoint)
			}
			if strings.Contains(line, "ATTRIBUTE_NAME") {
				v := strings.Fields(line)
				id = v[0]
				value = v[3]
				thresh = v[5]
				raw_value = v[9]
			}
			if strings.Contains(line, "0x00") && strings.Contains(line, "-") {
				val := strings.Fields(line)
				//LogRun(plu_name + "*****" + "smartkey: " + smartkey)
				//LogRun(plu_name + "*****" + "smartvalue : " + smartvalue)
				t := time.Now().Unix()
				timestamp := fmt.Sprintf("%d", t)
				tag1 := "type = " + id + "," + " disk = " + tags
				tag2 := "type = " + value + "," + " disk = " + tags
				tag3 := "type = " + thresh + "," + " disk = " + tags
				tag4 := "type = " + raw_value + "," + " disk = " + tags
				pushIt(val[0], timestamp, val[1], tag1, "", "GAUGE", endpoint)
				pushIt(val[3], timestamp, val[1], tag2, "", "GAUGE", endpoint)
				pushIt(val[5], timestamp, val[1], tag3, "", "GAUGE", endpoint)
				pushIt(val[9], timestamp, val[1], tag4, "", "GAUGE", endpoint)
			}
		}
		command.Wait()
	}
	cmd.Wait()
}
func Getip() string {
	address, err := net.InterfaceByName("br0")
	if err != nil {
		fmt.Println("failed to get  ip")
		os.Exit(2)
	}
	ip_info, err := address.Addrs()
	ip := strings.Split(ip_info[0].String(), "/")
	return ip[0]
}
