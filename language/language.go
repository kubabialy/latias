package language

import (
	"fmt"
	"math/rand"
	"time"
)

type Language struct {
  Name LanguageName
}

type LanguageName string;

// List of languages currently available.
const (
  PHP         LanguageName = "php"
  JavaScript  LanguageName = "javascript"
  TypeScript  LanguageName = "typescript"
  Java        LanguageName = "java"
  CSharp      LanguageName = "c#"
  CPP         LanguageName = "c++"
  Go          LanguageName = "go"
  Rust        LanguageName = "rust"
  Ruby        LanguageName = "ruby"
  Python      LanguageName = "python"
)

// Creates new instance of language with provided name.
func New(n LanguageName) Language {
    return Language{Name: n}
}

func SelectRandomLanguage() Language {
    l := [10]LanguageName{
        PHP,
        JavaScript,
        TypeScript,
        Java,
        CSharp,
        CPP,
        Go,
        Rust,
        Ruby,
        Python,
    }

    // Ditto have to initalize the global Source used by rand.Intn() and other functions of the rand package.
    rand.Seed(time.Now().Unix())
    ridx := rand.Intn(len(l))

    fmt.Println(ridx)
    return New(l[ridx])
}
