package rdx

//import "fmt";

var __RDX_actions = []int8{0, 1, 0, 1, 1, 1, 2, 1, 3, 1, 4, 1, 5, 1, 6, 1, 7, 1, 8, 2, 1, 3, 2, 1, 5, 2, 1, 7, 2, 1, 8, 2, 2, 3, 2, 2, 5, 2, 2, 7, 2, 2, 8, 2, 4, 0, 2, 4, 3, 2, 4, 5, 2, 4, 7, 2, 4, 8, 2, 6, 0, 2, 6, 3, 2, 6, 5, 2, 6, 7, 2, 6, 8, 0}
var __RDX_key_offsets = []int16{0, 0, 23, 30, 32, 34, 38, 42, 44, 50, 57, 63, 70, 76, 85, 91, 97, 103, 109, 132, 155, 164, 187, 201, 212, 235, 254, 269, 292, 313, 329, 345, 362, 0}
var __RDX_trans_keys = []byte{32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 34, 46, 57, 92, 120, 48, 49, 48, 57, 48, 57, 69, 101, 48, 57, 43, 45, 48, 57, 48, 57, 48, 57, 65, 70, 97, 102, 45, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 95, 48, 57, 65, 90, 97, 122, 48, 57, 65, 70, 97, 102, 34, 47, 92, 98, 102, 110, 114, 116, 117, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 58, 91, 93, 123, 125, 9, 13, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 46, 58, 69, 91, 93, 101, 123, 125, 9, 13, 48, 57, 32, 44, 58, 91, 93, 123, 125, 9, 13, 48, 57, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 45, 46, 58, 69, 91, 93, 101, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 45, 58, 91, 93, 95, 123, 125, 9, 13, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 58, 91, 93, 95, 123, 125, 9, 13, 48, 57, 65, 90, 97, 122, 32, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 43, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 0}
var __RDX_single_lengths = []int8{0, 9, 5, 0, 0, 2, 2, 0, 0, 1, 0, 1, 0, 9, 0, 0, 0, 0, 9, 9, 7, 9, 10, 7, 9, 11, 7, 9, 9, 8, 8, 9, 8, 0}
var __RDX_range_lengths = []int8{0, 7, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 0, 3, 3, 3, 3, 7, 7, 1, 7, 2, 2, 7, 4, 4, 7, 6, 4, 4, 4, 4, 0}
var __RDX_index_offsets = []int16{0, 0, 17, 24, 26, 28, 32, 36, 38, 42, 47, 51, 56, 60, 70, 74, 78, 82, 86, 103, 120, 129, 146, 159, 169, 186, 202, 214, 231, 247, 260, 273, 287, 0}
var __RDX_cond_targs = []int8{1, 2, 1, 1, 1, 18, 11, 1, 19, 1, 3, 25, 28, 11, 28, 11, 0, 20, 0, 0, 13, 0, 0, 2, 22, 0, 5, 0, 6, 6, 5, 0, 7, 7, 23, 0, 23, 0, 9, 9, 9, 0, 10, 9, 9, 9, 0, 26, 26, 26, 0, 29, 29, 29, 29, 0, 32, 9, 9, 0, 2, 2, 2, 2, 2, 2, 2, 2, 14, 0, 15, 15, 15, 0, 16, 16, 16, 0, 17, 17, 17, 0, 2, 2, 2, 0, 1, 2, 1, 1, 1, 18, 11, 1, 19, 1, 3, 25, 28, 11, 28, 11, 0, 1, 2, 1, 1, 1, 18, 11, 1, 19, 1, 3, 25, 28, 11, 28, 11, 0, 21, 21, 21, 21, 24, 21, 27, 21, 0, 21, 2, 21, 21, 21, 24, 11, 21, 27, 21, 3, 25, 28, 11, 28, 11, 0, 21, 21, 4, 21, 6, 21, 24, 6, 21, 27, 21, 22, 0, 21, 21, 21, 21, 24, 21, 27, 21, 23, 0, 21, 2, 21, 21, 21, 24, 11, 21, 27, 21, 3, 25, 28, 11, 28, 11, 0, 21, 21, 8, 4, 21, 31, 21, 24, 31, 21, 27, 21, 25, 30, 30, 0, 21, 21, 21, 21, 24, 21, 27, 21, 26, 26, 26, 0, 21, 2, 21, 21, 21, 24, 11, 21, 27, 21, 3, 25, 28, 11, 28, 11, 0, 21, 21, 8, 21, 21, 24, 29, 21, 27, 21, 28, 28, 29, 28, 29, 0, 21, 21, 21, 21, 24, 29, 21, 27, 21, 29, 29, 29, 0, 21, 21, 8, 21, 21, 24, 21, 27, 21, 30, 30, 30, 0, 21, 7, 21, 12, 21, 21, 24, 21, 27, 21, 30, 30, 30, 0, 21, 21, 10, 21, 21, 24, 21, 27, 21, 32, 9, 9, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 0}
var __RDX_cond_actions = []int8{0, 1, 15, 17, 11, 0, 0, 7, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 13, 58, 67, 70, 64, 13, 13, 61, 13, 13, 58, 58, 13, 13, 13, 13, 0, 9, 43, 52, 55, 49, 9, 9, 46, 9, 9, 43, 43, 9, 9, 9, 9, 0, 5, 37, 40, 34, 5, 31, 5, 5, 0, 0, 1, 15, 17, 11, 0, 0, 7, 0, 0, 1, 1, 0, 0, 0, 0, 0, 3, 25, 0, 28, 0, 22, 3, 0, 19, 3, 3, 0, 0, 0, 15, 17, 11, 0, 7, 0, 0, 0, 0, 13, 58, 67, 70, 64, 13, 13, 61, 13, 13, 58, 58, 13, 13, 13, 13, 0, 3, 25, 0, 0, 28, 0, 22, 3, 0, 19, 3, 3, 0, 0, 0, 0, 0, 15, 17, 11, 0, 7, 0, 0, 0, 0, 0, 0, 9, 43, 52, 55, 49, 9, 9, 46, 9, 9, 43, 43, 9, 9, 9, 9, 0, 0, 15, 0, 17, 11, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 15, 17, 11, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 15, 0, 17, 11, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0, 17, 11, 0, 7, 0, 0, 0, 0, 0, 0, 0, 15, 0, 17, 11, 0, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 3, 0, 13, 3, 0, 9, 0, 0, 0, 0, 0, 0}
var __RDX_eof_trans = []int16{301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319, 320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 0}
var _RDX_start int = 1
var _ = _RDX_start
var _RDX_first_final int = 20
var _ = _RDX_first_final
var _RDX_error int = 0
var _ = _RDX_error
var _RDX_en_main int = 1
var _ = _RDX_en_main

