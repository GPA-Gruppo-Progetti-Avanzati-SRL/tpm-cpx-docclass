package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/templateutil"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/fielddictionary"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
	"time"
)

const (
	DocPropertyValue int = iota
	ConstValue
	SysDate

	SemLogDocClass            = "doc-class"
	SemLogFieldName           = "field-name"
	SemLogPropertySegmentName = "segment-name"

	PackageTTL  = 3
	DocumentTTL = 3
)

type OnExportEntryDefinition struct {
	Name        string `mapstructure:"name" yaml:"name" json:"name"`
	Format      string `mapstructure:"format" yaml:"format" json:"format"`
	StringValue string `mapstructure:"value" yaml:"value" json:"value"`
	FieldName   string `mapstructure:"field-name" yaml:"field-name" json:"field-name"`
}

func (ed *OnExportEntryDefinition) Value() interface{} {

	var i interface{}
	switch ed.Format {
	case "%s":
		i = ed.StringValue
	case "%b":
		b, err := strconv.ParseBool(ed.StringValue)
		if err != nil {
			log.Error().Err(err).Send()
			i = ed.StringValue
		} else {
			i = b
		}
	default:
		log.Error().Str("format", ed.Format).Msg("unrecognized format")
		i = ed.StringValue
	}

	return i
}

type IndexEntryDefinitionValueExpressionTerm struct {
	TermType  int
	TermParam string
}

type IndexEntryDefinitionValueExpression struct {
	Terms []IndexEntryDefinitionValueExpressionTerm
}

type IndexEntryDefinition struct {
	Id           string                              `mapstructure:"id" yaml:"id" json:"id"`
	Name         string                              `mapstructure:"name" yaml:"name" json:"name"`
	NdxType      string                              `mapstructure:"type" yaml:"type" json:"type"`
	DataType     string                              `mapstructure:"data-type" yaml:"data-type" json:"data-type"`
	Format       string                              `mapstructure:"format" yaml:"format" json:"format"`
	Value        string                              `mapstructure:"value" yaml:"value" json:"value"`
	ValueExpr    IndexEntryDefinitionValueExpression `mapstructure:"-" yaml:"-" json:"-"`
	SourceFormat string                              `mapstructure:"source-format" yaml:"source-format" json:"source-format"`
	Required     bool                                `mapstructure:"required" yaml:"required"  json:"required"`
}

func (def *IndexEntryDefinition) postProcess(indexNumber int) error {

	var err error

	def.Id = fmt.Sprintf("i%d", indexNumber+1)

	// Set a few defaults in case have not been set.
	if def.DataType == "" {
		def.DataType = "string"
	} else if def.DataType == "date" {
		if def.SourceFormat == "" && def.Value != "sysdate" {
			err = errors.New("date datatype is missing the source-format definition")
		}
	}

	if def.Format == "" {
		def.Format = "%s"
	}

	if def.NdxType == "" {
		def.NdxType = "User"
	}

	return err
}

type IndexEntryValue struct {
	Id         string
	Name       string
	Value      string
	SourceName string
}

type DocClass struct {
	Id               string                    `mapstructure:"class-id" yaml:"class-id" json:"class-id"`
	Name             string                    `mapstructure:"name" yaml:"name"  json:"name"`
	ProducedClassIds []string                  `mapstructure:"produced-class-ids" yaml:"produced-class-ids"  json:"produced-class-ids"`
	Ext              string                    `mapstructure:"ext" yaml:"ext" json:"ext"`
	CodCliente       string                    `mapstructure:"cod-cliente" yaml:"cod-cliente" json:"cod-cliente"`
	MaxCpx           int                       `mapstructure:"max-cpx" yaml:"max-cpx" json:"max-cpx"`
	MaxDocs          int                       `mapstructure:"max-docs" yaml:"max-docs" json:"max-docs"`
	MaxSize          int                       `mapstructure:"max-size" yaml:"max-size" json:"max-size"`
	Platform         string                    `mapstructure:"platform" yaml:"platform" json:"platform"`
	Servizio         string                    `mapstructure:"servizio" yaml:"servizio" json:"servizio"`
	Procedura        string                    `mapstructure:"procedura" yaml:"procedura" json:"procedura"`
	Version          string                    `mapstructure:"version" yaml:"version" json:"version"`
	PackageLayout    string                    `mapstructure:"package-layout" yaml:"package-layout" json:"package-layout"`
	Index            []IndexEntryDefinition    `mapstructure:"index" yaml:"index" json:"index" json:"index"`
	OnExported       []OnExportEntryDefinition `mapstructure:"on-exported" yaml:"on-exported" json:"on-exported" json:"on-exported"`
	SqlQuery         string                    `mapstructure:"cos-query" yaml:"cos-query" json:"cos-query"`
	PackageTtl       int                       `mapstructure:"pkg-ttl" yaml:"pkg-ttl" json:"pkg-ttl"`
	DocumentTtl      int                       `mapstructure:"doc-ttl" yaml:"doc-ttl" json:"doc-ttl"`
	DistintaGED      bool                      `mapstructure:"distinta-ged" yaml:"distinta-ged" json:"distinta-ged"`
}

