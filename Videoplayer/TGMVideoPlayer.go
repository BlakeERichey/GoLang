package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	service := "omxplayer"
	hmdir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	} else {
		var files []string

		//get files
		dir := hmdir + "/TGM/Primary"
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if path != dir {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		} else { //play once
			for _, file := range files {
				args := make([]string, 0)
				args = append(args, "--aspect-mode fill")
				args = append(args, file)
				playVideo(service, 1, args...)
			}
		}

		//get files
		dir = hmdir + "/TGM/Secondary"
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
			for { //play forever
				for _, file := range files {
					args := make([]string, 0)
					args = append(args, "--aspect-mode fill")
					args = append(args, file)
					playVideo(service, 1, args...)
				}
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
			log.Println(cmd.Run())
		}
	}
}
