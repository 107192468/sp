package readexcel

import (
	"bufio"
	"fmt"
	"os"
)

func WriteResult(vals string, outfile string) error {

	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("writer", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(vals)
	writer.WriteString("\n")
	writer.Flush()

	return err
}
