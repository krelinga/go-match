package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type StructField struct {
	Name string
	Type string
}

type Generator struct {
	matchType  string
	outType    string
	outFile    string
	outPackage string
}

func main() {
	var (
		outFile    = flag.String("out", "", "Output .go file to generate")
		matchType  = flag.String("match_type", "", "Name of the Go type to match against")
		outType    = flag.String("out_type", "", "Name of the Go matcher type to generate")
		outPackage = flag.String("out_package", "", "Package name for the generated matcher (defaults to match_type's package)")
	)
	flag.Parse()

	if *outFile == "" || *matchType == "" || *outType == "" {
		log.Fatal("All flags -out, -match_type, and -out_type are required")
	}

	gen := &Generator{
		matchType:  *matchType,
		outType:    *outType,
		outFile:    *outFile,
		outPackage: *outPackage,
	}

	if err := gen.generate(); err != nil {
		log.Fatalf("Failed to generate matcher: %v", err)
	}

	fmt.Printf("Generated matcher for %s in %s\n", *matchType, *outFile)
}

func (g *Generator) generate() error {
	// Find the struct definition in the current package
	fields, packageName, err := findStructFields(g.matchType)
	if err != nil {
		return fmt.Errorf("failed to find struct fields: %w", err)
	}

	// Generate the matcher code
	code, err := g.generateMatcherCode(fields, packageName)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	// Write to output file
	if err := g.writeToFile(code); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func findStructFields(matchType string) ([]StructField, string, error) {
	// Parse all Go files in the current directory
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", func(info os.FileInfo) bool {
		return strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse package: %w", err)
	}

	var fields []StructField
	var packageName string

	// Look for the struct definition in all packages
	for pkgName, pkg := range pkgs {
		packageName = pkgName
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				switch node := n.(type) {
				case *ast.TypeSpec:
					if node.Name.Name == matchType {
						if structType, ok := node.Type.(*ast.StructType); ok {
							fields = extractFields(structType, fset)
							return false // Found it, stop searching
						}
					}
				}
				return true
			})
			if len(fields) > 0 {
				break
			}
		}
		if len(fields) > 0 {
			break
		}
	}

	if len(fields) == 0 {
		return nil, "", fmt.Errorf("struct type %s not found in current package", matchType)
	}

	return fields, packageName, nil
}

func extractFields(structType *ast.StructType, fset *token.FileSet) []StructField {
	var fields []StructField

	for _, field := range structType.Fields.List {
		// Only process exported fields (those starting with uppercase)
		for _, name := range field.Names {
			if name.IsExported() {
				fieldType := formatType(field.Type, fset)
				fields = append(fields, StructField{
					Name: name.Name,
					Type: fieldType,
				})
			}
		}
	}

	return fields
}

func formatType(expr ast.Expr, fset *token.FileSet) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		pkg := formatType(t.X, fset)
		return pkg + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + formatType(t.X, fset)
	case *ast.ArrayType:
		if t.Len == nil {
			// Slice
			return "[]" + formatType(t.Elt, fset)
		}
		// Array with length
		return "[" + formatExpr(t.Len, fset) + "]" + formatType(t.Elt, fset)
	case *ast.MapType:
		keyType := formatType(t.Key, fset)
		valueType := formatType(t.Value, fset)
		return "map[" + keyType + "]" + valueType
	case *ast.ChanType:
		dir := ""
		switch t.Dir {
		case ast.SEND:
			dir = "chan<- "
		case ast.RECV:
			dir = "<-chan "
		default:
			dir = "chan "
		}
		return dir + formatType(t.Value, fset)
	case *ast.InterfaceType:
		if len(t.Methods.List) == 0 {
			return "interface{}"
		}
		// For complex interfaces, we'll just use interface{}
		return "interface{}"
	case *ast.FuncType:
		return "func" + formatFuncSignature(t, fset)
	default:
		// Fallback: try to format as source code
		return formatExpr(expr, fset)
	}
}

func formatExpr(expr ast.Expr, fset *token.FileSet) string {
	var buf strings.Builder
	if err := format.Node(&buf, fset, expr); err != nil {
		return "interface{}" // Fallback
	}
	return buf.String()
}

