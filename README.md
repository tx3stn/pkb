# pkb

**P**ersonal **K**nowledge **B**ase

Config driven CLI to manage the notes and documents that make up your personal
knowledge base.

![pkb-demo](https://user-images.githubusercontent.com/14163530/197567638-f6909de6-833c-4854-8c12-bab9c639739a.gif)

## Contents

- [Why?](#why)
- [Install](#install)
  - [Download from GitHub](#download-from-github)
  - [Build it locally](#build-it-locally)
- [Commands](#commands)
- [Configuration](#configuration)
  - [Schema](#schema)
  - [Expanding values in templates](#expanding-values-in-templates)
  - [Custom file name formats](#custom-file-name-formats)
  - [Selecting an output directory](#selecting-an-output-directory)
- [Using with Obsidian](#using-with-obsidian)

## Why?

I use [Obsidian](https://obsidian.md/) for reading/reviewing and minor updates
but don't enjoy writing content there as much as I do in
[Neovim](https://neovim.io/).

I'm more likely to keep good notes and actually maintain documents if the
editing is done in Neovim rather than having to launch a specific app every
time I need to jot something down.

`pkb` is designed to be terminal first so you can create and edit documents
where you are most comfortable, but is fully compatible with Obsidian. See
[using with Obsidian](#using-with-obsidian) for specifics on how to set them up
to work together in perfect harmony.

Don't use Obsidian for your own knowledge base?
`pkb` is just creating markdown documents from templates, so it's **probably**
compatible with whatever you are using! If it's not but you would like it to be
just raise an issue and I'll see what I can do 🤞.

## Install

### Download from GitHub

Find the latest version for your system on the
[GitHub releases page](https://github.com/tx3stn/pkb/releases).

### Build it locally

If you have go installed, you can clone this repo and run:

```bash
make install
```

This will build the binary and then copy it to `/usr/bin/pkb` so it will be
available on your path. Nothing more to it.

## Commands

Run `pkb --help` for a full, up to date list of available commands.

### `new`

Create a new note/file from your defined templates.

You will be prompted to select the template (and if defined, any sub templates)
then the file will be opened in the editor specified in your config file.

Don't want to edit it right now? Just use `--no-edit`. The file will still be
created, just not opened.

![pkb-new-demo](https://user-images.githubusercontent.com/14163530/197568150-a6cea17a-18e2-4b78-aa0d-4b8939c141b6.gif)

### `edit`

Open your editor in your notes directory.

Want to pick a specific file through `pkb` rather than in your editor? Use the
`--pick` flag to select the specific file then open it for editing.

![pkb-edit-demo](https://user-images.githubusercontent.com/14163530/197568755-bff18878-dfb4-4423-8db1-7c42e00333b5.gif)

### `copy`

Copy the contents of a file to your system clipboard.

Useful if you want to write up notes about something in your editor, but then
need to share them somewhere for other people to read.

![pkb-copy-demo](https://user-images.githubusercontent.com/14163530/197569321-92c11700-7c43-4050-9e5b-123a1e8d38cf.gif)

**Coming soon:** even more things to simplify your knowledge base management.

## Configuration

Config driven means **you** control how `pkb` works with the options in a
config file.

By default `pkb` checks for the config file in your `$XDG_CONFIG_DIR`, or
`$HOME/.config`.

### Schema

You can see an example of the config file format in the
[example.config.json](./schema/example.config.json) in the schema directory.
Make sure you add the `$schema` keyword to the top of your config file to
for in editor validation and descriptions of what fields are used for.

### Expanding values in templates

The following values will be automatically expanded in templates:

- `{{.CustomDateFormat}}` - the current date in a custom format specified in the
template config. This must be a valid golang date format, with the exception
of day suffixes (e.g. 1st, 3rd etc). If your format contains a day suffix this
will be be handled so the suffix is correctly displayed. The format string needs
to be included in the template config to be able to use a custom date format.
- `{{.Date}}` - the current date in the format YYYY-MM-DD.
- `{{.Name}}` - the name of the created document.
- `{{.Time}}` - the time the file was created in the format HH:MM.
- `{{.Week}}` - the current week number.
- `{{.Year}}` - the current year.

### Custom file name formats

By default when you create a new file from a template you will be prompted to
enter the name value, however you can customise with `name_format`.

The following values are currently supported:

- `{{.Date}}` - the current date in the format `YYYY-MM-DD`.
- `{{.Prompt}}` - prompt for user input.
- `{{.Week}}` - the current week number.
- `{{.Year}}` - the current year.

You can combine the supported formats to use them together, e.g.:

```json
"name_format": "{{.Date}}-{{.Prompt}}"
```

In this example the document would be created with the current date then the
value you typed in the prompt, e.g. `2022-09-19-typed-value.md`.

### Selecting an output directory

If you don't want to hard code the output directory in your config you can use
the dynamic values to prompt you for input at time of creation.

The following values are currently supported:

- `{{.Prompt}}` - will let you type in a new directory name. If the directory
does not already exist it will be created.
- `{{.Select}}` - select from existing directories inside the parent.

## Using with Obsidian

To get the best out of `pkb` **and** `Obsidian`, you just need to tell them to
both look in the same place for your files.

If you already have an `Obsidian` vault just set the `directory` in your `pkb`
config file to the same location.

`pkb` will expect templates to be in the directory defined in the `templates_dir`
field of your config file inside the location you specified as the `directory`
You can make sure `Obsidian` is using the same location by going to
`Settings > Templates` and setting the `Template folder location` value.

![Template folder location](https://user-images.githubusercontent.com/14163530/197546420-02c0c607-93db-454b-9d38-743e23a879f3.png)
