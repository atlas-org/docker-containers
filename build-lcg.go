package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var g_hwaf_version = flag.String("hwaf-version", "20131204", "hwaf version to use")
var g_hwaf_variant = flag.String("hwaf-variant", "x86_64-slc6-gcc47-opt", "hwaf variant to use")
var g_siteroot = flag.String("siteroot", "/opt/atlas-sw", "where to install software")
var g_worch_profile = flag.String("profile", "build-lcg.cfg", "worch-profile to run for the build")
var g_docker_tag = flag.String("docker-tag", "binet/lcg-65", "tag to apply as a result of the build")

func main() {
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//script := "/build/build-lcg.sh"
	voldir := filepath.Join(pwd, "lcg", *g_hwaf_variant)

	cmdargs := []string{
		"docker",
		"build",
		"-t=" + *g_docker_tag + "-" + *g_hwaf_variant,
		//"run",
		//fmt.Sprintf("-v=%s:/build", voldir),
		//"binet/slc",
		voldir,
		//"/bin/sh",
		//script,
		//"/build",
		//*g_hwaf_version,
		//*g_hwaf_variant,
		//*g_siteroot,
		//*g_worch_profile,
	}

	fmt.Printf(">>> [%v]\n", cmdargs)

	docker := exec.Command("sudo", cmdargs...)

	docker.Stdout = os.Stdout
	docker.Stderr = os.Stderr
	docker.Stdin = os.Stdin

	err = docker.Run()
	if err != nil {
		panic(err)
	}
}

// EOF
