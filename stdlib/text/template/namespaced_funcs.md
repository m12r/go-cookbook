# Namespaced Functions in Templates

I was curious on how to create template functions with namespaces for
`text/template` and `html/template` in [Go][go]'s standard library. I knew,
that this must be possible, because namespaces to group template functions are
used in [Hugo][hugo]. So first I searched the web, because I did not want to
look into the humongous, but fantastic code base of Hugo. Neither the Web, nor
the documentation of `func (*template.Template) Funcs(funcMap FuncMap) *Template`
in `text/template` and `html/template` showed a useful example on how to
accomplish this. So I had a look at Hugo's code base anyways.

Have a look at the example code [namespaced_funcs.go](namespaced_funcs.go) and
tests [namespaced_funcs_test.go](namespaced_funcs_test.go).

Have fun and enjoy coding.

See you next time :)

---

Copyright © 2021, [Matthias Endler][me]. All rights reserved.


[go]: https://go.dev
[hugo]: https://gohugo.io
[me]: https://m12r.at
