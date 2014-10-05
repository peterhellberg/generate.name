# [generate.name](http://generate.name)

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

### Key

A key used to lock down editing of the generator.
(Optional, but a generator without a key can be edited by anyone)

You need to pass in the key as a query parameter in order to show the edit button:
`http://generate.name/<generator-slug>?key=<your-key>`

### Field1 - Field6

The fields used by the template.
