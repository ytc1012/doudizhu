package main

import (
    "fmt"
    "math/rand"
    "time"
    "bytes"
    "sort"
)

var FullPoker = []byte{
    0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, //黑桃 A - K
    0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, //红桃 A - K
    0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, //梅花 A - K
    0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, //方片 A - K
    0x4E, 0x4F, //小、大王
}

func main() {
    pf := &FullPoker
    fmt.Println("新牌：\n", FullPoker, "\n")
    res1 := shufflePoker(pf)
    fmt.Println("洗牌：\n", res1, "\n")
    p1, p2, p3, bo := dealPoker(res1)
    fmt.Println("玩家牌和底牌：\n", p1, "\n", p2, "\n", p3, "\n", bo, "\n")
    
    twoPokers := bytes.Repeat(FullPoker, 2)
    fmt.Println("两副：\n", twoPokers, "\n")
    res2 := shufflePoker(&twoPokers)
    fmt.Println("洗牌：\n", res2, "\n")
    ps1, ps2, ps3, ps4, bos := dealPokers(res2)
    fmt.Println("玩家牌和底牌：\n", ps1, "\n", ps2, "\n", ps3, "\n", ps4, "\n", bos, "\n")
    
    ceshi := []byte{0x03, 0x04, 0x03, 0x04, 0x03, 0x04, 0x05, 0x05, 0x05, 0x06, 0x09, 0x08}
    
    cardsType(&ceshi)
}

//洗牌
func shufflePoker(vals *[]byte) *[]byte {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    ret := make([]byte, len(*vals))
    perm := r.Perm(len(*vals))
    for i, randIndex := range perm {
        ret[i] = (*vals)[randIndex]
    }
    return &ret
}

//发牌
func dealPoker(vals *[]byte) (*[]byte, *[]byte, *[]byte, *[]byte) {
    pvals := *vals
    a, b, c, d := pvals[:17], pvals[17:34], pvals[34:51], pvals[51:]
    return &a, &b, &c, &d
}

func dealPokers(vals *[]byte) (*[]byte, *[]byte, *[]byte, *[]byte, *[]byte) {
    pvals := *vals
    a, b, c, d, e := pvals[:25], pvals[25:50], pvals[50:75], pvals[75:100], pvals[100:]
    return &a, &b, &c, &d, &e
}

