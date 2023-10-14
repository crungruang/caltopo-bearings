# CalTopo Bearings
This will take a JSON exported file from CalTopo and calculate bearing & distance between all present Markers

- This is a command-line utility only
- These instructions _assume_ (unless you save to another dir), you will run things from your Downloads directory: 
  - i.e. `/home/$username/Downloads/` for Linux
  - i.e. `/Users/$username/Downloads` 
  - i.e. `\Users\$username\Downloads\` for Windows

---

## Usage

### CalTopo Export

This is required to save a map locally on your workstation

1. Open desired CalTopo Map
2. Go to Export (top left of screen)
3. Save the file (`Export` button)

### OSX/Linux Users

1. Download the `calbear` file in this [repo](https://github.com/crungruang/caltopo-bearings/blob/main/calbear) the download icon should look like ![this](images/download_icon.png) 
2. Open a Terminal (OSX: Click on Launchpad Icon and type `terminal` and hit Enter`)
3. Navigate to your downloads directory
4. Type in `chmod +x calbear` to make the file executable (only needs to be done once)
5. Type in `./calbear`
6. You will be prompted to type in the name of the JSON file you exported from CalTopo

### Windows Users

1. Download the `calbear.exe` file in this [repo](https://github.com/crungruang/caltopo-bearings/blob/main/calbear.exe) the download icon should look like ![this](images/download_icon.png)
2. Open a Terminal (Start > Run > type `cmd` and hit Enter)
3. Navigate to your downloads directory (`cd Downloads`)
4. Type `calbear.exe`
5. You will be prompted to type in the name of the JSON file you exported from CalTopo

---

**See the videos folder for a [video demo](videos/calbearings_demo.mp4) with audio**

_If you have trouble opening up terminal on your workstation, maybe do some Googling -- there's a lot of ways to get there and covering all of them is out of scope for this README._
