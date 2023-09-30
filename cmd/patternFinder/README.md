# Usage
The program provides two main actions: visualize and countInstances.

### Action: Visualize
The visualize action generates a visualization of patterns in a source image based on a target image. It takes the following command-line arguments:

- action visualize: Specify the action as visualize.
- sourceImagePath <path>: Path to the source image file.
- targetImagePath <path>: Path to the target image file.
- tolerance <value>: Tolerance for extracting patterns from the target image (default is 1).
- outputFileName <filename>: Name of the file generated with the visualization function (default is "visualization.png").
  
Example:

```
go run . -action visualize -sourceImagePath source.png -targetImagePath target.png -tolerance 2 -outputFileName result.png
```

### Action: Count Instances
The countInstances action counts the instances of a pattern from a target image within a source image. It takes the following command-line arguments:

- action countInstances: Specify the action as countInstances.
- sourceImagePath <path>: Path to the source image file.
- targetImagePath <path>: Path to the target image file.
- tolerance <value>: Tolerance for extracting patterns from the target image (default is 1).
  
Example:

```
go run . -action visualize -sourceImagePath source.png -targetImagePath target.png -tolerance 2 -outputFileName result.png
```