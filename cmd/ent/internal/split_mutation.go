package internal

import (
    "bytes"
    "fmt"
    "github.com/chatpuppy/puppychat/utils"
    "golang.org/x/exp/slices"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "text/template"
)

func DeleteMutations() {
    p := "internal/ent"
    files, _ := os.ReadDir(p)
    for _, file := range files {
        if strings.HasPrefix(file.Name(), "mutation_") {
            _ = os.Remove(filepath.Join(p, file.Name()))
        }
    }
}

func SplitMutation() {
    p := "internal/ent"

    re := regexp.MustCompile(`(?m)/\*\* (.*)? START \*/\n/\*\n([\s\S]*?)\*/([\s\S]*?)/\*\* .*? END \*/`)
    mf := filepath.Join(p, "mutation.go")

    b, err := os.ReadFile(mf)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    t, err := template.ParseFiles("./cmd/ent/template/split_mutation.tmpl")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    type splitMutation struct {
        Imports []string
        Data    []byte
    }
    for _, matches := range re.FindAllSubmatch(b, -1) {
        var imports []string
        for _, bi := range strings.Split(string(matches[2]), "\n") {
            s := strings.TrimSpace(bi)

            if len(s) > 0 && !slices.Contains(imports, s) {
                imports = append(imports, s)
            }
        }
        sm := splitMutation{
            Imports: imports,
            Data:    matches[3],
        }
        f := filepath.Join(p, fmt.Sprintf("mutation_%s.go", utils.StrToSnakeCase(string(matches[1]))))
        w, _ := os.OpenFile(f, os.O_CREATE|os.O_WRONLY, 0755)
        _ = t.Execute(w, sm)
    }

    CleanMutation(re.ReplaceAll(b, []byte("")))

    os.Exit(0)
}

func CleanMutation(b []byte) {
    re := regexp.MustCompile(`(?m)[\S\s]*(const \([\S\s]*?\))[\s]+`)
    re.ReplaceAll(b, []byte(`$1\n`))
    buffer := &bytes.Buffer{}
    buffer.WriteString(`package ent
    
import (
    "entgo.io/ent"
)

`)
    buffer.Write(re.ReplaceAll(b, []byte("$1\n")))
    _ = os.WriteFile("internal/ent/mutation.go", buffer.Bytes(), 0755)
}
