package leetcode
//
///**
//zero
//one
//two
//three
//four
//five
//six
//seven
//eight
//nine
//
//0 -> z
//2 -> w
//4 -> u
//6 -> x
//8 -> g
//
//3, 8 -> h
//5, 4 -> f
//7, 6 -> s
//
//a b c d e f g h i j k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z
//0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
//*/
//func originalDigits(s string) string {
//
//    var (
//        res = make([]int, 26)
//        out = make([]int, 10)
//        result = make([]int, len(s))
//    )
//
//    for _, c := range s {
//        res[c-'a']++
//    }
//
//    out[0] = res[25]
//    out[1] = res[22]
//    out[4] = res[20]
//    out[6] = res[23]
//    out[8] = res[6]
//    out[3] = res[7] - out[8]
//    out[5] = res[5] - out[4]
//    out[7] = res[18] - out[6]
//    out[9] = res[8] - out[5] - out[6] - out[8]
//    out[1] = res[13] - out[7] - 2*out[9]
//
//    for i := 0; i < 10; i++ {
//        for j := 0; j < out[i]; j++ {
//            result = append(result, i)
//        }
//    }
//
//
//}