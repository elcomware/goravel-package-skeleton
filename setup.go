package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	gitName := run("git config user.name")
	authorName := ask("Author name", gitName)

	gitEmail := run("git config user.email")
	authorEmail := ask("Author email", gitEmail)
	authorUsername := ask("Author username", run("git config user.username"))

	vendorName := ask("Vendor name", authorName)
	vendorUsername := ask("Vendor username", slugify(vendorName))
	vendorSlug := slugify(vendorUsername)

	vendorNamespace := titleCase(vendorName)
	vendorNamespace = ask("Vendor namespace", vendorNamespace)

	currentDirectory := run("basename $(pwd)")
	packageName := ask("Package name", currentDirectory)
	packageSlug := slugify(packageName)
	packageSlugWithoutPrefix := strings.TrimPrefix(packageSlug, "goravel-")

	className := titleCase(packageName)
	className = ask("Class name", className)
	variableName := string(unicode.ToLower(rune(className[0]))) + className[1:]
	description := ask("Package description", fmt.Sprintf("This is my package %s", packageSlug))

	useGoLint := confirm("Enable Go Lint?", true)
	useDependabot := confirm("Enable Dependabot?", true)
	useChangelogWorkflow := confirm("Use automatic changelog updater workflow?", true)

	writeln("------")
	writeln(fmt.Sprintf("Author     : %s (%s, %s)", authorName, authorUsername, authorEmail))
	writeln(fmt.Sprintf("Vendor     : %s (%s)", vendorName, vendorSlug))
	writeln(fmt.Sprintf("Package    : %s <%s>", packageSlug, description))
	writeln(fmt.Sprintf("Namespace  : %s\\%s", vendorNamespace, className))
	writeln(fmt.Sprintf("Class name : %s", className))
	writeln("---")
	writeln("Packages & Utilities")
	writeln(fmt.Sprintf("Use Go Lint          : %v", useGoLint))
	writeln(fmt.Sprintf("Use Dependabot       : %v", useDependabot))
	writeln(fmt.Sprintf("Use Auto-Changelog   : %v", useChangelogWorkflow))
	writeln("------")

	writeln("This script will replace the above values in all relevant files in the project directory.")

	if !confirm("Modify files?", true) {
		os.Exit(1)
	}

	files := []string{
		"README.md",
		"go.mod",
		"main.go",
		"config/config.go",
		"routes/web.go",
	}

	replacements := map[string]string{
		":author_name":                 authorName,
		":author_username":             authorUsername,
		"author@domain.com":            authorEmail,
		":vendor_name":                 vendorName,
		"vendorName":                   vendorName,
		":vendor_slug":                 vendorSlug,
		"VendorName":                   vendorNamespace,
		":package_name":                packageName,
		"packageName":                  packageName,
		":package_slug":                packageSlug,
		":package_slug_without_prefix": packageSlugWithoutPrefix,
		"Skeleton":                     className,
		"skeleton":                     packageSlug,
		"migration_table_name":         strings.ReplaceAll(packageSlug, "-", "_"),
		"variable":                     variableName,
		":package_description":         description,
	}

	for _, file := range files {
		replaceInFile(file, replacements)
	}

	if !useGoLint {
		removeFile(".golangci.yml")
		removeFile(".github/workflows/golangci-lint.yml")
	}

	if !useDependabot {
		removeFile(".github/dependabot.yml")
		removeFile(".github/workflows/dependabot-auto-merge.yml")
	}

	if !useChangelogWorkflow {
		removeFile(".github/workflows/update-changelog.yml")
	}

	if confirm("Execute `go mod tidy` and run tests?", true) {
		run("go mod tidy")
		run("go test ./...")
	}

	if confirm("Let this script delete itself?", true) {
		err := os.Remove(os.Args[0])
		if err != nil {
			return
		}
	}
}

func run(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func ask(question string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (%s): ", question, defaultValue)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	if answer == "" {
		return defaultValue
	}
	return answer
}

func confirm(question string, defaultValue bool) bool {
	answer := ask(question, map[bool]string{true: "Y/n", false: "y/N"}[defaultValue])
	if answer == "" {
		return defaultValue
	}
	return strings.ToLower(answer) == "y"
}

func writeln(line string) {
	fmt.Println(line)
}

func slugify(subject string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9-]+")
	return strings.ToLower(reg.ReplaceAllString(subject, "-"))
}

func titleCase(subject string) string {
	subject = strings.ReplaceAll(subject, "-", " ")
	subject = strings.ReplaceAll(subject, "_", " ")
	words := strings.Fields(subject)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

func replaceInFile(file string, replacements map[string]string) {
	content, err := os.ReadFile(file)
	if err != nil {
		return
	}

	newContent := string(content)
	for oldText, newText := range replacements {
		newContent = strings.ReplaceAll(newContent, oldText, newText)
	}

	err = os.WriteFile(file, []byte(newContent), 0644)
	if err != nil {
		return
	}
}

func removeFile(filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			return
		}
	}
}
