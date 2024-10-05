package main
import "fmt"

var p [2]int = [2]int{1, 1}
var n [2]int = [2]int{0, 1}
var e [2]int = [2]int{-1, 0}
var H [2]int = [2]int{1, 1}
var He [2]int = [2]int{2, 4}
var Li [2]int = [2]int{3, 7}
var Be [2]int = [2]int{4, 9}
var B [2]int = [2]int{5, 11}
var C [2]int = [2]int{6, 12}
var N [2]int = [2]int{7, 14}
var O [2]int = [2]int{8, 16}
var F [2]int = [2]int{9, 19}
var Ne [2]int = [2]int{10, 20}
var Na [2]int = [2]int{11, 23}
var Mg [2]int = [2]int{12, 24}
var Al [2]int = [2]int{13, 27}
var Si [2]int = [2]int{14, 28}
var P [2]int = [2]int{15, 31}
var S [2]int = [2]int{16, 32}
var Cl [2]int = [2]int{17, 35}
var Ar [2]int = [2]int{18, 40}
var K [2]int = [2]int{19, 39}
var Ca [2]int = [2]int{20, 40}
var Sc [2]int = [2]int{21, 45}
var Ti [2]int = [2]int{22, 48}
var V [2]int = [2]int{23, 51}
var Cr [2]int = [2]int{24, 52}
var Mn [2]int = [2]int{25, 55}
var Fe [2]int = [2]int{26, 56}
var Co [2]int = [2]int{27, 59}
var Ni [2]int = [2]int{28, 59}
var Cu [2]int = [2]int{29, 64}
var Zn [2]int = [2]int{30, 65}
var Ga [2]int = [2]int{31, 70}
var Ge [2]int = [2]int{32, 73}
var As [2]int = [2]int{33, 75}
var Se [2]int = [2]int{34, 79}
var Br [2]int = [2]int{35, 80}
var Kr [2]int = [2]int{36, 84}
var Rb [2]int = [2]int{37, 85}
var Sr [2]int = [2]int{38, 88}
var Y [2]int = [2]int{39, 89}
var Zr [2]int = [2]int{40, 91}
var Nb [2]int = [2]int{41, 93}
var Mo [2]int = [2]int{42, 96}
var Tc [2]int = [2]int{43, 98}
var Ru [2]int = [2]int{44, 101}
var Rh [2]int = [2]int{45, 103}
var Pd [2]int = [2]int{46, 106}
var Ag [2]int = [2]int{47, 108}
var Cd [2]int = [2]int{48, 112}
var In [2]int = [2]int{49, 115}
var Sn [2]int = [2]int{50, 119}
var Sb [2]int = [2]int{51, 122}
var Te [2]int = [2]int{52, 128}
var I [2]int = [2]int{53, 127}
var Xe [2]int = [2]int{54, 131}
var Cs [2]int = [2]int{55, 133}
var Ba [2]int = [2]int{56, 137}
var La [2]int = [2]int{57, 139}
var Ce [2]int = [2]int{58, 140}
var Pr [2]int = [2]int{59, 141}
var Nd [2]int = [2]int{60, 144}
var Pm [2]int = [2]int{61, 145}
var Sm [2]int = [2]int{62, 150}
var Eu [2]int = [2]int{63, 152}
var Gd [2]int = [2]int{64, 157}
var Tb [2]int = [2]int{65, 159}
var Dy [2]int = [2]int{66, 162}
var Ho [2]int = [2]int{67, 165}
var Er [2]int = [2]int{68, 167}
var Tm [2]int = [2]int{69, 169}
var Yb [2]int = [2]int{70, 173}
var Lu [2]int = [2]int{71, 175}
var Hf [2]int = [2]int{72, 178}
var Ta [2]int = [2]int{73, 181}
var W [2]int = [2]int{74, 184}
var Re [2]int = [2]int{75, 186}
var Os [2]int = [2]int{76, 190}
var Ir [2]int = [2]int{77, 192}
var Pt [2]int = [2]int{78, 195}
var Au [2]int = [2]int{79, 197}
var Hg [2]int = [2]int{80, 201}
var Tl [2]int = [2]int{81, 204}
var Pb [2]int = [2]int{82, 207}
var Bi [2]int = [2]int{83, 209}
var Po [2]int = [2]int{84, 209}
var At [2]int = [2]int{85, 210}
var Rn [2]int = [2]int{86, 222}
var Fr [2]int = [2]int{87, 223}
var Ra [2]int = [2]int{88, 226}
var Ac [2]int = [2]int{89, 227}
var Th [2]int = [2]int{90, 232}
var Pa [2]int = [2]int{91, 231}
var U [2]int = [2]int{92, 238}
var Np [2]int = [2]int{93, 237}
var Pu [2]int = [2]int{94, 244}
var Am [2]int = [2]int{95, 243}
var Cm [2]int = [2]int{96, 247}
var Bk [2]int = [2]int{97, 247}
var Cf [2]int = [2]int{98, 251}
var Es [2]int = [2]int{99, 252}
var Fm [2]int = [2]int{100, 257}
var Md [2]int = [2]int{101, 258}
var No [2]int = [2]int{102, 259}
var Lr [2]int = [2]int{103, 262}
var Rf [2]int = [2]int{104, 267}
var Db [2]int = [2]int{105, 270}
var Sg [2]int = [2]int{106, 271}
var Bh [2]int = [2]int{107, 270}
var Hs [2]int = [2]int{108, 277}
var Mt [2]int = [2]int{109, 278}
var Ds [2]int = [2]int{110, 281}
var Rg [2]int = [2]int{111, 282}
var Cn [2]int = [2]int{112, 285}
var Nh [2]int = [2]int{113, 286}
var Fl [2]int = [2]int{114, 289}
var Mc [2]int = [2]int{115, 290}
var Lv [2]int = [2]int{116, 293}
var Ts [2]int = [2]int{117, 294}
var Og [2]int = [2]int{118, 294}

