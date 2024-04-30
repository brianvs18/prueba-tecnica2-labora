package delete_project

import (
	"bufio"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {

	input := "ebDV8ycct7Yg0on975MTu"
	scan := bufio.NewScanner(strings.NewReader(input))

	Run(scan)
}
