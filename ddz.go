package main

import (
    "fmt"
    "time"
    "math/rand"
    "strconv"
    //"sort"
    //"reflect"
)

const (
    playerNum = 4  //玩家数量
    cardsNum = 2   //牌副数
    laiziValue = 7 //癞子牌
)

type pokerCards struct {
    COLOR string
    SYMBOL string
    VALUE int
    LAIZI bool
}

//一副牌
type playCards struct {
    cards [54]pokerCards
}

type handCards struct {
    playerCards []*pokerCards
    bottomCards []*pokerCards
}

type showHandCards struct {
    playerCards [17]pokerCards
    bottomCards [3]pokerCards
    lordCards [20]pokerCards
}

//两副牌
type moreCards []*pokerCards

type mPlayCards struct {
    cards [108]pokerCards
}

type showHandCardsMore struct {
    playerCards1 [26]pokerCards
    playerCards2 [25]pokerCards
    bottomCards [6]pokerCards
    lordCards [32]pokerCards
}

func main() {
    res := CreateNew()
    fmt.Println("新牌：\n", res, "\n")
    
    ret := *res.shuffleCards()
    fmt.Println("洗牌后：\n", ret, "\n")
    //fmt.Println("r type:", reflect.TypeOf(ret))
    
    //发牌
    afterDealCards(ret)
    
    //两副牌
    cm := createMore(cardsNum, res)
    var mp = new(mPlayCards)
    mcards := mp.read(cm)
    fmt.Println("两副牌：", len(mcards.cards), "张\n", mcards, "\n")
    
    //mpshuffle := *mcards.shuffleCardsMore()
    //fmt.Println("洗牌后：", len(mpshuffle.cards), "张\n", mpshuffle, "\n")
    
    //发牌
    afterDealCardsMore(mcards)
    
    //有癞子的牌
    lzCards := *mcards.chooseLaiZi(laiziValue)
    fmt.Println("癞子是", laiziValue, "的牌：", len(lzCards.cards), "张\n", lzCards)
    
    //发牌
    afterDealCardsMore(lzCards)
}

//构造牌
func CreateNew() playCards {
    play := playCards{}
    color := ""
    value := 0
    symbol := ""
    for i := 0; i <= 53; i++ {
        if i == 52 {
            play.cards[i].COLOR = "Small"
            play.cards[i].VALUE = 16      //小王
            play.cards[i].SYMBOL = "S"
        } else if i == 53 {
            play.cards[i].COLOR = "Big"
            play.cards[i].VALUE = 17      //大王
            play.cards[i].SYMBOL = "B"
        } else {
            quotient := i/13
            remainder := i%13
            switch quotient {
            case 0:
                color = "Spade"
                value = remainder + 3
                if i == 11 {
                    symbol = "A"
                } else if i == 12 {
                    symbol = "2"
                } else {
                    symbol = strconv.Itoa(remainder+3)
                }
            case 1:
                color = "Heart"
                value = remainder + 3
                if i == 24 {
                    symbol = "A"
                } else if i == 25 {
                    symbol = "2"
                } else {
                    symbol = strconv.Itoa(remainder+3)
                }
            case 2:
                color = "Club"
                value = remainder + 3
                if i == 37 {
                    symbol = "A"
                } else if i == 38 {
                    symbol = "2"
                } else {
                    symbol = strconv.Itoa(remainder+3)
                }
            case 3:
                color = "Diamond"
                value = remainder + 3
                if i == 50 {
                    symbol = "A"
                } else if i == 51 {
                    symbol = "2"
                } else {
                    symbol = strconv.Itoa(remainder+3)
                }
            }
            play.cards[i].COLOR = color
            play.cards[i].VALUE = value
            play.cards[i].SYMBOL = symbol
        }
    }
    return play
}

//洗牌
func (p *playCards) shuffleCards() *playCards {
	rand.Seed(time.Now().Unix())
	for i := len(p.cards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		p.cards[i], p.cards[num] = p.cards[num], p.cards[i]
	}
    return p
}