var eq[2]int
var answ string
var l int = 0
var b [2]int
var x string
var y string
var u string
var q string
var o string
var a [2]int
var z [2]int
var w [2]int
var d [2]int

func main(){
  fmt.Print("\033[H\033[2J")
  fmt.Scan(&x)
  detect(x)
  fmt.Println("+")
  fmt.Scan(&y)
  detect(y)
  fmt.Println("->")
  fmt.Scan(&u)
  detect(u)
  fmt.Println("+")

  calculate(a, z)

  fmt.Print("\033[H\033[2J")

  fmt.Print(x)
  fmt.Print(" + ")
  fmt.Print(y)
  fmt.Print(" -> ")
  fmt.Print(u)
  fmt.Print(" + ")
  fmt.Print(answ)
  fmt.Println(" ")
}

func calculate(n [2]int, m [2]int){

  eq[0] = n[0] + m[0] - w[0]
  eq[1] = n[1] + m[1] - w[1]

  examine(eq)
}

func detect(h string){

  switch h {
  case "e":
    b = e
  case "n":
    b = n
  case "p":
    b = p
  case "H":
      b = H
  case "He":
      b = He
  case "Li":
      b = Li
  case "Be":
      b = Be
  case "B":
      b = B
  case "C":
      b = C
  case "N":
      b = N
  case "O":
      b = O
  case "F":
      b = F
  case "Ne":
      b = Ne
  case "Na":
      b = Na
  case "Mg":
      b = Mg
  case "Al":
      b = Al
  case "Si":
      b = Si
  case "P":
      b = P
  case "S":
      b = S
  case "Cl":
      b = Cl
  case "Ar":
      b = Ar
  case "K":
      b = K
  case "Ca":
      b = Ca
  case "Sc":
      b = Sc
  case "Ti":
      b = Ti
  case "V":
      b = V
  case "Cr":
      b = Cr
  case "Mn":
      b = Mn
  case "Fe":
      b = Fe
  case "Co":
      b = Co
  case "Ni":
      b = Ni
  case "Cu":
      b = Cu
  case "Zn":
      b = Zn
  case "Ga":
      b = Ga
  case "Ge":
      b = Ge
  case "As":
      b = As
  case "Se":
      b = Se
  case "Br":
      b = Br
  case "Kr":
      b = Kr
  case "Rb":
      b = Rb
  case "Sr":
      b = Sr
  case "Y":
      b = Y
  case "Zr":
      b = Zr
  case "Nb":
      b = Nb
  case "Mo":
      b = Mo
  case "Tc":
      b = Tc
  case "Ru":
      b = Ru
  case "Rh":
      b = Rh
  case "Pd":
      b = Pd
  case "Ag":
      b = Ag
  case "Cd":
      b = Cd
  case "In":
      b = In
  case "Sn":
      b = Sn
  case "Sb":
      b = Sb
  case "Te":
      b = Te
  case "I":
      b = I
  case "Xe":
      b = Xe
  case "Cs":
      b = Cs
  case "Ba":
      b = Ba
  case "La":
      b = La
  case "Ce":
      b = Ce
  case "Pr":
      b = Pr
  case "Nd":
      b = Nd
  case "Pm":
      b = Pm
  case "Sm":
      b = Sm
  case "Eu":
      b = Eu
  case "Gd":
      b = Gd
  case "Tb":
      b = Tb
  case "Dy":
      b = Dy
  case "Ho":
      b = Ho
  case "Er":
      b = Er
  case "Tm":
      b = Tm
  case "Yb":
      b = Yb
  case "Lu":
      b = Lu
  case "Hf":
      b = Hf
  case "Ta":
      b = Ta
  case "W":
      b = W
  case "Re":
      b = Re
  case "Os":
      b = Os
  case "Ir":
      b = Ir
  case "Pt":
      b = Pt
  case "Au":
      b = Au
  case "Hg":
      b = Hg
  case "Tl":
      b = Tl
  case "Pb":
      b = Pb
  case "Bi":
      b = Bi
  case "Po":
      b = Po
  case "At":
      b = At
  case "Rn":
      b = Rn
  case "Fr":
      b = Fr
  case "Ra":
      b = Ra
  case "Ac":
      b = Ac
  case "Th":
      b = Th
  case "Pa":
      b = Pa
  case "U":
      b = U
  case "Np":
      b = Np
  case "Pu":
      b = Pu
  case "Am":
      b = Am
  case "Cm":
      b = Cm
  case "Bk":
      b = Bk
  case "Cf":
      b = Cf
  case "Es":
      b = Es
  case "Fm":
      b = Fm
  case "Md":
      b = Md
  case "No":
      b = No
  case "Lr":
      b = Lr
  case "Rf":
      b = Rf
  case "Db":
      b = Db
  case "Sg":
      b = Sg
  case "Bh":
      b = Bh
  case "Hs":
      b = Hs
  case "Mt":
      b = Mt
  case "Ds":
      b = Ds
  case "Rg":
      b = Rg
  case "Cn":
      b = Cn
  case "Nh":
      b = Nh
  case "Fl":
      b = Fl
  case "Mc":
      b = Mc
  case "Lv":
      b = Lv
  case "Ts":
      b = Ts
  case "Og":
      b = Og 
  }

  if l == 0 {
    a = b
  } else if l == 1 {
    z = b
  } else if l == 2 {
    w = b
  }

  l = l + 1
}

