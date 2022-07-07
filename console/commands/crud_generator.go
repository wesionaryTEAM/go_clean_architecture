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

		if !regexp.MustCompile(`^[a-z_]*$`).MatchString(crudFileName) {
			l.Info("Provide name containing [a-z or _] to the crud!!!")
			return
		}

		modelName := cg.generateModelName(crudFileName)

		layers := []string{"repository", "service", "controller", "route"}

		for _, layer := range layers {
			l.Infof("--- Generating %s ---", layer)
			err := cg.fileGenerator(layer, crudFileName, modelName)
			if err != nil {
				l.Error(err)
				return
			}
			l.Infof("--- %s Generated ---", strings.Title(layer))
		}
	}
}

func (cg *CrudGeneratorCommand) fileGenerator(packageName, crudFileName, modelName string) error {
	dir := packageName
	if packageName != "repository" {
		dir += "s"
	}

	fileName := fmt.Sprintf("%s_%s.go", crudFileName, packageName)
	path := dir
	if dir == "controllers" || dir == "routes" {
		path = filepath.Join("api", dir)
	}
	file := filepath.Join(path, fileName)

	if _, err := os.Stat(file); err == nil {
		return fmt.Errorf("file already exists with name %s", fileName)
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

	packageName = dir
	err = t.Execute(f, Fields{packageName, "clean-architecture", modelName})
	if err != nil {
		return err
	}
	f.Close()
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
