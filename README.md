# Primordial Particle System

An implementation of the Primordial Particle System (PPS) described in:

Schmickl, T., Stefanec, M. & Crailsheim, K.
[How a life-like system emerges from a simplistic particle motion law.](http://www.nature.com/articles/srep37969)
Sci Rep 6, 37969 (2016).
https://doi.org/10.1038/srep37969

Explanation video by the authors of the article (not this implementation):

_How life emerges from a simple particle motion law: Introducing the Primordial Particle System_

[![Video: How life emerges from a simple particle motion law: Introducing the Primordial Particle System](https://img.youtube.com/vi/makaJpLvbow/0.jpg)](https://www.youtube.com/watch?v=makaJpLvbow)

## Installation

```
$ go install github.com/fzipp/pps/cmd/pps-simulate@latest
```

## Usage

```
$ pps-simulate
Visit http://localhost:8080 in a web browser
```

It uses the following parameter set (as described in the article) by default:

PPS =〈r=5, α=180°, β=17°, v=0.67〉

The parameters and the particle density (DPE) can be modified via command line flags:

```
$ pps-simulate -r=4 -alpha=180 -beta=17 -v=0.67 -dpe=0.08
```

The port is configurable as well:

```
$ pps-simulate -http=:6060
```

Get help about the flags:

```
$ pps-simulate -help
```

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
