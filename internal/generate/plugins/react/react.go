package react

import (
	_ "embed"
	"strconv"
	"text/template"

	"github.com/open-feature/cli/internal/generate"
	"github.com/open-feature/cli/internal/generate/types"

	"github.com/iancoleman/strcase"
)

type TmplData struct {
	*types.BaseTmplData
}

type genImpl struct {
}

// BaseTmplDataInfo provides the base template data for the codegen.
func (td *TmplData) BaseTmplDataInfo() *types.BaseTmplData {
	return td.BaseTmplData
}

// supportedFlagTypes is the flag types supported by the Go template.
var supportedFlagTypes = map[types.FlagType]bool{
	types.FloatType:  true,
	types.StringType: true,
	types.IntType:    true,
	types.BoolType:   true,
	types.ObjectType: false,
}

func (*genImpl) SupportedFlagTypes() map[types.FlagType]bool {
	return supportedFlagTypes
}

//go:embed react.tmpl
var reactTmpl string

func flagVarName(flagName string) string {
	return strcase.ToCamel(flagName)
}

func flagInitParam(flagName string) string {
	return strconv.Quote(flagName)
}

func defaultValueLiteral(flag *types.FlagTmplData) string {
	switch flag.Type {
	case types.StringType:
		return strconv.Quote(flag.DefaultValue)
	default:
		return flag.DefaultValue
	}
}

func typeString(flagType types.FlagType) string {
	switch flagType {
	case types.StringType:
		return "string"
	case types.IntType, types.FloatType:
		return "number"
	case types.BoolType:
		return "boolean"
	default:
		return ""
	}
}

func (g *genImpl) Generate(input types.Input) error {
	funcs := template.FuncMap{
		"FlagVarName":         flagVarName,
		"FlagInitParam":       flagInitParam,
		"DefaultValueLiteral": defaultValueLiteral,
		"TypeString":          typeString,
	}
	td := TmplData{
		BaseTmplData: input.BaseData,
	}
	return generate.GenerateFile(funcs, reactTmpl, &td)
}

// Params are parameters for creating a Generator
type Params struct {
}

// NewGenerator creates a generator for React.
func NewGenerator(params Params) types.Generator {
	return &genImpl{}
}
