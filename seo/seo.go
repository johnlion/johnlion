/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 上午11:56
 * File: seo.go
 * Desc:
 *********************************************/
package seo



import (

)

type Seo struct{
	domain string
	title string
	description string
	keywords string
	reg map[string]string
}


/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年4月24日 下午2:23
 * File:
 * Desc:构造函数
 *********************************************/
func seo() *Seo {
	var seo Seo
	return &seo
}

func ( s *Seo ) setKeywords( keywords string ){
	s.keywords = keywords
}

func ( s *Seo ) setDescription( description string ){
	s.description = description
}

func ( s *Seo ) setReg( reg map[string]string  ){
	s.reg = reg
}






