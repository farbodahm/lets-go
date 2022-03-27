package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/farbodahm/lets-go/monkeyLangInterpreter/lexer"
	"github.com/farbodahm/lets-go/monkeyLangInterpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprint(out, PROMPT)
		if err != nil {
			log.Fatalln("Error: Couldn't print to writer: ", err)
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err = fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				log.Fatalln("Error: Couldn't print to writer: ", err)
			}
		}
	}
}