//发牌
func (h *handCards) dealCards(order int, vals playCards) []*pokerCards {
    //var playerCards []int
    //var bottomCards []int
    for i := 0; i < len(vals.cards); i++ {
        if i > 50 {
            h.bottomCards = append(h.bottomCards, &vals.cards[i])
            bubbleSort(h.bottomCards)
        } else {
            switch i%3 == order {
            case true:
                h.playerCards = append(h.playerCards, &vals.cards[i])
                bubbleSort(h.playerCards)
            }
        }
    }
    switch order {
    case 0, 1, 2:
        return h.playerCards
    }
    return h.bottomCards
}

//默认玩家1为地主
func diZhu(hc1 []*pokerCards, hc4 []*pokerCards) []*pokerCards {
    for _, i := range hc4 {
        hc1 = append(hc1, i)
        bubbleSort(hc1)
    }
    return hc1
}

//各玩家牌和底牌
func afterDealCards(ret playCards) {
    var show = new(showHandCards)
    
    var h1 = new(handCards)
    h1.playerCards = h1.dealCards(0, ret)
    rpc1 := show.readPlayerCards(h1.playerCards)
    fmt.Println("玩家1的牌：\n", rpc1, "\n")
    
    var h2 = new(handCards)
    h2.playerCards = h2.dealCards(1, ret)
    rpc2 := show.readPlayerCards(h2.playerCards)
    fmt.Println("玩家2的牌：\n", rpc2, "\n")
    
    var h3 = new(handCards)
    h3.playerCards = h3.dealCards(2, ret)
    rpc3 := show.readPlayerCards(h3.playerCards)
    fmt.Println("玩家3的牌：\n", rpc3, "\n")
    
    var h4 = new(handCards)
    h4.bottomCards = h4.dealCards(3, ret)
    //checkCardsType(h4.bottomCards)
    rbc := show.readBottomCards(h4.bottomCards)
    fmt.Println("底牌：", rbc, "\n")
    
    diZhuCards := diZhu(h1.playerCards, h4.bottomCards)
    rlc := show.readLordCards(diZhuCards)
    fmt.Println("地主牌：\n", rlc, "\n")
}

//通过地址读取值(一副)
func (show *showHandCards) readPlayerCards(poker []*pokerCards) [17]pokerCards {
    for i, c := range poker {
        show.playerCards[i] = *c
    }
    return show.playerCards
}

func (show *showHandCards) readBottomCards(poker []*pokerCards) [3]pokerCards {
    for i, c := range poker {
        show.bottomCards[i] = *c
    }
    return show.bottomCards
}

func (show *showHandCards) readLordCards(poker []*pokerCards) [20]pokerCards {
    for i, c := range poker {
        show.lordCards[i] = *c
    }
    return show.lordCards
}

//两副牌
//构造两副牌
func (m moreCards) addCards(ms moreCards) moreCards {
    for _, card := range ms {
        m = append(m, card)
    }
    return m
}

func createMore(num int, p playCards) moreCards {
    m := moreCards{}
    for n := 0; n < num; n++ {
        for i, _ := range p.cards {
            m = m.addCards(moreCards{&p.cards[i]})
        }
    }
    return m
}

func (mp mPlayCards) read(m moreCards) mPlayCards {
    for i, c := range m {
        mp.cards[i] = *c
    }
    return mp
}

//洗牌
func (mp *mPlayCards) shuffleCardsMore() *mPlayCards {
	rand.Seed(time.Now().Unix())
	for i := len(mp.cards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		mp.cards[i], mp.cards[num] = mp.cards[num], mp.cards[i]
	}
    return mp
}

