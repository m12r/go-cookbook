# Go Cookbook

This is a growing collection of some [Go][go] recipes.

**Please, do not depend on this project via `go get`!** If you want to use
some code from here, then just use copy and paste.

## Recipes

### Standard Library

- `text/template`
  - [Namespaced Functions in Templates](stdlib/text/template/namespaced_funcs.md)
- `golang.org/x/term`
  - [How To Read Password From Stdin](x/term/password_from_stdin.go)
  
### Useful Tiny Techniques

- `github.com/m12r/go-cookbook/clock`
    - [A thin wrapper around time.Now, to make it testable](clock/README.md)

## License

This project is licensed under the [ISC License][isc]. See [LICENSE](LICENSE)
for details.

---

Copyright © 2021, [Matthias Endler][me]. All rights reserved.


[go]: https://go.dev
[isc]: https://opensource.org/licenses/ISC
[me]: https://m12r.at
