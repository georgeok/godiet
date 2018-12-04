package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	out, err := exec.Command("go", "build", "-work", "-o", "testExecutable").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove("testExecutable")
	output := string(out)
	if !strings.HasPrefix(output, "WORK=") {
		println("No build directory found")
	}
	txtLines := strings.Split(output, "\n")

	path := strings.Replace(txtLines[0], "WORK=", "", 1)
	path = strings.TrimSpace(path)
	println(path)
	linkFile := path + "/b001/importcfg.link"
	_, e := os.Stat(linkFile)
	if e != nil {
		panic(e)
		print("Link file doesn't exist")
		print(linkFile)
		os.Exit(1)
	}

	b, _ := ioutil.ReadFile(linkFile)
	lines := strings.Split(string(b), "\n")
	packageMap := make(map[string]string)
	for _, line := range lines {
		line = strings.Replace(line, "packagefile ", "", 1)
		if !strings.Contains(line, "=") {
			continue
		}
		split := strings.Split(line, "=")
		packageMap[split[0]] = split[1]
	}

	bySize := make(map[float64]string)

	fi, e := os.Stat("testExecutable")
	totalSize := float64(fi.Size())

	for k, v := range packageMap {
		f, _ := os.Stat(v)

		size := float64(f.Size())
		text := fmt.Sprintf("%-40s %4.0f kb \t%.2f %%", k, size/1024, size/totalSize*100) //kb
		if f.Size() > 1024*1024 {                                                         // more that 1 mb
			text = fmt.Sprintf("%-40s %4.2f mb \t%.2f %%", k, size/1024/1024, size/totalSize*100) //mb
		}
		bySize[float64(f.Size())] = text
	}

	// To store the keys in slice in sorted order
	var keys []float64
	for k, _ := range bySize {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	fmt.Printf("Total %d packages\n", len(packageMap))
	fmt.Printf("Output size: %4.2f mb\n", float64(fi.Size())/1024/1024)
	fmt.Println("=========================================================")
	// To perform the opertion you want
	for i, k := range keys {
		fmt.Printf("%-3d %s\n", i+1, bySize[k])
	}

}
