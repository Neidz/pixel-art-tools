# About
Pixel Art Tools is command-line utility wirten in Go that offers various functionalities for working with pixel arts and image processing. It's main focus is in analyzing and generating data using images and data created during event `r/place` event that takes place on reddit every now and then but `Visualize` and `Count Instances` modes can be used outside of this event since they are not tied to any specific pictures.

# Modes

## -mode visualize
The `visualize` mode generates a visualization of patterns in a source image based on a target image.

It takes the following command-line arguments:

- `sourceImagePath <path>`: Path to the source image file.
- `targetImagePath <path>`: Path to the target image file.
- `tolerance <value>`(default: 1): Tolerance for extracting patterns from the target image (default is 1).
- `outputFileName <filename>`(default: "visualization.png"): Name of the file generated with the visualization function.

Example:

Only required parameters
```
./pixelArtTools -mode visualize -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/reversedCrewmate.png 
```

All parameters
```
./pixelArtTools -mode visualize -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/reversedCrewmate.png -tolerance 2 -outputFileName result.png
```

## -mode countInstaces
The `countInstances` mode counts the instances of a pattern from a target image within a source image. 

It takes the following command-line arguments:

- `sourceImagePath <path>`: Path to the source image file.
- `targetImagePath <path>`: Path to the target image file.
- `tolerance <value>`(default: 1): Tolerance for extracting patterns from the target image.
  
Example:

Only required parameters
```
./pixelArtTools -mode countInstances -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/reversedCrewmate.png 
```

All parameters
```
./pixelArtTools -mode countInstances -sourceImagePath ./data/final_2023_place.png -targetImagePath ./data/reversedCrewmate.png -tolerance 2
```

## -mode imagesFromRplaceFeed
The `imagesFromRplaceFeed` mode in Pixel Art Tools is designed to generate images based on data from r/place CSV feeds. This mode provides a versatile way to visualize and process the data collected during r/place events, allowing you to create images that represent the evolution of the canvas over time.

It takes the following command-line arguments:

- `directoryPath <path>` (default: "./"): Path to the directory containing r/place CSV feed files.
- `baseName <name>` (default: "2023_place_canvas_history-"): Full name of the files without numbers.
- `numbersInName <value>` (default: 12): Amount of numbers present after the base name.
- `amountOfFiles <value>`(default: 12): Number of files that should be used for creating images.
- `verbose <true|false>`(default: false): Specifies whether the output should be verbose, providing extensive logging.
- `saveEveryHours <true|false>`(default: false): Specifies if images should be saved at hourly intervals.
saveEveryMinutes <true|false>: Specifies if images should be saved at minute intervals.
- `saveEveryValue <value>`(deafult: 1): Interval of hours or minutes at which images will be saved.
- `outputDir <path>`(default: "output"): Name or full path of the directory where generated images will be saved. If it doesn't exist, the tool will create it (default output).
  
Example:

Basic usage (parsing all files from 2023 r/place)
```
./pixelArtTools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -saveEveryMinutes true
```

Save images every hour from first 10 .csv files (parsing all files from 2023 r/place)
```
./pixelArtTools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 10 -saveEveryHours true -saveEveryValue 1
```

Save images every 10 minutes from first 3 .csv files while displaying informations about parsed data (parsing all files from 2023 r/place)
```
./pixelArtTools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 3 -saveEveryHours true -saveEveryValue 10 -verbose true
```

Save images every 3 hours from all 53 .csv files (parsing all files from 2023 r/place)
```
./pixelArtTools -mode imagesFromRplaceFeed -directoryPath ./data/rplace_data -baseName 2023_place_canvas_history- -numbersInPath 12 -amountOfFiles 53 -saveEveryHours true -saveEveryValue 3
```