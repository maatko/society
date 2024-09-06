# secrete
underground social media platform

## tailwind
for tailwind this project recommends using [gowind](https://github.com/maatko/gowind). It is really simple to use and provides direct access to the `tailwindcss` binary without needing nodejs installed on the system.

```sh
go install github.com/maatko/gowind@latest
```

### usage
To get the latest version of tailwind on the system you can run

```sh
gowind update
```

to start watching for styling changes you can run

```sh
make watch
```