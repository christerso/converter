# Image Format Converter
This is a command-line tool that converts image files to either JPEG or PNG format.

## Usage
The tool takes two optional command-line arguments:

target (default: "."): the directory containing the image files to convert.
format (default: "jpg"): the desired output format (either "jpg" or "png").
Here is an example of how to run the tool with the default options:

```shell
$ ./image-format-converter
```
To specify a target directory and an output format, use the following syntax:

```shell
$ ./image-format-converter -target /path/to/images -format png
```

## Supported Formats
The tool supports converting files with the extensions ".jpg", ".jpeg", and ".png". It skips files that are already in the desired format.

## Conversion Process
When converting an image file, the tool performs the following steps:

- Open the input file.
- Decode the input image to an image.Image object.
- Determine the output file name and format.
- Create the output file.
- Encode the image.Image object in the desired output format.
- Save the output to the new file with the same name but the new format.
- Print out the names of the original and converted files for each successful conversion.
## Output
At the end of the execution, the tool prints out the total number of files converted.