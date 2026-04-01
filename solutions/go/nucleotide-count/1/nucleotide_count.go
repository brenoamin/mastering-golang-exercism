package nucleotidecount

import "fmt"

type Histogram map[rune]int


// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.

type DNA string


func (d DNA) Counts() (Histogram, error) {
     h := Histogram{
        'A': 0,
        'C': 0,
        'G': 0,
        'T': 0,
    }

    for _, v := range d {
        switch v {
        case 'A', 'C', 'G', 'T':
            h[v]++
        default:
            return nil, fmt.Errorf("invalid nucleotide")
        }
    }

    return h, nil
}
