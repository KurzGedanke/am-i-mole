# Am-I-Mole

> Disclaimer: I'm an not endorse with mullvad.net in any way! This tool ist just for my personal lazyness and the love for the command line!

A command line tool to check if you are connect to mullvad.net

## Usage

```bash
$ am-i-mole
Hi! Getting data from am.i.mullvad.net!
You ARE connected to Mullvad!
IP:  185.65.135.171
Exit IP Hostname:  se21
Server Type:  OpenVPN
Country: Sweden
City <nil>
Your IP is NOT blacklisted!
Organisation:  Amagicom AB
```

```bash
$ am-i-mole -h
  -black
        Prints if blacklisted.
  -c
        Prints if connected to mullvad.
  -ct
        Prints the country connected in.
  -cty
        Prints the city connected in.
  -ip
        Prints your current IP.
  -o
        Prints the organization connecte to.
exit status 2
```

You can define different flags.

```bash
$ am-i-mole -black -c -ct -cty -ip -o
185.65.135.168
true
Sweden
<nil>
false
Amagicom AB
```

The values are printed not in order of the command line arguments.
The oder is:

1. IP: string
2. Connected: bool
3. Country: string
4. City: string
5. Blacklist: bool
6. Organization: string

## Installation

### Manual Instalaltion

Download the binary from the [Github Release Page](https://github.com/KurzGedanke/am-i-mole/releases) and put it manually in `usr/local/bin`. Then you can access it from the command line.

## License

MIT License

Copyright (c) 2018 Thore Jahn

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
