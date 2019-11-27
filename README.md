# Whisky Tango

Whisky Tango is a cross platform tool to install your applications using a config file. It allows your development teams to easily install similar environments and tooling on their laptops to get easily started. Only thing they need to do is fillout the config file and simply copy that to any environment they want to install.

## Contribute

To build the binary simply run `make`.

```bash
$ make
Download go.mod dependencies
Building binaries
```

To run Wisky Tango you can invoke the binary.

```bash
$ bin/wt
Using config file: /Users/marco/code/priv/whisky-tango/.wt.yml
Whisky Tango, automates your machine setup

Usage:
  wt [flags]
  wt [command]

Available Commands:
  config      Get current configuration as json output
  help        Help about any command
  install     installs your apps
  version     shows version information

Flags:
      --config string   config file (default is $HOME/.wt.yml)
  -h, --help            help for wt

Use "wt [command] --help" for more information about a command.
```