//判断牌型
func checkCardsType(cards *[]int) {
    pcards := *cards
    fmt.Println(pcards)
    switch len(pcards) {
    case 0:
        fmt.Println("没有牌\n")
    case 1:
        fmt.Println("单张\n")
    case 2:
        if pcards[0] == 0x4F && pcards[1] == 0x4E {
            fmt.Print("王炸\n")
        } else if pcards[0] & 0x0f == pcards[1] & 0x0f {
            fmt.Println("对子\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 3:
        if pcards[0] & 0x0f == pcards[1] & 0x0f && pcards[1] & 0x0f == pcards[2] & 0x0f {
            fmt.Print("三不带\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 4:
        //连队
        if isEvenPair(cards) {
            fmt.Println("连对\n")
        //炸弹
        } else if isBoom(cards) {
            fmt.Println("炸弹\n")
        //三带一
        } else if isThreeOne(cards) {
            fmt.Println("三带一\n")
        } else {
            fmt.Print("不符合规则\n")
        }
    case 5:
        // 顺子
        if isStraight(cards) {
            fmt.Println("顺子\n")
        } else if isThreePair(cards) {
            fmt.Println("三带对\n")
        } else if isBoom(cards) {
            fmt.Println("炸弹\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 6:
        if isPlane(cards) {
            fmt.Println("飞机不带\n")
        } else if isStraight(cards) {
            fmt.Println("顺子\n")
        } else if isFourTwo(cards) {
            fmt.Println("四带二单\n")
        } else if isEvenPair(cards) {
            fmt.Println("连对")
        } else if isBoom(cards) {
            fmt.Println("炸弹\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 8:
        if isStraight(cards) {
            fmt.Println("顺子\n")
        } else if isPlaneTwo(cards) {
            fmt.Println("飞机带二单\n")
        } else if isFourPair(cards) {
            fmt.Println("四带二对\n")
        } else if isEvenPair(cards) {
            fmt.Println("连对\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    default:
        if isStraight(cards) {
            fmt.Println("顺子\n")
        } else if isEvenPair(cards) {
            fmt.Println("连对\n")
        } else if isBoom(cards) {
            fmt.Println("炸弹\n")
        } else if isPlane(cards) {
            fmt.Println("飞机不带\n")
        } else if isThreePlaneThree(cards) {
            fmt.Println("三飞机带三单\n")
        } else if isThreePlanePair(cards) {
            fmt.Println("三飞机带三对\n")
        } else if isPlanePair(cards) {
            fmt.Println("飞机带二对\n")
        } else {
            fmt.Println("不符合规则")
        }
    }
}

//去重
func removeDup(a []int) []int {
    sort.Sort(sort.IntSlice(a))
    i := 0
    for j := 1; j < len(a); j ++ {
        if a[i] != a[j] {
            i ++
            a[i] = a[j]
        }
    }
    return a[:i+1]
}

//根据已有牌癞子可取值的集合
func laiZiValue(cards *[]int) []int {
    pcards := *cards
    slice := []int{}
    for _, v := range pcards {
        slice = append(slice, v)
        slice = append(slice, v-1)
        slice = append(slice, v+1)
    }
    a := removeDup(slice)
    return a
}

//先癞子替换再牌型验证
func cardsType(poker *[]byte) {
    cards := getPokerValue(poker)
    slice := laiziIndex(cards)
    lzValue := laiZiValue(cards)
    pcards := *cards
    
    switch len(slice) {
    case 0:
        checkCardsType(cards)

    case 1:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            checkCardsType(cards)
        }

    case 2:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                checkCardsType(cards)
            }
        }

    case 3:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    checkCardsType(cards)
                }
            }
        }

    case 4:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    for _, v3 := range lzValue {
                        pcards[slice[3]] = v3
                        checkCardsType(cards)
                    }
                }
            }
        }
    case 5:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    for _, v3 := range lzValue {
                        pcards[slice[3]] = v3
                        for _, v4 := range lzValue {
                            pcards[slice[4]] = v4
                            checkCardsType(cards)
                        }
                    }
                }
            }
        }
    case 6:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    for _, v3 := range lzValue {
                        pcards[slice[3]] = v3
                        for _, v4 := range lzValue {
                            pcards[slice[4]] = v4
                            for _, v5 := range lzValue {
                                pcards[slice[5]] = v5
                                checkCardsType(cards)
                            }
                        }
                    }
                }
            }
        }
    case 7:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    for _, v3 := range lzValue {
                        pcards[slice[3]] = v3
                        for _, v4 := range lzValue {
                            pcards[slice[4]] = v4
                            for _, v5 := range lzValue {
                                pcards[slice[5]] = v5
                                for _, v6 := range lzValue {
                                    pcards[slice[6]] = v6
                                    checkCardsType(cards)
                                }
                            }
                        }
                    }
                }
            }
        }
    case 8:
        for _, v := range lzValue {
            pcards[slice[0]] = v
            for _, v1 := range lzValue {
                pcards[slice[1]] = v1
                for _, v2 := range lzValue {
                    pcards[slice[2]] = v2
                    for _, v3 := range lzValue {
                        pcards[slice[3]] = v3
                        for _, v4 := range lzValue {
                            pcards[slice[4]] = v4
                            for _, v5 := range lzValue {
                                pcards[slice[5]] = v5
                                for _, v6 := range lzValue {
                                    pcards[slice[6]] = v6
                                    for _, v7 := range lzValue {
                                        pcards[slice[7]] = v7
                                        checkCardsType(cards)
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}

//取牌值
func getPokerValue(poker *[]byte) *[]int {
    ppoker := *poker
    newPoker := make([]int, len(ppoker))
    for i, _ := range ppoker {
        newPoker[i] = int(ppoker[i] & 0x0f)
    }
    return &newPoker
}

//炸弹
func isBoom(cards *[]int) bool {
    pcards := *cards
    a := pcards[0]
    for _, v := range pcards {
        if v == a {
            continue
        } else {
            return false
        }
    }
    return true
}

//连对
func isEvenPair(poker *[]int) bool {
    cards := descend(poker)
    if isIncludeJokerTwo(&cards) {
        return false
    } else {
        l := len(cards)
        if l > 3 && l < 25 && (l % 2 == 0) {
            for i := 0; i < l-1; {
                if cards[i] == cards[i+1] {
                    i += 2
                } else {
                    return false
                }
            }
        
            for i := 0; i < l-3; {
                if cards[i] - cards[i+2] == 1 {
                    i += 2
                } else {
                    return false
                }
            }
        } else {
            return false
        }
    }
    return true
}

//顺子
func isStraight(poker *[]int) bool {
    cards := descend(poker)
    if isIncludeJokerTwo(&cards) {
        return false
    } else {
        for i, _ := range cards {
            if cards[i] == 1 {
                cards[i] = 14
            }
        }
        card := descend(&cards)
        for i := 0; i < len(card)-1; i++ {
            if (card[i] - card[i+1]) == 1 {
                continue
            } else {
                return false
            }
        }
    }
    return true
}

//飞机不带
func isPlane(poker *[]int) bool {
    cards := descend(poker)
    if isIncludeJokerTwo(&cards) {
        return false
    } else {
        l := len(cards)
        if l > 5 && l < 25 && (l % 3 == 0) {
            for i := 0; i < l-2; {
                a := cards[i]
                b := cards[i+1]
                c := cards[i+2]
                if a == b && a == c{
                    i += 3
                } else {
                    return false
                }
            }
        
            for i := 0; i < l-5; {
                if cards[i] - cards[i+3] == 1 {
                    i += 3
                } else {
                    return false
                }
            }
        } else {
            return false
        }
    }
    return true
}

//三带单
func isThreeOne(poker *[]int) bool {
    cards := descend(poker)
    a := cards[0] == cards[1]
    b := cards[1] == cards[2]
    c := cards[2] == cards[3]
    if a && b && !c {
        return true
    } else if !a && b && c {
        return true
    }
    return false
}

//三带对
func isThreePair(poker *[]int) bool {
    cards := descend(poker)
    a := cards[0] == cards[1]
    b := cards[1] == cards[2]
    c := cards[2] == cards[3]
    d := cards[3] == cards[4]
    if a && b && !c && d {
        return true
    } else if a && !b && c && d {
        return true
    }
    return false
}

//四带二单
func isFourTwo(poker *[]int) bool {
    cards := descend(poker)
    a := cards[0] == cards[1]
    b := cards[1] == cards[2]
    c := cards[2] == cards[3]
    d := cards[3] == cards[4]
    e := cards[4] == cards[5]
    if a && b && c && !d && !e {
        return true
    } else if !a && b && c && d && !e {
        return true
    } else if !a && !b && c && d && e {
        return true
    }
    return false
}

//飞机带二单
func isPlaneTwo(poker *[]int) bool {
    cards := descend(poker)
    a := cards[0] == cards[1]
    b := cards[1] == cards[2]
    c := cards[2] == cards[3]
    d := cards[3] == cards[4]
    e := cards[4] == cards[5]
    f := cards[5] == cards[6]
    g := cards[6] == cards[7]
    
    dif1 := cards[2] - cards[3]
    dif2 := cards[3] - cards[4]
    dif3 := cards[4] - cards[5]
    
    if a && b && dif1 == 1 && d && e && !f && !g {
        return true
    } else if !a && b && c && dif2 == 1 && e && f && !g {
        return true
    } else if !a && !b && c && d && dif3 == 1 && f && g {
        return true
    }
    return false
}

//四带二对
func isFourPair(poker *[]int) bool {
    cards := descend(poker)
    a := cards[0] == cards[1]
    b := cards[1] == cards[2]
    c := cards[2] == cards[3]
    d := cards[3] == cards[4]
    e := cards[4] == cards[5]
    f := cards[5] == cards[6]
    g := cards[6] == cards[7]
    
    if a && b && c && !d && e && !f && g {
        return true
    } else if a && !b && c && d && e && !f && g {
        return true
    } else if a && !b && c && !d && e && f && g {
        return true
    }
    return false
}

//飞机带二对
func isPlanePair(poker *[]int) bool {
    if len(*poker) == 10 {
        cards := descend(poker)
        a := cards[0] == cards[1]
        b := cards[1] == cards[2]
        c := cards[2] == cards[3]
        d := cards[3] == cards[4]
        e := cards[4] == cards[5]
        f := cards[5] == cards[6]
        g := cards[6] == cards[7]
        h := cards[7] == cards[8]
        i := cards[8] == cards[9]
        
        dif1 := cards[2] - cards[3]  //AAABBB CC DD
        dif2 := cards[4] - cards[5]  //AA BBBCCC DD
        dif3 := cards[6] - cards[7]  //AA BB CCCDDD
        
        if a && b && dif1 == 1 && d && e && !f && g && !h && i {
            return true
        } else if a && !b && c && d && dif2 == 1 && f && g && !h && i {
            return true
        } else if a && !b && c && !d && e && f && dif3 == 1 && h && i {
            return true
        }
    }
    return false
}

//三飞机带三单
func isThreePlaneThree(poker *[]int) bool {
    if len(*poker) == 12 {
        cards := descend(poker)
        a := cards[0] == cards[1]
        b := cards[1] == cards[2]
        c := cards[2] == cards[3]
        d := cards[3] == cards[4]
        e := cards[4] == cards[5]
        f := cards[5] == cards[6]
        g := cards[6] == cards[7]
        h := cards[7] == cards[8]
        i := cards[8] == cards[9]
        j := cards[9] == cards[10]
        k := cards[10] == cards[11]
        
        dif1 := cards[2] - cards[3]
        dif2 := cards[5] - cards[6]
        
        dif3 := cards[3] - cards[4]
        dif4 := cards[6] - cards[7]
        
        dif5 := cards[4] - cards[5]
        dif6 := cards[7] - cards[8]
        
        dif7 := cards[5] - cards[6]
        dif8 := cards[8] - cards[9]
        
        if a && b && dif1==1 && d && e && dif2==1 && g && h && !i && !j && !k {
            return true
        } else if !a && b && c && dif3==1 && e && f && dif4==1 && h && i && !j && !k {
            return true
        } else if !a && !b && c && d && dif5==1 && f && g && dif6==1 && i && j && !k {
            return true
        } else if !a && !b && !c && d && e && dif7==1 && g && h && dif8==1 && j && k {
            return true
        }
    }
    return false
}

//三飞机带三对
func isThreePlanePair(poker *[]int) bool {
    if len(*poker) == 15 {
        cards := descend(poker)
        a := cards[0] == cards[1]
        b := cards[1] == cards[2]
        c := cards[2] == cards[3]
        d := cards[3] == cards[4]
        e := cards[4] == cards[5]
        f := cards[5] == cards[6]
        g := cards[6] == cards[7]
        h := cards[7] == cards[8]
        i := cards[8] == cards[9]
        j := cards[9] == cards[10]
        k := cards[10] == cards[11]
        l := cards[11] == cards[12]
        m := cards[12] == cards[13]
        n := cards[13] == cards[14]
        
        dif1 := cards[2] - cards[3]
        dif2 := cards[5] - cards[6]
        
        dif3 := cards[4] - cards[5]
        dif4 := cards[7] - cards[8]
        
        dif5 := cards[6] - cards[7]
        dif6 := cards[9] - cards[10]
        
        dif7 := cards[8] - cards[9]
        dif8 := cards[11] - cards[12]
        
        if a && b && dif1==1 && d && e && dif2==1 && g && h && !i && j && !k && l && !m && n {
            return true
        } else if a && !b && c && d && dif3==1 && f && g && dif4==1 && i && j && !k && l && !m && n {
            return true
        } else if a && !b && c && !d && e && f && dif5==1 && h && i && dif6==1 && k && l && !m && n {
            return true
        } else if a && !b && c && !d && e && !f && g && h && dif7==1 && j && k && dif8==1 && m && n {
            return true
        }
    }
    return false
}

//排序
func descend(cards *[]int) []int{
    pcards := *cards
    slice := make([]int, len(pcards)) //不能改变原始牌的顺序，根据原始牌中癞子牌的下标进行赋值，每赋值
    for i, _ := range pcards {        //一次，做一次排序、牌型判断
        slice[i] = pcards[i]
    }
    sort.Sort(sort.Reverse(sort.IntSlice(slice)))
    return slice
}

//是否有大王或小王或2
func isIncludeJokerTwo(cards *[]int) bool {
    pcards := *cards
    for _, v := range pcards {
        if v == 2 {
            return true
        } else if v == 14 {
            return true
        } else if v == 15 {
            return true
        }
    }
    return false
}

//确定癞子牌的下标,7为癞子
func laiziIndex(cards *[]int) []int {
    pcards := *cards
    slice := make([]int, 0)
    for i, _ := range pcards {
        if pcards[i] == 7 {
            slice = append(slice, i)
        }
    }
    return slice
}
