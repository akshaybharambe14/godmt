package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/godmt/pkg/godmt"

	"github.com/iancoleman/strcase"
)

func isEmbeddedStructForInheritance(field godmt.ScannedStructField) bool {
	return field.Kind == "struct" && field.Tag == ""
}

func getTypescriptCompatibleType(goType string) string {
	result, ok := goTypeScriptTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

func getSwiftCompatibleType(goType string) string {
	result, ok := goSwiftTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

func getRecordType(goMapType string) string {
	var result string

	result = strings.ReplaceAll(goMapType, "map[", "")
	recordTypes := strings.Split(result, "]")

	for i := range recordTypes {
		recordTypes[i] = getTypescriptCompatibleType(recordTypes[i])
	}

	result = strings.Join(recordTypes, ", ")

	return fmt.Sprintf("Record<%s>", result)
}

func getDictionaryType(goMapType string) string {
	var result string

	result = strings.ReplaceAll(goMapType, "map[", "")
	recordTypes := strings.Split(result, "]")

	for i := range recordTypes {
		recordTypes[i] = getSwiftCompatibleType(recordTypes[i])
	}

	result = strings.Join(recordTypes, ", ")

	return fmt.Sprintf("Dictionary<%s>", result)
}

func mapValuesToTypeScriptRecord(rawMap map[string]string) string {
	var entries []string
	for i := range rawMap {
		entries = append(entries, fmt.Sprintf("\t%s: %s", i, rawMap[i]))
	}

	return strings.Join(entries, ",\n")
}

func transformSliceTypeToTypeScript(rawSliceType string) string {
	var result string

	result = strings.ReplaceAll(rawSliceType, "[]", "")
	return fmt.Sprintf("%s[]", getTypescriptCompatibleType(result))
}

func transformSliceTypeToSwift(rawSliceType string) string {
	var result string

	result = strings.ReplaceAll(rawSliceType, "[]", "")
	return fmt.Sprintf("[%s]", getSwiftCompatibleType(result))
}

func toCamelCase(raw string) string {
	return strcase.ToLowerCamel(raw)
}
