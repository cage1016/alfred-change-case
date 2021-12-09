package main

import (
	"fmt"
	"log"
	"os"
	"unicode"

	aw "github.com/deanishe/awgo"
	changecase "github.com/ku/go-change-case"
	"golang.design/x/clipboard"

	"github.com/cage1016/alfred-change-case/handler"
)

var wf *aw.Workflow

var M = map[string]struct {
	Fn       func(string) string
	Icon     string
	Subtitle string
}{
	"camel":    {changecase.Camel, "A3BC3C97-6CC7-4E29-A317-119EC3DAAB8C.png", "Convert to a string with the separators denoted by having the next letter capitalized"},
	"constant": {changecase.Constant, "F9586C66-DF26-4A16-A635-CDC9C599E40A.png", "Convert to an upper case, underscore separated string"},
	"dot":      {changecase.Dot, "05DE2A2D-883F-4B22-B368-1638E60F3FE7.png", "Convert to a lower case, period separated string"},
	"lower":    {changecase.Lower, "C5E81DA7-A7B6-48F8-8DA0-49DA0197BB81.png", "Convert to a string in lower case"},
	"lcfirst":  {changecase.LcFirst, "76DBBBB2-DACC-4B35-9A93-C6073B9C84BA.png", "Convert to a string with the first character lower cased"},
	"no":       {changecase.No, "6A682A24-DBDA-4A68-9E3B-4E07FA3F4C24.png", "Convert the string without any casing (lower case, space separated)"},
	"param":    {changecase.Param, "35CB5FA6-BF41-4C43-8156-B18A6DBE0248.png", "Convert to a lower case, dash separated string"},
	"pascal":   {changecase.Pascal, "F869FC45-70F6-41C6-B046-060532172660.png", "Convert to a string denoted in the same fashion as camelCase, but with the first letter also capitalized"},
	"path":     {changecase.Path, "96E90A06-7D14-4E17-822E-F8956AF8D6A6.png", "Convert to a lower case, slash separated string"},
	"sentence": {changecase.Sentence, "25C0C74E-0E5B-4A7D-AEB8-0083E0D595B2.png", "Convert to a lower case, space separated string"},
	"snake":    {changecase.Snake, "9882ADC5-3C84-410B-AEC7-0F906C9EA223.png", "Convert to a lower case, underscore separated string"},
	"swap":     {changecase.Swap, "A608D078-BEC4-4FDF-B68C-CBEE6E6E6224.png", "Convert to a string with every character case reversed"},
	"title":    {changecase.Title, "E4825A44-DC65-4665-A4C8-3D70969B090F.png", "Convert to a space separated string with the first character of every word upper cased"},
	"upper":    {changecase.Upper, "22AEA759-6C80-44E7-B346-CD03F04A4437.png", "Convert to a string in upper case"},
	"ucfirst":  {changecase.UcFirst, "C4FAD516-20DF-4BEF-A034-E4335ABD6896.png", "Convert to a string with the first character upper cased"},
	"hashtag":  {handler.HashTag, "1DAD602D-927A-49A8-B430-D7DF8CF46921.png", "Convert to a string, space separated string with hashtag symbols"},
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
		"hashtag",
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
	"hashtag":  {"hashtag"},
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
	if query == "" {
		query = string(clipboard.Read(clipboard.FmtText))
	}

	for _, h := range hs {
		Fn(wf, M[h].Fn, M[h].Icon, M[h].Subtitle, query)
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
