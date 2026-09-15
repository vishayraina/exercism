package isbnverifier
import "strings"
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
    isbn = strings.TrimSpace(isbn)
    if len(isbn) != 10 {
        return false
    }
    if ((isbn[len(isbn)-1] < '0' || isbn[len(isbn)-1] > '9') && isbn[len(isbn)-1] != 'X') {
        return false
    }
    sum := 0
    for i := 0; i < len(isbn); i++ {
        if i == len(isbn)-1 && isbn[i] == 'X' {
            sum += 10
            continue
        } 
        if isbn[i] < '0' || isbn[i] > '9' {
            return false
        }
		sum += int(isbn[i]-'0') * (10-i)
        
    } 
    return sum%11 == 0
}
