package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"log"
	"time"
	"os"
	"unicode"
	"flag"
	"syscall"
)

const (
	VK_SHIFT              = 0x10
	VK_CTRL               = 0x11
	VK_ALT                = 0x12
	VK_LSHIFT             = 0xA0
	VK_RSHIFT             = 0xA1
	VK_LCONTROL           = 0xA2
	VK_RCONTROL           = 0xA3
	VK_LWIN               = 0x5B
	VK_RWIN               = 0x5C
	VK_ESC                = 0x1B
	KEYEVENTF_KEYUP       = 0x0002
	KEYEVENTF_SCANCODE    = 0x0008
	KEYEVENTF_EXTENDEDKEY = 0x0001
    VK_SPACE              = 0x20
	VK_MINUS              = 0xBD
	VK_EQUAL              = 0xBB
	VK_BACKSPACE          = 0x08
	VK_TAB                = 0x09
	VK_LEFTBRACE          = 0xDB
	VK_RIGHTBRACE         = 0xDD
	VK_ENTER              = 0x0D
	VK_SEMICOLON          = 0xBA
	VK_APOSTROPHE         = 0xDE
	VK_GRAVE              = 0xC0
	VK_BACKSLASH          = 0xDC
	VK_COMMA              = 0xBC
	VK_DOT                = 0xBE
	VK_SLASH              = 0xBF
	VK_KPASTERISK         = 0x6A
	VK_CAPSLOCK           = 0x14

	VK_0   = 0x30
	VK_1   = 0x31
	VK_2   = 0x32
	VK_3   = 0x33
	VK_4   = 0x34
	VK_5   = 0x35
	VK_6   = 0x36
	VK_7   = 0x37
	VK_8   = 0x38
	VK_9   = 0x39

	VK_A = 0x41
	VK_B = 0x42
	VK_C = 0x43
	VK_D = 0x44
	VK_E = 0x45
	VK_F = 0x46
	VK_G = 0x47
	VK_H = 0x48
	VK_I = 0x49
	VK_J = 0x4A
	VK_K = 0x4B
	VK_L = 0x4C
	VK_M = 0x4D
	VK_N = 0x4E
	VK_O = 0x4F
	VK_P = 0x50
	VK_Q = 0x51
	VK_R = 0x52
	VK_S = 0x53
	VK_T = 0x54
	VK_U = 0x55
	VK_V = 0x56
	VK_W = 0x57
	VK_X = 0x58
	VK_Y = 0x59
	VK_Z = 0x5A

	// Add more virtual key codes as needed...
)


func main() {
	a := app.New()
	w := a.NewWindow("Sendit V2")
	input := widget.NewEntry()
	content := container.NewVBox(
		widget.NewLabel("Send this"),
		input,
		widget.NewButton("Sendit", func() {
			log.Println("Preparing to Send")
			inputText := input.Text
			delay := 5 // Default delay in seconds

			// Parse delay from the command line argument if provided
			if len(os.Args) > 1 {
				delayFlag := flag.Int("delay", 5, "Delay in seconds")
				flag.Parse()
				delay = *delayFlag
			}

			// Start a Goroutine to send the input text with the specified delay
			go typeStringWithDelay(inputText, delay)
		}))
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 150))
	w.Show()

	a.Run()
}

func typeStringWithDelay(input string, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	typeString(input)
}


