replace [![Build Status](https://travis-ci.org/voutasaurus/replace.svg?branch=master)](https://travis-ci.org/voutasaurus/replace)
=======

Replace a string with another string in a file. Write to another file or modify the existing file.

Installation
============

```
$ go get github.com/voutasaurus/replace
```

Usage
=====

Write to a new file:
```
$ replace bad good dirty.txt clean.txt
```

Overwrite existing file:
```
$ replace lol ha story.txt
```

Contributions
=============

Issues and PRs welcome.