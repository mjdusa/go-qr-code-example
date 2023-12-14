# go-qr-code-example

[![Go Version][go_version_img]][go_dev_url]
[![Go Report Card][go_report_img]][go_report_url]
[![License][repo_license_img]][repo_license_url]

## Description
go-qr-code-example is a tool to sync all forks in the owner's account from upstream.

## Contributing
Please see our [Contributing](./CONTRIBUTING.md) for how to contribute to the project.

## Setting up for development...
git clone https://github.com/mjdusa/go-qr-code-example.git

## Pre-commit Hooks
When you clone this repository to your workstation, make sure to install the [pre-commit](https://pre-commit.com/) hooks. [GitHub Repo](https://github.com/pre-commit/pre-commit)

### Installing tools
```
brew install pre-commit
```

### Check installed versions.
```
pre-commit --version
pre-commit 3.3.2
```

### Update configured pre-commit plugins.  Updates repo versions in .pre-commit-config.yaml to the latest.
```
pre-commit autoupdate
```

### Install pre-commit into the local git.
```
pre-commit install --install-hooks
```

### Run pre-commit checks manually.
```
pre-commit run --all-files
```

## Running...
```
make release
...
./dist/go-qr-code-example
```


<!-- Go -->

[go_download_url]: https://golang.org/dl/
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none
[go_report_url]: https://goreportcard.com/report/github.com/mjdusa/go-qr-code-example

<!-- Repository -->

[repo_url]: https://github.com/mjdusa/go-qr-code-example
[repo_license_url]: https://github.com/mjdusa/go-qr-code-example/blob/main/LICENSE
[repo_license_img]: http://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge&logo=none

<!-- Project -->

<!-- Author -->

[author]: https://github.com/mjdusa
