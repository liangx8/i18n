package i18n_test

import (
	"testing"

	"github.com/liangx8/i18n"
)

func TestTr(t *testing.T) {
	if err := i18n.Register("zh_CN", "Simplified Chinese", i18n.NewJsonDecoder()); err != nil {
		t.Fatal(err)
	}
	if err := i18n.Register("zh_CN", "Simplified Chinese", i18n.NewJsonDecoder()); err == nil {
		t.Fatal("expected a error")
	}
	if err := i18n.AddResource("zh_CN", i18n.ResourceText(`[{"x":"交叉"},{"o":"圆圈"},{"fmt":"num:%d,str:%s"}]`)); err != nil {
		t.Fatal(err)
	}
	if err := i18n.SetLang("zh_CN"); err != nil {
		t.Fatalf("SetLang():%v", err)
	}
	if i18n.Tr("x") != "交叉" {
		t.Fatal("Tr() fail")
	}
	if i18n.Tr("fmt", 1, "str") != "num:1,str:str" {
		t.Fatal("Tr() fail")
	}
}
