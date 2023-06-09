package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Flag defines a Disco CLI flag which contains a list of
// subcommands the flag belongs to in `DefinedOn` field.
type Flag struct {
	Name      string
	Shorthand string
	Usage     string
	// Pointer to where the value of the flag will be stored
	Value interface{}
	// Default value of the flag. Needs to be set and determines the type of the flag that will be created
	DefValue           interface{}
	DefValuePerCommand map[string]interface{}
	FlagAddMethod      string
	DefinedOn          []string
	Hidden             bool
	pflag *pflag.Flag
}

// flagRegistry is a list of all Disco CLI flags.
// When adding a new flag to the registry, please specify the
// command/commands to which the flag belongs in `DefinedOn` field.
// If the flag is a global flag, or belongs to all the subcommands,
/// specify "all"
// FlagAddMethod is method which defines a flag value with specified
// name, default value, and usage string. e.g. `StringVar`, `BoolVar`
var flagRegistry = []Flag{
	{
		Name:          "config-dir",
		Shorthand:     "cd",
		Usage:         "disco <command> --config-dir /this/dir",
		Value:         &DiscoOptions.ConfigDir,
		DefValue:      "",
		FlagAddMethod: "",
		DefinedOn:     []string{"up"},
	},
	{
		Name:          "router",
		Shorthand:     "r",
		Usage:         "disco <command> --router frr",
		Value:         &DiscoOptions.Router,
		DefValue:      "",
		FlagAddMethod: "StringVar",
		DefinedOn: []string{"up"},
	},
}

func methodNameByType(v reflect.Value) string {
	t := v.Type().Kind()
	switch t {
	case reflect.Int:
		return "IntVar"
	case reflect.Bool:
		return "BoolVar"
	case reflect.String:
		return "StringVar"
	case reflect.Slice:
		return "StringSliceVar"
	case reflect.Struct:
		return "Var"
	case reflect.Ptr:
		return methodNameByType(reflect.Indirect(v))
	}
	return ""
}

// Flag return a pflag.Flag from the discoCMD-flag
func (fl *Flag) Flag() *pflag.Flag {
	if fl.pflag != nil {
		return fl.pflag
	}

	methodName := fl.FlagAddMethod
	if methodName == "" {
		methodName = methodNameByType(reflect.ValueOf(fl.Value))
	}
	inputs := []interface{}{fl.Value, fl.Name}
	if methodName != "Var" {
		inputs = append(inputs, fl.DefValue)
	}
	inputs = append(inputs, fl.Usage)

	fs := pflag.NewFlagSet(fl.Name, pflag.ContinueOnError)

	reflect.ValueOf(fs).MethodByName(methodName).Call(reflectValueOf(inputs))
	f := fs.Lookup(fl.Name)
	f.Shorthand = fl.Shorthand
	f.Hidden = fl.Hidden

	fl.pflag = f
	return f
}

func reflectValueOf(values []interface{}) []reflect.Value {
	var results []reflect.Value
	for _, v := range values {
		results = append(results, reflect.ValueOf(v))
	}
	return results
}

// ParseFlags - adds flags to the given cmd
func ParseFlags(cmd *cobra.Command, flags []*Flag) {
	// Update default values.
	for _, fl := range flags {
		flag := cmd.Flag(fl.Name)
		if fl.DefValuePerCommand != nil {
			if defValue, present := fl.DefValuePerCommand[cmd.Use]; present {
				if !flag.Changed {
					flag.Value.Set(fmt.Sprintf("%v", defValue))
				}
			}
		}
	}
}

// AddFlags adds to the command the common flags that are annotated with the command name.
func AddFlags(cmd *cobra.Command) {
	for i := range flagRegistry {
		fl := &flagRegistry[i]
		if !hasCmdAnnotation(cmd.Use, fl.DefinedOn) {
			continue
		}

		cmd.Flags().AddFlag(fl.Flag())

	}
}

func hasCmdAnnotation(cmdName string, annotations []string) bool {
	for _, a := range annotations {
		if cmdName == a || a == "all" {
			return true
		}
	}
	return false
}

