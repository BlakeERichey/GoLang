package main

import (
	"log"
	"os/exec"
)

//make TGM directory to hold videos
func main() {
	hmdir := "/home/pi"
	omx := exec.Command("apt-get", "install", "omxplayer")
	cmd := exec.Command("mkdir", hmdir+"/TGM")
	primary := exec.Command("mkdir", hmdir+"/TGM/Primary")
	secondary := exec.Command("mkdir", hmdir+"/TGM/Secondary")
	cp := exec.Command("cp", "TGMVideoPlayer", hmdir+"/TGM")
	log.Println(omx.Run())
	log.Println(cmd.Run())
	log.Println(primary.Run())
	log.Println(secondary.Run())
	log.Println(cp.Run())
}
