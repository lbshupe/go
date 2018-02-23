package main

import (
	"fmt"
	"strconv"
)

const SIDE = 10
const STATICMAX = 1000000
const STATICMIN = -1000000
const DEPTHLIMIT = 3
//const MAXLINE = 80

type Move struct {
	Row int
	Col int
	Mover string
}

type Moveresult struct {
	Row int
	Col int
	Mover string
	Score int32
}

type Board struct {
	board [SIDE][SIDE]string
	score int32
	movenum int32
}

func canmove(b *Board, muver string) bool {
	var can bool = false
	var muv Move

	muv.Mover = muver
	for row := 1; row < 9 && can == false; row++ {
		for col := 1; col < 9 && can == false; col++ {
			muv.Row = row
			muv.Col = col
			if b.board[row][col] == "-" && legal(b,&muv) {
				can = true
			}
		}
	}
	return can
}

func gameover(b *Board) bool {
	var me = true
	var you = true
	var over = false
	var availablespots = false
	var canyoumove = false
	var canmemove = false

	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if b.board[row][col] == "-" {
				availablespots = true
			}
		}
	}
	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if b.board[row][col] == "O" {
				you = false
			} 
		}
	}
	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if b.board[row][col] == "S" {
				me = false
			}
		}
	}
	if canmove(b,"O") {
		canyoumove = true
	}
	if canmove(b,"S") {
		canmemove = true
	}
	if me || you {
		over = true
	}
	if canyoumove == false || canmemove == false || availablespots == false {
		over = true
	}
	return over
}

func entermove(inputmove *Move) {
	var input string
	var first string
	var second string
	//var inputmove Move
	var legal bool = false

	inputmove.Row = 0
	inputmove.Col = 0
	//inputmove.Mover = "O"
	for legal == false {
		fmt.Print("Enter Move: ")
		fmt.Scanln(&input)
		if len(input) > 1 {
			first = input[0:1]
			second = input[1:2]
			if row, err := strconv.Atoi(first); err == nil {
				if row < 1 {
					fmt.Println("Row must be greater than 0 and less than 9")
				} else if row > 8 {
					fmt.Println("Row must be greater than 0 and less than 9")
				} else {
					legal = true
					inputmove.Row = row
				}
			}
			if col, err := strconv.Atoi(second); err == nil {
				if col < 1 {
					fmt.Println("Col must be greater than 0 and less than 9")
				} else if col > 8 {
					fmt.Println("Col must be greater than 0 and less than 9")
				} else {
					legal = true
					inputmove.Col = col
				}
			}
		}
	}
	//return inputmove
}

func printboard(b *Board) {
	var rowcol string
	for row := 1; row < 9; row++ {
		rowcol = b.board[row][1]+b.board[row][2]+b.board[row][3]+b.board[row][4]+b.board[row][5]+b.board[row][6]+b.board[row][7]+b.board[row][8]
		fmt.Println(rowcol)
	}
}

/*func makemovenewboard(b *Board, muv *Move) Board {
	b.board[muv.Row][muv.Col] = muv.Mover
}*/

func makemovecurboard(b *Board, muv *Move) {
	b.board[muv.Row][muv.Col] = muv.Mover
	j := N(b,muv)
	//fmt.Println("N =",j)
	fmt.Println("muv.Mover = ",muv.Mover)
	for i := 1; i <= j; i++ {
		b.board[muv.Row - i][muv.Col] = muv.Mover
	}
	j = NE(b,muv)
	//fmt.Println("NE =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row - i][muv.Col + i] = muv.Mover
	}
	j = E(b,muv)
	//fmt.Println("E =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row][muv.Col + i] = muv.Mover
	}
	j = SE(b,muv)
	//fmt.Println("SE =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row + i][muv.Col + i] = muv.Mover
	}
	j = S(b,muv)
	//fmt.Println("S =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row + i][muv.Col] = muv.Mover
	}
	j = SW(b,muv)
	//fmt.Println("SW =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row + i][muv.Col - i] = muv.Mover
	}
	j = W(b,muv)
	//fmt.Println("W =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row][muv.Col - i] = muv.Mover
	}
	j = NW(b,muv)
	//fmt.Println("NW =",j)
	for i := 1; i <= j; i++ {
		b.board[muv.Row - i][muv.Col - i] = muv.Mover
	}
}