func ReadDocClassYMLDefinition(fn string, fileContent []byte) (DocClass, error) {

	const semLogContext = "doc-class-registry::read-doc-class-from-yaml"
	log.Info().Str("fn", fn).Msg(semLogContext)
	var err error

	dc := struct {
		DocClass DocClass `yaml:"docClass"`
	}{}
	err = yaml.Unmarshal(fileContent, &dc)
	if err != nil {
		return DocClass{}, err
	}

	// do some computation on the loaded data.
	err = dc.DocClass.Finalize()
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
	}

	return dc.DocClass, nil
}

func (dc *DocClass) MapToIndex(sourceMap map[string]interface{}) []IndexEntryValue {

	m := make([]IndexEntryValue, 0)

	for _, ndx := range dc.Index {
		v := IndexEntryValue{Id: ndx.Id, SourceName: ndx.Value, Value: ndx.EvaluateIndexEntry(sourceMap)}
		m = append(m, v)
	}

	return m
}

func (dc *DocClass) SqlQueryText() string {
	const semLogContext = "doc-class::sql-query-text"

	durs := []time.Duration{0, -24 * 1 * time.Hour, -24 * 2 * time.Hour, -24 * 3 * time.Hour, -24 * 4 * time.Hour}

	q := dc.SqlQuery
	now := time.Now().UTC()
	for i := 1; i <= 4; i++ {
		nowMinusXDays := now.Add(durs[i]).Format(time.RFC3339Nano)
		q = strings.Replace(q, fmt.Sprintf("{now-%dd}", i), nowMinusXDays, -1)
	}

	log.Trace().Str(SemLogDocClass, dc.Id).Str("query", q).Msg(semLogContext)
	return q
}

func (dc *DocClass) Finalize() error {

	dc.setDefaultValues()

	for i, _ := range dc.Index {
		err := dc.Index[i].postProcess(i)
		if err != nil {
			log.Error().Err(err).Str(SemLogDocClass, dc.Id).Str(SemLogFieldName, dc.Index[i].Name).Msg("error in post processing field")
		}

		expr, err := parseIndexEntryDefinitionExpression(dc.Index[i].Value)
		if err != nil {
			log.Error().Err(err).Str(SemLogDocClass, dc.Id).Str(SemLogFieldName, dc.Index[i].Name).Msg("error in evaluating field name value expression")
		} else {
			dc.Index[i].ValueExpr = expr
		}
	}

	for j, _ := range dc.OnExported {
		if dc.OnExported[j].FieldName == "" {
			fm, err := fielddictionary.GetMapping(dc.OnExported[j].Name)
			if err != nil {
				log.Error().Err(err).Str("doc-class", dc.Id).Str(SemLogFieldName, dc.OnExported[j].Name).Msg("error in resolving on-exported field name")
			} else {
				dc.OnExported[j].FieldName = fm.DocumentMapping
			}
		}
	}

	return nil
}

func (dc *DocClass) setDefaultValues() {

	switch dc.PackageLayout {
	case "v1":
	case "v2":
		log.Info().Str(SemLogDocClass, dc.Id).Msg("v2 package layout not supported.... reverting to v1")
		dc.PackageLayout = "v1"
	case "v3":
	default:
		dc.PackageLayout = "v1"
	}

	if dc.PackageTtl == 0 {
		dc.PackageTtl = PackageTTL
	}

	if dc.DocumentTtl == 0 {
		dc.DocumentTtl = DocumentTTL
	}

	if len(dc.ProducedClassIds) == 0 {
		dc.ProducedClassIds = make([]string, 0)
		dc.ProducedClassIds = append(dc.ProducedClassIds, dc.Id)
	}
}

