Goflakes
===

Generating [Reiter snowflakes](http://www.patarnott.com/pdf/SnowCrystalGrowth.pdf) in go-lang.

Reiter published a set of rules for a cellular automaton that simulates snow-crystal growth. Each cell in the lattice has a value representing mass.  Values above an arbitrary threshold are considered "frozen."  For each discrete time step, cells are classified as either receptive or non-receptive based on their proximity to frozen cells, and values flow across the lattice towards the receptive cells which simply capture and accumulate mass.

This repository contains a go-lang script within a docker container which runs a parameterized version of Reiter's model.  At each discrete time step, an SVG image is written to the outputs directory.  After the model has run, the SVG images are converted to PNG images and ffmpeg is invoked to compile the images into an mp4 movie.

[![Watch the video](https://i.imgur.com/IVo1nfO.png)](https://youtu.be/38IR0x8xoxg)

Running locally
---

Start the container:

```sh
docker-compose up
```

Generate SVG stills:
```sh
docker-compose run goflakes /src/goflakes -width=400 -height=400 -iters=002000
```

To convert the SVG to PNG:
```sh
docker-compose run goflakes sh -c 'for input in /src/output/*.svg; do rsvg-convert -w 400 -h 400 $input -o $input.png; rm $input; done'
```

To convert the PNG to a video:
```sh
docker-compose run goflakes sh -c 'ffmpeg -framerate 60 -pattern_type glob -i "/src/output/*.png" -c:v libx264 -pix_fmt yuv420p /src/output/video.mp4'
```

To terminate the container:
```sh
docker-compose down
```

Changing model parameters
---

If you'd like to change the default flags, before building the container, change the flags in the Dockerfile:
```
CMD go run main.go -width=1000 -height=1000 -iters=000500 && tail -f /dev/null
```

If you'd like to change the default model parameters, before building the container, edit the constants in `constants/model.go`:
```
// Ice is the threshold value above which cells are considered "frozen"
const Ice = 0.9

// Additive is the addition constant 𝛾
const Additive = 0.0001

// BackgroundLevel is the initial background level β
const BackgroundLevel = 0.4
```

TODO
---

go does not support concurrent writes to a map without a lock.  Investigate storing the hexmap as an array instead of a map, concurrently update the values, and just amend the hex.Key().  This would require changing how neighborhoods are currently implemented unless both a map and an array are maintained.

```
Copyright (c) 2019 Nathaniel Wroblewski
I am making my contributions/submissions to this project solely in my personal
capacity and am not conveying any rights to any intellectual property of any
third parties.
```