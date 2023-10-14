# CalTopo Bearings
This will take a JSON exported file from CalTopo and calculate bearing & distance between all present Markers

- This is a command-line utility only
- These instructions _assume_ (unless you save to another dir), you will run things from your Downloads directory: 
  - i.e. `/home/$user/Downloads/` for OSX/Linux 
  - i.e. `$User\My Downloads\` for Windows

---

## Usage

### CalTopo Export

This is required to save a map locally on your workstation

1. Open desired CalTopo Map
2. Go to Export (top left of screen)
3. Save the file

### OSX/Linux Users

1. Open a Terminal (OSX: Click on Launchpad Icon and type `terminal` and hit Enter`)
2. Navigate to your downloads directory
3. Type in `chmod +x calbear` to make the file executable (only needs to be done once)
4. Type in `./calbear`
5. You will be prompted to type in the name of the JSON file you exported from CalTopo

### Windows Users

1. Open a Terminal (Start > Run > type `cmd` and hit Enter)
2. Navigate to your downloads directory
3. Type `calbear.exe`
4. You will be prompted to type in the name of the JSON file you exported from CalTopo

---

**See the videos folder for a video demo with audio**

_If you have trouble opening up terminal on your workstation, maybe do some Googling -- there's a lot of ways to get there and covering all of them is out of scope for this README._
