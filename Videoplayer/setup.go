package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//make TGM directory to hold videos
func main() {
	hmdir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(hmdir)
		omx := exec.Command("apt-get", "install", "omxplayer")
		cmd := exec.Command("mkdir", hmdir+"/TGM")
		primary := exec.Command("mkdir", hmdir+"/TGM/Primary")
		secondary := exec.Command("mkdir", hmdir+"/TGM/Secondary")
		cp := exec.Command("cp", "TGMVideoPlayer.exe", "/bin")
		log.Println(omx.Run())
		log.Println(cmd.Run())
		log.Println(primary.Run())
		log.Println(secondary.Run())
		log.Println(cp.Run())
	}
}
