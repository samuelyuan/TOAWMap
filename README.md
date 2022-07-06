# TOAWMap

This program will read scenario files (.sce) and render the maps.

The following games are supported:

* The Operational Art of War: Century of Warfare
* The Operational Art of War III
* The Operational Art of War IV

### Command-Line Usage

```
./TOAWMap.exe -input=[input filename] -output=[output filename]
```

Example
```
./TOAWMap.exe -input=scenario.sce -output=scenario.png
```

<div style="display:inline-block;">
<img src="https://raw.githubusercontent.com/samuelyuan/TOAWMap/master/screenshots/korea50.png" alt="korea50" width="145" height="300" />
<img src="https://raw.githubusercontent.com/samuelyuan/TOAWMap/master/screenshots/manchuria.png" alt="manchuria" width="300" height="300" />
</div>

### About

There have been many maps generated for The Operational Art of War (TOAW) series, but there isn't much information on the file formats and there is no way to parse the data. This project analyzes the different games and you can see how there is a lot of overlap between the various file formats.