func typeString(input string) {
	var capital bool
	for _, character := range input {
		if unicode.IsSpace(character) {
			capital = true
			typeKey(VK_SPACE, capital)
		} else {
			// Handle other characters
			switch string(character) {
			case "A":
				capital = true
				typeKey(VK_A, capital)
			case "a":
				capital = false
				typeKey(VK_A, capital)
			case "B":
				capital = true
				typeKey(VK_B, capital)
			case "b":
				capital = false
				typeKey(VK_B, capital)
			case "C":
				capital = true
				typeKey(VK_C, capital)
			case "c":
				capital = false
				typeKey(VK_C, capital)
			case "D":
				capital = true
				typeKey(VK_D, capital)
			case "d":
				capital = false
				typeKey(VK_D, capital)
			case "E":
				capital = true
				typeKey(VK_E, capital)
			case "e":
				capital = false
				typeKey(VK_E, capital)
			case "F":
				capital = true
				typeKey(VK_F, capital)
			case "f":
				capital = false
				typeKey(VK_F, capital)
			case "G":
				capital = true
				typeKey(VK_G, capital)
			case "g":
				capital = false
				typeKey(VK_G, capital)
			case "H":
				capital = true
				typeKey(VK_H, capital)
			case "h":
				capital = false
				typeKey(VK_H, capital)
			case "I":
				capital = true
				typeKey(VK_I, capital)
			case "i":
				capital = false
				typeKey(VK_I, capital)
			case "J":
				capital = true
				typeKey(VK_J, capital)
			case "j":
				capital = false
				typeKey(VK_J, capital)
			case "K":
				capital = true
				typeKey(VK_K, capital)
			case "k":
				capital = false
				typeKey(VK_K, capital)
			case "L":
				capital = true
				typeKey(VK_L, capital)
			case "l":
				capital = false
				typeKey(VK_L, capital)
			case "M":
				capital = true
				typeKey(VK_M, capital)
			case "m":
				capital = false
				typeKey(VK_M, capital)
			case "N":
				capital = true
				typeKey(VK_N, capital)
			case "n":
				capital = false
				typeKey(VK_N, capital)
			case "O":
				capital = true
				typeKey(VK_O, capital)
			case "o":
				capital = false
				typeKey(VK_O, capital)
			case "P":
				capital = true
				typeKey(VK_P, capital)
			case "p":
				capital = false
				typeKey(VK_P, capital)
			case "Q":
				capital = true
				typeKey(VK_Q, capital)
			case "q":
				capital = false
				typeKey(VK_Q, capital)
			case "R":
				capital = true
				typeKey(VK_R, capital)
			case "r":
				capital = false
				typeKey(VK_R, capital)
			case "S":
				capital = true
				typeKey(VK_S, capital)
			case "s":
				capital = false
				typeKey(VK_S, capital)
			case "T":
				capital = true
				typeKey(VK_T, capital)
			case "t":
				capital = false
				typeKey(VK_T, capital)
			case "U":
				capital = true
				typeKey(VK_U, capital)
			case "u":
				capital = false
				typeKey(VK_U, capital)
			case "V":
				capital = true
				typeKey(VK_V, capital)
			case "v":
				capital = false
				typeKey(VK_V, capital)
			case "W":
				capital = true
				typeKey(VK_W, capital)
			case "w":
				capital = false
				typeKey(VK_W, capital)
			case "X":
				capital = true
				typeKey(VK_X, capital)
			case "x":
				capital = false
				typeKey(VK_X, capital)
			case "Y":
				capital = true
				typeKey(VK_Y, capital)
			case "y":
				capital = false
				typeKey(VK_Y, capital)
			case "Z":
				capital = true
				typeKey(VK_Z, capital)
			case "z":
				capital = false
				typeKey(VK_Z, capital)
			case "0":
				capital = false
				typeKey(VK_0, capital)
			case "1":
				capital = false
				typeKey(VK_1, capital)
			case "2":
				capital = false
				typeKey(VK_2, capital)
			case "3":
				capital = false
				typeKey(VK_3, capital)
			case "4":
				capital = false
				typeKey(VK_4, capital)
			case "5":
				capital = false
				typeKey(VK_5, capital)
			case "6":
				capital = false
				typeKey(VK_6, capital)
			case "7":
				capital = false
				typeKey(VK_7, capital)
			case "8":
				capital = false
				typeKey(VK_8, capital)
			case "9":
				capital = false
				typeKey(VK_9, capital)
			case "!":
				capital = true
				typeKey(VK_1, capital)
			case "@":
				capital = true
				typeKey(VK_2, capital)
			case "#":
				capital = true
				typeKey(VK_3, capital)
			case "$":
				capital = true
				typeKey(VK_4, capital)
			case "%":
				capital = true
				typeKey(VK_5, capital)
			case "^":
				capital = true
				typeKey(VK_6, capital)
			case "&":
				capital = true
				typeKey(VK_7, capital)
			case "*":
				capital = true
				typeKey(VK_8, capital)
			case "(":
				capital = true
				typeKey(VK_9, capital)
			case ")":
				capital = true
				typeKey(VK_0, capital)
			case ",":
				capital = false
				typeKey(VK_COMMA, capital)
			case ".":
				capital = false
				typeKey(VK_DOT, capital)
			case "/":
				capital = false
				typeKey(VK_SLASH, capital)
			case "<":
				capital = true
				typeKey(VK_COMMA, capital)
			case ">":
				capital = true
				typeKey(VK_DOT, capital)
			case "?":
				capital = true
				typeKey(VK_SLASH, capital)
			case ";":
				capital = false
				typeKey(VK_SEMICOLON, capital)
			case ":":
				capital = true
				typeKey(VK_SEMICOLON, capital)
			case "'":
				capital = false
				typeKey(VK_APOSTROPHE, capital)
			case "[":
				capital = false
				typeKey(VK_LEFTBRACE, capital)
			case "{":
				capital = true
				typeKey(VK_LEFTBRACE, capital)
			case "]":
				capital = false
				typeKey(VK_RIGHTBRACE, capital)
			case "}":
				capital = true
				typeKey(VK_RIGHTBRACE, capital)
			case "\\":
				capital = false
				typeKey(VK_BACKSLASH, capital)
			case "|":
				capital = true
				typeKey(VK_BACKSLASH, capital)
			case "-":
				capital = false
				typeKey(VK_MINUS, capital)
			case "_":
				capital = true
				typeKey(VK_MINUS, capital)
			case "=":
				capital = false
				typeKey(VK_EQUAL, capital)
			case "+":
				capital = true
				typeKey(VK_EQUAL, capital)
			case "`":
				capital = false
				typeKey(VK_GRAVE, capital)
			case "~":
				capital = true
				typeKey(VK_GRAVE, capital)
			}
		}
	}
}


func typeKey(key int, capital bool) {
	var dll = syscall.NewLazyDLL("user32.dll")
	var procKeyBd = dll.NewProc("keybd_event")
	if capital == true {
		procKeyBd.Call(uintptr(VK_SHIFT), uintptr(0), uintptr(0), 0)
		procKeyBd.Call(uintptr(key), uintptr(0), uintptr(0), 0)
		procKeyBd.Call(uintptr(key), uintptr(0), uintptr(KEYEVENTF_KEYUP), 0)
		procKeyBd.Call(uintptr(VK_SHIFT), uintptr(0), uintptr(KEYEVENTF_KEYUP), 0)
	} else {
		procKeyBd.Call(uintptr(key), uintptr(0), uintptr(0), 0)
		procKeyBd.Call(uintptr(key), uintptr(0), uintptr(KEYEVENTF_KEYUP), 0)
	}
}

func delayType(delayseconds int) {
	time.Sleep(time.Duration(delayseconds) * time.Second)
}
