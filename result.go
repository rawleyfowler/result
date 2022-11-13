package result

type variant bool

const (
	v_error variant = false
	v_ok variant = true
)

type ok[T any] struct {
	value T
}

type err[T any] struct {
	error T
}

type Result[T any, G any] struct {
	ok ok[T]
	err err[G]
	v variant
}

func Ok[T, G any](v T) *Result[T, G] {
	return &Result[T, G]{
		ok: ok[T]{ value: v },
		v: v_ok,
	}
}

func Error[T, G any](e G) *Result[T, G] {
	return &Result[T, G]{
		err: err[G]{ error: e },
		v: v_error,
	}
}

func (r *Result[_, _]) IsOk() bool {
	return r.v == v_ok
}

func (r *Result[_, _]) IsError() bool {
	return r.v == v_error
}

/**
  Forcibly unwraps the result, this can cause a panic!
  Not safe! Please checkout the UnwrapOr functions!
 */
func (r *Result[T, _]) Unwrap() T {
	if r.IsOk() {
		return r.ok.value
	} else {
		panic("Failed to unwrap. Please see result.UnwrapOr*")
	}
}

func (r *Result[T, _]) UnwrapOr(or T) T {
	if r.IsOk() {
		return r.Unwrap()
	} else {
		return or
	}
}

func (r *Result[T, _]) UnwrapOrDefault() T {
	if r.IsOk() {
		return r.Unwrap()
	} else {
		return *new(T)
	}
}

func (r *Result[T, G]) UnwrapOrElse(fn func(G) T) T {
	if r.IsOk() {
		return r.Unwrap()
	} else {
		return fn(r.err.error)
	}
}

func Map[T, G, N any](r *Result[T, G], fn func(T) N) *Result[N, G]  {
	if r.IsOk() {
		return Ok[N, G](fn(r.Unwrap()))
	} else {
		return Error[N, G](r.err.error)
	}
}

func Bind[T, G, N any](r *Result[T, G], fn func(T) *Result[N, G]) *Result[N, G] {
	if r.IsOk() {
		return fn(r.Unwrap())
	} else {
		return Error[N, G](r.err.error)
	}
}
