package tools

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
)

func ToIntSlice(filename string) ([]int,error)  {
	b,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(b))

	scanner.Split(bufio.ScanLines)
	var res []int

	for scanner.Scan() {
		v,err :=strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res,v)
	}
	return res,nil
}

func ToStringSlice(filename string) ([]string,error) {
	b,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(b))

	scanner.Split(bufio.ScanLines)
	var res []string

	for scanner.Scan() {

		res = append(res,scanner.Text())
	}
	return res,nil
}