func initboard(brd *Board) {
	for i := 0; i < SIDE; i++ {
		brd.board[i][0] = "#"
		brd.board[0][i] = "#"
		brd.board[9][i] = "#"
		brd.board[i][9] = "#"
	}
	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			brd.board[i][j] = "-"
		}
	}
	brd.board[4][4] = "O"
	brd.board[4][5] = "S"
	brd.board[5][4] = "S"
	brd.board[5][5] = "O"
	brd.movenum = 0
}

func copyboard(brd *Board) (newboard Board) {
	newboard = Board{};

	for i := 0; i < SIDE; i++ {
		for j := 0; j < SIDE; j++ {
			newboard.board[i][j] = brd.board[i][j]
		}
	}
	return
}

func N(b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row - i][muv.Col] != "-") &&
			(b.board[muv.Row - i][muv.Col] != muv.Mover) &&
			(b.board[muv.Row - i][muv.Col] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row - value - 1][muv.Col] != muv.Mover) {
		value = 0
	}
	return value
} 

func NE (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row - i][muv.Col + i] != "-") &&
			(b.board[muv.Row - i][muv.Col + i] != muv.Mover) &&
			(b.board[muv.Row - i][muv.Col + i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row - value - 1][muv.Col + value + 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func E (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row][muv.Col + i] != "-") &&
			(b.board[muv.Row][muv.Col + i] != muv.Mover) &&
			(b.board[muv.Row][muv.Col + i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row][muv.Col + value + 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func SE (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row + i][muv.Col + i] != "-") &&
			(b.board[muv.Row + i][muv.Col + i] != muv.Mover) &&
			(b.board[muv.Row + i][muv.Col + i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row + value + 1][muv.Col + value + 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func S (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row + i][muv.Col] != "-") &&
			(b.board[muv.Row + i][muv.Col] != muv.Mover) &&
			(b.board[muv.Row + i][muv.Col] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row + value + 1][muv.Col] != muv.Mover) {
		value = 0
	}
	return value
} 

func SW (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row + i][muv.Col - i] != "-") &&
			(b.board[muv.Row + i][muv.Col - i] != muv.Mover) &&
			(b.board[muv.Row + i][muv.Col - i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row + value + 1][muv.Col - value - 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func W (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row][muv.Col - i] != "-") &&
			(b.board[muv.Row][muv.Col - i] != muv.Mover) &&
			(b.board[muv.Row][muv.Col - i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row][muv.Col - value - 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func NW (b *Board, muv *Move) int {
	value := 0
	loop := true

	for i := 1; i < 10 && loop == true; i++ {
		if (b.board[muv.Row - i][muv.Col - i] != "-") &&
			(b.board[muv.Row - i][muv.Col - i] != muv.Mover) &&
			(b.board[muv.Row - i][muv.Col - i] != "#") {
			value = i
		} else {
			loop = false
		}
	}
	if (b.board[muv.Row - value - 1][muv.Col - value - 1] != muv.Mover) {
		value = 0
	}
	return value
} 

func legal (b *Board, muv *Move) bool {
	if (b.board[muv.Row][muv.Col] != "-") {
		return false
	}
	if (N(b,muv) > 0) {
		return true
	}
	if (NE(b,muv) > 0) {
		return true
	}
	if (E(b,muv) > 0) {
		return true
	}
	if (SE(b,muv) > 0) {
		return true
	}
	if (S(b,muv) > 0) {
		return true
	}
	if (SW(b,muv) > 0) {
		return true
	}
	if (W(b,muv) > 0) {
		return true
	}
	if (NW(b,muv) > 0) {
		return true
	} else {
		return false
	}
}

func nextmover (b *Board, muv *Move) string {
	var mover string

	if muv.Mover == "S" {
		mover = "O"
	} else {
		mover = "S"
	}
	/*if canmove(b,mover) {  
		return mover
	} else {
		return "S"
	}*/
	return mover
}

func score (b *Board) int {
	var sum = 0
	var self = true
	var other = true

	if b.movenum < 8 {
		for i := 3; i < 7; i++ {
			if b.board[i][3] == "S" {
				sum = sum + 10
			}
			if b.board[3][i] == "S" {
				sum = sum + 10
			}
			if b.board[6][i] == "S" {
				sum = sum + 10
			}
			if b.board[i][6] == "S" {
				sum = sum + 10
			}
		}
	}
	for row := 1; row < 9; row++ {
		for col := 1; col < 9; col++ {
			if b.board[row][col] == "S" {
				sum++
				other = false
			} else if b.board[row][col] == "O" {
                                sum--
                                self = false
			}
		}
	}
	if self {
		sum = STATICMAX
	} else if other {
		sum = STATICMIN
	} else {
		for ro := 1; ro < 9; {
			if b.board[2][2] == "S" {
				sum = sum - 1000
			}
			if b.board[2][7] == "S" {
				sum = sum - 1000
			}
			if b.board[7][2] == "S" {
				sum = sum - 1000
			}
			if b.board[7][7] == "S" {
				sum = sum - 1000
			}
			for co := 3; co <= 6 && b.movenum < 50; co++ {
				if b.board[ro][co] == "S" {
					sum = sum + 10
					if b.board[ro][co-1] == "O" && b.board[ro][co+1] == "O" {
						sum = sum + 100
					}
				} else if b.board[ro][co] == "O" {
					sum = sum - 10
					if b.board[ro][co-1] == "S" && b.board[ro][co+1] == "S" {
						sum = sum - 100
					}
				}
				if b.board[ro][co] == "S" {
					sum = sum + 10
					if b.board[ro-1][co] == "O" && b.board[ro+1][co] == "O" {
						sum = sum + 100
					}
				} else if b.board[ro][co] == "O" {
					sum = sum - 10
					if b.board[ro-1][co] == "S" && b.board[ro+1][co] == "S" {
						sum = sum - 100
					}
				}
			}
			ro = ro + 7
		}
		
	}
	for r := 1; r <= 8 && b.movenum < 50; {
		for c := 1; c <=8; {
			if b.board[r][c] == "O" {
				sum = sum - 100000
			} else if b.board[r][c] == "S" {
				sum = sum + 100000
			}
			c = c + 7
		}
		r = r + 7
	}
	return sum
}

func selectmove (b *Board, muv *Move) {
	//var breadth = 5
	//var depth = 5
	//var bdepth = 0
	//var balpha = STATICMIN
	//var bbeta = STATICMAX
	//var tempmove Move
	//var tempscore = 0
	var can = false
	for row := 1; row < 9 && can == false; row++ {
		for col := 1; col < 9 && can == false; col++ {
			muv.Row = row
			muv.Col = col
			if b.board[row][col] == "-" && legal(b,muv) {
				can = true
			}
		}
	}
	//tempmove,tempscore = minimax(muv.Mover, 2, b)

}

//func deepenough (b *Board, d int) bool {
//	var flag bool = true
//
//	for x := 1; x <=8; x++ {
//		for y := 1; y <= 8; y++ {
//			if b.Board[x][y] == "-" {
//				flag = false
//			}
//		}
//	}
//	if d == DEPTHLIMIT {
//		flag = true
//	}
//	return flag
//}

//func minimax (b Board, int depth, int alpha, in beta, Moveresult *reslt) {

//	if deepenough(&b, depth) {
		//reslt.Mover = "*"
		//reslt.Row = 1
		//reslt.Col = 1
		//reslt.Score = score(&b)
	//} else {
	//}
//}



func minimax(player string, depth int, brd Board) (m Move, boardscore int) {

	m = Move{}

	boardscore = score(&brd)

	return
}

func main() {
	var gameboard Board
	var usermove Move

	initboard(&gameboard)
	printboard(&gameboard)
	usermove.Mover = "O"
	entermove(&usermove)
	for (! gameover(&gameboard)) {
		if canmove(&gameboard,usermove.Mover) {
			printboard(&gameboard)
		}
		if legal(&gameboard, &usermove) {
			//fmt.Println("Move is Legal!")
			makemovecurboard(&gameboard, &usermove)
			gameboard.movenum++
			usermove.Mover = nextmover(&gameboard,&usermove)
		} else {
			fmt.Println("Move is Not Legal!")
		}
		printboard(&gameboard)
		//fmt.Println("Score is")
		//fmt.Println(score(&gameboard))
		//fmt.Println("move number is")
		//fmt.Println(gameboard.movenum)
		if usermove.Mover == "O" {
			entermove(&usermove)
		} else {
			selectmove(&gameboard,&usermove)
		}
	}
	printboard(&gameboard)
}
