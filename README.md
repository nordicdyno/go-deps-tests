# go-deps-tests

consumer code for go deps cases tests

## More info

it only uses public API of `github.com/nordicdyno/go-deps-A1`.

- with `go-deps-A1@v0.0.1` dependency binary size is ~2Mb
- with `go-deps-A1@v0.1.0` dependency binary size is still ~2Mb, but go.sum now contains internal API's dependencies hash sums (`gorm.io/gorm`, `mattn/go-sqlite3`, `jinzhu/inflection`, `jinzhu/now`)
