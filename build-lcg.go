package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"
)

var g_hwaf_version = flag.String("hwaf-version", "20131204", "hwaf version to use")
var g_hwaf_variant = flag.String("hwaf-variant", "x86_64-slc6-gcc47-opt", "hwaf variant to use")
var g_siteroot = flag.String("siteroot", "/opt/atlas-sw", "where to install software")
var g_worch_profile = flag.String("profile", "build-lcg.cfg", "worch-profile to run for the build")
var g_docker_tag = flag.String("docker-tag", "binet/lcg-65", "tag to apply as a result of the build")

type Dockerfile struct {
	Container   string
	Siteroot    string
	HwafVersion string
	HwafVariant string
	Profile     string
	LcgBranch   string
	DockerTag   string
}

const tmpl = `
## a Dockerfile to install the whole LCG stack
FROM {{.Container}}
MAINTAINER binet@cern.ch

ENV HOME     /root
RUN export   HOME

ENV SITEROOT {{.Siteroot}}
RUN export SITEROOT

ENV HWAF_VARIANT {{.HwafVariant}}
ENV HWAF_VERSION {{.HwafVersion}}
ENV HWAF_ROOT    $SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64
ENV PATH         $HWAF_ROOT/bin:$PATH

RUN export HWAF_VARIANT
RUN export HWAF_VERSION
RUN export HWAF_ROOT
RUN export PATH

VOLUME ["/build"]
ADD build-lcg.sh /build/build-lcg.sh

RUN /build/build-lcg.sh /build $HWAF_VERSION $HWAF_VARIANT $SITEROOT {{.LcgBranch}} {{.Profile}}

ENV MODULEPATH $SITEROOT/sw/modules/$HWAF_VARIANT:$MODULEPATH
RUN export MODULEPATH

## EOF
`

func main() {
	flag.Parse()

	start := time.Now()
	defer func() {
		fmt.Printf(">>> build: [%v]\n", time.Since(start))
	}()

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("", "go-docker-build-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	//script := "/build/build-lcg.sh"
	voldir := filepath.Join(pwd, "lcg", *g_hwaf_variant)

	err = copytree(dir, voldir)
	if err != nil {
		panic(err)
	}

	cfg := Dockerfile{
		Container:   "binet/slc",
		Siteroot:    *g_siteroot,
		HwafVersion: *g_hwaf_version,
		HwafVariant: *g_hwaf_variant,
		Profile:     *g_worch_profile,
		LcgBranch:   "lcg-65-branch",
		DockerTag:   *g_docker_tag + "-" + *g_hwaf_variant,
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

func copytree(dstdir, srcdir string) error {
	var err error

	if !path_exists(dstdir) {
		err = os.MkdirAll(dstdir, 0755)
		if err != nil {
			return err
		}
	}

	err = filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		rel := ""
		rel, err = filepath.Rel(srcdir, path)
		out := filepath.Join(dstdir, rel)
		fmode := info.Mode()
		if fmode.IsDir() {
			err = os.MkdirAll(out, fmode.Perm())
			if err != nil {
				return err
			}
		} else if fmode.IsRegular() {
			dst, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR, fmode.Perm())
			if err != nil {
				return nil
			}
			src, err := os.Open(path)
			if err != nil {
				return nil
			}
			_, err = io.Copy(dst, src)
			if err != nil {
				return nil
			}
		} else if (fmode & os.ModeSymlink) != 0 {
			rlink, err := os.Readlink(path)
			if err != nil {
				return err
			}
			err = os.Symlink(rlink, out)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unhandled mode (%v) for path [%s]", fmode, path)
		}
		return nil
	})
	return err
}

func path_exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// EOF
