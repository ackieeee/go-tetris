package tetris

var tminoTypes = map[string][][]int{
	"type1": {
		{0, 2, 0},
		{2, 2, 2},
		{0, 0, 0},
	},
	"type2": {
		{0, 2, 0},
		{0, 2, 2},
		{0, 2, 0},
	},
	"type3": {
		{0, 0, 0},
		{2, 2, 2},
		{0, 2, 0},
	},
	"type4": {
		{0, 2, 0},
		{2, 2, 0},
		{0, 2, 0},
	},
}

type TMino struct {
	BaseMino
}

func NewTMino(x, y int, t string) Mino {
	return &TMino{
		BaseMino{x, y, t, tminoTypes},
	}
}

func (m *TMino) LeftRotate(field *Field) {
	switch m.Type {
	case "type1":
		if CanMoveMinoByType(m, field, "type4", 0, 0) {
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 1, 0) {
			m.X += 1
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 1, -1) {
			m.X += 1
			m.Y += -1
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 0, 2) {
			m.Y += 2
			m.Type = "type4"
			return
		}
		if CanMoveMinoByType(m, field, "type4", 1, 2) {
			m.X += 1
			m.Y += 2
			m.Type = "type4"
			return
		}
	case "type2":
		if CanMoveMinoByType(m, field, "type1", 0, 0) {
			m.Type = "type1"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 1, 0) {
			m.X += 1
			m.Type = "type1"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 1, 1) {
			m.X += 1
			m.Y += 1
			m.Type = "type1"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 0, -2) {
			m.Y += -2
			m.Type = "type1"
			return
		}
		if CanMoveMinoByType(m, field, "type1", 1, -2) {
			m.X += 1
			m.Y += -2
			m.Type = "type1"
			return
		}
	case "type3":
		if CanMoveMinoByType(m, field, "type2", 0, 0) {
			m.Type = "type2"
			return
		}
		if CanMoveMinoByType(m, field, "type2", -1, 0) {
			m.X += -1
			m.Type = "type2"
			return
		}
		if CanMoveMinoByType(m, field, "type2", -1, -1) {
			m.X += -1
			m.Y += -1
			m.Type = "type2"
			return
		}
		if CanMoveMinoByType(m, field, "type2", 0, 2) {
			m.Y += 2
			m.Type = "type2"
			return
		}
		if CanMoveMinoByType(m, field, "type2", -1, 2) {
			m.X += -1
			m.Y += 2
			m.Type = "type2"
			return
		}
	case "type4":
		if CanMoveMinoByType(m, field, "type3", 0, 0) {
			m.Type = "type3"
			return
		}
		if CanMoveMinoByType(m, field, "type3", -1, 0) {
			m.X += -1
			m.Type = "type3"
			return
		}
		if CanMoveMinoByType(m, field, "type3", -1, 1) {
			m.X += -1
			m.Y += 1
			m.Type = "type3"
			return
		}
		if CanMoveMinoByType(m, field, "type3", 0, -2) {
			m.Y += -2
			m.Type = "type3"
			return
		}
		if CanMoveMinoByType(m, field, "type3", -1, -2) {
			m.X += -1
			m.Y += -2
			m.Type = "type3"
			return
		}
	}
}

func (m *TMino) RightRotate(field *Field) {
}
func (m *TMino) Copy() Mino {
	return &TMino{
		BaseMino{m.X, m.Y, m.Type, m.Types},
	}
}
