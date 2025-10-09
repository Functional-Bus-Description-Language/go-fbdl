[![Tests](https://github.com/Functional-Bus-Description-Language/go-fbdl/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/Functional-Bus-Description-Language/go-fbdl/actions?query=master)

# go-fbdl

Functional Bus Description Language compiler front-end written in Go.

## Installation

### go
```
go install github.com/Functional-Bus-Description-Language/go-fbdl/cmd/fbdl@latest
```

Go installation installs to go configured path.

### Manual

```
git clone https://github.com/Functional-Bus-Description-Language/go-fbdl.git
make
make install
```

Manual installation installs to `/usr/local/bin`.

## Citation

If you find fbdl useful, and write any academic publication on a project utilizing fbdl please consider citing [How Shifting Focus from Register to Data Functionality Can Enhance Register and Bus Management](https://www.mdpi.com/2079-9292/13/4/719).

```
@Article{electronics13040719,
  AUTHOR = {Kruszewski, Michał},
  TITLE = {How Shifting Focus from Register to Data Functionality Can Enhance Register and Bus Management},
  JOURNAL = {Electronics},
  VOLUME = {13},
  YEAR = {2024},
  NUMBER = {4},
  ARTICLE-NUMBER = {719},
  URL = {https://www.mdpi.com/2079-9292/13/4/719},
  ISSN = {2079-9292},
  DOI = {10.3390/electronics13040719}
}
```
