# Preface

`godir` is a small util that walks through directories and finds 
matches based on arguments. If there is only one match, it prints the
match to `stdout`, otherwise prints a list of matched directories
to `stderr` and lets the user select one.

It is useful for quick navigation when used in combination with the 
shell.

Arguments are either absolute paths or regular expressions.
If an absolute path is encountered, directory traversal starts
from that directory, otherwise it is treated as a regular expression.
If no absolute path is provided, start directory for search is the 
current directory. If an argument is not an absolute path and 
contains path separator (e.g. `/`), it is split to multiple
arguments.

# Install

Install the binary:

    go install github.com/korneil/godir
    
Add functions to your shell initialization, e.g. to `.bashrc`/`.profile`/`.zshrc` 
(see **Shell examples** for example functions).

# Example

Given the following folder structure in `/opt/local`:

    ├── etc
    │   ├── mc
    │   ├── openssl
    │   └── xml
    ├── include
    │   ├── curl
    │   ├── leveldb
    ├── lib
    │   ├── texinfo
    ├── libexec
    │   ├── mc
    └── var
        ├── cache
        └── state

`godir /opt/local e url` outputs on `/opt/local/include/curl` on `stdout`.

Or `godir /opt/local e` (or `cd /opt/local; godir e`) gives the following output on `stderr`:

    0: /opt/local/etc
    1: /opt/local/include
    2: /opt/local/libexec    
    
After selecting `1`, output on `stdout`:
    
    /opt/local/include

# Shell example

Given the following shell functions:

    c(){ cd `godir $@` }

    cdp(){ cd `godir ~/Projects $@` }   

    cdol(){ cd `godir /opt/local $@` }

`cdol e` in the shell will cd to `/opt/local/include`.