//发牌
func (h *handCards) dealCardsMore(order int, playerNum int, vals mPlayCards) []*pokerCards {
    for i := 0; i < len(vals.cards); i++ {
        if i > 50 + 51 * (cardsNum - 1) {
            h.bottomCards = append(h.bottomCards, &vals.cards[i])
            bubbleSort(h.bottomCards)
        } else {
            switch i % playerNum == order {
            case true:
                h.playerCards = append(h.playerCards, &vals.cards[i])
                bubbleSort(h.playerCards)
            }
        }
    }
    switch order {
    case playerNum:
        return h.bottomCards
    }
    return h.playerCards
}

//各玩家的牌和底牌（两副）
func afterDealCardsMore(ret mPlayCards) {
    var show = new(showHandCardsMore)
    
    var h1 = new(handCards)
    h1.playerCards = h1.dealCardsMore(0, playerNum, ret)
    rpc1 := show.readPlayerCardsMore1(h1.playerCards)
    fmt.Println("玩家1的牌：", len(h1.playerCards), "张\n", rpc1, "\n")
    
    var h2 = new(handCards)
    h2.playerCards = h2.dealCardsMore(1, playerNum, ret)
    rpc2 := show.readPlayerCardsMore1(h2.playerCards)
    fmt.Println("玩家2的牌：", len(h2.playerCards), "张\n", rpc2, "\n")
    
    var h3 = new(handCards)
    h3.playerCards = h3.dealCardsMore(2, playerNum, ret)
    rpc3 := show.readPlayerCardsMore2(h3.playerCards)
    fmt.Println("玩家3的牌：", len(h3.playerCards), "张\n", rpc3, "\n")
    
    var h4 = new(handCards)
    h4.playerCards = h4.dealCardsMore(3, playerNum, ret)
    rpc4 := show.readPlayerCardsMore2(h4.playerCards)
    fmt.Println("玩家4的牌：", len(h4.playerCards), "张\n", rpc4, "\n")
    
    //重新赋值测试(最多24张)
    h4.playerCards[:][0].COLOR = "Spade"
    h4.playerCards[:][0].SYMBOL = "3"
    h4.playerCards[:][0].VALUE = 3
    h4.playerCards[:][0].LAIZI = false
    h4.playerCards[:][1].COLOR = "Spade"
    h4.playerCards[:][1].SYMBOL = "4"
    h4.playerCards[:][1].VALUE = 4
    h4.playerCards[:][1].LAIZI = false
    h4.playerCards[:][2].COLOR = "Spade"
    h4.playerCards[:][2].SYMBOL = "3"
    h4.playerCards[:][2].VALUE = 3
    h4.playerCards[:][2].LAIZI = false
    h4.playerCards[:][3].COLOR = "Spade"
    h4.playerCards[:][3].SYMBOL = "4"
    h4.playerCards[:][3].VALUE = 4
    h4.playerCards[:][3].LAIZI = false
    h4.playerCards[:][4].COLOR = "Spade"
    h4.playerCards[:][4].SYMBOL = "7"
    h4.playerCards[:][4].VALUE = 7
    h4.playerCards[:][4].LAIZI = true
    h4.playerCards[:][5].COLOR = "Spade"
    h4.playerCards[:][5].SYMBOL = "8"
    h4.playerCards[:][5].VALUE = 8
    h4.playerCards[:][5].LAIZI = false
    h4.playerCards[:][6].COLOR = "Spade"
    h4.playerCards[:][6].SYMBOL = "9"
    h4.playerCards[:][6].VALUE = 9
    h4.playerCards[:][6].LAIZI = false
    h4.playerCards[:][7].COLOR = "Spade"
    h4.playerCards[:][7].SYMBOL = "8"
    h4.playerCards[:][7].VALUE = 8
    h4.playerCards[:][7].LAIZI = false
    h4.playerCards[:][8].COLOR = "Spade"
    h4.playerCards[:][8].SYMBOL = "9"
    h4.playerCards[:][8].VALUE = 9
    h4.playerCards[:][8].LAIZI = false
    h4.playerCards[:][9].COLOR = "Spade"
    h4.playerCards[:][9].SYMBOL = "7"
    h4.playerCards[:][9].VALUE = 7
    h4.playerCards[:][9].LAIZI = true
    
    rpctest := show.readPlayerCardsMore2(h4.playerCards)
    fmt.Println("测试牌：\n", rpctest[:10])
    /*
    checkCardsType(h4.playerCards[:1])
    checkCardsType(h4.playerCards[:2])
    checkCardsType(h4.playerCards[:3])
    checkCardsType(h4.playerCards[:4])
    checkCardsType(h4.playerCards[:5])
    checkCardsType(h4.playerCards[:6])
    checkCardsType(h4.playerCards[:7])
    checkCardsType(h4.playerCards[:8])
    checkCardsType(h4.playerCards[:9])
    */
    checkCardsType(h4.playerCards[:10])
    fmt.Println("癞子组成的牌是：\n", *h4.playerCards[:10][0], *h4.playerCards[:10][1], *h4.playerCards[:10][2], *h4.playerCards[:10][3], *h4.playerCards[:10][4], *h4.playerCards[:10][5], *h4.playerCards[:10][6], *h4.playerCards[:10][7], *h4.playerCards[:10][8], *h4.playerCards[:10][9], "\n")
    
    var h5 = new(handCards)
    h5.bottomCards = h5.dealCardsMore(4, playerNum, ret)
    rbc := show.readBottomCardsMore(h5.bottomCards)
    fmt.Println("底牌：", rbc, "\n")

    //重新赋值测试（最多6张）
    h5.bottomCards[:][0].COLOR = "Spade"
    h5.bottomCards[:][0].SYMBOL = "10"
    h5.bottomCards[:][0].VALUE = 10
    h5.bottomCards[:][0].LAIZI = false
    h5.bottomCards[:][1].COLOR = "Spade"
    h5.bottomCards[:][1].SYMBOL = "10"
    h5.bottomCards[:][1].VALUE = 10
    h5.bottomCards[:][1].LAIZI = false
    h5.bottomCards[:][2].COLOR = "Heart"
    h5.bottomCards[:][2].SYMBOL = "7"
    h5.bottomCards[:][2].VALUE = 7
    h5.bottomCards[:][2].LAIZI = true
    h5.bottomCards[:][3].COLOR = "Spade"
    h5.bottomCards[:][3].SYMBOL = "11"
    h5.bottomCards[:][3].VALUE = 11
    h5.bottomCards[:][3].LAIZI = false
    h5.bottomCards[:][4].COLOR = "Heart"
    h5.bottomCards[:][4].SYMBOL = "11"
    h5.bottomCards[:][4].VALUE = 11
    h5.bottomCards[:][4].LAIZI = false
    h5.bottomCards[:][5].COLOR = "Spade"
    h5.bottomCards[:][5].SYMBOL = "7"
    h5.bottomCards[:][5].VALUE = 7
    h5.bottomCards[:][5].LAIZI = true
    rbctest := show.readBottomCardsMore(h5.bottomCards)
    fmt.Println("测试牌：\n", rbctest)

    //checkCardsType(h5.bottomCards[:1])
    
    //checkCardsType(h5.bottomCards[:2])
    
    //checkCardsType(h5.bottomCards[:3])
    
    //checkCardsType(h5.bottomCards[:4])
    
    //checkCardsType(h5.bottomCards[:5])
    
    //判断牌型
    checkCardsType(h5.bottomCards)
    //打印癞子值改变后组成的牌型
    fmt.Println("癞子组成的牌是：\n", *h5.bottomCards[:][0], *h5.bottomCards[:][1], *h5.bottomCards[:][2], *h5.bottomCards[:][3], *h5.bottomCards[:][4], *h5.bottomCards[:][5], "\n")
    
    diZhuCards := diZhu(h1.playerCards, h5.bottomCards)
    rlc := show.readLordCardsMore(diZhuCards)
    fmt.Println("地主牌：", len(diZhuCards), "张\n", rlc, "\n")
}

