# Civ5ToHexEmpire3

This program will convert a Civ 5 map (.Civ5Map) to a Hex Empire 3 map (.he3).

### Command-Line Usage

```
./Civ5ToHexEmpire3.exe -input=[input filename] -output=[output filename]
```

Example
```
./Civ5ToHexEmpire3.exe -input=earth.Civ5Map -output=earth.he3
```

### About

Hex Empire 3 comes with a map editor, but there aren't that many user generated maps. I wanted to see if there was a way to reuse maps from Civilization 5 and convert them to be used for Hex Empire 3. Recreating the maps tile by tile in Hex Empire 3 would be time consuming. Since I had information on both the file formats, I could write a program that automatically converted maps from Civ 5 to Hex Empire 3, which would save a lot of time. The remaining work would be to open up the converted map in the Hex Empire 3 map editor to convert some of the towns into factories or cities and set the capitals for each player.

The maximum size of a Hex Empire 3 map (large) is 60x60, but I was able to load much larger maps into the game using this conversion tool. However the game wasn't designed for larger maps and will freeze if the map is too large. Some of the maps had to be resized to fit the 60x60 map.

### Examples

Europe

<div style="display:inline-block;">
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/europe_civ5.png" alt="europe_civ5" width="415" height="300" />
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/europe_he3.png" alt="europe_he3" width="350" height="300" />
</div>

India
<div style="display:inline-block;">
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/india_civ5.png" alt="europe_civ5" width="300" height="300" />
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/india_he3.png" alt="europe_he3" width="300" height="300" />
</div>


Stalingrad

<div style="display:inline-block;">
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/stalingrad_civ5.png" alt="stalingrad_civ5" width="400" height="300" />
<img src="https://raw.githubusercontent.com/samuelyuan/Civ5ToHexEmpire3/master/screenshots/stalingrad_he3.png" alt="stalingrad_he3" width="400" height="300" />
</div>