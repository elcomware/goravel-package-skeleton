package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"
)

// FileSystem is an interface for file operations.
type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
	Rename(oldPath, newPath string) error
	Remove(filename string) error
	Stat(filename string) (os.FileInfo, error)
}

// RealFileSystem implements FileSystem using the real file system.
type RealFileSystem struct{}

func (fs RealFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (fs RealFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (fs RealFileSystem) Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func (fs RealFileSystem) Remove(filename string) error {
	return os.Remove(filename)
}

func (fs RealFileSystem) Stat(filename string) (os.FileInfo, error) {
	return os.Stat(filename)
}

// Commander is an interface for executing shell commands.
type Commander interface {
	Run(command string) string
}

// RealCommander implements Commander using the real shell.
type RealCommander struct{}

func (c RealCommander) Run(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

var (
	fs        FileSystem = RealFileSystem{}
	commander Commander  = RealCommander{}
)

// Ask prompts the user for input.
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

// confirm prompts the user for a yes/no confirmation.
func confirm(question string, defaultValue bool) bool {
	// Ask the user for input, with the default choices as the default value
	answer := ask(question, map[bool]string{true: "Y/n", false: "y/N"}[defaultValue])

	// If the user presses Enter without typing anything, return the default value
	if answer == "" {
		return defaultValue
	}
	if answer == "y/N" {
		return false
	}
	if answer == "Y/n" {
		return true
	}
	// Normalize the answer to lowercase and check if it starts with 'y'
	return strings.ToLower(answer)[0] == 'y'
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
		words[i] = cases.Title(language.English).String(word)
	}
	return strings.Join(words, "")
}

// replaceInFile replaces placeholders in a file with actual values.
func replaceInFile(file string, replacements map[string]string) {
	content, err := fs.ReadFile(file)
	if err != nil {
		writeln(fmt.Sprintf("Error reading file %s: %v\n", file, err))
		return
	}

	newContent := string(content)
	for oldText, newText := range replacements {
		newContent = strings.ReplaceAll(newContent, oldText, newText)
	}

	err = fs.WriteFile(file, []byte(newContent), 0644)
	if err != nil {
		writeln(fmt.Sprintf("Error writing file %s: %v\n", file, err))
	}
}

// renameFile renames a file.
func renameFile(oldPath, newPath string) {
	err := fs.Rename(oldPath, newPath)
	if err != nil {
		writeln(fmt.Sprintf("Error renaming file %s to %s: %v\n", oldPath, newPath, err))
	}
}

func removeFile(filename string) {
	if _, err := fs.Stat(filename); err == nil {
		err := fs.Remove(filename)
		if err != nil {
			writeln(fmt.Sprintf("Error removing file %s : %v\n", filename, err))
			return
		}
	}
}

func main() {
	initializePackage()
}

func initializePackage() {
	gitName := commander.Run("git config user.name")
	authorName := ask("Author name", gitName)

	gitEmail := commander.Run("git config user.email")
	authorEmail := ask("Author email", gitEmail)
	authorUsername := ask("Author username", commander.Run("git config user.username"))

	vendorName := ask("Vendor name", authorName)
	vendorUsername := ask("Vendor username", slugify(vendorName))
	vendorSlug := slugify(vendorUsername)

	vendorNamespace := titleCase(vendorName)
	vendorNamespace = ask("Vendor namespace", vendorNamespace)

	currentDirectory := commander.Run("basename $(pwd)")
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
		"config/packageName.go",
		"contracts/packageName.go",
		"facades/packageName.go",
		"commands/packageName.go",
		"packageName.go",
	}

	replacements := map[string]string{
		":author_name":      authorName,
		":author_username":  authorUsername,
		"author@domain.com": authorEmail,

		":vendor_name": vendorName,
		"vendorName":   vendorName,
		":vendor_slug": vendorSlug,
		"VendorName":   vendorNamespace,

		"PackageName":                  packageName,
		"packageName":                  packageSlug,
		":package_name":                packageName,
		":package_slug":                packageSlug,
		":package_slug_without_prefix": packageSlugWithoutPrefix,
		":package_description":         description,

		"Skeleton":             className,
		"migration_table_name": strings.ReplaceAll(packageSlug, "-", "_"),
		"variable":             variableName,
	}

	// Replace names
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

	// Rename files
	renameFile("config/packageName.go", fmt.Sprintf("config/%s.go", packageSlug))
	renameFile("commands/packageName.go", fmt.Sprintf("commands/%s.go", packageSlug))
	renameFile("packageName.go", fmt.Sprintf("%s.go", packageSlug))
	renameFile("service_provider.go", fmt.Sprintf("%s_service_provider.go", packageSlug))

	// Clean up
	if confirm("Execute `go mod tidy` and run tests?", true) {
		commander.Run("go mod tidy")
		commander.Run("go test ./...")
	}

	if confirm("Let this script delete itself?", true) {
		err := fs.Remove("setup.go")
		if err != nil {
			return
		}
	}

	writeln("Package initialized successfully!")
}
