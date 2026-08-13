# jaft - just a f***ing timer

Annoyed by lack of freaking fast, multi-platform and accessible timers you can start in seconds? Use `jaft`!
After countdown you will hear a sound, and you can cancel the sound with `CTRL+C` like any other terminal utility.
# Installing

The only requirement is `go` installed as well as adding go bin path to your PATH, then run:
```bash
go install github.com/gornius/jaft`
```

# Usage
```bash
jaft DURATION
```
Duration is go `time` compatible string.
## Example usages
```bash
jaft 1m
```
```bash
jaft 1m45s
```
```bash
jaft 0.2h
```
