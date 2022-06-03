package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32, n int32) string {
    runes := []rune(s)

    for i := 0; i < int(n); i++ {
        value := int32(runes[i])
        if 'a' <= runes[i] && runes[i] <='z' {
            value += k
            for value > int32('z'){
                value = value - int32('z') + int32('a') - 1
            }
            runes[i] = rune(int32(value))
        }else if 'A' <= runes[i] && runes[i] <='Z' {
            value += k
            for value > int32('Z'){
                value = value - int32('Z') + int32('A') - 1
            }
            runes[i] = rune(int32(value))
        }
    }
    result := string(runes)

    return result

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := caesarCipher(s, k, n)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
