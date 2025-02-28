# :package_description

[![Latest Version on Packagist](https://img.shields.io/packagist/v/:vendor_slug/:package_slug.svg?style=flat-square)](https://packagist.org/packages/:vendor_slug/:package_slug)
[![GitHub Tests Action Status](https://img.shields.io/github/actions/workflow/status/:vendor_slug/:package_slug/run-tests.yml?branch=main&label=tests&style=flat-square)](https://github.com/:vendor_slug/:package_slug/actions?query=workflow%3Arun-tests+branch%3Amain)
[![GitHub Code Style Action Status](https://img.shields.io/github/actions/workflow/status/:vendor_slug/:package_slug/fix-php-code-style-issues.yml?branch=main&label=code%20style&style=flat-square)](https://github.com/:vendor_slug/:package_slug/actions?query=workflow%3A"Fix+PHP+code+style+issues"+branch%3Amain)
[![Total Downloads](https://img.shields.io/packagist/dt/:vendor_slug/:package_slug.svg?style=flat-square)](https://packagist.org/packages/:vendor_slug/:package_slug)
<!--delete-->
---
This repo can be used to scaffold a Goravel package. Follow these steps to get started:

1. Press the "Use this template" button at the top of this repo to create a new repo with the contents of this skeleton.
2. Run "go run setup.go" to run a script that will replace all placeholders throughout all the files.
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

```bash
return [
];
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
