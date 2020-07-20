package main

import (
    "bufio"
    "log"
    "flag"
    "os"
    "image/png"
    "botsandus"
)

var inputFile = flag.String("input", "", "input file")

func main() {
    flag.Parse()
    if *inputFile == "" {
        flag.PrintDefaults()
        return
    }

    file, err := os.Open(*inputFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        encoded, checksummed, err := botsandus.EncodeNumber(line)
        image := botsandus.CreateImageFromBinaryString(encoded)

        f, err := os.Create(checksummed + ".png")
        defer f.Close()
        if err != nil {
            log.Fatal(err)
        }
        if err := png.Encode(f, image); err != nil {
            log.Fatal(err)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