//通过地址读取值(两副)
func (show *showHandCardsMore) readPlayerCardsMore1(poker []*pokerCards) [26]pokerCards {
    for i, c := range poker {
        show.playerCards1[i] = *c
    }
    return show.playerCards1
}

func (show *showHandCardsMore) readPlayerCardsMore2(poker []*pokerCards) [25]pokerCards {
    for i, c := range poker {
        show.playerCards2[i] = *c
    }
    return show.playerCards2
}

func (show *showHandCardsMore) readBottomCardsMore(poker []*pokerCards) [6]pokerCards {
    for i, c := range poker {
        show.bottomCards[i] = *c
    }
    return show.bottomCards
}

func (show *showHandCardsMore) readLordCardsMore(poker []*pokerCards) [32]pokerCards {
    for i, c := range poker {
        show.lordCards[i] = *c
    }
    return show.lordCards
}

//是否有大王或小王或2
func isIncludeJokerTwo(cards []*pokerCards) bool {
    for _, c := range cards {
        if c.VALUE == 16 {
            return true
        } else if c.VALUE == 17 {
            return true
        } else if c.VALUE == 15 {
            return true
        }
    }
    return false
}

//是否有大小王
func isIncludeJoker(cards []*pokerCards) bool {
    for _, c := range cards {
        if c.VALUE == 16 {
            return true
        } else if c.VALUE == 17 {
            return true
        }
    }
    return false
}

