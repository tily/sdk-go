package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "generate model json files",
	Long:  `generate model json files`,
	Run:   generate,
}

type Generator struct {
	Metadata                     Metadata
	HTTP                         HTTP
	OperationList                []OperationHTML
	APIModel                     APIModel
	DocsModel                    DocsModel
	RequestShapeSelector         string
	RequestShapeNameSelector     string
	RequestShapeTypeSelector     string
	RequestShapeRequiredSelector string
	RequestShapeRequiredRegexp   string
	RequestShapeIgnoreSelector   string
	RequestShapeIgnoreRegexp     string
	ResultShapeSelector          string
	ResultShapeNameSelector      string
	ResultShapeTypeSelector      string
	ResultExampleSelector        string
}

func init() {
	RootCommand.AddCommand(generateCommand)
}

func generate(cmd *cobra.Command, args []string) {
	var err error
	var services []Service

	servicesJson, err := ioutil.ReadFile("gen/config/services.json")
	err = json.Unmarshal(servicesJson, &services)
	panicIfErr(err)

	for _, service := range services {
		if options.ServiceID != "" && options.ServiceID != service.ID {
			continue
		}

		generator := service.Generator
		generator.APIModel.Metadata = generator.Metadata
		generator.APIModel.Operations = make(map[string]Operation)
		generator.APIModel.Shapes = make(map[string]Shape)

		generator.DocsModel.Version = "2"
		generator.DocsModel.Service = service.ID
		generator.DocsModel.Operations = make(map[string]string)
		generator.DocsModel.Shapes = make(map[string]DocsShape)

		fmt.Println("Loading oeration list")
		generator.loadOperationList(service.ID)

		fmt.Println("Processing operation list")
		generator.processOperationList()

		b, err := json.MarshalIndent(generator.APIModel, "", "  ")
		panicIfErr(err)

		outPath := fmt.Sprintf("models/apis/%s/%s/api-2.json", service.ID, generator.Metadata.APIVersion)
		prepareDir(outPath)
		err = ioutil.WriteFile(outPath, b, 0644)
		panicIfErr(err)

		b2, err := json.MarshalIndent(generator.DocsModel, "", "  ")
		panicIfErr(err)

		outPath2 := fmt.Sprintf("models/apis/%s/%s/docs-2.json", service.ID, generator.Metadata.APIVersion)
		prepareDir(outPath2)
		err = ioutil.WriteFile(outPath2, b2, 0644)
	}
}

func (g *Generator) loadOperationList(serviceID string) {
	var list []OperationHTML
	paths, err := filepath.Glob(fmt.Sprintf("gen/html/%s/*.html", serviceID))
	panicIfErr(err)
	for _, path := range paths {
		basename := filepath.Base(path)

		if regexp.MustCompile(`^_`).MatchString(basename) {
			continue
		}

		ext := filepath.Ext(basename)
		opName := basename[0 : len(basename)-len(ext)]
		opHTML := OperationHTML{OpName: opName, Path: path}
		list = append(list, opHTML)
	}
	g.OperationList = list
}

func (g *Generator) processOperationList() {
	for _, opHTML := range g.OperationList {
		if options.OpName != "" && options.OpName != opHTML.OpName {
			continue
		}

		if opHTML.OpName == "DescribeNASSecurityGroup" || opHTML.OpName == "DescribeNASInstance" {
			opHTML.OpName = opHTML.OpName + "s"
		}

		fmt.Printf("Processing %s\n", opHTML.OpName)
		doc, err := g.loadOperationHTML(opHTML.Path)
		panicIfErr(err)

		for _, shape := range g.generateRequestShapes(opHTML.OpName, doc) {
			g.APIModel.Shapes[shape.ShapeName] = shape
		}

		for _, shape := range g.generateResultShapes(opHTML.OpName, doc) {
			g.APIModel.Shapes[shape.ShapeName] = shape
		}

		g.APIModel.Operations[opHTML.OpName] = g.generateOperation(opHTML.OpName)

		for name, _ := range g.APIModel.Operations {
			g.DocsModel.Operations[name] = name
		}

		for name, _ := range g.APIModel.Shapes {
			g.DocsModel.Shapes[name] = DocsShape{Base: name}
		}
	}
}

