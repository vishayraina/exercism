package luhn

import "strings"

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
        return false
    }
    var sum int
    for i := len(id)-1; i >= 0; i-- {
        pos := len(id)-1-i
        if id[i] < '0' || id[i] > '9' {
            return false
        } 
        d := int(id[i] - '0')
        if pos%2 == 1 {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }
        sum += d
    }
    return sum%10 == 0
}
