package main

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"tpm-cpx-docclass/docclass/crawlingrules"
	"tpm-cpx-docclass/docclass/fielddictionary"
	"tpm-cpx-docclass/docclass/registry"
)

func main() {

	const semLogContext = "sql-scripts::main"

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fmt.Printf("-- delete tables \n")
	fmt.Printf("delete from cpx_registry;\n")
	fmt.Printf("delete from cpx_out_action;\n")
	fmt.Printf("delete from cpx_ndx_item;\n")
	fmt.Printf("delete from cpx_doc_class;\n")
	fmt.Printf("delete from cpx_field_dictionary;\n")

	cfgs, err := crawlingrules.ReadYamlConfig(crawlingrules.RulesConfigYaml)
	if err != nil {
		log.Fatal().Err(err).Msg(semLogContext)
	}

	fmt.Printf("--- Begin of insert into cpx_registry\n")
	for _, cfg := range cfgs {
		fmt.Printf("insert into cpx_registry (name, class_id, enabled) values('%s', '%s', %t);\n", "default", cfg.Name, !cfg.Disabled)
	}
	fmt.Printf("--- End of insert into cpx_registry: num-inserts=%d\n", len(cfgs))

	fd, err := fielddictionary.ReadDictionaryFromYamlData(fielddictionary.EmbeddedDictionary)
	if err != nil {
		log.Fatal().Err(err).Msg(semLogContext)
	}

	fmt.Printf("--- Begin of insert into cpx_field_dictionary\n")
	for n, mapping := range fd {
		fmt.Printf("insert into cpx_field_dictionary(name, document_mapping) values('%s', '%s');\n", n, mapping.DocumentMapping)
	}
	fmt.Printf("--- End of insert into cpx_field_dictionary: num-inserts=%d\n", len(fd))

	numEntries, err := registry.ReadRegistryFromEmbeddedData(registry.EmbeddedDocClassFS, registry.EmbeddedDocClassFSRootPath)
	if err != nil {
		log.Fatal().Err(err).Msg(semLogContext)
	}

	log.Info().Int("num entries", numEntries).Msg(semLogContext)

	r := registry.GetRegistry()
	fmt.Printf("--- Begin of insert into cpx_doc_class\n")
	for n, dc := range r {
		classId := n
		name := ToSqlString(dc.Name)
		extension := ToSqlString(dc.Ext)
		codCliente := ToSqlString(dc.CodCliente)
		maxCpx := fmt.Sprintf("%d", dc.MaxCpx)
		maxDocs := fmt.Sprintf("%d", dc.MaxDocs)
		maxSize := fmt.Sprintf("%d", dc.MaxSize)
		platform := ToSqlString(dc.Platform)
		servizio := ToSqlString(dc.Servizio)
		procedura := ToSqlString(dc.Procedura)
		version := ToSqlString(dc.Version)
		layout := ToSqlString(dc.PackageLayout)
		distintaGED := fmt.Sprintf("%t", dc.DistintaGED)
		query := ToSqlString(util.StripDuplicateWhiteSpaces(dc.SqlQuery))
		fmt.Printf("insert into cpx_doc_class (class_id, name, childIds, extension, cod_cliente, max_cpx, max_docs, max_size, platform, servizio, procedura, version, pkg_layout, sql_query, distinta_ged) values(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) \n",
			classId, name, "NULL", extension, codCliente, maxCpx, maxDocs, maxSize, platform, servizio, procedura, version, layout, query, distintaGED)
	}
	fmt.Printf("--- End of insert into cpx_field_dictionary: num-inserts=%d\n", len(r))

	fmt.Printf("--- Begin of insert into cpx_ndx_item\n")
	totalRows := 0
	for n, dc := range r {
		fmt.Printf("--- Begin of insert into cpx_ndx_item for class %s\n", n)
		for _, x := range dc.Index {
			classId := n
			ndxId := ToSqlString(x.Id)
			name := ToSqlString(x.Name)
			typ := ToSqlString(x.NdxType)
			dataTyp := ToSqlString(x.DataType)

			format := "NULL"
			if x.Format != "" {
				format = ToSqlString(x.Format)
			}

			value := ToSqlString(x.Value)
			sourceFormat := "NULL"
			if x.SourceFormat != "" {
				sourceFormat = ToSqlString(x.SourceFormat)
			}
			required := fmt.Sprintf("%t", x.Required)
			fmt.Printf("insert into cpx_ndx_item (class_id, ndx_id, name, type, data_type, format, value, source_format, required) values(%s, %s, %s, %s, %s, %s, %s, %s, %s) \n",
				classId, ndxId, name, typ, dataTyp, format, value, sourceFormat, required)
		}
		fmt.Printf("--- End of insert into cpx_ndx_item for class %s: num-inserts=%d\n", n, len(dc.Index))
		totalRows += len(dc.Index)
	}
	fmt.Printf("--- End of insert into cpx_ndx_item: num-inserts=%d\n", totalRows)

}

func ToSqlString(s string) string {
	var sb strings.Builder
	sb.WriteRune('\'')
	for _, c := range s {
		sb.WriteRune(c)
		if c == '\'' {
			sb.WriteRune(c)
		}
	}
	sb.WriteRune('\'')
	return sb.String()
}
