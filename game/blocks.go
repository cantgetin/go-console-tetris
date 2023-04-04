package game

type Block int

const (
	IBlockType Block = 0
	JBlockType       = 1
	LBlockType       = 2
	OBlockType       = 3
	SBlockType       = 4
	TBlockType       = 5
	ZBlockType       = 6
)

var IBlock = [4][4]int{
	{0, 0, 0, 0},
	{1, 1, 1, 1},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var JBlock = [4][4]int{
	{1, 0, 0, 0},
	{1, 1, 1, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var LBlock = [4][4]int{
	{0, 0, 0, 1},
	{1, 1, 1, 1},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var OBlock = [4][4]int{
	{1, 1, 0, 0},
	{1, 1, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var SBlock = [4][4]int{
	{0, 1, 1, 0},
	{1, 1, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var TBlock = [4][4]int{
	{0, 1, 0, 0},
	{1, 1, 1, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var ZBlock = [4][4]int{
	{1, 1, 0, 0},
	{0, 1, 1, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
}

var Blocks = [7][4][4]int{
	IBlockType: IBlock,
	JBlockType: JBlock,
	LBlockType: LBlock,
	OBlockType: OBlock,
	SBlockType: SBlock,
	TBlockType: TBlock,
	ZBlockType: ZBlock,
}
