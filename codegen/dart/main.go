package main

import (
	"github.com/tada-team/tdproto/codegen"
)

const dartHeader = `
import 'package:freezed_annotation/freezed_annotation.dart';`

const dartClassTemplate = `

`

func generateDart(tdprotoInfo *codegen.TdInfo) {

}

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(err)
	}

	generateDart(tdprotoInfo)
}