//判断牌型前先排序 降序
func bubbleSort(cards []*pokerCards) []*pokerCards {
    length := len(cards)
    slice := make([]*pokerCards, length) //不能改变原始牌的顺序，根据原始牌中癞子牌的下标进行赋值，每赋值
    for i := 0; i < length; i++ {        //一次，做一次排序、牌型判断
        slice[i] = cards[i]
    }
    for i := 0; i < length-1; i++ {
        for j := 0; j < length - i -1; j++ {
            if slice[j].VALUE < slice[j+1].VALUE {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
    return slice
}

//判断牌型
func checkCardsType(cards []*pokerCards) {
    switch len(cards) {
    case 0:
        fmt.Println("没有牌\n")
    case 1:
        fmt.Println("单张\n")
    case 2:
        if cards[0].VALUE == 17 && cards[1].VALUE == 16 {
            fmt.Print("王炸\n")
        } else if cards[0].VALUE == cards[1].VALUE {
            fmt.Println("对子\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 3:
        if cards[0].VALUE == cards[1].VALUE && cards[1].VALUE == cards[2].VALUE {
            fmt.Print("三不带\n")
        } else {
            fmt.Println("不符合规则\n")
        }
    case 4:
        //连队
        if isEvenPair(cards) {
            fmt.Println("连对\n")
        //炸弹
        } else if cards[0].VALUE == cards[1].VALUE && cards[1].VALUE == cards[2].VALUE && cards[2].VALUE == cards[3].VALUE {
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
        } else {
            fmt.Println("不符合规则\n")
        }
    case 6:
        if isThreePlane(cards) {
            fmt.Println("飞机不带\n")
        } else if isStraight(cards) {
            fmt.Println("顺子\n")
        } else if isFourTwo(cards) {
            fmt.Println("四带二单\n")
        } else if isEvenPair(cards) {
            fmt.Println("连对")
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
        }
        
        if isEvenPair(cards) {
            fmt.Println("连对\n")
        }
        
        if isPlanePair(cards) {
            fmt.Println("飞机带二对\n")
        } else {
            fmt.Println("不符合规则")
        }
    }
}

//确定癞子牌的下标
func laiziIndex(cards []*pokerCards) []int {
    slice := make([]int, 0)
    for i, _ := range cards {
        if cards[i].LAIZI == true {
            slice = append(slice, i)
        }
    }
    return slice
}

//没有大小王、2、癞子的情况下是不是连对
func isOfEvenPair(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    l := len(cardSlice)
    if l > 3 && l < 25 && (l % 2 == 0) {
        for i := 0; i < l-1; {
            if cardSlice[i].VALUE == cardSlice[i+1].VALUE {
                i += 2
            } else {
                return false
            }
        }
    
        for i := 0; i < l-3; {
            if cardSlice[i].VALUE - cardSlice[i+2].VALUE == 1 {
                i += 2
            } else {
                return false
            }
        }
    } else {
        return false
    }
    return true
}

//是否连对
func isEvenPair(cards []*pokerCards) bool {
    if isIncludeJokerTwo(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfEvenPair(cards) {
                return true
            }
            return false

        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfEvenPair(cards) {
                    return true
                }
            }
            

        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfEvenPair(cards) {
                        return true
                    }
                }
            }

        case 3:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfEvenPair(cards) {
                            return true
                        }
                    }
                }
            }

        case 4:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        for l := 3; l < 15; l++ {
                            cards[slice[3]].VALUE = l
                            if isOfEvenPair(cards) {
                                return true
                            }
                        }
                    }
                }
            }
        }
    }
    return false
}

