# Preface

`godir` is a small util that walks through directories and finds
matches based on arguments. If there is only one match, it prints the
match to `stdout`, otherwise prints a list of matched directories
to `stderr` sorted by the and lets the user select one.

It is useful for quick navigation when used in combination with the
shell.

Arguments are either absolute paths or string expressions.
If an absolute path is encountered, directory traversal starts
from that directory, otherwise it is treated as a regular string,
and a fuzzy match is performed against subdirectories.
If no absolute path is provided, start directory for search is the
current directory. If an argument is not an absolute path and
contains path separator (e.g. `/`), it is split to multiple
arguments.

# Install

Install the binary:

    go get -u github.com/korneil/godir
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

    # cd to dir in ~/Projects
    cdp(){ cd `godir ~/Projects $@` }

    # active python environment
    pe(){ source `godir ~/.venv $@`/bin/activate; }

    # create python environment
    cpe(){ virtualenv -p python ~/.venv/$@; }

    cdol(){ cd `godir /opt/local $@` }

`cdol e` in the shell will cd to `/opt/local/include`.

`cpe new_python_env; pe newp` will create and activate a new
python virtual env. After creation (`cpe`), `pe n` will activate
it.
