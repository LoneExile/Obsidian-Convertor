# Obsidian Markdown Converter (oc)

A command-line tool that converts Obsidian markdown files to regular markdown
files.

## Background

This project was born out of personal necessity, because I needed a way to use
my [Obsidian](https://obsidian.md/) markdown files with my
[Astro-built](https://astro.build/) blog. Astro likes regular markdown files and
separate image directories, and this tool makes that happen.

For example, the tool converts Obsidian's `![[image.png]]` syntax into the
regular markdown image link format `![image](path/to/image.png)`.

Right now, it's mainly all about converting image lines. But hey, I'm always
looking for ways to make this tool better. If you have an idea for a new feature
or if you're interested in contributing to the project, I encourage you to open
an issue or submit a pull request.

## Installation

<!-- You can install Obsidian-Convertor in two ways: -->

<!-- ### 1. Cloning the repository -->

Cloning the repository

```bash
git clone https://github.com/LoneExile/Obsidian-Convertor.git
cd Obsidian-Convertor
```

The binary is already built in the repository, so you can run it with `./oc`

or you can build it yourself

```bash
go build -o oc main.go
```

<!-- ### 2. Via Go -->

<!-- First, make sure you have Go installed on your machine. If not, you can download -->
<!-- it from the [official Go website](https://golang.org/dl/). -->

<!-- Once Go is installed, run the following command: -->

<!-- ``` -->
<!-- go get github.com/LoneExile/Obsidian-Convertor -->
<!-- ``` -->

<!-- This will download the repository and install the `oc` command in your -->
<!-- `$GOPATH/bin` directory. Ensure that `$GOPATH/bin` is added to your `$PATH` for -->
<!-- the `oc` command to be globally accessible. -->

## Usage

The Obsidian Markdown Converter is used with the following command:

```bash
# with the repository cloned method run: ./oc
oc <input-path> <image-path> <output-path> <output-image-path> --custom-image-path <custom-image-path>
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

## Example

Here's an example of how to use the tool with the example directories in the
repository:

```bash

# with the repository cloned method run: ./oc

# run without <custom-image-path>
oc example/SecondBranin/Blog example/SecondBranin/Assets/image example/output/blogs/ example/output/images/

# run with <custom-image-path>
oc example/SecondBranin/Blog example/SecondBranin/Assets/image example/output/blogs/ example/output/images/ image/blog/

```

After running these commands, you should find the converted markdown files in
the `example/output/blogs/` directory and the copied images in the
`example/output/images/` directory. The paths to the images in the markdown
files will be relative to the location of the markdown file itself. If you used
the `<custom-image-path>` option, the paths to the images(in the markdown file)
will instead use this custom path.
