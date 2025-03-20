<div align="center">

<img src="https://github.com/elcomware/goravel-package-skeleton/blob/main/skeleton.png" width="300" alt="Logo">

[![Doc](https://pkg.go.dev/badge/github.com/:vendor_slug/:package_slug)](https://pkg.go.dev/github.com/:vendor_slug/:package_slug)
[![Go](https://img.shields.io/github/go-mod/go-version/:vendor_slug/:package_slug)](https://go.dev/)
[![Release](https://img.shields.io/github/release/:vendor_slug/:package_slug.svg)](https://github.com/:vendor_slug/:package_slug/releases)
[![Test](https://github.com/:vendor_slug/:package_slug/actions/workflows/test.yml/badge.svg)](https://github.com/:vendor_slug/:package_slug/actions)
[![Report Card](https://goreportcard.com/badge/github.com/:vendor_slug/:package_slug)](https://goreportcard.com/report/github.com/:vendor_slug/:package_slug)
[![Codecov](https://codecov.io/gh/:vendor_slug/:package_slug/branch/master/graph/badge.svg)](https://codecov.io/gh/:vendor_slug/:package_slug)
![License](https://img.shields.io/github/license/:vendor_slug/:package_slug)

</div>

English | [中文](./README_zh.md)

## About :package_slug

:package_slug is a Goravel application framework package tool with complete functions and good scalability. As a starting scaffolding to help
Gopher quickly build their own Goravel packages.

The framework style is consistent with [Laravel-spatie-laravel-package-tools](https://github.com/spatie/laravel-package-tools), let Php developer don't need to learn a
new framework, but also happy to play around Golang! In tribute to Laravel!

Welcome to star, PR and issues！

<!--delete-->

## Getting started

---
This repo can be used to scaffold a Goravel package. Follow these steps to get started:

1. Press the "Use this template" button at the top of this repo to create a new repo with the contents of this skeleton.
2. Run "go run main/setup.go" to run a script that will replace all placeholders throughout all the files.
3. Have fun creating your package.
4. If you need help creating a package, consider picking up our <a href="https://www.goravel.dev/digging-deeper/package-development.html">Goralvel Package Deveopment</a> Documentation.
---
<!--/delete-->
This is where your description should go. Limit it to a paragraph or two. Consider adding a small example.

## Support us

[<img src="https://github-ads.s3.eu-central-1.amazonaws.com/:package_name.jpg?t=1" width="419px" />](https://spatie.be/github-ad-click/:package_name)

We invest a lot of resources into creating [best in class open source packages](https://github/elcomware). You can support us by [buying one of our paid products](https://#).

## Installation / Add package

1. You can install the package via go get -u:

```bash
# Install the latest version of the goravel installer:

go get -u github.com/:vendor_slug/:package_slug
```

2. You can publish the config file with:

```bash
go run . artisan vendor:publish --package=github.com/:vendor_slug/:package_slug

```

You can publish and run the migrations with:

```bash
go run . artisan vendor:publish --tag=":package_slug-migrations"
go run . artisan migrate
```

You can publish the config file with:

```bash
go run . artisan vendor:publish --tag=":package_slug-config"

```

This is the contents of the published config file:

```
func init() {
    config := facades.Config()
    config.Add("go-secure", map[string]any{
        "name": "go-secure",
    })
}
```

Optionally, you can publish the views using

```bash
go run . artisan vendor:publish --tag=":package_slug-views"
```

## Usage

```bash

// main.go
import examplefacades "github.com/goravel/example-package/facades"

fmt.Println(examplefacades.Hello().World())
The console will print Welcome To Goravel Package.
```

```bash
Register service provider
// config/app.go
import examplepackage "github.com/goravel/example-package"

"providers": []foundation.ServiceProvider{
    ...
    &examplepackage.ServiceProvider{},
}
```

## Testing

```bash
composer test

```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

- [:author_name](https://github.com/:author_username)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
