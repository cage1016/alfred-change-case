package main

import (
	"fmt"
	"log"
	"os"
	"unicode"

	aw "github.com/deanishe/awgo"
	changecase "github.com/ku/go-change-case"
	"golang.design/x/clipboard"
)

var wf *aw.Workflow

var M = map[string]struct {
	Fn       func(string) string
	Icon     string
	Subtitle string
}{
	"camel":    {changecase.Camel, "icons/camel.png", "Convert to a string with the separators denoted by having the next letter capitalized"},
	"constant": {changecase.Constant, "icons/constant.png", "Convert to an upper case, underscore separated string"},
	"dot":      {changecase.Dot, "icons/dot.png", "Convert to a lower case, period separated string"},
	"lower":    {changecase.Lower, "icons/lower.png", "Convert to a string in lower case"},
	"lcfirst":  {changecase.LcFirst, "icons/lcfirst.png", "Convert to a string with the first character lower cased"},
	"no":       {changecase.No, "icons/no.png", "Convert the string without any casing (lower case, space separated)"},
	"param":    {changecase.Param, "icons/param.png", "Convert to a lower case, dash separated string"},
	"pascal":   {changecase.Pascal, "icons/pascal.png", "Convert to a string denoted in the same fashion as camelCase, but with the first letter also capitalized"},
	"path":     {changecase.Path, "icons/path.png", "Convert to a lower case, slash separated string"},
	"sentence": {changecase.Sentence, "icons/sentence.png", "Convert to a lower case, space separated string"},
	"snake":    {changecase.Snake, "icons/snake.png", "Convert to a lower case, underscore separated string"},
	"swap":     {changecase.Swap, "icons/swap.png", "Convert to a string with every character case reversed"},
	"title":    {changecase.Title, "icons/title.png", "Convert to a space separated string with the first character of every word upper cased"},
	"upper":    {changecase.Upper, "icons/upper.png", "Convert to a string in upper case"},
	"ucfirst":  {changecase.UcFirst, "icons/ucfirst.png", "Convert to a string with the first character upper cased"},
}

var handlers = map[string][]string{
	"commands": {
		"camel",
		"constant",
		"dot",
		"lower",
		"lcfirst",
		"no",
		"param",
		"pascal",
		"path",
		"sentence",
		"snake",
		"swap",
		"title",
		"upper",
		"ucfirst",
	},
	"camel":    {"camel"},
	"constant": {"constant"},
	"dot":      {"dot"},
	"lower":    {"lower"},
	"lcfirst":  {"lcfirst"},
	"no":       {"no"},
	"param":    {"param"},
	"pascal":   {"pascal"},
	"path":     {"path"},
	"sentence": {"sentence"},
	"snake":    {"snake"},
	"swap":     {"swap"},
	"title":    {"title"},
	"upper":    {"upper"},
	"ucfirst":  {"ucfirst"},
}

func Fn(wf *aw.Workflow, h func(string) string, icon, subtitle, query string) {
	wf.NewItem(h(query)).Subtitle(subtitle).Icon(&aw.Icon{Value: icon}).Arg(h(query)).Valid(true)
}

func init() {
	wf = aw.New()
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		exitWithError("please provide some input 👀")
	}

	hs, found := handlers[args[0]]
	if !found {
		exitWithError("command not recognized 👀")
	}

	query := args[1]
	if query != "" {
		for _, h := range hs {
			Fn(wf, M[h].Fn, M[h].Icon, M[h].Subtitle, args[1])
		}
	} else {
		for _, h := range hs {
			Fn(wf, M[h].Fn, M[h].Icon, M[h].Subtitle, string(clipboard.Read(clipboard.FmtText)))
		}
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}

func exitWithError(msg string) {
	fmt.Print(capitalize(msg))
	log.Print(msg)
	os.Exit(1)
}

func capitalize(msg string) string {
	return string(unicode.ToUpper(rune(msg[0]))) + msg[1:]
}
