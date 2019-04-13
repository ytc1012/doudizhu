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
    playerNum = 4
    cardsNum = 2
)

type pokerCards struct {
    COLOR string
    SYMBOL string
    VALUE int
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
    fmt.Println("地主牌：\n", rlc)
}

//判断牌型前先排序 降序
func bubbleSort(cards []*pokerCards) []*pokerCards {
    length := len(cards)
    for i := 0; i < length; i++ {
        for j := 0; j < length - i -1; j++ {
            if cards[j].VALUE < cards[j+1].VALUE {
                cards[j], cards[j+1] = cards[j+1], cards[j]
            }
        }
    }
    return cards
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
    
    //重新赋值测试
    h4.playerCards[:][0].COLOR = "Spade"
    h4.playerCards[:][0].SYMBOL = "3"
    h4.playerCards[:][0].VALUE = 3
    h4.playerCards[:][1].COLOR = "Spade"
    h4.playerCards[:][1].SYMBOL = "3"
    h4.playerCards[:][1].VALUE = 4
    h4.playerCards[:][2].COLOR = "Spade"
    h4.playerCards[:][2].SYMBOL = "3"
    h4.playerCards[:][2].VALUE = 5
    h4.playerCards[:][3].COLOR = "Spade"
    h4.playerCards[:][3].SYMBOL = "3"
    h4.playerCards[:][3].VALUE = 6
    h4.playerCards[:][4].COLOR = "Spade"
    h4.playerCards[:][4].SYMBOL = "3"
    h4.playerCards[:][4].VALUE = 7
    h4.playerCards[:][5].COLOR = "Spade"
    h4.playerCards[:][5].SYMBOL = "3"
    h4.playerCards[:][5].VALUE = 8
    h4.playerCards[:][6].COLOR = "Spade"
    h4.playerCards[:][6].SYMBOL = "3"
    h4.playerCards[:][6].VALUE = 9
    h4.playerCards[:][7].COLOR = "Spade"
    h4.playerCards[:][7].SYMBOL = "3"
    h4.playerCards[:][7].VALUE = 10
    h4.playerCards[:][8].COLOR = "Spade"
    h4.playerCards[:][8].SYMBOL = "3"
    h4.playerCards[:][8].VALUE = 11
    h4.playerCards[:][9].COLOR = "Spade"
    h4.playerCards[:][9].SYMBOL = "3"
    h4.playerCards[:][9].VALUE = 12
    rpctest := show.readPlayerCardsMore2(h4.playerCards)
    fmt.Println("测试牌：\n", rpctest)
    checkCardsType(h4.playerCards[:1])
    checkCardsType(h4.playerCards[:2])
    checkCardsType(h4.playerCards[:3])
    checkCardsType(h4.playerCards[:4])
    checkCardsType(h4.playerCards[:5])
    checkCardsType(h4.playerCards[:6])
    checkCardsType(h4.playerCards[:7])
    checkCardsType(h4.playerCards[:8])
    checkCardsType(h4.playerCards[:9])
    checkCardsType(h4.playerCards[:10])
    
    var h5 = new(handCards)
    h5.bottomCards = h5.dealCardsMore(4, playerNum, ret)
    rbc := show.readBottomCardsMore(h5.bottomCards)
    fmt.Println("底牌：", rbc, "\n")

    //重新赋值测试（6张）
    h5.bottomCards[:][0].COLOR = "Spade"
    h5.bottomCards[:][0].SYMBOL = "8"
    h5.bottomCards[:][0].VALUE = 7
    h5.bottomCards[:][1].COLOR = "Spade"
    h5.bottomCards[:][1].SYMBOL = "7"
    h5.bottomCards[:][1].VALUE = 5
    h5.bottomCards[:][2].COLOR = "Spade"
    h5.bottomCards[:][2].SYMBOL = "6"
    h5.bottomCards[:][2].VALUE = 6
    h5.bottomCards[:][3].COLOR = "Spade"
    h5.bottomCards[:][3].SYMBOL = "5"
    h5.bottomCards[:][3].VALUE = 8
    h5.bottomCards[:][4].COLOR = "Spade"
    h5.bottomCards[:][4].SYMBOL = "4"
    h5.bottomCards[:][4].VALUE = 9
    h5.bottomCards[:][5].COLOR = "Spade"
    h5.bottomCards[:][5].SYMBOL = "3"
    h5.bottomCards[:][5].VALUE = 4
    rbctest := show.readBottomCardsMore(h5.bottomCards)
    fmt.Println("测试牌：\n", rbctest)
    checkCardsType(h5.bottomCards[:1])
    checkCardsType(h5.bottomCards[:2])
    checkCardsType(h5.bottomCards[:3])
    checkCardsType(h5.bottomCards[:4])
    checkCardsType(h5.bottomCards[:5])
    checkCardsType(h5.bottomCards[:])
    
    
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

//是否有大王或小王
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

//判断牌型
func checkCardsType(cards []*pokerCards) {
    //bubbleSort(cards)
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
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
            if cards[0].VALUE == cards[1].VALUE && cards[1].VALUE == cards[2].VALUE {
                fmt.Print("三不带\n")
            } else {
                fmt.Println("不符合规则\n")
            }
        }
    case 4:
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
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
        }
    case 5:
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
            // 顺子
            if isStraight(cards) {
                fmt.Println("顺子\n")
            } else if isThreeTwo(cards) {
                fmt.Println("三带对\n")
            } else {
                fmt.Println("不符合规则\n")
            }
        }
    case 6:
        a := cards[0].VALUE == cards[1].VALUE
        b := cards[1].VALUE == cards[2].VALUE
        c := cards[3].VALUE == cards[4].VALUE
        d := cards[4].VALUE == cards[5].VALUE
        e := cards[0].VALUE != cards[3].VALUE
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
            if a == true && a && b && c && d && e {
                fmt.Println("飞机不带\n")
            } else if isStraight(cards) {
                fmt.Println("顺子\n")
            } else if isFourTwo(cards) {
                fmt.Println("四带二单\n")
            } else {
                fmt.Println("不符合规则\n")
            }
        }
    case 8:
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
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
        }
    default:
        if isIncludeJoker(cards) {
            fmt.Println("不符合规则\n")
        } else {
            if isStraight(cards) {
                fmt.Println("顺子\n")
            } else if isEvenPair(cards) {
                fmt.Println("连对\n")
            } else if isPlanePair(cards) {
                fmt.Println("飞机带二对\n")
            } else {
                fmt.Println("不符合规则")
            }
        }
    }
}

