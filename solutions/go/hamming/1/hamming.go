package hamming
import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b){ return 0, errors.New("strings are not same length, cant calculate hamming")}
    var count = 0
	for i := range a{
        if a[i] != b[i]{
            count += 1
        }
    }
    return count, nil
}
