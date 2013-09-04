
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(1),		/* BoolExpr */
			shift(2),		/* BoolExpr1 */
			shift(3),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* true */
			shift(6),		/* false */
			shift(7),		/* CompareExpr */
			shift(8),		/* SubStringExpr */
			shift(9),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(10),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			shift(11),		/* & */
			shift(12),		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: BoolExpr */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: BoolExpr1 */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(2),		/* &, reduce: BoolExpr1 */
			reduce(2),		/* |, reduce: BoolExpr1 */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(13),		/* BoolExpr */
			shift(14),		/* BoolExpr1 */
			shift(15),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* true */
			shift(18),		/* false */
			shift(19),		/* CompareExpr */
			shift(20),		/* SubStringExpr */
			shift(21),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(22),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Val */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(6),		/* &, reduce: Val */
			reduce(6),		/* |, reduce: Val */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Val */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(7),		/* &, reduce: Val */
			reduce(7),		/* |, reduce: Val */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Val */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(8),		/* &, reduce: Val */
			reduce(8),		/* |, reduce: Val */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Val */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(9),		/* &, reduce: Val */
			reduce(9),		/* |, reduce: Val */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			shift(23),		/* < */
			shift(24),		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			shift(25),		/* in */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(26),		/* BoolExpr */
			shift(27),		/* BoolExpr1 */
			shift(3),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* true */
			shift(6),		/* false */
			shift(7),		/* CompareExpr */
			shift(8),		/* SubStringExpr */
			shift(9),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(10),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(26),		/* BoolExpr */
			shift(28),		/* BoolExpr1 */
			shift(3),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(4),		/* ( */
			nil,		/* ) */
			shift(5),		/* true */
			shift(6),		/* false */
			shift(7),		/* CompareExpr */
			shift(8),		/* SubStringExpr */
			shift(9),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(10),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			shift(29),		/* & */
			shift(30),		/* | */
			nil,		/* ( */
			shift(31),		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			reduce(1),		/* ), reduce: BoolExpr */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(2),		/* &, reduce: BoolExpr1 */
			reduce(2),		/* |, reduce: BoolExpr1 */
			nil,		/* ( */
			reduce(2),		/* ), reduce: BoolExpr1 */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(32),		/* BoolExpr */
			shift(14),		/* BoolExpr1 */
			shift(15),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* true */
			shift(18),		/* false */
			shift(19),		/* CompareExpr */
			shift(20),		/* SubStringExpr */
			shift(21),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(22),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(6),		/* &, reduce: Val */
			reduce(6),		/* |, reduce: Val */
			nil,		/* ( */
			reduce(6),		/* ), reduce: Val */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(7),		/* &, reduce: Val */
			reduce(7),		/* |, reduce: Val */
			nil,		/* ( */
			reduce(7),		/* ), reduce: Val */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(8),		/* &, reduce: Val */
			reduce(8),		/* |, reduce: Val */
			nil,		/* ( */
			reduce(8),		/* ), reduce: Val */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(9),		/* &, reduce: Val */
			reduce(9),		/* |, reduce: Val */
			nil,		/* ( */
			reduce(9),		/* ), reduce: Val */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			shift(33),		/* < */
			shift(34),		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			shift(35),		/* in */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			shift(36),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			shift(37),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(38),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			shift(11),		/* & */
			shift(12),		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: BoolExpr1 */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: BoolExpr1 */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(39),		/* BoolExpr */
			shift(40),		/* BoolExpr1 */
			shift(15),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* true */
			shift(18),		/* false */
			shift(19),		/* CompareExpr */
			shift(20),		/* SubStringExpr */
			shift(21),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(22),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			shift(39),		/* BoolExpr */
			shift(41),		/* BoolExpr1 */
			shift(15),		/* Val */
			nil,		/* & */
			nil,		/* | */
			shift(16),		/* ( */
			nil,		/* ) */
			shift(17),		/* true */
			shift(18),		/* false */
			shift(19),		/* CompareExpr */
			shift(20),		/* SubStringExpr */
			shift(21),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(22),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: BoolExpr1 */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(5),		/* &, reduce: BoolExpr1 */
			reduce(5),		/* |, reduce: BoolExpr1 */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			shift(29),		/* & */
			shift(30),		/* | */
			nil,		/* ( */
			shift(42),		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			shift(43),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			shift(44),		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			nil,		/* & */
			nil,		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			shift(45),		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: CompareExpr */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(10),		/* &, reduce: CompareExpr */
			reduce(10),		/* |, reduce: CompareExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: CompareExpr */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(11),		/* &, reduce: CompareExpr */
			reduce(11),		/* |, reduce: CompareExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: SubStringExpr */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(12),		/* &, reduce: SubStringExpr */
			reduce(12),		/* |, reduce: SubStringExpr */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			shift(29),		/* & */
			shift(30),		/* | */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			reduce(3),		/* ), reduce: BoolExpr1 */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(1),		/* &, reduce: BoolExpr */
			reduce(1),		/* |, reduce: BoolExpr */
			nil,		/* ( */
			reduce(4),		/* ), reduce: BoolExpr1 */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(5),		/* &, reduce: BoolExpr1 */
			reduce(5),		/* |, reduce: BoolExpr1 */
			nil,		/* ( */
			reduce(5),		/* ), reduce: BoolExpr1 */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(10),		/* &, reduce: CompareExpr */
			reduce(10),		/* |, reduce: CompareExpr */
			nil,		/* ( */
			reduce(10),		/* ), reduce: CompareExpr */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(11),		/* &, reduce: CompareExpr */
			reduce(11),		/* |, reduce: CompareExpr */
			nil,		/* ( */
			reduce(11),		/* ), reduce: CompareExpr */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* S' */
			nil,		/* BoolExpr */
			nil,		/* BoolExpr1 */
			nil,		/* Val */
			reduce(12),		/* &, reduce: SubStringExpr */
			reduce(12),		/* |, reduce: SubStringExpr */
			nil,		/* ( */
			reduce(12),		/* ), reduce: SubStringExpr */
			nil,		/* true */
			nil,		/* false */
			nil,		/* CompareExpr */
			nil,		/* SubStringExpr */
			nil,		/* int_lit */
			nil,		/* < */
			nil,		/* > */
			nil,		/* string_lit */
			nil,		/* in */
			
		},

	},
	
}