func (dc *DocClass) ToJson() ([]byte, error) {

	b, err := json.Marshal(dc)
	if err != nil {
		log.Error().Err(err).Str(SemLogDocClass, dc.Id).Msg("error serializing doc-class")
		return b, err
	}

	return b, nil
}

func (dc *DocClass) GetPackageName(packageId string, refTime time.Time) (string, string, error) {

	var err error
	mida := packageId
	if dc.PackageLayout == "v1" || dc.PackageLayout == "v3" {
		mida, err = ComputeMida(dc.Platform, packageId)
	}

	// MAI: stringa formato cambiata da 20060102150405 --> 060102150405
	tmplData := map[string]string{
		"id":   packageId,
		"mida": mida,
		"zeta": dc.CodCliente,
		"ts":   refTime.Format("060102150405"),
	}

	tmpl := templateutil.MustParse([]templateutil.Info{{
		Name:    "doc-class-name",
		Content: dc.Name,
	}}, nil)

	n, err := templateutil.Process(tmpl, tmplData, false)
	if err != nil {
		return "", mida, err
	}

	return string(n), mida, nil
}

func parseIndexEntryDefinitionExpression(v string) (IndexEntryDefinitionValueExpression, error) {

	expr := IndexEntryDefinitionValueExpression{Terms: nil}
	vals := make([]IndexEntryDefinitionValueExpressionTerm, 0)

	sarr := strings.Split(v, ",")
	for _, s := range sarr {
		if strings.HasPrefix(s, "lookup:") {
			lookupName := s[len("lookup:"):]
			fm, err := fielddictionary.GetMapping(lookupName)
			if err != nil {
				log.Error().Err(err).Str("lookup-name", lookupName).Msg("lookup error for field name")
				return expr, err
			}

			vals = append(vals, IndexEntryDefinitionValueExpressionTerm{TermType: DocPropertyValue, TermParam: fm.DocumentMapping})
		}

		if strings.HasPrefix(s, "const:") {
			vals = append(vals, IndexEntryDefinitionValueExpressionTerm{TermType: ConstValue, TermParam: s[len("const:"):]})
		}

		if strings.HasPrefix(s, "sysdate") {
			vals = append(vals, IndexEntryDefinitionValueExpressionTerm{TermType: SysDate})
		}
	}

	expr.Terms = vals
	return expr, nil
}

func (def *IndexEntryDefinition) EvaluateIndexEntry(sourceMap map[string]interface{}) string {
	if !def.ValueExpr.IsEmpty() {
		return def.ValueExpr.EvaluateIndexEntry(def.DataType, def.Format, def.SourceFormat, sourceMap)
	}

	return ""
}

func (expr *IndexEntryDefinitionValueExpression) EvaluateIndexEntry(dtaType string, format string, sourceFormat string, sourceMap map[string]interface{}) string {

	r := make([]string, 0)
	for _, t := range expr.Terms {
		switch t.TermType {
		case DocPropertyValue:
			pv := util.GetProperty(sourceMap, t.TermParam)
			r = append(r, formatProperty(pv, dtaType, format, sourceFormat))
		case ConstValue:
			r = append(r, t.TermParam)
		case SysDate:
			r = append(r, formatTimeProperty(time.Now(), format))
		}
	}

	return strings.Join(r, "")
}

func (expr *IndexEntryDefinitionValueExpression) IsEmpty() bool {
	return len(expr.Terms) == 0
}

func formatProperty(v interface{}, dataType string, format string, sourceFormat string) string {

	if v == nil {
		return ""
	}

	s := fmt.Sprintf("%v", v)
	switch dataType {
	case "date":
		if s != "" {
			t, err := time.Parse(sourceFormat, s)
			if err != nil {
				log.Error().Err(err).Str("data-type", dataType).Str("format", format).Str("source-format", sourceFormat).Str("value", s /*fmt.Sprintf("%v", v)*/).Msg("error formatting property")
			} else {
				s = t.Format(format)
			}
		}
	case "number":
		if b, ok := v.(bool); ok {
			if b {
				s = "1"
			} else {
				s = "0"
			}
		}
	}

	return s
}

func formatTimeProperty(t time.Time, format string) string {
	return t.Format(format)
}
