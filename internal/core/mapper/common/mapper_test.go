package common

import (
	"reflect"
	"testing"
)

type (
	BasicDst struct {
		Name string
		Age  int64
	}

	PointerDst struct {
		Name *string
		Age  *int64
	}

	NestedStruct struct {
		Inner struct {
			Value string
		}
	}

	MapDst struct {
		ExtAttrs map[string]any
	}

	MixedDst struct {
		Name     string
		Details  *InnerDetails
		ExtAttrs map[string]any
	}

	InnerDetails struct {
		City *string
	}
)

func TestMapFields_BasicAssign(t *testing.T) {
	src := &struct {
		Name string
		Age  int
	}{
		Name: "Person123",
		Age:  25,
	}

	dst := &BasicDst{}

	fieldMap := map[string]string{
		"Name": "Name",
		"Age":  "Age",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != "Person123" {
		t.Fatalf("expected Name=Person123, got %s", dst.Name)
	}

	if dst.Age != 25 {
		t.Fatalf("expected Age=25, got %d", dst.Age)
	}
}

func TestMapFields_PointerAssign(t *testing.T) {
	src := &struct {
		Name string
		Age  int
	}{
		Name: "NYC",
		Age:  10,
	}

	dst := &PointerDst{}

	fieldMap := map[string]string{
		"Name": "Name",
		"Age":  "Age",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name == nil || *dst.Name != "NYC" {
		t.Fatalf("expected Name pointer to NYC")
	}

	if dst.Age == nil || *dst.Age != 10 {
		t.Fatalf("expected Age pointer to 10")
	}
}

func TestMapFields_PointerToNonPointer(t *testing.T) {
	name := "NYC"
	age := int64(10)
	src := &struct {
		Name *string
		Age  *int64
	}{
		Name: &name,
		Age:  &age,
	}

	dst := &BasicDst{}

	fieldMap := map[string]string{
		"Name": "Name",
		"Age":  "Age",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != "NYC" {
		t.Fatalf("expected Name=NYC, got %s", dst.Name)
	}

	if dst.Age != 10 {
		t.Fatalf("expected Age=10, got %d", dst.Age)
	}
}

func TestMapFields_NestedStruct(t *testing.T) {
	src := &struct {
		Value string
	}{
		Value: "hello",
	}

	dst := &NestedStruct{}

	fieldMap := map[string]string{
		"Value": "Inner.Value",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Inner.Value != "hello" {
		t.Fatalf("expected nested value hello, got %s", dst.Inner.Value)
	}
}

func TestMapFields_MapSimple(t *testing.T) {
	src := &struct {
		Site string
	}{
		Site: "NYC",
	}

	dst := &MapDst{}

	fieldMap := map[string]string{
		"Site": "ExtAttrs.Site",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.ExtAttrs["Site"] != "NYC" {
		t.Fatalf("expected ExtAttrs[Site]=NYC")
	}
}

func TestMapFields_MapNested(t *testing.T) {
	src := &struct {
		Site string
	}{
		Site: "NYC",
	}

	dst := &MapDst{}

	fieldMap := map[string]string{
		"Site": "ExtAttrs.Site.Value",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	site, ok := dst.ExtAttrs["Site"].(map[string]any)
	if !ok {
		t.Fatalf("expected Site to be map")
	}

	if site["Value"] != "NYC" {
		t.Fatalf("expected nested Value=NYC")
	}
}

func TestMapFields_MapAutoCreate(t *testing.T) {
	src := &struct {
		Key string
	}{
		Key: "val",
	}

	dst := &MapDst{
		ExtAttrs: nil, // ensure nil map handled
	}

	fieldMap := map[string]string{
		"Key": "ExtAttrs.Level1.Level2",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	level1 := dst.ExtAttrs["Level1"].(map[string]any)
	if level1["Level2"] != "val" {
		t.Fatalf("nested map creation failed")
	}
}

func TestMapFields_PointerNestedStruct(t *testing.T) {
	src := &struct {
		City string
	}{
		City: "Delhi",
	}

	dst := &MixedDst{}

	fieldMap := map[string]string{
		"City": "Details.City",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Details == nil {
		t.Fatalf("Details should be initialized")
	}

	if dst.Details.City == nil || *dst.Details.City != "Delhi" {
		t.Fatalf("expected pointer City=Delhi")
	}
}

func TestMapFields_ConvertibleTypes(t *testing.T) {
	src := &struct {
		Age int
	}{
		Age: 42,
	}

	dst := &BasicDst{}

	fieldMap := map[string]string{
		"Age": "Age",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Age != 42 {
		t.Fatalf("conversion int->int64 failed")
	}
}

func TestMapFields_IgnoreZeroValues(t *testing.T) {
	src := &struct {
		Name string
	}{
		Name: "", // zero value
	}

	dst := &BasicDst{
		Name: "existing",
	}

	fieldMap := map[string]string{
		"Name": "Name",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != "existing" {
		t.Fatalf("zero value should not overwrite existing")
	}
}

func TestMapFields_InvalidPath(t *testing.T) {
	src := &struct {
		Name string
	}{
		Name: "test",
	}

	dst := &BasicDst{}

	fieldMap := map[string]string{
		"Name": "Invalid.Field",
	}

	// should not panic
	_ = MapFields(src, dst, fieldMap)
}

func TestMapFields_InterfaceMap(t *testing.T) {
	src := &struct {
		Key string
	}{
		Key: "value",
	}

	dst := &struct {
		Data map[string]any
	}{}

	fieldMap := map[string]string{
		"Key": "Data.Level1.Level2",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	level1 := dst.Data["Level1"].(map[string]any)
	if level1["Level2"] != "value" {
		t.Fatalf("interface map nesting failed")
	}
}

func TestMapFields_DeepCompare(t *testing.T) {
	src := &struct {
		Site string
	}{
		Site: "NYC",
	}

	dst := &MapDst{}

	fieldMap := map[string]string{
		"Site": "ExtAttrs.Site.Value",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := map[string]any{
		"Site": map[string]any{
			"Value": "NYC",
		},
	}

	if !reflect.DeepEqual(dst.ExtAttrs, expected) {
		t.Fatalf("expected %+v, got %+v", expected, dst.ExtAttrs)
	}
}

func TestMapFields_MapToMap_SameType(t *testing.T) {
	src := &struct {
		Data map[string]string
	}{
		Data: map[string]string{"k": "v"},
	}

	dst := &struct {
		Data map[string]string
	}{}

	fieldMap := map[string]string{
		"Data": "Data",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Data["k"] != "v" {
		t.Fatalf("expected map value copied")
	}
}

func TestMapFields_MapToPointerMap(t *testing.T) {
	src := &struct {
		Data map[string]string
	}{
		Data: map[string]string{"k": "v"},
	}

	dst := &struct {
		Data *map[string]string
	}{}

	fieldMap := map[string]string{
		"Data": "Data",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Data == nil || (*dst.Data)["k"] != "v" {
		t.Fatalf("expected pointer map to be assigned")
	}
}

func TestMapFields_MapIncompatibleTypes(t *testing.T) {
	src := &struct {
		Data map[string]string
	}{
		Data: map[string]string{"k": "abc"},
	}

	dst := &struct {
		Data map[string]int
	}{}

	fieldMap := map[string]string{
		"Data": "Data",
	}

	err := MapFields(src, dst, fieldMap)
	if err == nil {
		t.Fatalf("expected error for incompatible map types")
	}
}

func TestMapFields_GetByPathNestedStruct(t *testing.T) {
	src := &struct {
		Nios struct {
			Creator string
		}
	}{
		Nios: struct{ Creator string }{Creator: "admin"},
	}

	dst := &struct {
		Creator string
	}{}

	fieldMap := map[string]string{
		"Nios.Creator": "Creator",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Creator != "admin" {
		t.Fatalf("expected nested source mapping to work")
	}
}

func TestMapFields_GetByPathMapNested(t *testing.T) {
	src := &struct {
		ExtAttrs map[string]any
	}{
		ExtAttrs: map[string]any{
			"Site": map[string]any{
				"Value": "BLR",
			},
		},
	}

	dst := &struct {
		Site string
	}{}

	fieldMap := map[string]string{
		"ExtAttrs.Site.Value": "Site",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Site != "BLR" {
		t.Fatalf("expected nested map read via getByPath")
	}
}

func TestMapFields_GetByPathMissingField(t *testing.T) {
	src := &struct {
		Name string
	}{
		Name: "abc",
	}

	dst := &BasicDst{}

	fieldMap := map[string]string{
		"Invalid.Field": "Name",
	}

	// should not panic
	_ = MapFields(src, dst, fieldMap)
}

func TestMapFields_NilPointerSource(t *testing.T) {
	src := &struct {
		Name *string
	}{
		Name: nil,
	}

	dst := &PointerDst{}

	fieldMap := map[string]string{
		"Name": "Name",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != nil {
		t.Fatalf("nil source pointer should not set destination")
	}
}

func TestMapFields_AutoMapSameName(t *testing.T) {
	src := &struct {
		Name    string
		Comment string
	}{
		Name:    "test",
		Comment: "auto-mapped",
	}

	dst := &struct {
		Name    string
		Comment string
	}{}

	// Only map Name explicitly, Comment should auto-map
	fieldMap := map[string]string{
		"Name": "Name",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != "test" {
		t.Fatalf("expected Name=test, got %s", dst.Name)
	}

	if dst.Comment != "auto-mapped" {
		t.Fatalf("expected Comment to be auto-mapped, got %s", dst.Comment)
	}
}

func TestMapFields_AutoMapSkipsExplicitlyMapped(t *testing.T) {
	src := &struct {
		Name string
	}{
		Name: "original",
	}

	dst := &struct {
		Name      string
		OtherName string
	}{}

	// Explicitly map Name to OtherName
	fieldMap := map[string]string{
		"Name": "OtherName",
	}

	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.OtherName != "original" {
		t.Fatalf("expected OtherName=original, got %s", dst.OtherName)
	}

	// Name should NOT be auto-mapped since it was explicitly mapped
	if dst.Name != "" {
		t.Fatalf("expected Name to be empty, got %s", dst.Name)
	}
}

func TestMapFields_NilNestedStructSkipsGracefully(t *testing.T) {
	src := &struct {
		NIOS *struct {
			Creator string
		}
	}{
		NIOS: nil, // nil nested struct
	}

	dst := &struct {
		Creator string
	}{
		Creator: "existing",
	}

	fieldMap := map[string]string{
		"NIOS.Creator": "Creator",
	}

	// Should not error, just skip
	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should remain unchanged
	if dst.Creator != "existing" {
		t.Fatalf("expected Creator to remain existing, got %s", dst.Creator)
	}
}

func TestMapFields_AutoMapSkipsMissingDstField(t *testing.T) {
	src := &struct {
		Name  string
		Extra string // Not in dst
	}{
		Name:  "test",
		Extra: "ignored",
	}

	dst := &struct {
		Name string
	}{}

	fieldMap := map[string]string{}

	// Should not error even though Extra doesn't exist in dst
	err := MapFields(src, dst, fieldMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if dst.Name != "test" {
		t.Fatalf("expected Name=test, got %s", dst.Name)
	}
}
