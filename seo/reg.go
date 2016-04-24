package seo

import (
	"regexp"
)

/* ****************************************
 * version: 0.1
 * author: chandlerxue
 * date: 2016年4月19日
 * func:
 * desc: 返回[]string
 * ****************************************/
func Reg( text string, target string , protocol string ) []byte {
	/* 实例化 Seo */
	objSeo := seo()
	/* 取得seo数据 */
	objSeo.setKeywords( "this is a keys for objSeo" )
	objSeo.setDescription( "this is a desps for objSeo" )
	//objSeo.reg = map[string]string{ "A":"This is a", "B":"This is b"}
	objSeo.setReg( map[string]string{
		`href="`:`href=[\"\']` +  protocol + "://" + target ,
		"charset=utf-8": `charset=[a-z0-9]{0,10}`,
	})





	//fmt.Println( "[seo][Reg]" + objSeo.keywords )
	//fmt.Println( "[seo][Reg]" + objSeo.description )
	//fmt.Println( "[seo][Reg]" + objSeo.reg["A"] )
	//fmt.Println( "[seo][Reg]" + objSeo.reg["B"] )

	for i,val := range objSeo.reg{
		reg := regexp.MustCompile( val )
		text = reg.ReplaceAllString(text,  i  )

	}


	return  []byte(text)




	//fmt.Println("[Reg Replaced Text] " + reg.ReplaceAllString(*text, "/yyyy" ) )
}