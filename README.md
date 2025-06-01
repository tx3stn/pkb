# pkb

**P**ersonal **K**nowledge **B**ase

Config driven CLI to manage the notes and documents that make up your personal
knowledge base.

![pkb-demo](https://github.com/user-attachments/assets/686ba6d8-fb42-4dab-a740-cb897231c0aa)

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
  - [Accessible mode](#accessible-mode)
- [Using with Obsidian](#using-with-obsidian)
  - [Expanding Obsidian Template Variables](#expanding-obsidian-template-variables)
- [Using with multiple vaults or computers](#using-with-multiple-vaults-or-computers)

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

This will build the binary and then copy it to `/usr/local/bin/pkb` so it will be
available on your path. Nothing more to it.

## Commands

Run `pkb --help` for a full, up to date list of available commands.

### `new`

Create a new note/file from your defined templates.

You will be prompted to select the template (and if defined, any sub templates)
then the file will be opened in the editor specified in your config file.

Don't want to edit it right now? Just use `--no-edit`. The file will still be
created, just not opened.

![pkb-new-demo](https://github.com/user-attachments/assets/63cce0d5-726f-4903-addf-dc0c254b2c08)

### `edit`

Select an existing note and open it in your editor.

![pkb-edit-demo](https://github.com/user-attachments/assets/d81b65ca-1afa-46e9-8e1e-e5345334997e)

### `copy`

Copy the contents of a file to your system clipboard.

Useful if you want to write up notes about something in your editor, but then
need to share them somewhere for other people to read.

![pkb-copy-demo](https://user-images.githubusercontent.com/14163530/197569321-92c11700-7c43-4050-9e5b-123a1e8d38cf.gif)

### `open`

Open your notes directory in your specified editor.


## Configuration

Config driven means **you** control how `pkb` works with the options in a
config file.

By default `pkb` checks for a config file called `pkb.json` in your `$XDG_CONFIG_DIR`, 
or `$HOME/.config` in a directory called `pkb`, e.g: `$HOME/.config/pkb/pkb.json`.

See  [Using with multiple vaults or computers](#using-with-multiple-vaults-or-computers)
if you want to support multiple different config files for different uses.

### Schema

You can see an example of the config file format in the
[pkb.json](./.schema/pkb.json) in the schema directory.

> [!TIP]
> Make sure you add the `$schema` keyword to the top of your config file to
> for in editor validation and descriptions of what fields are used for.

### Expanding values in templates

The following values will be automatically expanded in templates:

- `{{.CustomDateFormat}}` - the current date in a custom format specified in the
template config. This must be a valid golang date format, with the exception
of day suffixes (e.g. 1st, 3rd etc). If your format contains a day suffix this
will be be handled so the suffix is correctly displayed. The format string needs
to be included in the template config to be able to use a custom date format.
- `{{.Date}}` - the current date in the format YYYY-MM-DD.
- `{{.Directory}}` - the name of the parent directory.
- `{{.Name}}` - the name of the created document.
- `{{.SelectFromList "filename.json"}}` - multiple select options from a json file
in the same directory as the template, then render the selected options as a comma
separated list if multiple are selected.
- `{{.Time}}` - the time the file was created in the format HH:MM.
- `{{.Week}}` - the current week number.
- `{{.Year}}` - the current year.

See [Expanding Obsidian Template Variables](#expanding-obsidian-template-variables)
for details on using the default Obsidian template values.

### Custom file name formats

By default when you create a new file from a template you will be prompted to
enter the name value, however you can customise with `name_format`.

The following values are currently supported:

- `{{.Date}}` - the current date in the format `YYYY-MM-DD`.
- `{{.Prompt}}` - prompt for user input.
- `{{.Week}}` - the current week number.
- `{{.Year}}` - the current year.

> [!TIP]
> You can combine the supported formats to use them together, e.g.:
> 
> ```json
> "name_format": "{{.Date}}-{{.Prompt}}"
> ```
> 
> In this example the document would be created with the current date then the
> value you typed in the prompt, e.g. `2022-09-19-typed-value.md`.

### Selecting an output directory

If you don't want to hard code the output directory in your config you can use
the dynamic values to prompt you for input at time of creation.

The following values are currently supported:

- `{{.Prompt}}` - will let you type in a new directory name. If the directory
does not already exist it will be created.
- `{{.Select}}` - select from existing directories inside the parent.
- `{{.Year}}` - the current year.

### Accessible mode

Some of the options when creating a file spawn interactive TUI elements. 
If you would prefer for them to use standard prompt elements which should work 
better with screen reader tools, you can set `"accessible_mode": true` in your 
config file.

![accessible mode](https://github.com/user-attachments/assets/e6638d47-311a-4a88-875c-e3c6499e67cc)

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

### Expanding Obsidian Template Variables

For ease of compatibility `pkb` can expand [the default Obsidian template values](https://help.obsidian.md/plugins/templates).

- `{{date}}` - today's date (in the format `YYYY-MM-DD`).
- `{{time}}` - the current time (in the format `HH:mm`).
- `{{title}}` - the title of the created note.

It does **not** currently support the MomentJS format tokens you can specify to
customise the `date` and `time` values, as the syntax for this breaks the standard
go template syntax used by `pkb`, so the templates would need to be pre-processed
before they are rendered. This *may* be implemented in a future update, so if it's
something you would like to see please create an issue requesting it.

## Using with multiple vaults or computers

If you use `Obsidian` with multiple vaults, or on multiple computers, you may
wish to have different templates and therefore different `pkb` config files for
each vault or computer.

To support this `pkb` offers the `--vault` flag, which you can use to specify
the name of the config file in your config directory, without needing to pass
the full file path with the `--config` flag. e.g.:

If your `pkb` config directory looked like this:

```bash
pkb/
	./pkb.json
	./personal.json
	./work.json
```

You could run the command:

```bash
pkb --vault work new
```

To use the `work.json` config file to create a new document.

> [!TIP] 
> `pkb.json` is the name of the default config file.
> If no file with this name exists in your config directory and you don't
> specify a `--vault` flag, `pkb` will error.
> If you want to keep your configs aligned to your vaults you will need to 
> either always pass the `--vault` flag, or make one of them your default and
> keep that vault's config in `pkb.json`.
