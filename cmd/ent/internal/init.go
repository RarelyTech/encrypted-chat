package internal

import (
    "bytes"
    _ "embed"
    "entgo.io/ent/entc/gen"
    "fmt"
    "github.com/chatpuppy/puppychat/utils"
    "github.com/spf13/cobra"
    "log"
    "os"
    "path/filepath"
    "strings"
    "text/template"
)

const (
    defaultSchema = "./internal/ent/schema"
)

// schema template for the "init" command.
var tmpl = template.Must(template.New("schema").
    Parse(`package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "github.com/chatpuppy/puppychat/internal/ent/internal"
)

// {{ .name }} holds the schema definition for the {{ .name }} entity.
type {{ .name }} struct {
    ent.Schema
}

// Annotations of the {{ .name }}.
func ({{ .name }}) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "{{ .tableName }}"},
    }
}

// Fields of the {{ .name }}.
func ({{ .name }}) Fields() []ent.Field {
    return []ent.Field{}
}

// Edges of the {{ .name }}.
func ({{ .name }}) Edges() []ent.Edge {
    return []ent.Edge{}
}

func ({{ .name }}) Mixin() []ent.Mixin {
    return []ent.Mixin{
        internal.SnowflakeIDMixin{},
        internal.TimeMixin{},
    }
}

func ({{ .name }}) Indexes() []ent.Index {
    return []ent.Index{}
}
`))

func InitCmd() *cobra.Command {
    var target, customtmpl string
    cmd := &cobra.Command{
        Use:   "init [flags] [schemas]",
        Short: "initialize an environment with zero or more schemas",
        Example: examples(
            "ent init Example",
            "ent init --target entv1/schema OperatorName Group",
            "ent init --target entv1/schema --template tmpl/default.tmpl OperatorName Group",
        ),
        Run: func(cmd *cobra.Command, names []string) {
            if err := initEnv(target, names); err != nil {
                log.Fatalln(fmt.Errorf("ent/init: %w", err))
            }
        },
    }
    cmd.Flags().StringVar(&target, "target", defaultSchema, "target directory for schemas")
    cmd.Flags().StringVar(&customtmpl, "template", "", "target template for schemas")
    return cmd
}

func initEnv(target string, names []string) error {
    if err := createDir(target); err != nil {
        return fmt.Errorf("create dir %s: %w", target, err)
    }

    for _, name := range names {
        name = utils.StrToFirstUpper(name)
        if err := gen.ValidSchemaName(name); err != nil {
            return fmt.Errorf("init schema %s: %w", name, err)
        }
        b := bytes.NewBuffer(nil)
        if err := tmpl.Execute(b, map[string]string{
            "name":      name,
            "tableName": strings.ToLower(utils.StrToSnakeCase(name)),
        }); err != nil {
            return fmt.Errorf("executing template %s: %w", name, err)
        }
        newFileTarget := filepath.Join(target, strings.ToLower(name+".go"))
        if err := os.WriteFile(newFileTarget, b.Bytes(), 0644); err != nil {
            return fmt.Errorf("writing file %s: %w", newFileTarget, err)
        }
    }
    return nil
}
