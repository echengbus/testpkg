package testpkg

import (
    "fmt"
)

func TestpkgPrint(txt string) {
    fmt.Println(txt);
}

func main() {
    TestpkgPrint("I am testpkg");
}