//没有大小王、2、癞子的情况下是不是顺子
func isOfStraight(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    for i := 0; i < len(cardSlice)-1; i++ {
        if (cardSlice[i].VALUE - cardSlice[i+1].VALUE) == 1 {
            continue
        } else {
            return false
        }
    }
    return true
}

//是否顺子
func isStraight(cards []*pokerCards) bool {
    if isIncludeJokerTwo(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfStraight(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfStraight(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfStraight(cards) {
                        return true
                    }
                }
            }
        case 3:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfStraight(cards) {
                            return true
                        }
                    }
                }
            }
        case 4:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        for l := 3; l < 15; l++ {
                            cards[slice[3]].VALUE = l
                            if isOfStraight(cards) {
                                return true
                            }
                        }
                    }
                }
            }
        }
    }
    return false
}

//没有大小王、2、癞子的情况下判断是不是飞机不带
func isOfThreePlane(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    l := len(cardSlice)
    if l > 5 && l < 25 && (l % 3 == 0) {
        for i := 0; i < l-2; {
            a := cardSlice[i].VALUE
            b := cardSlice[i+1].VALUE
            c := cardSlice[i+2].VALUE
            if a == b && a == c{
                i += 3
            } else {
                return false
            }
        }
    
        for i := 0; i < l-5; {
            if cardSlice[i].VALUE - cardSlice[i+3].VALUE == 1 {
                i += 3
            } else {
                return false
            }
        }
    } else {
        return false
    }
    return true
}

//是否飞机不带
func isThreePlane(cards []*pokerCards) bool {
    if isIncludeJokerTwo(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfThreePlane(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfThreePlane(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j <15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfThreePlane(cards) {
                        return true
                    }
                }
            }
        case 3:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfThreePlane(cards) {
                            return true
                        }
                    }
                }
            }
        }
    }
    return false
}

//没有癞子情况下是否三带单
func isOfThreeOne(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    if a && b && !c {
        return true
    } else if !a && b && c {
        return true
    }
    return false
}

//是否三带单
func isThreeOne(cards []*pokerCards) bool {
    slice := laiziIndex(cards)
    switch len(slice) {
    case 0:
        if isOfThreeOne(cards) {
            return true
        }
    case 1:
        for i := 3; i < 15; i++ {
            cards[slice[0]].VALUE = i
            if isOfThreeOne(cards) {
                return true
            }
        }
    case 2:
        for i := 3; i < 15; i++ {
            cards[slice[0]].VALUE = i
            for j := 3; j < 15; j++ {
                cards[slice[1]].VALUE = j
                if isOfThreeOne(cards) {
                    return true
                }
            }
        }
    }
    return false
}

//没有癞子、大小王情况下是否三带对
func isOfThreePair(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    d := cardSlice[3].VALUE == cardSlice[4].VALUE
    if a && b && !c && d {
        return true
    } else if a && !b && c && d {
        return true
    }
    return false
}

