package utils

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "github.com/denisbrodbeck/machineid"
    "log"
    "regexp"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var (
    matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
    matchAllCap   = regexp.MustCompile("([a-z\\d])([A-Z])")
)

func GetMachineID() string {
    id, err := machineid.ID()
    if err != nil {
        log.Fatalf("MachineID get failed: %v", err)
    }
    return id
}

func GetTimestampNanoString() string {
    return fmt.Sprintf("%d", time.Now().UnixNano())
}

func StrToFirstUpper(str string) string {
    if len(str) == 0 {
        return ""
    }
    tmp := []rune(str)
    tmp[0] = unicode.ToUpper(tmp[0])
    return string(tmp)
}

func StrToSnakeCase(str string) string {
    snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
    snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
    return strings.ToLower(snake)
}

// Md5String hashes using md5 algorithm
func Md5String(b []byte) string {
    algorithm := md5.New()
    algorithm.Write(b)
    return hex.EncodeToString(algorithm.Sum(nil))
}

func ParseID(str string) uint64 {
    id, _ := strconv.ParseUint(str, 10, 64)
    return id
}
