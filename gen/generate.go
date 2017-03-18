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
	RequestShapeTypeRegexp       string
	RequestShapeRequiredSelector string
	RequestShapeRequiredRegexp   string
	RequestShapeIgnoreSelector   string
	RequestShapeIgnoreRegexp     string
	ResultShapeSelector          string
	ResultShapeNameSelector      string
	ResultShapeTypeSelector      string
	ResultShapeTypeRegexp        string
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

		for opName, _ := range g.APIModel.Operations {
			g.DocsModel.Operations[opName] = opName
		}

		for shapeName, _ := range g.APIModel.Shapes {
			g.DocsModel.Shapes[shapeName] = DocsShape{Base: shapeName}
		}
	}
}

func (g *Generator) generateOperation(opName string) (operation Operation) {
	operation.OpName = opName
	operation.Input = &ShapeRef{ShapeName: fmt.Sprintf("%sRequest", operation.OpName)}

	outputShapeName := fmt.Sprintf("%sResult", operation.OpName)
	// create output shape ref only when there is a result shape
	if g.APIModel.Shapes[outputShapeName].ShapeName == outputShapeName {
		operation.Output = &ShapeRef{ShapeName: outputShapeName}
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

func (g *Generator) generateRequestShapes(operationName string, doc *goquery.Document) (shapes Shapes) {
	shape := Shape{}
	shape.ShapeName = fmt.Sprintf("%sRequest", operationName)
	shape.Type = "structure"
	shape.Members = map[string]ShapeRef{}

	required := []string{}

	doc.Find(g.RequestShapeSelector).Each(func(_ int, s *goquery.Selection) {
		member := ShapeRef{}
		shapeName := s.Find(g.RequestShapeNameSelector).First().Text()
		shapeName = regexp.MustCompile(`(\s|\t|\n)`).ReplaceAllString(shapeName, "")
		if shapeName == "" {
			return
		}

		ignoreText := s.Find(g.RequestShapeIgnoreSelector).First().Text()
		if regexp.MustCompile(g.RequestShapeIgnoreRegexp).MatchString(ignoreText) {
			return
		}

		typeText := s.Find(g.RequestShapeTypeSelector).First().Text()
		if g.RequestShapeTypeRegexp != "" {
			r := regexp.MustCompile(g.RequestShapeTypeRegexp)
			m := r.FindStringSubmatch(typeText)
			result := make(map[string]string)
			for i, name := range r.SubexpNames() {
				if i != 0 && i <= len(m) {
					result[name] = m[i]
				}
			}
			if result["type"] == "" {
				return
			}
			typeText = result["type"]
		}

		shapeType := ""
		switch {
		case regexp.MustCompile(`^(数値|Long|int|Integer|xsd:(int|Int))$`).MatchString(typeText):
			shapeType = "Integer"
		case regexp.MustCompile(`^(文字列|String|xsd:string|Sring|string|String )$`).MatchString(typeText):
			shapeType = "String"
		case regexp.MustCompile(`^(真偽値|boolean|Boolean|bBoolean)$`).MatchString(typeText):
			shapeType = "Boolean"
		case regexp.MustCompile(`^Double$`).MatchString(typeText):
			shapeType = "Double"
		case regexp.MustCompile(`^(日付|日時)$`).MatchString(typeText):
			shapeType = "TStamp"
		case regexp.MustCompile(`^文字配列$`).MatchString(typeText):
			shapeType = "List"
		}

		shapeName, shapeType = g.parseRequestParam(shapeName, shapeType, &shapes)
		member.ShapeName = shapeType

		requiredText := s.Find(g.RequestShapeRequiredSelector).First().Text()
		if regexp.MustCompile(g.RequestShapeRequiredRegexp).MatchString(requiredText) {
			required = append(required, shapeName)
		}

		shape.Members[shapeName] = member
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
		if g.ResultShapeTypeRegexp != "" {
			r := regexp.MustCompile(g.ResultShapeTypeRegexp)
			m := r.FindStringSubmatch(typeText)
			result := make(map[string]string)
			for i, name := range r.SubexpNames() {
				if i != 0 && i <= len(m) {
					result[name] = m[i]
				}
			}
			if result["type"] == "" {
				return
			}
			typeText = result["type"]
		}
		switch {
		case regexp.MustCompile(`^(数値|Long|xsd?:\s?(Integer|integer|integer |int|intger|int\(A 16-bit unsigned\)|long|short))$`).MatchString(typeText):
			shape := Shape{ShapeName: shapeName, Type: "integer"}
			shapes = append(shapes, shape)
		case regexp.MustCompile(`^xsd?:(Double|double)$`).MatchString(typeText):
			shape := Shape{ShapeName: shapeName, Type: "double"}
			shapes = append(shapes, shape)
		case regexp.MustCompile(`^(文字列|xsd?::?\s?(string|String|stringint))$`).MatchString(typeText):
			shape := Shape{ShapeName: shapeName, Type: "string"}
			shapes = append(shapes, shape)
		case regexp.MustCompile(`^(真偽値|boolean|xsd?:(xs:)?\s?(boolean|Boolean))$`).MatchString(typeText):
			shape := Shape{ShapeName: shapeName, Type: "boolean"}
			shapes = append(shapes, shape)
		case regexp.MustCompile(`^(日時|xsd?:\s?(dateTime|DateTime|datetime|calendar|Calendar)(　?\(yyyy-mm-ddThh:mm:ssZ\))?)$`).MatchString(typeText):
			shape := Shape{ShapeName: shapeName, Type: "timestamp"}
			shapes = append(shapes, shape)
		case regexp.MustCompile(`^リスト$`).MatchString(typeText):
			shape := g.generateResultListShape(shapeName, shapeNames, hintXML)
			if shape.ShapeName != "" {
				shapes = append(shapes, shape)
			}
		case regexp.MustCompile(`^(－|\-|\x{00a0})$`).MatchString(typeText):
			shape := g.generateResultStructShape(shapeName, shapeNames, hintXML)
			shapes = append(shapes, shape)
		}
	})
	shapes = append(shapes, Shape{ShapeName: "Integer", Type: "integer"})
	shapes = append(shapes, Shape{ShapeName: "String", Type: "string"})
	shapes = append(shapes, Shape{ShapeName: "Boolean", Type: "boolean"})
	shapes = append(shapes, Shape{ShapeName: "Double", Type: "double"})
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

func (g *Generator) generateResultListShape(shapeName string, shapeNames ShapeNames, hintXML *goquery.Document) (shape Shape) {
	child := goquery.NodeName(hintXML.Find(strings.ToLower(shapeName)).Children().First())
	if pos := shapeNames.pos(child); pos != -1 {
		shape = Shape{
			ShapeName: shapeName,
			Type:      "list",
			Member:    &ShapeRef{ShapeName: shapeNames[pos], LocationName: shapeNames[pos]},
		}
	}
	return shape
}

func (g *Generator) generateResultStructShape(shapeName string, shapeNames ShapeNames, hintXML *goquery.Document) (shape Shape) {
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
	return Shape{
		ShapeName: shapeName,
		Type:      "structure",
		Members:   members,
	}
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

func (g *Generator) parseRequestParam(param string, shapeType string, shapes *Shapes) (shapeName string, retShapeType string) {
	if regexp.MustCompile(`\.`).MatchString(param) {
		r := regexp.MustCompile(`([a-zA-Z]+)((\.member)?\.(n|m|l|1))?`)
		parts := r.FindAllStringSubmatch(param, -1)
		if parts[0][2] != "" {
			shapeName = fmt.Sprintf("%sList", parts[0][1])
			retShapeType = shapeName
		} else {
			shapeName = parts[0][1]
			retShapeType = fmt.Sprintf("%sStruct", parts[0][1])
		}
		for i, _ := range parts {
			if i == len(parts)-1 {
				currentShape := Shape{}
				currentShapeName := ""
				if parts[i][2] != "" {
					currentShapeName = fmt.Sprintf("%sList", parts[i][1])
					currentShape = Shape{ShapeName: currentShapeName, Type: "list", Member: &ShapeRef{ShapeName: shapeType}}
				} else {
					currentShapeName = parts[i][1]
					currentShape = Shape{ShapeName: currentShapeName, Type: strings.ToLower(shapeType)}
				}
				*shapes = append(*shapes, currentShape)
			} else if parts[i][2] != "" {
				currentShapeName := fmt.Sprintf("%sList", parts[i][1])
				currentShape := Shape{ShapeName: currentShapeName, Type: "list"}
				refShapeName := ""
				if parts[i+1][2] != "" {
					refShapeName = fmt.Sprintf("%sList", parts[i+1][1])
				} else {
					if i == len(parts)-2 {
						refShapeName = parts[i+1][1]
					} else {
						refShapeName = fmt.Sprintf("%sStruct", parts[i+1][1])
					}
				}
				currentShape.Member = &ShapeRef{ShapeName: refShapeName}
				*shapes = append(*shapes, currentShape)
			} else {
				currentShapeName := fmt.Sprintf("%sStruct", parts[i][1])
				currentShape := *shapes.findShapeByName(currentShapeName)
				flg := false
				if currentShape.ShapeName == "" {
					flg = true
					currentShape = Shape{ShapeName: currentShapeName, Type: "structure", Members: map[string]ShapeRef{}}
				}
				refShapeName := ""
				if parts[i+1][2] != "" {
					refShapeName = fmt.Sprintf("%sList", parts[i+1][1])
				} else {
					if i == len(parts)-2 {
						refShapeName = parts[i+1][1]
					} else {
						refShapeName = fmt.Sprintf("%sStruct", parts[i+1][1])
					}
				}
				currentShape.Members[refShapeName] = ShapeRef{ShapeName: refShapeName}
				if flg {
					*shapes = append(*shapes, currentShape)
				}
			}
		}
	} else {
		shapeName = param
		retShapeType = shapeType
		if shapeType == "TStamp" {
			tstampShape := Shape{
				ShapeName: shapeType,
				Type:      "timestamp",
			}
			*shapes = append(*shapes, tstampShape)
		}
	}

	return shapeName, retShapeType
}
