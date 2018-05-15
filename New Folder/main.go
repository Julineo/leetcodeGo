/*
Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

Example 1:

Input: S = "loveleetcode", C = 'e'
Output: [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]
 

Note:

S string length is in [1, 10000].
C is a single character, and guaranteed to be in string S.
All letters in S and C are lowercase.
*/

func shortestToChar(S string, C byte) []int {
    //initiate slice
    sl := make([]int,len(S))
    for i := 0; i < len(sl); i++ {
        sl[i] = 10000
    }
    
    //mark zero point
    for i := 0; i < len(sl); i++ {
        if S[i] == C {
            sl[i] = 0
        }
    }
    
    //foreward increase with min
    for i := 1; i < len(sl); i ++ {
        sl[i] = min(sl[i], sl[i-1] + 1)
    }
    
    //backward increase with min
    for i := len(sl) - 2; i >= 0; i-- {
        sl[i] = min(sl[i], sl[i+1] + 1)
    }

    
    return sl
}

func min (a, b int) int {
    if a > b {
        return b
    }
    return a
}
