// generated by stringer -type=Command; DO NOT EDIT

package lwgo

import "fmt"

const _Command_name = "OnOffDimIncreaseDecreaseMoodAllOnAllOff"

var _Command_index = [...]uint8{0, 2, 5, 8, 16, 24, 28, 33, 39}

func (i Command) String() string {
	if i < 0 || i+1 >= Command(len(_Command_index)) {
		return fmt.Sprintf("Command(%d)", i)
	}
	return _Command_name[_Command_index[i]:_Command_index[i+1]]
}
