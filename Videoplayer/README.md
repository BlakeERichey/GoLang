# Installation Instructions

Ensure setup and TGMVideoPlayer are executable  
`chmod u+x setup`  
`chmod u+x TGMVideoPlayer`  

Afterwards, run setup
`./setup`  
The output should contain no "errors".

This will create a folder in `/home/pi` called `TGM`. 
Inside `TGM` will be two folders:  
* Primary  
All videos to be run once and only once should be put in here  
* Secondary  
Videos placed here will be run indefinitely  

# Autorun on Bootup instructions
To enable the script to autorun on bootup simply run the command `crontab -e` and 
add `@reboot /home/pi/TGM/TGMVideoPlayer` to end.  
Then reboot the system `sudo reboot`.