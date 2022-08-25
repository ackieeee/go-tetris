package tetris

var IminoType = map[string][][]int{
	"type1": {
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	},
	"type2": {
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
	},
	"type3": {
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
	},
	"type4": {
		{0, 0, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 1},
	},
	"type5": {
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	"type6": {
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	},
	"type7": {
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	},
	"type8": {
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{1, 1, 1, 1},
	},
}

type IMino struct {
	BaseMino
}

func NewIMino(x, y int, t string) Mino {
	return &IMino{
		BaseMino{x, y, t, IminoType},
	}
}

func (m *IMino) LeftRotate(field *Field) {
	switch m.Type {
	case "type2":
		if CanMoveMinoByType(m, field, "type6", 0, 0) {
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type6", -2, 0) {
			m.X += -2
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type6", 1, 0) {
			m.X += 1
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type8", 1, 0) {
			m.X += 1
			m.Type = "type8"
			return
		}
		if CanMoveMinoByType(m, field, "type5", -2, 0) {
			m.X += -2
			m.Type = "type5"
			return
		}
	case "type3":
		if CanMoveMinoByType(m, field, "type6", 0, 0) {
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type6", 2, 0) {
			m.X += 2
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type6", -1, 0) {
			m.X += -1
			m.Type = "type6"
			return
		}
		if CanMoveMinoByType(m, field, "type5", 2, 0) {
			m.X += 2
			m.Type = "type5"
			return
		}
		if CanMoveMinoByType(m, field, "type8", -1, 0) {
			m.X += -1
			m.Type = "type8"
			return
		}
	case "type6":
		// pattern0
		if CanMoveMinoByType(m, field, "type2", 0, 0) {
			m.Type = "type2"
			return
		}
		// pattern1
		if CanMoveMinoByType(m, field, "type1", 0, 0) {
			m.Type = "type1"
			return
		}
		// pattern2
		if CanMoveMinoByType(m, field, "type4", 0, 0) {
			m.Type = "type4"
			return
		}
		// pattern3
		if CanMoveMinoByType(m, field, "type1", 0, -2) {
			m.Y = m.Y - 2
			m.Type = "type1"
			return
		}
		// pattern4
		if CanMoveMinoByType(m, field, "type4", 0, 1) {
			m.Y = m.Y + 1
			m.Type = "type4"
			return
		}
	case "type7":
		if CanMoveMinoByType(m, field, "type3", 0, 0) {
			m.Type = "type3"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 0, 0) {
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 0, 0) {
			m.Type = "type1"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 0, 2) {
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 0, -1) {
			m.Type = "type1"
			return
		}
	}
}

func (m *IMino) RightRotate(field *Field) {
}

func (m *IMino) Copy() Mino {
	return &IMino{
		BaseMino{
			m.X,
			m.Y,
			m.Type,
			m.Types,
		},
	}
}
