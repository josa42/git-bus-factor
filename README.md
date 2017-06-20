# git-bus-factor

A golang implementation of [zats/github_bus_factor](https://github.com/zats/github_bus_factor)

```
$ git bus-factor atom/atom
🍴  6545 forks.
🔭  2105 watchers.
🌟  37723 stars.
📆  Created about over 5 years ago; last push about 5 hours ago.
🍻  200 PRs: 100 closed; 100 open; 50.00% are closed.
🛠️  Deletions to additions ratio: 96.65% (-2737092/2831991).
📦  100 releases; latest release "1.19.0-beta1": about 11 hours ago.
🚌  Bus factor: 50% (2 impactful contributors out of 100).
```

**Work in progress**

## TODOs

- [ ] Support none github repos (bitbucket, git fallback)
  - [ ] bitbucket
  - [ ] git fallback
- [ ] Use tags as releases
- [ ] Add some colors and `--no-colors` option
- [ ] Add `--no-emojis` option
- [ ] Add screenshot
- [ ] Add `--json` option
- [ ] Add options to skip steps

## Installation

**Homebrew (macOS)**

```
brew tap josa42/homebrew-git-tools
brew install git-bus-factor
```

**Other**

```
go get github.com/josa42/git-bus-factor
```

## Usage

```
Usage:
  git-bus-factor

Options:
  -h --help          Show this screen.
  --version          Show version.
```

## License

MIT (See [license.md](LICENSE.md))