//是否三带对
func isThreePair(cards []*pokerCards) bool {
    if isIncludeJoker(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfThreePair(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfThreePair(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfThreePair(cards) {
                        return true
                    }
                }
            }
        }
    }
    return false
}

//没有大小王、癞子的情况下是不是四带二单
func isOfFourTwo(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    d := cardSlice[3].VALUE == cardSlice[4].VALUE
    e := cardSlice[4].VALUE == cardSlice[5].VALUE
    if a && b && c && !d && !e {
        return true
    } else if !a && b && c && d && !e {
        return true
    } else if !a && !b && c && d && e {
        return true
    }
    return false
}

//是否四带二单
func isFourTwo(cards []*pokerCards) bool {
    if isIncludeJoker(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfFourTwo(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfFourTwo(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfFourTwo(cards) {
                        return true
                    }
                }
            }
        }
    }
    return false
}

//没有大小王、癞子的情况下是否飞机带二单
func isOfPlaneTwo(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    d := cardSlice[3].VALUE == cardSlice[4].VALUE
    e := cardSlice[4].VALUE == cardSlice[5].VALUE
    f := cardSlice[5].VALUE == cardSlice[6].VALUE
    g := cardSlice[6].VALUE == cardSlice[7].VALUE
    
    dif1 := cardSlice[2].VALUE - cardSlice[3].VALUE
    dif2 := cardSlice[3].VALUE - cardSlice[4].VALUE
    dif3 := cardSlice[4].VALUE - cardSlice[5].VALUE
    
    if a && b && dif1 == 1 && d && e && !f && !g {
        return true
    } else if !a && b && c && dif2 == 1 && e && f && !g {
        return true
    } else if !a && !b && c && d && dif3 == 1 && f && g {
        return true
    }
    return false
}

//是否飞机带二单
func isPlaneTwo(cards []*pokerCards) bool {
    if isIncludeJoker(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfPlaneTwo(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfPlaneTwo(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfPlaneTwo(cards) {
                        return true
                    }
                }
            }
        case 3:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfPlaneTwo(cards) {
                            return true
                        }
                    }
                }
            }
        }
    }
    return false
}

//没有大小王、癞子的情况下是否四带二对
func isOfFourPair(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    d := cardSlice[3].VALUE == cardSlice[4].VALUE
    e := cardSlice[4].VALUE == cardSlice[5].VALUE
    f := cardSlice[5].VALUE == cardSlice[6].VALUE
    g := cardSlice[6].VALUE == cardSlice[7].VALUE
    
    if a && b && c && !d && e && !f && g {
        return true
    } else if a && !b && c && d && e && !f && g {
        return true
    } else if a && !b && c && !d && e && f && g {
        return true
    }
    return false
}

//是否四带二对
func isFourPair(cards []*pokerCards) bool {
    if isIncludeJoker(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfFourPair(cards) {
                return true
            }
        case 1:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfFourPair(cards) {
                    return true
                }
            }
        case 2:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfFourPair(cards) {
                        return true
                    }
                }
            }
        case 3:
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfFourPair(cards) {
                            return true
                        }
                    }
                }
            }
        }
    }
    return false
}

/*
//判断切片内元素是否相同
func isEqual(p []*pokerCards) bool {
    firstValue := p[0].VALUE
    for i := 1; i < len(p); i++ {
        if p[i].VALUE == firstValue {
            continue
        } else {
            return false
        }
    }
    return true
}
*/

