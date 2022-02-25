package pakke

import "rsc.io/quote"

func TestQuote() string {
	return quote.Glass() + "\n" + quote.Go() + "\n" + quote.Opt() + "\n" + quote.Hello();

}

