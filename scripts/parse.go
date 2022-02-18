package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/slices"
)

// outfile is the filename to write out the generated types.
const outfile = "types.go"

// body describes the schema's JSON structure.
type body struct {
	Schema struct {
		QueryType struct {
			Name string `json:"name"`
		} `json:"queryType"`
		MutationType struct {
			Name string `json:"name"`
		} `json:"mutationType"`
		SubscriptionType interface{} `json:"subscriptionType"`
		Types            []gqlType   `json:"types"`
	} `json:"__schema"`
}

// enumValue describes a GraphQL enum value.
type enumValue struct {
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	IsDeprecated      bool        `json:"isDeprecated"`
	DeprecationReason interface{} `json:"deprecationReason"`
}

// gqlType describes a GraphQL type.
type gqlType struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Fields      []struct {
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Args        []interface{} `json:"args"`
		Type        struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
			// We're going to nest a bunch of discrete OfTypes to avoid the
			// possibility of nil dereferences. In practice, these can go three
			// levels deep 🙃
			OfType struct {
				Kind   string `json:"kind"`
				Name   string `json:"name"`
				OfType struct {
					Kind   string `json:"kind"`
					Name   string `json:"name"`
					OfType struct {
						Kind string `json:"kind"`
						Name string `json:"name"`
					}
				} `json:"ofType,omitempty"`
			} `json:"ofType"`
		} `json:"type"`
		IsDeprecated      bool        `json:"isDeprecated"`
		DeprecationReason interface{} `json:"deprecationReason"`
	} `json:"fields"`
	EnumValues []enumValue `json:"enumValues"`
}

// basicType describes the essential information required to construct a Go type
// from the JSON type descriptor.
type basicType struct {
	PropertyName string
	Name         string
	Description  string
	Comment      string
	GoType       string
}

// A list of enums to skip during generation, since it gets mildly excessive.
var enumSkip = []string{
	"CountryCode", "CurrencyCode", "WeightUnit",
	"UnitPriceMeasurementMeasuredType", "UnitPriceMeasurementMeasuredUnit",
}

// typeMap maps GraphQL/Storefront types to Go's types.
var typeMap = map[string]string{
	"Boolean":  "bool",
	"Decimal":  "float64",
	"Float":    "float64",
	"ID":       "string",
	"HTML":     "string",
	"Int":      "int",
	"JSON":     "map[string]interface{}",
	"Money":    "string",
	"String":   "string",
	"DateTime": "time.Time",
	"URL":      "string",
}

// acronymMap acts as a LUT to correctly case generated identifiers.
var acronymMap = map[string]string{
	"Html":    "HTML",
	"Seo":     "SEO",
	"3d":      "3D",
	"Png":     "PNG",
	"Jpg":     "JPG",
	"Webp":    "WebP",
	"Url":     "URL",
	"Jcb":     "JCB",
	"Sku":     "SKU",
	"Ssl":     "SSL",
	"Youtube": "YouTube",
	"Zip":     "ZIP",
}

