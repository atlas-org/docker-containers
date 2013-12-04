package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var g_hwaf_version = flag.String("hwaf-version", "20131203", "hwaf version to use")
var g_hwaf_variant = flag.String("hwaf-variant", "x86_64-slc6-gcc47-opt", "hwaf variant to use")
var g_siteroot = flag.String("siteroot", "/opt/atlas-sw", "where to install software")

func main() {
	flag.Parse()

	script := "/build/build-lcg.sh"
	fmt.Printf(">>> [%s]\n", script)

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	voldir := filepath.Join(pwd, "lcg", *g_hwaf_variant)

	docker := exec.Command(
		"sudo",
		"docker",
		"run",
		fmt.Sprintf("-v=%s:/build", voldir),
		"binet/slc",
		"/bin/sh",
		script,
		*g_hwaf_variant,
		*g_hwaf_version,
		*g_siteroot,
	)
	docker.Stdout = os.Stdout
	docker.Stderr = os.Stderr
	docker.Stdin = os.Stdin

	err = docker.Run()
	if err != nil {
		panic(err)
	}
}

// EOF
