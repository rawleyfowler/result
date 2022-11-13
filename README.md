# result
Simple implementation of the `Result` monad from OCaml.

## How to use
Create a new result with `result.Error[T, G]` or `result.Ok[T, G]`.

Then you can use a variety of methods to `Unwrap` or `Map` or `Bind` the
result.

## License
This library is available under the MIT license. View the `LICENSE` file at
the root of the directory for more information.