func formatFuncSignature(funcType *ast.FuncType, fset *token.FileSet) string {
	var parts []string

	// Parameters
	if funcType.Params != nil {
		var params []string
		for _, field := range funcType.Params.List {
			fieldType := formatType(field.Type, fset)
			if len(field.Names) == 0 {
				params = append(params, fieldType)
			} else {
				for range field.Names {
					params = append(params, fieldType)
				}
			}
		}
		parts = append(parts, "("+strings.Join(params, ", ")+")")
	} else {
		parts = append(parts, "()")
	}

	// Results
	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		var results []string
		for _, field := range funcType.Results.List {
			fieldType := formatType(field.Type, fset)
			if len(field.Names) == 0 {
				results = append(results, fieldType)
			} else {
				for range field.Names {
					results = append(results, fieldType)
				}
			}
		}
		if len(results) == 1 {
			parts = append(parts, " "+results[0])
		} else {
			parts = append(parts, " ("+strings.Join(results, ", ")+")")
		}
	}

	return strings.Join(parts, "")
}

func (g *Generator) generateMatcherCode(fields []StructField, packageName string) (string, error) {
	var builder strings.Builder

	// Determine which package to use
	targetPackage := packageName
	if g.outPackage != "" {
		targetPackage = g.outPackage
	}

	// Package declaration
	builder.WriteString(fmt.Sprintf("package %s\n\n", targetPackage))

	// Determine needed imports
	needsTime := false
	needsSourcePackage := false

	for _, field := range fields {
		if strings.Contains(field.Type, "time.Time") {
			needsTime = true
		}
	}

	// Check if we need to import the source package for the match type
	if g.outPackage != "" && g.outPackage != packageName {
		needsSourcePackage = true
	}

	// Imports
	builder.WriteString("import (\n")
	if needsTime {
		builder.WriteString("\t\"time\"\n")
	}
	if needsSourcePackage {
		// We need to import the source package to reference the match type
		// For now, we'll assume it's in the current module path
		builder.WriteString("\t\"github.com/krelinga/go-match\"\n")
	}
	builder.WriteString("\t\"github.com/krelinga/go-match/matchfmt\"\n")
	builder.WriteString(")\n\n")

	// Struct definition
	matcherTypeRef := "Matcher"
	if needsSourcePackage {
		matcherTypeRef = packageName + ".Matcher"
	}
	builder.WriteString(fmt.Sprintf("type %s struct {\n", g.outType))
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("\t%s %s[%s]\n", field.Name, matcherTypeRef, field.Type))
	}
	builder.WriteString("}\n\n")

	// Match method
	matchTypeRef := g.matchType
	if needsSourcePackage {
		matchTypeRef = packageName + "." + g.matchType
	}
	builder.WriteString(fmt.Sprintf("func (m *%s) Match(got %s) (bool, string) {\n", g.outType, matchTypeRef))
	builder.WriteString("\tvar details []string\n")
	builder.WriteString("\tallMatched := true\n\n")

	// Generate field matching logic
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("\tif m.%s != nil {\n", field.Name))
		builder.WriteString(fmt.Sprintf("\t\tmatched, explanation := m.%s.Match(got.%s)\n", field.Name, field.Name))
		builder.WriteString("\t\tif !matched {\n")
		builder.WriteString("\t\t\tallMatched = false\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t\tdetails = append(details, explanation)\n")
		builder.WriteString("\t}\n\n")
	}

	builder.WriteString(fmt.Sprintf("\treturn allMatched, matchfmt.Explain(allMatched, \"%s\", details...)\n", g.outType))
	builder.WriteString("}\n")

	return builder.String(), nil
}

func (g *Generator) writeToFile(code string) error {
	// Parse and format the generated code
	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse generated code: %w", err)
	}

	// Create output directory if it doesn't exist
	dir := filepath.Dir(g.outFile)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Open output file
	file, err := os.Create(g.outFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Format and write the code
	if err := format.Node(file, fset, parsed); err != nil {
		return fmt.Errorf("failed to format and write code: %w", err)
	}

	return nil
}
