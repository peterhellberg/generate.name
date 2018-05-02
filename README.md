# [generate.name](http://generate.name)

[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/generate.name)](https://goreportcard.com/report/github.com/peterhellberg/generate.name)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/generate.name/server)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/generate.name#license-mit)

A Generator-Generator

## Fields on the edit page

Each generator has `Name`, `Template`, `Key` and `Field1` to `Field6` input fields.

### Name

The name field is used to generate the slug, the slug is not changed if you update the name.

### Template

The template field should contain a Golang [text/template](http://golang.org/pkg/text/template/).
Example: `{{.Field1}}, {{.Field2}}`

You can also use a few predefined tags in both the template and each individual field:

 - *[BR]* - Inserts a `<br>`
 - *[D]* - Random number from 1-9
 - *[ROMAN10]* - Random roman numeral from 1-10
 - *[END]* - End the generated string
 - *[GENERATE slug]* - Generates string based on generator identified by `slug`

### Key

A key used to lock down editing of the generator.
(Optional, but a generator without a key can be edited by anyone)

You need to pass in the key as a query parameter in order to show the edit button:
`http://generate.name/<generator-slug>?key=<your-key>`

### Field1 - Field6

The fields used by the template.

## License (MIT)

Copyright (c) 2015-2018 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="230" height="230">

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
