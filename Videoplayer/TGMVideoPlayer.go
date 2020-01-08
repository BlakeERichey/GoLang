package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	service := "omxplayer"
	hmdir := "/home/pi"
	var files []string

	//get files
	dir := hmdir + "/TGM/Primary"
	fmt.Println("Dir:", dir)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path != dir {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	} else { //play once
		fmt.Println("Primary Files:", files)
		for _, file := range files {
			args := make([]string, 0)
			args = append(args, "--aspect-mode")
			args = append(args, "fill")
			args = append(args, file)
			playVideo(service, 1, args...)
		}
	}

	//get files
	dir = hmdir + "/TGM/Secondary"
	fmt.Println("Dir:", dir)
	files = make([]string, 0)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path != dir {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Secondary Files:", files)
		for { //play forever
			for _, file := range files {
				args := make([]string, 0)
				args = append(args, "--aspect-mode")
				args = append(args, "fill")
				args = append(args, file)
				playVideo(service, 1, args...)
			}
		}
	}
}

func playVideo(service string, iterations int, args ...string) {
	if iterations > 0 {
		for i := 0; i < iterations; i++ {
			cmd := exec.Command(service, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			disableScreen() //turn off background
			log.Println(cmd.Run())
		}
	}
}

func disableScreen() {
	args := []string{"dpms", "force", "off"}
	cmd := exec.Command("xset", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Println(cmd.Run())
}
