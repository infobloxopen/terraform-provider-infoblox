package utils

import (
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

const (
	NaiveDatetimeLayout string = "2006-01-02T15:04:05"
)

// DataSourceResultAttributes combines resource attributes into datasource attributes.
// Converts all fields to computed for datasource use.
func DataSourceResultAttributes(attrs map[string]resourceschema.Attribute) map[string]datasourceschema.Attribute {
	result := make(map[string]datasourceschema.Attribute)
	for k, v := range attrs {
		result[k] = toDataSourceAttribute(v)
	}
	return result
}

func toDataSourceAttribute(val resourceschema.Attribute) datasourceschema.Attribute {
	switch a := val.(type) {
	case resourceschema.StringAttribute:
		return datasourceschema.StringAttribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.BoolAttribute:
		return datasourceschema.BoolAttribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Int32Attribute:
		return datasourceschema.Int32Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Int64Attribute:
		return datasourceschema.Int64Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.Float64Attribute:
		return datasourceschema.Float64Attribute{Computed: true, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.MapAttribute:
		return datasourceschema.MapAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.ListAttribute:
		return datasourceschema.ListAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.SetAttribute:
		return datasourceschema.SetAttribute{Computed: true, ElementType: a.ElementType, MarkdownDescription: a.MarkdownDescription}
	case resourceschema.ListNestedAttribute:
		return datasourceschema.ListNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	case resourceschema.SetNestedAttribute:
		return datasourceschema.SetNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	case resourceschema.SingleNestedAttribute:
		return datasourceschema.SingleNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			Attributes:          nestedAttrsToDataSource(a.Attributes),
		}
	case resourceschema.MapNestedAttribute:
		return datasourceschema.MapNestedAttribute{
			Computed:            true,
			MarkdownDescription: a.MarkdownDescription,
			NestedObject: datasourceschema.NestedAttributeObject{
				Attributes: nestedAttrsToDataSource(a.NestedObject.Attributes),
			},
		}
	default:
		return nil
	}
}

func nestedAttrsToDataSource(attrs map[string]resourceschema.Attribute) map[string]datasourceschema.Attribute {
	result := make(map[string]datasourceschema.Attribute)
	for k, v := range attrs {
		result[k] = toDataSourceAttribute(v)
	}
	return result
}
