package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Provider name and prefix on every object (e.g. infoblox_record_a).
const providerName = "infoblox"

// Backends emitted in order; skipped when no example file exists.
var backends = []struct {
	dir   string // example path segment
	label string // heading in the docs
}{
	{dir: "nios", label: "NIOS Backend"},
	{dir: "uddi", label: "UDDI Backend"},
}

// kind is a documentable object category (resources / data-sources / ...).
type kind struct {
	dir         string // docs/<dir> and examples/<dir>
	exampleFile string // example file name inside each backend folder
}

var kinds = []kind{
	{dir: "resources", exampleFile: "resource.tf"},
	{dir: "data-sources", exampleFile: "data-source.tf"},
	{dir: "list-resources", exampleFile: "list-resource.tfquery.hcl"},
}

// Subcategory labels for known groups. Unknown groups are upper-cased.
var subcategoryLabels = map[string]string{
	"acl":              "ACL",
	"cloud":            "CLOUD",
	"dhcp":             "DHCP",
	"discovery":        "DISCOVERY",
	"dns":              "DNS",
	"dtc":              "DTC",
	"federatedrealms":  "FEDERATED REALMS",
	"grid":             "GRID",
	"ipam":             "IPAM",
	"microsoft":        "MICROSOFT",
	"microsoftserver":  "MICROSOFT SERVER",
	"misc":             "MISC",
	"notification":     "NOTIFICATION",
	"parentalcontrol":  "PARENTAL CONTROL",
	"rir":              "RIR",
	"rpz":              "RPZ",
	"security":         "SECURITY",
	"smartfolder":      "SMART FOLDER",
	"threatinsight":    "THREAT INSIGHT",
	"threatprotection": "THREAT PROTECTION",
}

func subcategory(group string) string {
	if label, ok := subcategoryLabels[group]; ok {
		return label
	}
	return strings.ToUpper(group)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gendocs:", err)
		os.Exit(1)
	}
}

func run() error {
	total := 0
	for _, k := range kinds {
		n, err := processKind(k)
		if err != nil {
			return err
		}
		total += n
	}
	fmt.Printf("gendocs: post-processed %d doc page(s)\n", total)
	return nil
}

// processKind updates every docs/<kind>/*.md in place.
func processKind(k kind) (int, error) {
	docsDir := filepath.Join("docs", k.dir)
	entries, err := os.ReadDir(docsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	count := 0
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		short := strings.TrimSuffix(e.Name(), ".md")
		group := lookupGroup(short)
		if group == "" {
			fmt.Printf("gendocs: %s/%s.md (no group found, left as-is)\n", k.dir, short)
			continue
		}

		docPath := filepath.Join(docsDir, e.Name())
		raw, err := os.ReadFile(docPath)
		if err != nil {
			return count, err
		}

		out := injectSubcategory(string(raw), subcategory(group))
		out = injectExamples(out, k, group, short)

		if err := os.WriteFile(docPath, []byte(out), 0o644); err != nil {
			return count, err
		}
		fmt.Printf("gendocs: updated %s (%s)\n", docPath, subcategory(group))
		count++
	}
	return count, nil
}

// Manual group overrides for pages that don't match the service layout
// (e.g. plural doc name vs. singular Go file name).
var shortGroupOverrides = map[string]string{
	"next_available_ips":            "ipam",
	"next_available_subnets":        "ipam",
	"next_available_address_blocks": "ipam",
}

// lookupGroup returns the internal/service/<group> that owns this object.
// Objects renamed via 'tfName' keep an unseparated Go file name
// (network_container -> networkcontainer_*.go), so retry without underscores.
func lookupGroup(short string) string {
	if g, ok := shortGroupOverrides[short]; ok {
		return g
	}
	if g := findGroup(short); g != "" {
		return g
	}
	return findGroup(strings.ReplaceAll(short, "_", ""))
}

// findGroup returns the group holding a <base><suffix>.go file, or "".
func findGroup(base string) string {
	serviceRoot := filepath.Join("internal", "service")
	groups, err := os.ReadDir(serviceRoot)
	if err != nil {
		return ""
	}
	suffixes := []string{"_resource.go", "_data_source.go", "_list.go"}
	for _, g := range groups {
		if !g.IsDir() {
			continue
		}
		for _, sfx := range suffixes {
			if _, err := os.Stat(filepath.Join(serviceRoot, g.Name(), base+sfx)); err == nil {
				return g.Name()
			}
		}
	}
	return ""
}

// injectSubcategory adds a `subcategory:` line right after page_title,
// unless the page already has one.
func injectSubcategory(doc, label string) string {
	if strings.Contains(doc, "\nsubcategory:") {
		return doc
	}
	lines := strings.Split(doc, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "page_title:") {
			injected := append([]string{}, lines[:i+1]...)
			injected = append(injected, fmt.Sprintf("subcategory: %q", label))
			injected = append(injected, lines[i+1:]...)
			return strings.Join(injected, "\n")
		}
	}
	return doc
}

// injectExamples inserts an "## Example Usage" section (one block per backend
// with an example file) just before the schema. No-op if the page already
// has one or if no examples exist.
func injectExamples(doc string, k kind, group, short string) string {
	if strings.Contains(doc, "## Example Usage") {
		return doc
	}

	full := providerName + "_" + short
	var sec strings.Builder
	sec.WriteString("## Example Usage\n")
	found := false
	for _, b := range backends {
		exPath := filepath.Join("examples", k.dir, group, full, b.dir, k.exampleFile)
		content, err := os.ReadFile(exPath)
		if err != nil {
			continue
		}
		found = true
		fmt.Fprintf(&sec, "\n### %s\n\n```terraform\n%s\n```\n",
			b.label, strings.TrimRight(string(content), "\n"))
	}
	if !found {
		return doc
	}
	sec.WriteString("\n")

	// Insert before the schema marker, or the first "## " heading if missing.
	lines := strings.Split(doc, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "<!-- schema generated by tfplugindocs -->") ||
			strings.HasPrefix(line, "## ") {
			out := append([]string{}, lines[:i]...)
			out = append(out, sec.String())
			out = append(out, lines[i:]...)
			return strings.Join(out, "\n")
		}
	}
	return doc + "\n" + sec.String()
}
