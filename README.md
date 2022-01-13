[![Test](https://github.com/latentd/morph/actions/workflows/test.yml/badge.svg)](https://github.com/latentd/morph/actions?query=workflow%3ATest)
[![Go Reference](https://pkg.go.dev/badge/github.com/latentd/morph.svg)](https://pkg.go.dev/github.com/latentd/morph)

# Morph

Morph is a simple CLI tool to generate files from templates.

## Features 

* Replaces `{{ .Values.xxx }}` in templates
* File and directory names can also be templated
* Remote git repository supported

## Installation

```
$ go install github.com/latentd/morph
```

## Usage

### Using local template files

Any type of file can be used as a template.

Curly brackets `{{ .Values.xxx }}` will be replaced when generating files from templates.

```
$ tree ./templates
./templates
├── sample.txt
└── {{ .Values.key1 }}
    └── {{ .Values.key2 }}.txt

$ cat ./templates/sample.txt
{{ .Values.key3 }}
```

Morph command will generate files in directory specified with `-o` option. (Default: current directory)

Values in templates will be replaced with parameters set by `--set` argument.

```
$ morph -t templates -o dst --set key1=value1,key2=value2,key3=value3

$ tree ./dst
./dst
├── sample.txt
└── value1
    └── value2.txt

$ cat ./dst/sample.txt
value3
```

### Using remote templates

Remote git repository can also be used as templates with `-r` option.

```
$ morph -r https://github.com/latentd/morph -t templates --set key1=value1,key2=value2,key3=value3
$ tree .
.
├── sample.txt
└── value1
    └── value2.txt
```

## License

Apache License Version 2.0
