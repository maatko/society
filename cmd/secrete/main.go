package main

import (
	"context"
	"os"

	"github.com/maatko/secrete/web/template"
)

func main() {
	component := template.Index("Welcome Brother")
	component.Render(context.Background(), os.Stdout)
}