//是否连对
func isEvenPair(cards []*pokerCards) bool {
    bubbleSort(cards)
    l := len(cards)
    if l > 3 && l < 25 && (l % 2 == 0) {
        for i := 0; i < l; {
            if cards[i].VALUE == cards[i+1].VALUE {
                i += 2
            } else {
                return false
            }
        }
        
        for i := 0; i < l; {
            if cards[i].VALUE - cards[i+2].VALUE == 1 {
                i += 4
            } else {
                return false
            }
        }
    }
    return true
}

//是否顺子
func isStraight(cards []*pokerCards) bool {
    bubbleSort(cards)
    for i := 0; i < len(cards)-1; i++ {
        if (cards[i].VALUE - cards[i+1].VALUE) == 1 {
            continue
        } else {
            return false
        }
    }
    return true
}

//是否三带一
func isThreeOne (cards []*pokerCards) bool {
    a := cards[0].VALUE == cards[1].VALUE
    b := cards[1].VALUE == cards[2].VALUE
    c := cards[2].VALUE == cards[3].VALUE
    if a && b && !c {
        return true
    } else if !a && b && c {
        return true
    }
    return false
}

//是否三带对
func isThreeTwo(cards []*pokerCards) bool {
    if cards[0].VALUE == cards[1].VALUE && cards[1].VALUE == cards[2].VALUE {
        if cards[3].VALUE == cards[4].VALUE && cards[0].VALUE != cards[3].VALUE {
            return true
        }
    } else if cards[0].VALUE == cards[1].VALUE {
        if cards[2].VALUE == cards[3].VALUE && cards[3].VALUE == cards[4].VALUE && cards[0].VALUE != cards[2].VALUE {
            return true
        }
    }
    return false
}

//是否四带二单
func isFourTwo(cards []*pokerCards) bool {
    a := cards[0].VALUE == cards[1].VALUE
    b := cards[1].VALUE == cards[2].VALUE
    c := cards[2].VALUE == cards[3].VALUE
    d := cards[3].VALUE == cards[4].VALUE
    e := cards[4].VALUE == cards[5].VALUE
    if a { 
        if a && b && c && d == false && e == false {
            return true
        }
    } else if b {
        if b && c && d && a == false && e == false {
            return true
        }
    } else if c {
        if c && d && e && a == false && b == false {
            return true
        }
    }
    return false
}

//是否飞机带二单
func isPlaneTwo(cards []*pokerCards) bool {
    a := cards[0].VALUE == cards[1].VALUE
    b := cards[1].VALUE == cards[2].VALUE
    c := cards[2].VALUE == cards[3].VALUE
    d := cards[3].VALUE == cards[4].VALUE
    e := cards[4].VALUE == cards[5].VALUE
    f := cards[5].VALUE == cards[6].VALUE
    g := cards[6].VALUE == cards[7].VALUE
    if a {
        if a && b && c == false && d && e && f == false && g == false {
            return true
        }
    } else if b {
        if a == false && b && c && d == false && e && f && g == false {
            return true
        }
    } else if c {
        if a == false && b == false && c && d && e == false && f && g {
            return true
        }
    }
    return false
}

//是否四带二对
func isFourPair(cards []*pokerCards) bool {
    a := cards[0].VALUE == cards[1].VALUE
    b := cards[1].VALUE == cards[2].VALUE
    c := cards[2].VALUE == cards[3].VALUE
    d := cards[3].VALUE == cards[4].VALUE
    e := cards[4].VALUE == cards[5].VALUE
    f := cards[5].VALUE == cards[6].VALUE
    g := cards[6].VALUE == cards[7].VALUE
    if a {
        if a && b && c && d == false && e && f == false && g {
            return true
        }
    } else if c {
        if a && b == false && c && d && e && f == false && g {
            return true
        }
    } else if e {
        if a && b == false && c && d == false && e && f && g {
            return true
        }
    }
    return false
}

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

//是否飞机带二对
func isPlanePair(cards []*pokerCards) bool {
    if isEqual(cards[:3]) && !isEqual(cards[2:4]) && isEqual(cards[3:6]) && !isEqual(cards[5:7]) && isEqual(cards[6:8]) && !isEqual(cards[7:9]) && isEqual(cards[8:]) {
        return true
    } else if isEqual(cards[:2]) && !isEqual(cards[1:3]) && isEqual(cards[2:5]) && !isEqual(cards[4:6]) && isEqual(cards[5:8]) && !isEqual(cards[7:9]) && isEqual(cards[8:]) {
        return true
    } else if isEqual(cards[:2]) && !isEqual(cards[1:3]) && isEqual(cards[2:4]) && !isEqual(cards[3:5]) && isEqual(cards[4:7]) && !isEqual(cards[6:8]) && isEqual(cards[7:]) {
        return true
    }
    return false
}
