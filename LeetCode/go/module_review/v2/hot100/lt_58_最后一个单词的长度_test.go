package hot100

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	ret:=lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(ret)
}