func main() {
	flag.Parse()

	inFiles := flag.Args()
	var inFile string

	if len(inFiles) == 0 {
		inFile = path.Join("schema", "2022-01", "schema.json")
	} else {
		inFile = inFiles[0]
	}

	for k, v := range acronymMap {
		strcase.ConfigureAcronym(k, v)
	}

	for _, enum := range enumSkip {
		// We're going to skip generating these enums, so we want our generated
		// types for these enums to be string.
		typeMap[enum] = "string"
	}

	bs, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err)
	}

	var b body
	if err := json.Unmarshal(bs, &b); err != nil {
		log.Fatal(err)
	}

	out := jen.NewFile("storefront")

	for _, t := range b.Schema.Types {
		// Types that begin with __ are internal, I think. Omit them.
		if strings.HasPrefix(t.Name, "__") {
			continue
		}

		switch t.Kind {
		case "OBJECT", "INTERFACE":
			// Generic *Connection and *Edge types are pre-implemented, as their
			// inclusion during generation would lead to illegal cycles.
			if strings.HasSuffix(t.Name, "Connection") || strings.HasSuffix(t.Name, "Edge") {
				continue
			}

			var fields []basicType

			for _, f := range t.Fields {
				// If we're to leverage the built-in Go type, look it up from the map.
				// If the typename ends in "Connection" or "Edge", we'll specify the
				// pre-defined generic type.
				//
				// To break cyclic references, we implement special-case handling for
				// ProductVariant.Product and for Blog.
				//
				// Failing that, use the existing type name.
				typ := f.Type.Name
				if typ == "" {
					typ = f.Type.OfType.Name
				}

				if val, ok := typeMap[typ]; ok {
					typ = val
				}

				// If the type is a list, traverse leaves further down and determine
				// the element type.
				if f.Type.Kind == "LIST" || f.Type.OfType.Kind == "LIST" {
					elementTypeName := f.Type.OfType.OfType.Name
					if elementTypeName == "" {
						elementTypeName = f.Type.OfType.OfType.OfType.Name // really?
					}

					if val, ok := typeMap[elementTypeName]; ok {
						typ = "[]" + val
					} else {
						typ = "[]" + elementTypeName
					}
				}

				// For simplicity, just treat unions as strings.
				if f.Type.Kind == "UNION" || f.Type.OfType.Kind == "UNION" {
					typ = "string"
				}

				// Special cases:
				if (t.Name == "ProductVariant" && f.Name == "product") || (t.Name == "Article" && f.Name == "blog") {
					log.Printf("Note: %s.%s type set to interface{}", t.Name, f.Name)
					typ = "interface{}"
				}

				name := strcase.ToCamel(f.Name)

				// Replace any instances in acronymMap with the correctly-cased variants.
				for k, v := range acronymMap {
					if strings.Contains(typ, k) {
						typ = strings.Replace(typ, k, v, -1)
					}

					if strings.Contains(name, k) {
						name = strings.Replace(name, k, v, -1)
					}
				}

				comment := fmt.Sprintf("%s: %s", name, f.Description)
				if f.IsDeprecated {
					comment = fmt.Sprintf("%s: %s: %s", "Deprecated", name, f.Description)
				}

				fields = append(fields, basicType{
					PropertyName: name,
					Name:         f.Name,
					Description:  f.Description,
					Comment:      comment,
					GoType:       typ,
				})
			}

			var props []jen.Code

			for _, f := range fields {
				name := jen.Id(f.PropertyName)
				tag := jen.Tag(map[string]string{"json": fmt.Sprintf("%s,omitempty", f.Name)})

				qual := jen.Id(f.GoType)
				if f.GoType == "time.Time" {
					qual = jen.Qual("time", "Time")
				}
				if strings.HasSuffix(f.GoType, "Connection") {
					// For *Connection properties, we're going to strip the "Connection"
					// and specify the generic type.
					before, _, _ := strings.Cut(f.GoType, "Connection")
					if before == "String" {
						before = "string"
					}

					qual = jen.Id("Connection").Types(jen.Id(before))
				}

				props = append(
					props,
					jen.Comment(transformFieldComment(f.PropertyName, f.Description)),
					jen.Add(name, qual, tag),
				)
			}

			typeName := replaceAcronyms(t.Name)

			out.Commentf("%s: %s", typeName, t.Description)
			out.Type().Id(typeName).Struct(props...).Line()
		case "ENUM":
			if slices.Contains(enumSkip, t.Name) {
				continue
			}

			out.Commentf("%s: %s", t.Name, t.Description)
			out.Add(jen.Type().Id(t.Name).String())
			var props []jen.Code

			for _, e := range t.EnumValues {
				name := t.Name + strcase.ToCamel(strings.ToLower(e.Name))

				name = replaceAcronyms(name)
				props = append(props, jen.Id(name).Qual("", t.Name).Op("=").Lit(e.Name))
			}

			out.Const().Defs(props...)
		}
	}

	if err := out.Save(outfile); err != nil {
		log.Fatal(err)
	}
}

var fetchReturnsOrFindRgx = regexp.MustCompile(`^(Fetch|Returns|Find\w*)\s(an|a|the)`)
var mutationTermRgx = regexp.MustCompile(`^(Creates|Updates|Adds|Removes|Completes|Associates|Disassociates|Applies|Appends|Sets|Activates|Sends|Resets\w*)\s(an|a|the)`)
var whetherRegex = regexp.MustCompile(`^(Whether\w*)\s(an|a|the)`)
var specialsRegex = regexp.MustCompile(`^(Stripped|Image\w*)`)

// transformFieldComment accepts an identifier and the GQL description,
// attempting to hack together a more Go-like doc comment.
func transformFieldComment(identifier, s string) string {
	// the "connecting word", i.e. the word inserted betwixt the identifier and
	// the normalized comment.
	const cw = "is"

	if fetchReturnsOrFindRgx.MatchString(s) {
		return fmt.Sprintf("%s %s %s", identifier, cw, fetchReturnsOrFindRgx.ReplaceAllString(s, "$2"))
	}
	if mutationTermRgx.MatchString(s) {
		return fmt.Sprintf(
			"%s %s",
			identifier,
			mutationTermRgx.ReplaceAllStringFunc(s, func(ss string) string {
				return strings.ToLower(ss)
			}))
	}
	if whetherRegex.MatchString(s) {
		return fmt.Sprintf("%s %s %s", identifier, cw, whetherRegex.ReplaceAllString(s, "whether $2"))
	}
	if specialsRegex.MatchString(s) {
		return fmt.Sprintf(
			"%s %s %s",
			identifier,
			cw,
			specialsRegex.ReplaceAllStringFunc(s, func(ss string) string {
				return "the " + strings.ToLower(ss)
			}))
	}
	if strings.HasPrefix(s, "List of") {
		return fmt.Sprintf("%s %s %s", identifier, cw, strings.Replace(s, "List of", "a list of", 1))
	}
	if strings.HasPrefix(s, "Indicates") {
		return fmt.Sprintf("%s %s", identifier, strings.Replace(s, "Indicates", "indicates", 1))
	}
	if strings.HasPrefix(s, "The") {
		return fmt.Sprintf("%s %s %s", identifier, cw, strings.Replace(s, "The", "the", 1))
	}
	if strings.HasPrefix(s, "A") {
		return fmt.Sprintf("%s %s %s", identifier, cw, strings.Replace(s, "A", "a", 1))
	}

	// Just return something that is, hopefully, *technically* correct.
	var b strings.Builder

	b.WriteString(strings.ToLower(string(s[0])))
	b.WriteString(s[1:])

	return fmt.Sprintf("%s %s the %s", identifier, cw, b.String())
}

func replaceAcronyms(s string) string {
	for k, v := range acronymMap {
		if strings.Contains(s, k) {
			return strings.Replace(s, k, v, -1)
		}
	}

	return s
}
