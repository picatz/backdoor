# backdoor

Demo for backdooring golang's `http.DefaultServeMux`

> **Note:** To see an example victim scenario check out [`backdoor-victim`](https://github.com/picatz/backdoor)

## Inspiration

> Any package you import, directly or through other dependencies, has access to `http.DefaultServeMux` and might register routes you don’t expect.
>
> -- [Filippo Valsorda](https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/)

## Attack

To take advantage of this feature, this is a demo of a package that registers a naughty route in `http.DefaultServeMux` located at `/backdoor` which serves up environemnt variables. This could easily exec commands, serve up files or directories and other nefarious business.

I think the main advantage to this attack is that it can be made *real sneaky-like*. However, it's probably no more impressive than other third-party library attacks.

However, the network traffic associated with this kind of backdoor can more easily blend-in to normal process behavior which can be an advantage for attackers not looking to spawn other connections/processes which can be easily caught because it may go outside the lines of "expected" behavior by tuned IPS/IDS systems.

If you have any thoughts on it, feel free to share!
