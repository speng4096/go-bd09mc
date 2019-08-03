package bd09mc

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

var vm *goja.Runtime

func getVM() (*goja.Runtime, error) {
	if vm != nil {
		return vm, nil
	}
	tmpVM := goja.New()
	_, err := tmpVM.RunString(script)
	if err != nil {
		return nil, errors.Wrap(err, "加载转换脚本失败")
	}
	vm = tmpVM
	return vm, nil
}

func parseLL(s string) (float64, float64, error) {
	item := strings.Split(s, "|")
	if len(item) != 2 {
		return 0, 0, errors.Errorf("转换结果格式错误: %v", item)
	}
	lng, err := strconv.ParseFloat(item[0], 10)
	if err != nil {
		return 0, 0, errors.Errorf("转换结果无法转换到float64型: %s", item[0])
	}
	lat, err := strconv.ParseFloat(item[1], 10)
	if err != nil {
		return 0, 0, errors.Errorf("转换结果无法转换到float64型: %s", item[1])
	}
	return lng, lat, nil
}

// 百度经纬度坐标 -> 百度墨卡托
func LL2MC(lng, lat float64) (float64, float64, error) {
	vm, err := getVM()
	if err != nil {
		return 0, 0, err
	}
	s := fmt.Sprintf("map.convertLL2MC({lng: %f, lat: %f})", lng, lat)
	value, err := vm.RunString(s)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "转换失败")
	}
	return parseLL(value.String())
}

// 百度墨卡托 -> 百度经纬度坐标
func MC2LL(lng, lat float64) (float64, float64, error) {
	vm, err := getVM()
	if err != nil {
		return 0, 0, err
	}
	s := fmt.Sprintf("map.convertMC2LL({lng: %f, lat: %f})", lng, lat)
	value, err := vm.RunString(s)
	if err != nil {
		return 0, 0, errors.Wrapf(err, "转换失败")
	}
	return parseLL(value.String())
}
