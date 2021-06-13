package main

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestIntMinBasic(t *testing.T) {
//	min := IntMin(2, -2)
//	if min != -2 {
//		t.Errorf("IntMin(2, -2) = %d; want -2", min)
//	}
//}
//
//func TestIntMinTableDriven(t *testing.T) {
//	var tests = []struct{
//		a, b int
//		want int
//	} {
//		{0,1,0},
//		{1,0,0},
//		{1,2,1},
//		{-1,0,-1},
//	}
//
//	for _, tt := range tests {
//		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
//		t.Run(testname, func(t *testing.T) {
//			ans := IntMin(tt.a, tt.b)
//			if ans != tt.want {
//				t.Errorf("IntMin(%d, %d) = %d; want %d", tt.a, tt.b, ans, tt.want)
//			}
//		})
//	}
//}
