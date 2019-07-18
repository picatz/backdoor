# backdoor

Demo for backdooring golang's `http.DefaultServeMux`

## Inspiration

> Any package you import, directly or through other dependencies, has access to `http.DefaultServeMux` and might register routes you don’t expect.
>
> -- [Filippo Valsorda](https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/)

