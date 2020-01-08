# Installation Instructions

Ensure setup and TGMVideoPlayer are executable  
`chmod u+x setup`  
`chmod u+x TGMVideoPlayer`  

Afterwards, run setup
`./setup`  
The output should contain no "errors".

This will create a folder in `/home/pi` called `TGM`. 
Additionally, `TGMVideoPlayer` should have been copied into this folder, automatically.  

Inside `TGM` will be two folders:  
* Primary  
All videos to be run once and only once should be put in here  
* Secondary  
Videos placed here will be run indefinitely   

# Autorun on Bootup instructions  
Ensure that the binary works correctly before initializing autorun. 
To do so, go to `/home/pi/TGM` and run `./TGMVideoPlayer`. If the program executes, 
then it is safe to proceed.  

To enable the script to autorun on bootup simply  
run the command `crontab -e` and add `@reboot /home/pi/TGM/TGMVideoPlayer` to the end.  
Then reboot the system `sudo reboot`.

# Autorun from USB  
The following is slightly different from the previous bootup instructions.

Ensure that the binary works correctly before initializing autorun. 
To do so, go to `/home/pi/TGM` and run `./TGMVideoPlayer`. If the program executes, 
then it is safe to proceed. This should still run in `home/pi/TGM`.  

To enable the script to autorun on bootup simply run the command  `crontab -e`  
and add  
```@reboot sleep 60 && /home/pi/TGM/TGMVideoPlayer -path=/media/pi/USB_DRIVE_NAME```  
to the end. This will give the USB time to mount before running the program.
Then reboot the system `sudo reboot`.  

The USB Direcotory should be structured similarly to `/home/pi` containing a folder 
called `TGM` along with two subdirectories inside TGM called `Primary` and `Secondary`