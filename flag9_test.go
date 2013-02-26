package flag9

import "fmt"

func ExampleArgs() {
	s := []string{"-ab", "-carg", "--", "-d"}
	args := NewArgs(s)
	for args.Next() {
		switch args.Argc() {
		case 'a', 'b':
			fmt.Printf("%c\n", args.Argc())
		case 'c':
			argf, ok := args.Argf()
			if ok {
				fmt.Println("c", argf)
			} else {
				fmt.Println("c without argument")
			}
		case 'd':
			fmt.Println("not reached")
		}
	}
	fmt.Println(args.Argv())
	// Output:
	// a
	// b
	// c arg
	// [-d]
}
