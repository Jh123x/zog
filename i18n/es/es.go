package es

import (
	"github.com/Oudwins/zog/zconst"
)

var Map zconst.LangMap = map[zconst.ZogType]map[zconst.ZogIssueCode]string{
	zconst.TypeString: {
		zconst.IssueCodeRequired:        "Es obligatorio",
		zconst.IssueCodeNotNil:          "No debe estar vacio",
		zconst.IssueCodeMin:             "Cadena debe contener al menos {{min}} caracter(es)",
		zconst.IssueCodeMax:             "Cadena debe contener como máximo {{max}} caracter(es)",
		zconst.IssueCodeLen:             "Cadena debe tener exactamente {{len}} caracter(es)",
		zconst.IssueCodeEmail:           "Debe ser un correo electrónico válido",
		zconst.IssueCodeUUID:            "Debe ser un UUID válido",
		zconst.IssueCodeMatch:           "Cadena no es válida",
		zconst.IssueCodeURL:             "Debe ser una URL válida",
		zconst.IssueCodeHasPrefix:       "Cadena debe comenzar con {{prefix}}",
		zconst.IssueCodeHasSuffix:       "Cadena debe terminar con {{suffix}}",
		zconst.IssueCodeContains:        "Cadena debe contener {{contained}}",
		zconst.IssueCodeContainsDigit:   "Cadena debe contener al menos un dígito",
		zconst.IssueCodeContainsUpper:   "Cadena debe contener al menos una letra mayúscula",
		zconst.IssueCodeContainsLower:   "Cadena debe contener al menos una letra minúscula",
		zconst.IssueCodeContainsSpecial: "Cadena debe contener al menos un carácter especial",
		zconst.IssueCodeOneOf:           "Cadena debe ser una de las siguientes: {{one_of_options}}",
		zconst.IssueCodeFallback:        "Cadena no es válida",
	},
	zconst.TypeBool: {
		zconst.IssueCodeRequired: "Es obligatorio",
		zconst.IssueCodeNotNil:   "No debe estar vacio",
		zconst.IssueCodeTrue:     "Debe ser verdadero",
		zconst.IssueCodeFalse:    "Debe ser falso",
		zconst.IssueCodeFallback: "Valor no es válido",
	},
	zconst.TypeNumber: {
		zconst.IssueCodeRequired: "Es obligatorio",
		zconst.IssueCodeNotNil:   "No debe estar vacio",
		zconst.IssueCodeLTE:      "Número debe ser menor o igual a {{lte}}",
		zconst.IssueCodeLT:       "Número debe ser menor que {{lt}}",
		zconst.IssueCodeGTE:      "Número debe ser mayor o igual a {{gte}}",
		zconst.IssueCodeGT:       "Número debe ser mayor que {{gt}}",
		zconst.IssueCodeEQ:       "Número debe ser igual a {{eq}}",
		zconst.IssueCodeOneOf:    "Número debe ser uno de los siguientes: {{options}}",
		zconst.IssueCodeFallback: "Número no es válido",
	},
	zconst.TypeTime: {
		zconst.IssueCodeRequired: "Es obligatorio",
		zconst.IssueCodeNotNil:   "No debe estar vacio",
		zconst.IssueCodeAfter:    "Fecha debe ser posterior a {{after}}",
		zconst.IssueCodeBefore:   "Fecha debe ser anterior a {{before}}",
		zconst.IssueCodeEQ:       "Fecha debe ser igual a {{eq}}",
		zconst.IssueCodeFallback: "Fecha no es válida",
	},
	zconst.TypeSlice: {
		zconst.IssueCodeRequired: "Es obligatorio",
		zconst.IssueCodeNotNil:   "No debe estar vacio",
		zconst.IssueCodeMin:      "Lista debe contener al menos {{min}} elementos",
		zconst.IssueCodeMax:      "Lista debe contener como máximo {{max}} elementos",
		zconst.IssueCodeLen:      "Lista debe contener exactamente {{len}} elementos",
		zconst.IssueCodeContains: "Lista debe contener {{contained}}",
		zconst.IssueCodeFallback: "Lista no es válida",
	},
	zconst.TypeStruct: {
		zconst.IssueCodeRequired: "Es obligatorio",
		zconst.IssueCodeNotNil:   "No debe estar vacio",
		zconst.IssueCodeFallback: "Estructura no es válida",
		// JSON
		zconst.IssueCodeInvalidJSON: "JSON no válido",
		// ZHTTP ISSUES
		zconst.IssueCodeZHTTPInvalidForm:  "Formulario no válido",
		zconst.IssueCodeZHTTPInvalidQuery: "Parámetros de consulta no válidos",
	},
}
