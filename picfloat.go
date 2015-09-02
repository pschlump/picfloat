// Copyright (C) Philip Schlump, 2014.
/*

Implementation of Picture format for floating point numbers

	##,###,##0.000

	format with commas and 0.

*/
package picfloat

import (
	"fmt"
	"strings"
)

const (
	Version = "0.4.0"
)

/*
	#,###0.000
	1. Round LSD or not?

	# - ' ' or digit
	, -	',' if digit to left fond
	0 - '0' or digit or - if negative
	. - decimal point
	+ - '+' or '-' if negavie - else the +- will cuddle beginning of number

*/

func picFmt(pic string, f float64) string {

	whPic := ""
	fpPic := ""

	hd := false // has decimal point in format
	if len(pic) <= 0 {
		return ""
	}
	if strings.Index(pic, ".") >= 0 {
		pp := strings.Split(pic, ".")
		whPic = pp[0]
		fpPic = ""
		if len(pp) > 1 {
			hd = true
			fpPic = pp[1]
		}
	} else {
		whPic = pic
	}
	fpNd := 0 // # of digits in the fractional part
	for _, v := range fpPic {
		if v == '0' || v == '#' {
			fpNd++
		}
	}
	whNd := 0 // # of digits in the whole part
	for _, v := range whPic {
		if v == '0' || v == '#' {
			whNd++
		}
	}

	fmtS := fmt.Sprintf("%%.%df", fpNd)
	s := fmt.Sprintf(fmtS, f)
	// fmt.Printf ( "s after printing is ->%s<-\n", s )
	p := strings.Split(s, ".")
	wh := p[0]
	fp := ""
	if fpNd > 0 {
		fp = p[1]
	}

	rv := ""

	// fmt.Printf ( "wh  ->%s<- fp ->%s<- ,,, fpNd=%d whNd=%d \n" , wh , fp, fpNd, whNd )
	// fmt.Printf ( "pic ->%s<- whPic ->%s<- fpPic ->%s<- \n" , pic, whPic, fpPic )

	if len(wh) > whNd { // If the number is too large then just return the format.
		return pic
	}

	wh = PadStr(whNd, " ", wh)
	// wh = PadStr ( whNd, wh )

	j := 0
	ds := false
	for i := 0; i < len(whPic); i++ {
		switch whPic[i : i+1] {
		case ",":
			if ds {
				rv += ","
			} else {
				rv += " "
			}
		case "#":
			if !ds && strings.Index("0123456789", wh[j:j+1]) >= 0 {
				ds = true
			}
			rv += wh[j : j+1]
			j++
		case "0":
			if !ds && strings.Index("0123456789", wh[j:j+1]) >= 0 {
				ds = true
			}
			if wh[j:j+1] == " " {
				rv += "0"
			} else {
				rv += wh[j : j+1]
			}
			j++
		}
	}

	rv = strings.Replace(rv, "- ", " -", 1)

	if hd { // If it has a decimal point
		rv += "." // add the d.p. to the string
		j = 0
		for i := 0; i < len(fpPic); i++ {
			switch fpPic[i : i+1] {
			case ",":
				rv += ","
			case "#": // Kind of odd to use '#' to the right of the d.p., but...
				// maybee we should treat exactly same as '0' after decimal point?
				rv += fp[j : j+1]
				j++
			case "0":
				if fp[j:j+1] == " " {
					rv += "0"
				} else {
					rv += fp[j : j+1]
				}
				j++
			}
		}
	}

	return rv

}

func Format(format string, value float64) string {
	rv := picFmt(format, value)
	return rv
}

func PadStr2(n int, s string) (r string) {
	r = fmt.Sprintf(fmt.Sprintf("%%%ds", n), s)
	return
}

func PadStr(l int, w string, s string) string {
	if len(s) >= l {
		return s
	}
	k := l - len(s)
	t := ""
	for i := 0; i < k; i++ {
		t += w
	}
	return t + s
}

/* vim: set noai ts=4 sw=4: */
