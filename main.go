package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-bus-factor/busFactor"
	stringutils "github.com/josa42/go-stringutils"
)

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-bus-factor

		Options:
		  -h --help          Show this screen.
		  --version          Show version.

		Legend:
		  🍴   Forks. Might mean people planning are fixing bugs or adding features.
		  🔭   Watchers. Shows number of people interested in project changes.
		  🌟   Stars. Might mean it is a good project or that it was featured in a mailing list. Some people use 🌟  as a "Like".
		  🗓   Age. Mature projects might mean battle tested project. Recent pushes might mean project is actively maintained.
		  🍻   Pull Requests. Community contributions to the project. Many closed PRs usually is a good sign, while no PRs usual is bad.
		  🛠   Refactoring. Balance between added and deleted code. Crude value not including semantic understanding of the code.
		  📦   Releases. Might mean disciplined maintainer. Certain dependency managers rely on releases to be present.
		  🚌   Bus factor. Chances of the project to become abandoned once current collaborators stop updating it. The higher - the worse.
  `)

	// // arguments, _ :=
	docopt.Parse(usage, nil, true, "Git Bus Factor 0.0.0", false)

	busFactor.Print()
}
