package commands

import (
	"clean-architecture/lib"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type CrudGeneratorCommand struct {
	Name string
}

func NewCrudGenerator() *CrudGeneratorCommand {
	return &CrudGeneratorCommand{}
}

func (cg *CrudGeneratorCommand) Short() string {
	return "generate a crud"
}

func (cg *CrudGeneratorCommand) Setup(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&cg.Name, "name", "n", "", "generate user crud")
}

func (cg *CrudGeneratorCommand) Run() lib.CommandRunner {
	return func(l lib.Logger) {
		l.Info("running crud generator command")

		if cg.Name == "" {
			l.Info("Provide name to the crud!!!")
			return
		}

		crudFileName := strings.ToLower(cg.Name)

		if !regexp.MustCompile(`^[a-z]+[a-z_]*[a-z]+$`).MatchString(crudFileName) {
			l.Info("Provide name that contains only alphabets and can have underscore(_) in middle!!!")
			return
		}

		modelName := cg.generateModelName(crudFileName)

		layers := []string{"model", "repository", "service", "controller", "route"}

		for _, layer := range layers {
			l.Infof("--- Generating %s", layer)
			err := cg.fileGenerator(layer, crudFileName, modelName)
			if err != nil {
				l.Error(err)
				return
			}
			l.Infof("--- %s Generated", strings.Title(layer)) //nolint
		}
	}
}

func (cg *CrudGeneratorCommand) fileGenerator(packageName, crudFileName, modelName string) error {
	dir := packageName
	if packageName != "repository" {
		dir += "s"
	}

	if packageName == "model" {
		crudFileName += ".go"
	} else {
		crudFileName = fmt.Sprintf("%s_%s.go", crudFileName, packageName)
	}

	path := dir
	if dir == "controllers" || dir == "routes" {
		path = filepath.Join("api", dir)
	}
	file := filepath.Join(path, crudFileName)

	if _, err := os.Stat(file); err == nil {
		return fmt.Errorf("file already exists with name %s", crudFileName)
	}

	t, err := template.ParseFiles(fmt.Sprintf("templates/crud_templates/generate_%s.txt", packageName))
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	type Fields struct {
		PackageName,
		ModuleName,
		ModelName string
	}

	err = t.Execute(f, Fields{dir, "clean-architecture", modelName})
	if err != nil {
		return err
	}
	f.Close()

	if packageName != "model" {
		moduleFilepath := filepath.Join(path, fmt.Sprintf("%s.go", dir))
		moduleFileData, err := os.ReadFile(moduleFilepath)
		if err != nil {
			return err
		}

		if packageName == "route" {
			packageName = dir
		}
		finalData := strings.ReplaceAll(string(moduleFileData), ",\n)", fmt.Sprintf(",\n    fx.Provide(New%s%s),\n)", modelName, strings.Title(packageName))) //nolint

		moduleFileRW, err := os.OpenFile(moduleFilepath, os.O_RDWR, 0o755)
		if err != nil {
			return err
		}

		_, err = moduleFileRW.WriteString(finalData)
		if err != nil {
			return err
		}
		_ = moduleFileRW.Close()
	}

	return nil
}

func (cg *CrudGeneratorCommand) generateModelName(name string) (modelName string) {
	i := 0
	for _, c := range name {
		if i == 0 {
			modelName += strings.ToUpper(string(c))
			i++
			continue
		}
		if c == '_' {
			i = 0
			continue
		}
		modelName += string(c)
	}
	return modelName
}
