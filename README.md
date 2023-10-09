# About
Pixel Art Tools is command-line utility wirten in Go that offers various functionalities for working with pixel arts and image processing. It's main focus is in analyzing and generating data using images and data created during event `r/place` event that takes place on reddit every now and then but `Visualize` and `Count Instances` modes can be used outside of this event since they are not tied to any specific pictures.

# Showcase
### Counting amount of pixel arts in different pixel art (-mode countInstances)
https://github.com/Neidz/pixel-art-tools/assets/68917960/ae41995b-f46a-48c8-a725-8bfd9f12abc3

### Visualizing pixel arts in different pixel art (-mode visualize)
https://github.com/Neidz/pixel-art-tools/assets/68917960/6af008b0-9177-4989-a0ba-137722682edb

### Parsing r/place .csv data and creating images on specified intervals (-mode imagesFromRplaceFeed)
https://github.com/Neidz/pixel-art-tools/assets/68917960/431ba84f-f07e-4637-b803-85a139dde3be

# Permissions
On Unix-based systems you will have to give executable permission to run

``` bash
chmod +x ./pixel-art-tools-os
```

# Link to .csv data
All pixel data was made public in [this post on reddit](https://www.reddit.com/r/place/comments/15bjm5o/rplace_2023_data/).

# Modes

## -mode visualize
The `visualize` mode generates a visualization of patterns in a source image based on a target image.

It takes the following command-line arguments:

- `sourceImagePath <path>`: Path to the source image file.
- `targetImagePath <path>`: Path to the target image file.
- `tolerance <value>`(default: 1): Tolerance for extracting patterns from the target image (default is 1).
- `outputFileName <filename>`(default: "visualization.png"): Name of the file generated with the visualization function.
- `targetColor <value>`(default: #000000 (black)): Hex code of the color that will be used to search for target pattern.

Example:

Only required parameters
``` bash
./pixel-art-tools -mode visualize -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/crewmate.png 
```

All parameters
``` bash
./pixel-art-tools -mode visualize -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/crewmate.png -tolerance 2 -outputFileName result.png -targetColor #0000FF
```

## -mode countInstaces
The `countInstances` mode counts the instances of a pattern from a target image within a source image. 

It takes the following command-line arguments:

- `sourceImagePath <path>`: Path to the source image file.
- `targetImagePath <path>`: Path to the target image file.
- `tolerance <value>`(default: 1): Tolerance for extracting patterns from the target image.
- `targetColor <value>`(default: #000000): Hex code of the color that will be used to search for target pattern.
  
Example:

Only required parameters
``` bash
./pixel-art-tools -mode countInstances -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/crewmate.png 
```

All parameters
``` bash
./pixel-art-tools -mode countInstances -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/crewmate.png -tolerance 2 -targetColor #0000FF
```

## -mode imagesFromRplaceFeed
The `imagesFromRplaceFeed` mode in Pixel Art Tools is designed to generate images based on data from r/place CSV feeds. This mode provides a versatile way to visualize and process the data collected during r/place events, allowing you to create images that represent the evolution of the canvas over time.

It takes the following command-line arguments:

- `directoryPath <path>` (default: "./"): Path to the directory containing r/place CSV feed files.
- `baseName <name>` (default: "2023_place_canvas_history-"): Full name of the files without numbers.
- `numbersInName <value>` (default: 12): Amount of numbers present after the base name.
- `amountOfFiles <value>`(default: 12): Number of files that should be used for creating images.
- `verbose`(default: false): Specifies whether the output should be verbose, providing extensive logging.
- `saveEveryHours`(default: false): Specifies if images should be saved at hourly intervals.
- `saveEveryMinutes`(default: false): Specifies if images should be saved at minute intervals.
- `saveEveryValue <value>`(deafult: 1): Interval of hours or minutes at which images will be saved.
- `outputDir <path>`(default: "output"): Name or full path of the directory where generated images will be saved. If it doesn't exist, the tool will create it (default output).
  
Example:

Basic usage (parsing all files from 2023 r/place)
``` bash
./pixel-art-tools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -saveEveryMinutes
```

Save images every hour from first 10 .csv files (parsing all files from 2023 r/place)
``` bash
./pixel-art-tools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 10 -saveEveryValue 1 -saveEveryHours
```

Save images every 10 minutes from first 3 .csv files while displaying informations about parsed data (parsing all files from 2023 r/place)
``` bash
./pixel-art-tools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 3 -saveEveryValue 10 -saveEveryMinutes -verbose
```

Save images every 3 hours from all 53 .csv files (parsing all files from 2023 r/place)
``` bash
./pixel-art-tools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 53 -saveEveryValue 3 -saveEveryHours
```

# Notes

### General
- **if you want to create images from entire r/place event (17.9GB of .csv files from 2023) then you should really consider location of those files. Using HDD instead of SSD is going to be very slow, so move your data to your SSD since it will be abottleneck.**
- imagesFromRplaceFeed mode assumes that all [data provided by reddit](https://www.reddit.com/r/place/comments/15bjm5o/rplace_2023_data/) is unpacked and in the same folder. Amount of numbers and file names are default values for `numbersInName` and `baseName`, so if raw data is provided without any changes (other than unpacking) then those arguments don't have to be passed.
- both visualize and countInstances use go routines, so the speed of finding patterns heavily depends on amount of threads of your cpu.
- verbose parameter in imagesFromRplaceFeed mode doesn't silence warnings, errors and important informations. Some problems won't stop program but will be printed to the console like for example failing to parse the line in .csv file.

### Rules for matching in visualize and countInstances modes
- all pixels surrounding target have to be different color (big square of the same color won't be matched).
- all pixels of matched target have to be the same color.
- target has to be exact match.
