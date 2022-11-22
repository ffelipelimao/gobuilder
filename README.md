# Generate builders for Go

[![GoDoc](https://pkg.go.dev/badge/github.com/ffelipelimao/gobuilder)](https://pkg.go.dev/github.com/ffelipelimao/gobuilder)

gobuilder is a library for create files,
at this moment we are creating only from the design pattern Builder.

[Builder Design Pattern](https://refactoring.guru/design-patterns/builder)



## Installation
```
$ go install github.com/ffelipelimao/gobuilder@latest
```

## Usage


##### Create a builder file
```bash
gobuilder gen -n=Test -f=Foo-string,Bar-string
```

##### Command Line syntax flag

- **-n** : Create a name to your mock builder struct ex: -n=Test 

- **-f** : Create a name the fields to your struct
     the fields are separated by "," for you to define a type for your field, you separate it by "-"
     ex: -f=Foo-string,Bar-string
    in this case Foo is the field and string is the type of this field

- **-d** : Create the destination folder, he need start with ./ 
    

## Examples

```bash
gobuilder gen -n=Test -f=Foo-string,Bar-string
```

With Pointer: use letter p

```bash
gobuilder gen -n=Test -f=Foo-pstring,Bar-string
```

With Array: use letter a

```bash
gobuilder gen -n=Test -f=Foo-astring,Bar-string
```

## Help commands

```bash
gobuilder -h
```
or

```bash
gobuilder gen -h
```

## Issues and Contributing

If you find an issue with this library, please report an issue. If you'd
like, we welcome any contributions. Fork this library and submit a pull
request.

Versions used with gobuilder must follow [SemVer](http://semver.org/).

## License
[MIT](https://choosealicense.com/licenses/mit/)