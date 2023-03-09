package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/urfave/cli"
)

const (
	BASE             = 62
	DIGIT_OFFSET     = 48
	LOWERCASE_OFFSET = 61
	UPPERCASE_OFFSET = 55
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func char2ord(char string) (int, error) {
	fmt.Printf("Char is %s\n", char)
	if matched, _ := regexp.MatchString("[0-9]", char); matched {
		/*
		 * rune('1') = 49 which is ASCII val of 1
		 */
		return int([]rune(char)[0] - DIGIT_OFFSET), nil //int(rune('A') - DIGIT_OFFSET)  or int([]rune("ABC")[0] - DIGIT_OFFSET) both works
	} else if matched, _ := regexp.MatchString("[A-Z]", char); matched {
		return int([]rune(char)[0] - UPPERCASE_OFFSET), nil
	} else if matched, _ := regexp.MatchString("[a-z]", char); matched {
		return int([]rune(char)[0] - LOWERCASE_OFFSET), nil
	} else {
		return -1, fmt.Errorf("%s is not a valid character", char)
	}
}

func ord2char(ord int) (string, error) {
	switch {
	case ord < 10:
		return string(ord + DIGIT_OFFSET), nil // string(1+48)= string(49)=1 => 49 is ASCII/Unicode val of 1
	case ord >= 10 && ord <= 35:
		return string(ord + UPPERCASE_OFFSET), nil
	case ord >= 36 && ord < 62:
		return string(ord + LOWERCASE_OFFSET), nil
	default:
		return "", fmt.Errorf("%d is not a valid integer in the range of base %d", ord, BASE)
	}
}

func Decode(str string) (int, error) {
	sum := 0
	for i, c := range reverse(str) {
		if d, err := char2ord(string(c)); err == nil {
			sum = sum + d*int(math.Pow(BASE, float64(i)))
		} else {
			return -1, err
		}
	}
	return sum, nil
}

func Encode(digits int) (string, error) {
	if digits == 0 {
		return "0", nil
	}

	str := ""
	for digits >= 0 {
		remainder := digits % BASE
		if s, err := ord2char(remainder); err != nil {
			return "", err
		} else {
			str = s + str
		}

		if digits == 0 {
			break
		}
		digits = int(digits / BASE)
	}
	return str, nil
}

func main() {

	/*
	 *	fmt.Println("Encode= ", string(49)) = 1 , string() takes ASCII code and returns int val
	 *	fmt.Println("decode= ", rune('1'))  = 49, rune('') takes chars and returns ASCII val
	 */
	var decode, encode bool
	app := cli.NewApp()

	app.Name = "Base62Convertor"
	app.Usage = "Encode digits to short strings or decode short strings to digits"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "decode, d",
			Usage:       "Use decode function",
			Destination: &decode,
		},
		cli.BoolFlag{
			Name:        "encode, e",
			Usage:       "Use encode function",
			Destination: &encode,
		},
	}

	app.Action = func(c *cli.Context) error {
		for _, arg := range c.Args() {
			fmt.Println("decode = ", decode)
			if decode {
				if d, err := Decode(arg); err != nil {
					return err
				} else {
					fmt.Printf("Decode: %s => %d\n", arg, d)
				}
			} else {
				d, err := strconv.Atoi(arg)
				if err != nil {
					return err
				}
				if s, err := Encode(d); err != nil {
					return err
				} else {
					fmt.Printf("Encode: %s => %s\n", arg, s)
				}
			}
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

/*
normal encode & decode
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"regexp"
)

var SET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

const (
	BASE       = 62
	NUM_OFFSET = 48
	CAP_OFFSET = 55
	SMA_OFFSET = 61
)

func main() {
	en := 123
	val1 := ""
	if val, err := encode(en); err == nil {
		val1 = val
		fmt.Printf("Encoding of  %v is %v\n", en, val)
	} else {

		fmt.Println(err)
	}

	if de, err := decode(val1); err == nil {
		fmt.Printf("Decoding of %v is %v\n", val1, de)
	}
}

func int2char(val int) (string, error) {
	fmt.Println("char is : ", val)
	if 0 < val && val < 10 {
		return string(val + NUM_OFFSET), nil
	} else if 10 <= val && val < 35 {
		return string(val + CAP_OFFSET), nil
	} else if 35 <= val && val < 62 {
		return string(val + SMA_OFFSET), nil
	} else {
		fmt.Println("wrong input")
		return "", fmt.Errorf("wrong input")
	}

}

func encode(en int) (string, error) {
	str := ""
	for en > 0 {
		rem := en % BASE
		if s, err := int2char(rem); err != nil {
			return "", err
		} else {
			str = s + str
		}
		if en == 0 {
			break
		}
		en = en / BASE

	}

	return str, nil
}
func reverse(str1 string) string {
	str := []rune(str1)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

func char2int(ch string) (int, error) {
	if matched, _ := regexp.MatchString("[0-9]", ch); matched {
		return int([]rune(ch)[0] - NUM_OFFSET), nil
	} else if matched, _ := regexp.MatchString("[A-Z]", ch); matched {
		return int([]rune(ch)[0] - CAP_OFFSET), nil
	} else if matched, _ := regexp.MatchString("[a-z]", ch); matched {
		return int([]rune(ch)[0] - SMA_OFFSET), nil
	} else {
		return -1, fmt.Errorf("Not a valid char")
	}
}

func decode(str string) (int, error) {
	fmt.Println("str is=", str)
	decVal := 0
	for i, ch := range reverse(str) {
		if val, err := char2int(string(ch)); err != nil {
			return -1, err
		} else {
			decVal = decVal + val*int(math.Pow(BASE, float64(i)))
		}
	}
	return decVal, nil
}



*/
