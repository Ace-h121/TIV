Here’s an in-depth README file for your program:

---

# ASCII Image Viewer (TIV)

## Overview

The **ASCII Image Viewer (TIV)** is a Go program that converts images into ASCII art. It reads an image file, processes it, and displays an ASCII representation of the image in the terminal. The program uses customizable scaling factors to control the resolution of the ASCII art.

---

## Features

- **Supports multiple image formats**: Works with PNG, JPEG, and GIF files.
- **Customizable resolution**: Users can specify vertical and horizontal scaling factors to adjust the ASCII art's granularity.
- **Lightweight and fast**: Efficient processing of images using Go's standard libraries.

---

## Installation

1. Install [Go](https://go.dev/doc/install) if you haven't already.
2. Clone this repository:
   ```bash
   git clone <repository_url>
   cd <repository_directory>
   ```
3. Build the program:
   ```bash
   go build -o tiv main.go
   ```

---

## Usage

### Command Syntax

```bash
tiv <image_path> <scale_y> <scale_x>
```

### Arguments

- `<image_path>`: Path to the image file to be converted to ASCII art.
- `<scale_y>`: Vertical scale factor. Must be a positive integer.
- `<scale_x>`: Horizontal scale factor. Must be a positive integer.

### Example

```bash
tiv example.png 10 5
```

This command processes `example.png` and generates ASCII art with a vertical scale of `10` and a horizontal scale of `5`.

### Help

To display usage instructions:

```bash
tiv help
```

---

## How It Works

1. **Image Loading**: The program loads the image using the `image` package.
2. **Grayscale Conversion**: Each pixel's grayscale value is calculated using weighted contributions of red, green, and blue components.
3. **Pixel Averaging**: Pixels are grouped into blocks defined by the scaling factors, and the average grayscale value of each block is calculated.
4. **ASCII Mapping**: The average grayscale value is mapped to an ASCII character from a predefined ramp.
5. **Display**: The ASCII art is printed to the terminal.

---

## Notes

- **Grayscale Mapping**: The program uses the luminosity method to convert RGB values to grayscale:
  \[
  \text{Grayscale} = 0.299 \times R + 0.587 \times G + 0.114 \times B
  \]
- **Character Ramp**: The ramp determines the ASCII character for a given grayscale value. Darker characters correspond to lower grayscale values, and lighter characters correspond to higher values.

---

## Common Errors

- **Missing Arguments**: If no arguments are provided, the program prompts for usage instructions.
- **Invalid Scale Factors**: If the scale factors are not integers, the program exits with an error.
- **File Errors**: If the image file cannot be loaded, the program exits with a panic.

---

## Improvements and Future Features

- Add support for color ASCII art.
- Include a flag for exporting the ASCII art to a text file.
- Optimize pixel averaging for large images.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

Let me know if you'd like to tweak anything!
