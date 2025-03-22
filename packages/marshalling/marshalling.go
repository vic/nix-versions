package marshalling

import (
	"bufio"
	"os"
	"path"
	"regexp"
	"strings"

	lib "github.com/vic/nix-versions/packages/versions"
)

var (
	RE_NIX_TOOL              *regexp.Regexp
	RE_NIX_INSTALLABLE       *regexp.Regexp
	RE_NIX_INSTALLABLE_SHORT *regexp.Regexp
	TOOL_FILE_READERS        map[string](func(string) (*string, error))
	TOOLS_FILE_READERS       map[string](func(string) ([]string, error))
)

func init() {
	RE_NIX_TOOL = regexp.MustCompile(`^([^ #]+[^ ]+)[ ]*(# ([^ #]+)@([^ #]+))?`)
	RE_NIX_INSTALLABLE = regexp.MustCompile(`^([^#]+/)([^#]+)#([^# ]+)(#([^ #]+)#([^ #]+))?`)
	RE_NIX_INSTALLABLE_SHORT = regexp.MustCompile(`^([^#/]+)#([^# ]+)(#([^ #]+)#([^ #]+))?`)

	TOOL_FILE_READERS = map[string](func(string) (*string, error)){
		".node-version": readConstraint,
		".java-version": readConstraint,
		".ruby-version": readConstraint,
	}

	TOOLS_FILE_READERS = map[string](func(string) ([]string, error)){
		".nix-tools":     readNixTools,
		".tool-versions": readAsdfToolVersions,
	}
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func readConstraint(path string) (*string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	constraint := strings.TrimSpace(string(bytes))
	return &constraint, nil
}

func readAsdfToolVersions(file string) ([]string, error) {
	return nil, nil
}

func readNixTools(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var names []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := RE_NIX_TOOL.FindStringSubmatch(line)
		if len(match) > 1 {
			name := match[1]
			if IsInstallable(name) {
				if len(match) > 4 && match[4] != "" {
					name = name + "#" + match[3] + "#" + match[4]
				}
				names = append(names, name)
			} else {
				read, err := ReadPackagesFromFile(name)
				if err != nil {
					return nil, err
				}
				names = append(names, read...)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return names, nil
}

func ReadPackagesFromFile(arg string) ([]string, error) {
	if strings.Contains(arg, "@") && isFile(arg[strings.Index(arg, "@")+1:]) {
		name := arg[:strings.Index(arg, "@")]
		file := arg[strings.Index(arg, "@")+1:]

		tools_reader, is_tools_file := TOOLS_FILE_READERS[path.Base(file)]
		tool_reader, is_tool_file := TOOL_FILE_READERS[path.Base(file)]

		if name == "" {
			if is_tools_file {
				return tools_reader(file)
			}
			return readNixTools(file)
		}

		if is_tool_file {
			constraint, err := tool_reader(file)
			if err != nil {
				return nil, err
			}
			return []string{name + "@" + *constraint}, nil
		}

		tools_reader, is_tools_file = TOOLS_FILE_READERS[name]
		if is_tools_file {
			return tools_reader(file)
		}

		constraint, err := readConstraint(file)
		if err != nil {
			return nil, err
		}
		return []string{name + "@" + *constraint}, nil
	}
	return []string{arg}, nil
}

func IsSimpleName(str string) bool {
	return !strings.ContainsAny(str, ":/#@ ")
}

func IsInstallable(str string) bool {
	return strings.ContainsAny(str, ":/#") && !strings.HasPrefix(str, "bin/")
}

func FromInstallableStr(str string) lib.Version {
	var (
		flake    string
		revision string
		attr     string
		version  string
		name     string
	)
	match := RE_NIX_INSTALLABLE.FindStringSubmatch(str)
	if len(match) > 0 {
		if len(match) > 6 {
			version = match[6]
		}
		if len(match) > 5 {
			name = match[5]
		}
		if len(match) > 3 {
			attr = match[3]
		}
		if len(match) > 2 {
			revision = match[2]
		}
		flake = strings.TrimRight(match[1], "/")
	} else {
		match := RE_NIX_INSTALLABLE_SHORT.FindStringSubmatch(str)
		if len(match) > 5 {
			version = match[5]
		}
		if len(match) > 4 {
			name = match[4]
		}
		if len(match) > 2 {
			attr = match[2]
		}
		flake = match[1]
		revision = "HEAD"
	}
	result := lib.Version{
		Name:      name,
		Version:   version,
		Flake:     flake,
		Revision:  revision,
		Attribute: attr,
	}
	return result
}
