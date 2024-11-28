func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    sArray, tArray := []rune(s), []rune(t)
    sort.Slice(sArray, func(i, j int) bool {
        return sArray[i] < sArray[j]
    })

    sort.Slice(tArray, func(i, j int) bool {
        return tArray[i] < tArray[j]
    })

    for i := range sArray {
        if sArray[i] != tArray[i] {
            return false
        }
    }
    return true
    }
