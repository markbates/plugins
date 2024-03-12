package plugcmd

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/markbates/plugins"
)

// Print will try and print a helpful Usage printing
// of the plugin and any plugins that are provided.
//
//	$ foobar
//	---------
//
//	Usage of foobar:
//	-h, --help   print this help
//
//	Available Commands:
//		foobar info     Print diagnostic information (useful for debugging)
//		foobar version  Print the version information
func Print(w io.Writer, main plugins.Plugin) error {

	header := strings.TrimSpace(cmdName(main))
	header = fmt.Sprintf("$ %s", header)
	fmt.Fprintln(w, header)
	for i := 0; i < len(header); i++ {
		fmt.Fprint(w, "-")
	}
	fmt.Fprintf(w, "\n%s\n", typeName(main))

	if u, ok := main.(UsagePrinter); ok {
		fmt.Fprintln(w)
		if err := u.PrintUsage(w); err != nil {
			return err
		}
		fmt.Fprintln(w)
	}

	if a, ok := main.(Aliaser); ok {
		aliases := a.CmdAliases()
		if len(aliases) != 0 {
			const al = "\nAliases:\n"
			fmt.Fprint(w, al)
			fmt.Fprintln(w, strings.Join(aliases, ", "))
		}
	}

	if err := printFlags(w, main); err != nil {
		return err
	}

	if err := printCommands(w, main); err != nil {
		return err
	}

	if err := printPlugins(w, main); err != nil {
		return err
	}

	return nil
}

func printFlags(w io.Writer, p plugins.Plugin) error {
	if u, ok := p.(FlagPrinter); ok {
		// fmt.Fprintln(w)
		fmt.Fprintln(w, "\nFlags:")
		return u.PrintFlags(w)
	}

	if u, ok := p.(Flagger); ok {
		flags, err := u.Flags()
		if err != nil {
			return err
		}

		ow := flags.Output()
		flags.SetOutput(w)
		fmt.Fprintln(w)
		flags.Usage()
		flags.SetOutput(ow)
	}

	return nil
}

func printPlugins(w io.Writer, main plugins.Plugin) error {
	mm := map[string]plugins.Plugin{}

	usingPlugins(main, mm)

	if len(mm) == 0 {
		return nil
	}

	plugs := make(plugins.Plugins, 0, len(mm))
	for _, p := range mm {
		plugs = append(plugs, p)
	}

	sort.Slice(plugs, func(i, j int) bool {
		return plugs[i].PluginName() < plugs[j].PluginName()
	})

	fmt.Fprintln(w, "\nUsing Plugins:")
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "\t%s\t%s\t%s\n", "Name", "Description", "Type")
	fmt.Fprintf(tw, "\t%s\t%s\t%s\n", "----", "-----------", "----")
	for _, p := range plugs {
		fmt.Fprintf(tw, "\t%s\t%s\t%s\n", p.PluginName(), desc(p), typeName(p))
	}

	tw.Flush()
	return nil
}

func typeName(p plugins.Plugin) string {
	rv := reflect.Indirect(reflect.ValueOf(p))
	rt := reflect.TypeOf(rv.Interface())

	bb := &bytes.Buffer{}

	t := fmt.Sprintf("%T", p)
	if strings.HasPrefix(t, "*") {
		fmt.Fprint(bb, "*")
		t = strings.TrimPrefix(t, "*")
	}
	fmt.Fprintf(bb, "%s/", path.Dir(rt.PkgPath()))
	fmt.Fprint(bb, t)

	return bb.String()
}

func usingPlugins(plug plugins.Plugin, mm map[string]plugins.Plugin) error {
	if mm == nil {
		return fmt.Errorf("mm cannot be nil")
	}

	wp, ok := plug.(plugins.Scoper)
	if !ok {
		return nil
	}

	plugs := wp.ScopedPlugins()

	for _, p := range plugs {
		mm[p.PluginName()] = p
	}

	return nil
}

func printCommands(w io.Writer, main plugins.Plugin) error {
	sc, ok := main.(SubCommander)
	if !ok {
		return nil
	}

	plugs := sc.SubCommands()
	if len(plugs) == 0 {
		return nil
	}

	sort.Slice(plugs, func(i, j int) bool {
		return plugs[i].PluginName() < plugs[j].PluginName()
	})

	const ac = "\nAvailable Commands:\n"
	fmt.Fprint(w, ac)

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "\t%s\t%s\n", "Command", "Description")
	fmt.Fprintf(tw, "\t%s\t%s\n", "-------", "-----------")
	for _, c := range plugs {
		line := fmt.Sprintf("\t%s", cmdName(c))
		if d := desc(c); d != "" {
			line = fmt.Sprintf("%s\t%s", line, d)
		}
		fmt.Fprintln(tw, line)
	}

	tw.Flush()
	return nil
}

func desc(p plugins.Plugin) string {
	if d, ok := p.(Describer); ok {
		return d.Description()
	}
	return ""
}

func cmdName(p plugins.Plugin) string {
	if d, ok := p.(Namer); ok {
		return d.CmdName()
	}
	return path.Base(p.PluginName())
}
