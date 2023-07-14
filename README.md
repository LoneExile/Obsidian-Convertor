# Obsidian Markdown Converter

A command-line tool that converts Obsidian markdown files to regular markdown
files.

## Background

This project was born out of personal necessity, because I needed a way to use
my [Obsidian](https://obsidian.md/) markdown files with my
[Astro-built](https://astro.build/) blog. Astro likes regular markdown files and
separate image directories, and this tool makes that happen.

For example, the tool converts Obsidian's `![[image.png]]` syntax into the
regular markdown image link format `![image](path/to/image.png)`. It also has
the ability to convert images to different formats (jpg, png, avif) and optimize
them by specifying the quality.

Right now, it's mainly all about converting image lines. But hey, I'm always
looking for ways to make this tool better. If you have an idea for a new feature
or if you're interested in contributing to the project, I encourage you to open
an issue or submit a pull request.

<details>
  <summary><b>🖼️&nbsp;&nbsp;Example image</b></summary>

![2023-06-21_16-41](https://github.com/LoneExile/obsidian-convertor/assets/82561297/2c796f6d-a850-45b3-8d22-736c6aa48f98)

![2023-06-21_16-35](https://github.com/LoneExile/obsidian-convertor/assets/82561297/be99c56b-6c3c-4c7a-bfd3-19199e974f9e)

</details>

## Installation

> ❗Dependencies: <https://github.com/davidbyttow/govips#dependencies>

You can install Obsidian-Convertor in two ways:

### 1. Cloning the repository

```bash
git clone https://github.com/LoneExile/Obsidian-Convertor.git
cd Obsidian-Convertor
```

The binary is already built in the repository, so you can run it with
`./obsidian-convertor`

or you can build it yourself

```bash
go build -o obsidian-convertor main.go
```

### 2. Via Go

First, make sure you have Go installed on your machine. If not, you can download
it from the [official Go website](https://golang.org/dl/).

Once Go is installed, run the following command:

```
go install github.com/LoneExile/obsidian-convertor@v0.1.6
```

This will download the repository and install the `obsidian-convertor` command
in your `$GOPATH/bin` directory. Ensure that `$GOPATH/bin` is added to your
`$PATH` for the `obsidian-convertor` command to be globally accessible.

## Usage

The Obsidian Markdown Converter is used with the following command:

```bash
# with the repository cloned method run: ./obsidian-convertor
obsidian-convertor convert <input-path> \
<image-path> \
<output-path> \
<output-image-path> \
[<custom-image-path>] \
[--format <format>] \
[--quality <quality>]

```

Where:

- `<input-path>` is the path to the directory containing the Obsidian markdown
  files you want to convert.
- `<image-path>` is the path to the directory containing the images referenced
  in your Obsidian markdown files.
- `<output-path>` is the path to the directory where you want to save the
  converted markdown files.
- `<output-image-path>` is the path to the directory where you want to save the
  copied images.
- `<custom-image-path>` (optional) is a custom path for images in the converted
  markdown files. If not provided, the path to the images in the converted
  markdown files will be the same as `<output-image-path>`.
- `<format>` (optional) is the output format for the images (options: jpg, png,
  avif, same). Default is 'same', which means the images will keep their
  original format.
- `<quality>` (optional) is the quality for output images (1-100). Default
  is 100.

## Example

Here's an example of how to use the tool with the example directories in the
repository:

```bash
# with the repository cloned method run: ./obsidian-convertor

# run without <custom-image-path>
obsidian-convertor convert example/SecondBrain/Blog \
example/SecondBrain/Assets/image \
example/output/blogs/ \
example/output/images/ \
--format jpg \
--quality 85

# run with <custom-image-path>
obsidian-convertor convert example/SecondBrain/Blog \
example/SecondBrain/Assets/image \
example/output/blogs/ \
example/output/images/ \
image/blog/ \ # <custom-image-path>
--format png \
--quality 90

```

After running these commands, you should find the converted markdown files in
the `example/output/blogs/` directory and the copied images in the
`example/output/images/` directory. The paths to the images in the markdown
files will be relative to the location of the markdown file itself. If you used
the `<custom-image-path>` option, the paths to the images(in the markdown file)
will instead use this custom path.

## TODO

- [ ] Convert back to Obsidian format
- [ ] Add tests
- [ ] integrate with [Viper](https://github.com/spf13/viper) for config file?
- [ ] integrate with [Charm](https://github.com/charmbracelet)?