func ParseRDX(data []byte) (rdx *RDX, err error) {

	var mark [RdxMaxNesting]int
	nest, cs, p, pe, eof := 0, 0, 0, len(data), len(data)

	rdx = &RDX{}

	{
		cs = int(_RDX_start)

	}
	{
		var _klen int
		var _trans uint = 0
		var _keys int
		var _acts int
		var _nacts uint
	_resume:
		{

		}
		if p == pe && p != eof {
			goto _out

		}
		if p == eof {
			if __RDX_eof_trans[cs] > 0 {
				_trans = uint(__RDX_eof_trans[cs]) - 1

			}

		} else {
			_keys = int(__RDX_key_offsets[cs])

			_trans = uint(__RDX_index_offsets[cs])
			_klen = int(__RDX_single_lengths[cs])
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + _klen - 1
				var _mid int
				for {
					if _upper < _lower {
						_keys += _klen
						_trans += uint(_klen)
						break

					}
					_mid = _lower + ((_upper - _lower) >> 1)
					if (data[p]) < __RDX_trans_keys[_mid] {
						_upper = _mid - 1

					} else if (data[p]) > __RDX_trans_keys[_mid] {
						_lower = _mid + 1

					} else {
						_trans += uint((_mid - _keys))
						goto _match

					}

				}

			}
			_klen = int(__RDX_range_lengths[cs])
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + (_klen << 1) - 2
				var _mid int
				for {
					if _upper < _lower {
						_trans += uint(_klen)
						break

					}
					_mid = _lower + (((_upper - _lower) >> 1) & ^1)
					if (data[p]) < __RDX_trans_keys[_mid] {
						_upper = _mid - 2

					} else if (data[p]) > __RDX_trans_keys[_mid+1] {
						_lower = _mid + 2

					} else {
						_trans += uint(((_mid - _keys) >> 1))
						break

					}

				}

			}
		_match:
			{

			}

		}
		cs = int(__RDX_cond_targs[_trans])
		if __RDX_cond_actions[_trans] != 0 {
			_acts = int(__RDX_cond_actions[_trans])

			_nacts = uint(__RDX_actions[_acts])
			_acts += 1
			for _nacts > 0 {
				switch __RDX_actions[_acts] {
				case 0:
					{
						mark[nest] = p
					}

				case 1:
					{
						rdx.RdxType = RdxInt
						rdx.Text = data[mark[nest]:p]
					}

				case 2:
					{
						rdx.RdxType = RdxString
						rdx.Text = data[mark[nest]:p]
					}

				case 3:
					{
						n := rdx.Nested
						n = append(n, RDX{Parent: rdx})
						rdx.Nested = n
						rdx.RdxType = RdxMap
						rdx = &n[len(n)-1]
						nest++
					}

				case 4:
					{
						nest--
						rdx = rdx.Parent
					}

				case 5:
					{
						n := rdx.Nested
						n = append(n, RDX{Parent: rdx})
						rdx.Nested = n
						rdx.RdxType = RdxArray
						rdx = &n[len(n)-1]
						nest++
					}

				case 6:
					{
						nest--
						rdx = rdx.Parent
					}

				case 7:
					{
						if rdx.Parent == nil {
							{
								p += 1
								goto _out

							}

						}
						n := rdx.Parent.Nested
						n = append(n, RDX{Parent: rdx.Parent})
						rdx.Parent.Nested = n
						rdx = &n[len(n)-1]
					}

				case 8:
					{
						n := rdx.Parent.Nested
						n = append(n, RDX{Parent: rdx.Parent})
						rdx.Parent.Nested = n
						rdx = &n[len(n)-1]
					}

				}
				_nacts -= 1
				_acts += 1

			}

		}
		if p == eof {
			if cs >= 20 {
				goto _out

			}

		} else {
			if cs != 0 {
				p += 1
				goto _resume

			}

		}
	_out:
		{

		}

	}
	if cs < _RDX_first_final {
		err = ErrBadRdx
	}

	return
}