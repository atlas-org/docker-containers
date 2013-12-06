package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

var g_hwaf_version = flag.String("hwaf-version", "", "hwaf version to install")
var g_siteroot = flag.String("siteroot", "/opt/atlas-sw", "where to install software")

type Dockerfile struct {
	Container   string
	Siteroot    string
	HwafVersion string
	DockerTag   string
}

const tmpl = `
## a Dockerfile to install hwaf
FROM {{.Container}}
MAINTAINER Sebastien Binet <binet@cern.ch>

ENV SITEROOT {{.Siteroot}}
RUN export SITEROOT

ENV HWAF_VERSION {{.HwafVersion}}
ENV HWAF_ROOT    $SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64
ENV PATH         $HWAF_ROOT/bin:$PATH

RUN export HWAF_VERSION
RUN export HWAF_ROOT
RUN export PATH

RUN /bin/mkdir -p $HWAF_ROOT
RUN (curl -L http://cern.ch/hwaf/downloads/tar/hwaf-$HWAF_VERSION-linux-amd64.tar.gz | tar -C $HWAF_ROOT -zxf -)

## EOF
`

func main() {
	flag.Parse()

	start := time.Now()
	defer func() {
		fmt.Printf(">>> build: [%v]\n", time.Since(start))
	}()

	if *g_hwaf_version == "" {
		*g_hwaf_version = start.Format("20060102")
	}

	dir, err := ioutil.TempDir("", "go-docker-build-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	cfg := Dockerfile{
		Container:   "binet/slc-dev",
		Siteroot:    *g_siteroot,
		HwafVersion: *g_hwaf_version,
		DockerTag:   "binet/hwaf:" + *g_hwaf_version,
	}

	docker_tmpl, err := template.New("Dockerfile").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	docker_file, err := os.Create(filepath.Join(dir, "Dockerfile"))
	if err != nil {
		panic(err)
	}
	defer docker_file.Close()

	err = docker_tmpl.Execute(docker_file, cfg)
	if err != nil {
		panic(err)
	}
	err = docker_file.Sync()
	if err != nil {
		panic(err)
	}
	err = docker_file.Close()
	if err != nil {
		panic(err)
	}

	cmdargs := []string{
		"docker",
		"build",
		"-rm",
		"-t=" + cfg.DockerTag,
		dir,
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
