```text
 ____    _    ____ _   _ _____ _____ _____ _____
| __ )  / \  / ___| | | | ____|_   _|_   _| ____|
|  _ \ / _ \| |  _| | | |  _|   | |   | | |  _|
| |_) / ___ \ |_| | |_| | |___  | |   | | | |___
|____/_/   \_\____|\___/|_____| |_|   |_| |_____|
```
# Baguette

> A personal CLI toolbox written in Go.

Baguette is my personal collection of small command-line tools,
built to make everyday development and system tasks a little easier.

It started as a personal project, but it is open to anyone curious
enough to use it.

## Commands

### `fs`

Find every occurrence of a string in a file and display the
corresponding line.

### `loadbench` 🚧

Work in progress.

A tool to help determine the CPU and memory requirements of a
containerized application under load.

## Installation

Coming soon.

For now, build Baguette from source.

## Usage

### `fs`

```text
baguette fs -f "path/to/file" -s "search expression"
```

### `loadbench`

```text
baguette loadbench config
baguette loadbench load
```

## Documentation

Additional documentation and architecture notes can be found in
[`docs/`](./docs/).

## Development

Baguette is written in Go and uses [Cobra](https://github.com/spf13/cobra)
for its CLI.

## Contributing

Baguette is primarily a personal project, but contributions,
suggestions and ideas are welcome.

If you find a bug or have an idea for a useful tool, feel free to
open an issue or a pull request.

## License
Baguette is licensed under the MIT License.
See [`LICENSE`](./LICENSE) for details.