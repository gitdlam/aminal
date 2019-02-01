package terminal

import "fmt"

var charSets = map[rune]*map[rune]rune{
	'0': &decSpecGraphics,
	'B': nil, // ASCII
	// @todo 1,2,A
}

var decSpecGraphics = map[rune]rune{
	0x5f: 0x00A0, // NO-BREAK SPACE
	0x60: 0x25C6, // BLACK DIAMOND
	0x61: 0x2592, // MEDIUM SHADE
	0x62: 0x2409, // SYMBOL FOR HORIZONTAL TABULATION
	0x63: 0x240C, // SYMBOL FOR FORM FEED
	0x64: 0x240D, // SYMBOL FOR CARRIAGE RETURN
	0x65: 0x240A, // SYMBOL FOR LINE FEED
	0x66: 0x00B0, // DEGREE SIGN
	0x67: 0x00B1, // PLUS-MINUS SIGN
	0x68: 0x2424, // SYMBOL FOR NEWLINE
	0x69: 0x240B, // SYMBOL FOR VERTICAL TABULATION
	0x6a: 0x2518, // BOX DRAWINGS LIGHT UP AND LEFT
	0x6b: 0x2510, // BOX DRAWINGS LIGHT DOWN AND LEFT
	0x6c: 0x250C, // BOX DRAWINGS LIGHT DOWN AND RIGHT
	0x6d: 0x2514, // BOX DRAWINGS LIGHT UP AND RIGHT
	0x6e: 0x253C, // BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
	0x6f: 0x23BA, // HORIZONTAL SCAN LINE-1
	0x70: 0x23BB, // HORIZONTAL SCAN LINE-3
	0x71: 0x2500, // BOX DRAWINGS LIGHT HORIZONTAL
	0x72: 0x23BC, // HORIZONTAL SCAN LINE-7
	0x73: 0x23BD, // HORIZONTAL SCAN LINE-9
	0x74: 0x251C, // BOX DRAWINGS LIGHT VERTICAL AND RIGHT
	0x75: 0x2524, // BOX DRAWINGS LIGHT VERTICAL AND LEFT
	0x76: 0x2534, // BOX DRAWINGS LIGHT UP AND HORIZONTAL
	0x77: 0x252C, // BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
	0x78: 0x2502, // BOX DRAWINGS LIGHT VERTICAL
	0x79: 0x2264, // LESS-THAN OR EQUAL TO
	0x7a: 0x2265, // GREATER-THAN OR EQUAL TO
	0x7b: 0x03C0, // GREEK SMALL LETTER PI
	0x7c: 0x2260, // NOT EQUAL TO
	0x7d: 0x00A3, // POUND SIGN
	0x7e: 0x00B7, // MIDDLE DOT
}

func scs0Handler(pty chan rune, terminal *Terminal) error {
	return scsHandler(pty, terminal, 0)
}

func scs1Handler(pty chan rune, terminal *Terminal) error {
	return scsHandler(pty, terminal, 1)
}

func scsHandler(pty chan rune, terminal *Terminal, which int) error {
	b := <-pty

	cs, ok := charSets[b]
	if ok {
		terminal.logger.Debugf("Selected charset %v into G%v", string(b), which)
		terminal.terminalState.Charsets[which] = cs
		return nil
	}
	terminal.terminalState.Charsets[which] = nil
	return fmt.Errorf("Unknown SCS charset code: 0x%02X [%v]", b, string(b))
}