//没有大小王、癞子情况下是否飞机带二对
func isOfPlanePair(cards []*pokerCards) bool {
    cardSlice := bubbleSort(cards)
    a := cardSlice[0].VALUE == cardSlice[1].VALUE
    b := cardSlice[1].VALUE == cardSlice[2].VALUE
    c := cardSlice[2].VALUE == cardSlice[3].VALUE
    d := cardSlice[3].VALUE == cardSlice[4].VALUE
    e := cardSlice[4].VALUE == cardSlice[5].VALUE
    f := cardSlice[5].VALUE == cardSlice[6].VALUE
    g := cardSlice[6].VALUE == cardSlice[7].VALUE
    h := cardSlice[7].VALUE == cardSlice[8].VALUE
    i := cardSlice[8].VALUE == cardSlice[9].VALUE
    
    dif1 := cardSlice[2].VALUE - cardSlice[3].VALUE  //AAABBB CC DD
    dif2 := cardSlice[4].VALUE - cardSlice[5].VALUE  //AA BBBCCC DD
    dif3 := cardSlice[6].VALUE - cardSlice[7].VALUE  //AA BB CCCDDD
    
    if a && b && dif1 == 1 && d && e && !f && g && !h && i {
        return true
    } else if a && !b && c && d && dif2 == 1 && f && g && !h && i {
        return true
    } else if a && !b && c && !d && e && f && dif3 == 1 && h && i {
        return true
    }
    return false
}

//是否飞机带二对
func isPlanePair(cards []*pokerCards) bool {
    if isIncludeJoker(cards) {
        return false
    } else {
        slice := laiziIndex(cards)
        switch len(slice) {
        case 0:
            if isOfPlanePair(cards) {
                return true
            }

        case 1:
            //changeLaiZiValue(slice, isOfPlanePair, cards)
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                if isOfPlanePair(cards) {
                    return true
                }
            }
            
        case 2:
            //changeLaiZiValue2(slice, isOfPlanePair, cards)
            
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    if isOfPlanePair(cards) {
                        return true
                    }
                }
            }
            
        case 3:
            //changeLaiZiValue3(slice, isOfPlanePair, cards)
            
            for i := 3; i < 15; i++ {
                cards[slice[0]].VALUE = i
                for j := 3; j < 15; j++ {
                    cards[slice[1]].VALUE = j
                    for k := 3; k < 15; k++ {
                        cards[slice[2]].VALUE = k
                        if isOfPlanePair(cards) {
                            return true
                        }
                    }
                }
            }
            
        }
    }
    return false
}

//选择癞子牌
func (mp *mPlayCards) chooseLaiZi(val int) *mPlayCards {
    for i, _ := range mp.cards {
        if mp.cards[i].VALUE == val {
            mp.cards[i].LAIZI = true
        }
    }
    return mp
}
/*
//一张癞子
func changeLaiZiValue(slice []int, f func([]*pokerCards) bool, cards []*pokerCards) bool {
    for i := 3; i < 15; i++ {
        cards[slice[0]].VALUE = i
        go f(cards)
    }
    return false
}

//2张
func changeLaiZiValue2(slice []int, f func([]*pokerCards) bool, cards []*pokerCards) bool {
    for i := 3; i < 15; i++ {
        cards[slice[0]].VALUE = i
        for j := 3; j < 15; j++ {
            cards[slice[1]].VALUE = j
            if f(cards) {
                return true
            }
        }
    }
    return false
}

//3张
func changeLaiZiValue3(slice []int, f func([]*pokerCards) bool, cards []*pokerCards) bool {
    for i := 3; i < 15; i++ {
        cards[slice[0]].VALUE = i
        for j := 3; j < 15; j++ {
            cards[slice[1]].VALUE = j
            for k := 3; k < 15; k++ {
                cards[slice[2]].VALUE = k
                if f(cards) {
                    return true
                }
            }
        }
    }
    return false
}

//4张
func changeLaiZiValue4(slice []int, f func([]*pokerCards) bool, cards []*pokerCards) bool {
    for i := 3; i < 15; i++ {
        cards[slice[0]].VALUE = i
        for j := 3; j < 15; j++ {
            cards[slice[1]].VALUE = j
            for k := 3; k < 15; k++ {
                cards[slice[2]].VALUE = k
                for l := 3; l < 15; l++ {
                    cards[slice[3]].VALUE = l
                    if f(cards) {
                        return true
                    }
                }
            }
        }
    }
    return false
}
*/