func (g *Generator) generateOperation(name string) (operation Operation) {
	operation.OpName = name
	operation.Input = &ShapeRef{ShapeName: fmt.Sprintf("%sRequest", operation.OpName)}

	outputRefName := fmt.Sprintf("%sResult", operation.OpName)
	// create output shape ref only when there is a result shape
	if g.APIModel.Shapes[outputRefName].ShapeName == outputRefName {
		operation.Output = &ShapeRef{ShapeName: outputRefName}
	}

	if g.HTTP.Method != "" && g.HTTP.RequestURI != "" {
		operation.HTTP = g.HTTP
	}
	return operation
}

func (g *Generator) loadOperationHTML(path string) (doc *goquery.Document, err error) {
	f, err := os.Open(path)
	if err != nil {
		return doc, err
	}

	reader := bufio.NewReader(f)
	doc, err = goquery.NewDocumentFromReader(reader)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

func (g *Generator) generateRequestShapes(operationName string, doc *goquery.Document) (shapes []Shape) {
	shape := Shape{}
	shape.ShapeName = fmt.Sprintf("%sRequest", operationName)
	shape.Type = "structure"
	shape.Members = map[string]ShapeRef{}

	required := []string{}

	doc.Find(g.RequestShapeSelector).Each(func(_ int, s *goquery.Selection) {
		member := ShapeRef{}
		name := s.Find(g.RequestShapeNameSelector).First().Text()
		if name == "" {
			return
		}

		ignoreText := s.Find(g.RequestShapeIgnoreSelector).First().Text()
		if regexp.MustCompile(g.RequestShapeIgnoreRegexp).MatchString(ignoreText) {
			return
		}

		typeText := s.Find(g.RequestShapeTypeSelector).First().Text()
		var _type string
		switch {
		case regexp.MustCompile(`\.member\.N\..+$`).MatchString(name):
			r := regexp.MustCompile(`^(.+)\.member\.N\.(.+)$`)
			match := r.FindAllStringSubmatch(name, -1)
			name = match[0][1]
			value := match[0][2]

			_type = fmt.Sprintf("%sStructList", name)
			structShapeName := fmt.Sprintf("%sStruct", name)
			structListShape := Shape{
				ShapeName: _type,
				Type:      "list",
				Member:    &ShapeRef{ShapeName: structShapeName},
			}

			shapes = append(shapes, structListShape)
			structShape := Shape{}
			for _, shape := range shapes {
				if shape.ShapeName == structShapeName {
					structShape = shape
					break
				}
			}
			member := ShapeRef{ShapeName: "String"}
			if structShape.ShapeName == "" {
				structShape = Shape{
					ShapeName: structShapeName,
					Type:      "structure",
					Members:   map[string]ShapeRef{value: member},
				}
				shapes = append(shapes, structShape)
			} else {
				structShape.Members[value] = member
			}
		case typeText == "数値" || typeText == "Long":
			_type = "Integer"
		case typeText == "文字列":
			_type = "String"
		case typeText == "真偽値" || typeText == "boolean":
			_type = "Boolean"
		case typeText == "日付" || typeText == "日時":
			_type = "TStamp"
			tstampShape := Shape{
				ShapeName: _type,
				Type:      "timestamp",
			}
			shapes = append(shapes, tstampShape)
		case typeText == "文字配列":
			name = regexp.MustCompile(`\.member\.N$`).ReplaceAllString(name, "")
			_type = fmt.Sprintf("%sStringList", name)
			stringListShape := Shape{
				ShapeName: _type,
				Type:      "list",
				Member:    &ShapeRef{ShapeName: "String"},
			}
			shapes = append(shapes, stringListShape)
		}
		member.ShapeName = _type

		requiredText := s.Find(g.RequestShapeRequiredSelector).First().Text()
		if regexp.MustCompile(g.RequestShapeRequiredRegexp).MatchString(requiredText) {
			required = append(required, name)
		}

		shape.Members[name] = member
		shape.Required = required
	})
	shapes = append(shapes, shape)
	return shapes
}

func (g *Generator) generateResultShapes(operationName string, doc *goquery.Document) (shapes []Shape) {
	hintXML := g.extractHintXML(doc)
	shapeNames := g.extractShapeNames(doc)

	doc.Find(g.ResultShapeSelector).Each(func(i int, s *goquery.Selection) {
		shapeName := s.Find(g.ResultShapeNameSelector).First().Text()
		shapeName = regexp.MustCompile(`(\s)+$`).ReplaceAllString(shapeName, "")

		// skip if it's the xml root tag
		if i == 0 && regexp.MustCompile(`Response$`).MatchString(shapeName) {
			return
		}
		// skip if it's the `ResponseMetadata`
		if shapeName == "ResponseMetadata" {
			return
		}

		typeText := s.Find(g.ResultShapeTypeSelector).First().Text()
		switch {
		case typeText == "数値" || typeText == "Long":
			shape := Shape{ShapeName: shapeName, Type: "integer"}
			shapes = append(shapes, shape)
		case typeText == "文字列":
			shape := Shape{ShapeName: shapeName, Type: "string"}
			shapes = append(shapes, shape)
		case typeText == "真偽値" || typeText == "boolean":
			shape := Shape{ShapeName: shapeName, Type: "boolean"}
			shapes = append(shapes, shape)
		case typeText == "日時":
			shape := Shape{ShapeName: shapeName, Type: "timestamp"}
			shapes = append(shapes, shape)
		case typeText == "リスト":
			child := goquery.NodeName(hintXML.Find(strings.ToLower(shapeName)).Children().First())
			if pos := shapeNames.pos(child); pos != -1 {
				shape := Shape{
					ShapeName: shapeName,
					Type:      "list",
					Member:    &ShapeRef{ShapeName: shapeNames[pos], LocationName: shapeNames[pos]},
				}
				shapes = append(shapes, shape)
			}
		case typeText == "－" || typeText == "-" || typeText == "\u00a0":
			children := hintXML.Find(strings.ToLower(shapeName)).Children()
			members := map[string]ShapeRef{}
			children.Each(func(_ int, s *goquery.Selection) {
				nodeName := goquery.NodeName(s)
				if nodeName == "responsemetadata" {
					return
				}
				if pos := shapeNames.pos(nodeName); pos != -1 {
					memberShapeName := shapeNames[pos]
					members[memberShapeName] = ShapeRef{ShapeName: memberShapeName}
				}
			})
			shape := Shape{
				ShapeName: shapeName,
				Type:      "structure",
				Members:   members,
			}
			shapes = append(shapes, shape)
		}
	})
	shapes = append(shapes, Shape{ShapeName: "Integer", Type: "integer"})
	shapes = append(shapes, Shape{ShapeName: "String", Type: "string"})
	shapes = append(shapes, Shape{ShapeName: "Boolean", Type: "boolean"})
	return shapes
}

func (g *Generator) extractShapeNames(doc *goquery.Document) (names ShapeNames) {
	doc.Find(g.ResultShapeSelector).Each(func(_ int, s *goquery.Selection) {
		name := s.Find(g.ResultShapeNameSelector).First().Text()
		names = append(names, name)
	})
	return names
}

func (g *Generator) extractHintXML(doc *goquery.Document) (hintXML *goquery.Document) {
	text := doc.Find(g.ResultExampleSelector).Text()
	reader := strings.NewReader(text)
	hintXML, err := goquery.NewDocumentFromReader(reader)
	panicIfErr(err)
	return hintXML
}

type ShapeNames []string

// get position in strint slice case-insensitively
func (names ShapeNames) pos(value string) int {
	for p, v := range names {
		if strings.ToLower(v) == value {
			return p
		}
	}
	return -1
}
