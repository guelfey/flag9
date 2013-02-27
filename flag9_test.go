package flag9

import (
	"fmt"
	"testing"
)

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

func TestFlag9(t *testing.T) {
	type parg struct {
		c       rune
		hasArgf bool
		argf    string
	}
	type test struct {
		args  []string
		pargs []parg
		argv  []string
	}
	tests := []test{
		test{
			[]string{"-bfoo", "-a", "bar", "-", "-d"},
			[]parg{
				parg{'b', true, "foo"},
				parg{'a', true, "bar"},
			},
			[]string{"-", "-d"},
		},
		test{
			[]string{},
			[]parg{},
			[]string{},
		},
		test{
			[]string{"--"},
			[]parg{},
			[]string{},
		},
		test{
			[]string{"-"},
			[]parg{},
			[]string{"-"},
		},
		test{
			[]string{"foo"},
			[]parg{},
			[]string{"foo"},
		},
		test{
			[]string{"-abcfoo", "-c", "bar", "-c", "baz"},
			[]parg{
				parg{'a', false, ""},
				parg{'b', false, ""},
				parg{'c', true, "foo"},
				parg{'c', true, "bar"},
				parg{'c', false, ""},
			},
			[]string{"baz"},
		},
	}
	for i, v := range tests {
		args := NewArgs(v.args)
		for args.Next() {
			if len(v.pargs) == 0 {
				t.Fatalf("test %d: too many parsed arguments", i)
			}
			c := args.Argc()
			if c != v.pargs[0].c {
				t.Fatalf("test %d: got option %c, wanted %c", i, c, v.pargs[0].c)
			}
			if v.pargs[0].hasArgf {
				argf, ok := args.Argf()
				if !ok {
					t.Fatalf("test %d: expected argf, but got none", i)
				}
				if argf != v.pargs[0].argf {
					t.Fatalf("test %d: got argf %q, but expected %q", i, argf,
						v.pargs[0].argf)
				}
			}
			v.pargs = v.pargs[1:]
		}
		if len(v.pargs) != 0 {
			t.Fatalf("test %d: %d unparsed options", i, len(v.pargs))
		}
		argv := args.Argv()
		if len(argv) != len(v.argv) {
			t.Fatalf("test %d: argv count mismatch (got %d, expected %d)", i,
				len(argv), len(v.argv))
		}
		for n := range argv {
			if argv[n] != v.argv[n] {
				t.Fatalf("test %d: argv[%d] mismatch (got %q, expected %q)", i,
					n, argv[n], v.argv[n])
			}
		}
	}
}
