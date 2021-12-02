package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	changecase "github.com/ku/go-change-case"
	"golang.design/x/clipboard"
)

type Case struct {
	Fn       func(string) string
	Icon     string
	Subtitle string
}

var wf *aw.Workflow

var cases = []Case{
	{
		Fn:       changecase.Camel,
		Icon:     "icons/camel.png",
		Subtitle: "Convert to a string with the separators denoted by having the next letter capitalised",
	},
	{
		Fn:       changecase.Constant,
		Icon:     "icons/constant.png",
		Subtitle: "Convert to an upper case, underscore separated string",
	},
	{
		Fn:       changecase.Dot,
		Icon:     "icons/dot.png",
		Subtitle: "Convert to a lower case, period separated string",
	},
	{
		Fn:       changecase.Lower,
		Icon:     "icons/lower.png",
		Subtitle: "Convert to a string in lower case",
	},
	{
		Fn:       changecase.LcFirst,
		Icon:     "icons/lcfirst.png",
		Subtitle: "Convert to a string with the first character lower cased",
	},
	{
		Fn:       changecase.No,
		Icon:     "icons/no.png",
		Subtitle: "Convert the string without any casing (lower case, space separated)",
	},
	{
		Fn:       changecase.Param,
		Icon:     "icons/param.png",
		Subtitle: "Convert to a lower case, dash separated string",
	},
	{
		Fn:       changecase.Pascal,
		Icon:     "icons/pascal.png",
		Subtitle: "Convert to a string denoted in the same fashion as camelCase, but with the first letter also capitalised",
	},
	{
		Fn:       changecase.Path,
		Icon:     "icons/path.png",
		Subtitle: "Convert to a lower case, slash separated string",
	},
	{
		Fn:       changecase.Sentence,
		Icon:     "icons/sentence.png",
		Subtitle: "Convert to a lower case, space separated string",
	},
	{
		Fn:       changecase.Snake,
		Icon:     "icons/snake.png",
		Subtitle: "Convert to a lower case, underscore separated string",
	},
	{
		Fn:       changecase.Swap,
		Icon:     "icons/swap.png",
		Subtitle: "Convert to a string with every character case reversed",
	},
	{
		Fn:       changecase.Title,
		Icon:     "icons/title.png",
		Subtitle: "Convert to a space separated string with the first character of every word upper cased",
	},
	{
		Fn:       changecase.Upper,
		Icon:     "icons/upper.png",
		Subtitle: "Convert to a string in upper case",
	},
	{
		Fn:       changecase.UcFirst,
		Icon:     "icons/ucfirst.png",
		Subtitle: "Convert to a string with the first character upper cased",
	},
}

func init() {
	wf = aw.New()
}

func run() {
	wf.Args()
	flag.Parse()
	query := flag.Arg(0)

	if query == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}

	if query != "" {
		wf.Filter(query)
	}

	for c := range cases {
		wf.NewItem(cases[c].Fn(query)).
			Subtitle(cases[c].Subtitle).
			Icon(&aw.Icon{Value: cases[c].Icon}).
			Arg(cases[c].Fn(query)).
			Valid(true)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