func examine(d [2]int){

  switch d[0] {
  case -1:
    o = "e"
  case 0:
    o = "n"
  case 1:
    o = "p/H"
  case 2:
    o = "He"
  case 3:
    o = "Li"
  case 4:
    o = "Be"
  case 5:
    o = "B"
  case 6:
    o = "C"
  case 7:
    o = "N"
  case 8:
    o = "O"
  case 9:
    o = "F"
  case 10:
    o = "Ne"
  case 11:
    o = "Na"
  case 12:
    o = "Mg"
  case 13:
    o = "Al"
  case 14:
    o = "Si"
  case 15:
    o = "P"
  case 16:
    o = "S"
  case 17:
    o = "Cl"
  case 18:
    o = "Ar"
  case 19:
    o = "K"
  case 20:
    o = "Ca"
  case 21:
    o = "Sc"
  case 22:
    o = "Ti"
  case 23:
    o = "V"
  case 24:
    o = "Cr"
  case 25:
    o = "Mn"
  case 26:
    o = "Fe"
  case 27:
    o = "Co"
  case 28:
    o = "Ni"
  case 29:
    o = "Cu"
  case 30:
    o = "Zn"
  case 31:
    o = "Ga"
  case 32:
    o = "Ge"
  case 33:
    o = "As"
  case 34:
    o = "Se"
  case 35:
    o = "Br"
  case 36:
    o = "Kr"
  case 37:
    o = "Rb"
  case 38:
    o = "Sr"
  case 39:
    o = "Y"
  case 40:
    o = "Zr"
  case 41:
    o = "Nb"
  case 42:
    o = "Mo"
  case 43:
    o = "Tc"
  case 44:
    o = "Ru"
  case 45:
    o = "Rh"
  case 46:
    o = "Pd"
  case 47:
    o = "Ag"
  case 48:
    o = "Cd"
  case 49:
    o = "In"
  case 50:
    o = "Sn"
  case 51:
    o = "Sb"
  case 52:
    o = "Te"
  case 53:
    o = "I"
  case 54:
    o = "Xe"
  case 55:
    o = "Cs"
  case 56:
    o = "Ba"
  case 57:
    o = "La"
  case 58:
    o = "Ce"
  case 59:
    o = "Pr"
  case 60:
    o = "Nd"
  case 61:
    o = "Pm"
  case 62:
    o = "Sm"
  case 63:
    o = "Eu"
  case 64:
    o = "Gd"
  case 65:
    o = "Tb"
  case 66:
    o = "Dy"
  case 67:
    o = "Ho"
  case 68:
    o = "Er"
  case 69:
    o = "Tm"
  case 70:
    o = "Yb"
  case 71:
    o = "Lu"
  case 72:
    o = "Hf"
  case 73:
    o = "Ta"
  case 74:
    o = "W"
  case 75:
    o = "Re"
  case 76:
    o = "Os"
  case 77:
    o = "Ir"
  case 78:
    o = "Pt"
  case 79:
    o = "Au"
  case 80:
    o = "Hg"
  case 81:
    o = "Tl"
  case 82:
    o = "Pb"
  case 83:
    o = "Bi"
  case 84:
    o = "Po"
  case 85:
    o = "At"
  case 86:
    o = "Rn"
  case 87:
    o = "Fr"
  case 88:
    o = "Ra"
  case 89:
    o = "Ac"
  case 90:
    o = "Th"
  case 91:
    o = "Pa"
  case 92:
    o = "U"
  case 93:
    o = "Np"
  case 94:
    o = "Pu"
  case 95:
    o = "Am"
  case 96:
    o = "Cm"
  case 97:
    o = "Bk"
  case 98:
    o = "Cf"
  case 99:
    o = "Es"
  case 100:
    o = "Fm"
  case 101:
    o = "Md"
  case 102:
    o = "No"
  case 103:
    o = "Lr"
  case 104:
    o = "Rf"
  case 105:
    o = "Db"
  case 106:
    o = "Sg"
  case 107:
    o = "Bh"
  case 108:
    o = "Hs"
  case 109:
    o = "Mt"
  case 110:
    o = "Ds"
  case 111:
    o = "Rg"
  case 112:
    o = "Cn"
  case 113:
    o = "Nh"
  case 114:
    o = "Fl"
  case 115:
    o = "Mc"
  case 116:
    o = "Lv"
  case 117:
    o = "Ts"
  case 118:
    o = "Og"
  }

  answ = o

  fmt.Println(answ)
}

